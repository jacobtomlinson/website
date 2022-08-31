---
author: Jacob Tomlinson
date: 2017-01-16T00:00:00+00:00
categories:
  - Monitoring
tags:
  - telegraf
  - freebsd
  - influxdata
  - 32bit
  - i386
  - docker
thumbnail: influxdata
title: Building Telegraf for 32bit FreeBSD
aliases:
  - /2017/01/16/build-telegraf-for-i386-32bit-pfsense-freebsd/
---


## Introduction

Currently InfluxData do not provide a 32bit FreeBSD build of [Telegraf](https://www.influxdata.com/time-series-platform/telegraf/) as part of their standard packages. Luckily it is easy to build yourself.

The only requirement for the following steps is golang 1.5+. To make my life simpler I'm going to create an interactive session in a [golang Docker container](https://hub.docker.com/_/golang/), this way I know exactly what dependancies I have available to me. I'll also create a volume in a location where I want the packaged compiled binary to be left once the container is gone.

```
docker run --rm -ti -v /path/to/output/directory:/output golang /bin/bash
```

Now that we are in an environment with golang 1.5+ we should make a note of the version we want to build. As this article may go stale you should just update this number with the current stable version.

Set the version we want to build.
```
export VERSION=1.1.2
```

## Compiling the binary

Get the source, checkout the version we want.
```
go get github.com/influxdata/telegraf
cd /go/src/github.com/influxdata/telegraf
git checkout ${VERSION}
```

Build for FreeBSD i386.
```
env GOOS=freebsd GOARCH=386 make
```

## Packaging the binary

Change to the output directory.
```
cd /output
```

Get the 64bit package which includes the log rotation, startup script, etc and extract it.
```
wget https://dl.influxdata.com/telegraf/releases/telegraf-${VERSION}_freebsd_amd64.tar.gz
tar -xvzf telegraf-${VERSION}_freebsd_amd64.tar.gz
rm telegraf-${VERSION}_freebsd_amd64.tar.gz
```

Replace the 64bit binary with the 32bit one and archive it again.
```
mv /go/bin/freebsd_386/telegraf telegraf/usr/bin/telegraf
tar -cvzf telegraf-${VERSION}_freebsd_i386.tar.gz telegraf
rm -rf telegraf
```

## Conclusion
You should now have a tar archive in your output directory ready to move onto your FreeBSD machine and extract.
