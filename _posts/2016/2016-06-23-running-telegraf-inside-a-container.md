---
title: 'Running Telegraf inside a docker container'
author: Jacob Tomlinson
layout: post
category: Monitoring
thumbnail: influxdata
tags:
  - telegraf
  - tick
  - influxdata
  - monitoring
  - docker
---

[Telegraf][telegraf] is an application for collecting server and application telemetry and metrics and sending them to a time series datastore like [InfluxDB][influxdb]. Like me you may prefer running all of your applications in [Docker][docker] containers, however this means Telegraf will only collect data for the container. This article will cover the configuration options to allow Telegraf to collect host metrics from inside a container.

## Prerequisites

This article assumes you have already done the following:

 * Installed Docker
 * Installed and configured InfluxDB somewhere
 * Installed and configured you preferred dashboard for InfluxDB ([Grafana][grafana] or [Chronograf][chronograf] are good options)

## Generate config

First we need a copy of the default Telegraf config file to edit. Let's create a directory to store it in and use the Telegraf docker image to run the `-sample-config` command to generate the file and direct the output into our directory.

```
mkdir telegraf
docker run --rm telegraf -sample-config > telegraf/telegraf.conf
```

## Configure output

Open the configuration file `telegraf/telegraf.conf` in you favourite text editor and find the `[[outputs.influxdb]]` section. Change the `url` option to point at your influxdb instance. _If you do not have a dns address for your influxdb host then leave it as `influxdb` and we'll configure it in docker._

```
urls = ["http://influxdb:8086"] # required
```

## Configure inputs

The configuration file will also have configuration sections for the different inputs you wish to collect. By default things like system cpu/memory/network usage are already enabled.

To enable collection of Docker data uncomment `[[inputs.docker]]` and the default `entrypoint` line.

## Run the container

As Telegraf will be run inside a container we need to pass some host resources through including the docker socket, `/proc`, `/sys` and `/etc`. This will allow Telegraf to collect data from the whole host, not just what is visible to the container. We also need to give the container a specific hostname as this will be passed to InfluxDB, if we don't set this the container ID will be used.

_docker command_
```
docker run -d --restart=always --add-host="influxdb:52.18.79.191" --hostname=myhostname -e "HOST_PROC=/rootfs/proc" -e "HOST_SYS=/rootfs/sys" -e "HOST_ETC=/rootfs/etc" -v $(pwd)/telegraf/telegraf.conf:/etc/telegraf/telegraf.conf:ro -v /var/run/docker.sock:/var/run/docker.sock:ro -v /sys:/rootfs/sys:ro -v /proc:/rootfs/proc:ro -v /etc:/rootfs/etc:ro telegraf
```

_docker compose_
```
telegraf:
  image: telegraf
  restart: always
  extra_hosts:
   - "influxdb:52.18.79.191"
  environment:
    HOST_PROC: /rootfs/proc
    HOST_SYS: /rootfs/sys
    HOST_ETC: /rootfs/etc
  hostname: myhostname
  volumes:
   - ./telegraf/telegraf.conf:/etc/telegraf/telegraf.conf:ro
   - /var/run/docker.sock:/var/run/docker.sock:ro
   - /sys:/rootfs/sys:ro
   - /proc:/rootfs/proc:ro
   - /etc:/rootfs/etc:ro
```

## Checking in InfluxDB

Connect to the InfluxDB web GUI at `http://your-influxdb-host:8083/` and select the `telegraf` database from the dropdown at the top. Then from the Query Templates dropdown select `Show Measurements` or run the query `SHOW MEASUREMENTS`.

You should then see a list of the measurements Telegraf is collection that looks like this.

![InfluxDB Telegraf Measurements](http://i.imgur.com/NVBIMLd.png)

[chronograf]: https://influxdata.com/time-series-platform/chronograf/
[docker]: https://www.docker.com/
[grafana]: http://grafana.org/
[influxdb]: https://influxdata.com/
[telegraf]: https://influxdata.com/time-series-platform/telegraf/
