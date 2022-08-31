---
title: "Monitoring Dask + RAPIDS with Prometheus + Grafana"
date: 2021-04-09T00:00:00+00:00
draft: false
author: "Jacob Tomlinson"
canonical: https://medium.com/rapids-ai/monitoring-dask-rapids-with-prometheus-grafana-96eaf6b8f3a0
canonical_title: the NVIDIA RAPIDS Blog
categories:
  - blog
tags:
  - RAPIDS
  - Dask
  - Monitoring
  - Prometheus
  - Grafana
---

[Prometheus](https://prometheus.io) is a popular monitoring tool within the cloud community. It has out-of-the-box integration with popular platforms including [Kubernetes](https://kubernetes.io/), [Open Stack](https://www.openstack.org/), and the [major cloud vendors](https://prometheus.io/docs/prometheus/latest/configuration/configuration/), and integrates with dashboarding tools like [Grafana](https://grafana.com/).

In this post, we will explore installing and configuring a minimal Prometheus service with Grafana as our front end and using it to monitor [RAPIDS](https://rapids.ai/).

## Prometheus overview

![A diagram of the prometheus stack. Showing prometheus at the centre with pull metrics gathering data, push alerts sending out warnings and notifications and Grafana querying the database for information.](https://miro.medium.com/max/2702/0*OV5dDGK3XZU3CUn3)

Source: <https://prometheus.io/docs/introduction/overview/>

At its core, Prometheus is a time-series database for storing system and application metrics. It gathers metrics by polling metric exporters periodically and then allows you to query those metrics with [PromQL](https://prometheus.io/docs/prometheus/latest/querying/basics/).

It also has additional services such as **pushgateway**, for short-lived jobs, and **alertmanager** for notifying operators of issues based on metric rules.

## Exporting metrics

Exporting metrics from a system or application is either done by a standalone exporter or by the application itself.

Commonly, metrics are made available in a [text format](https://prometheus.io/docs/instrumenting/exposition_formats/#text-based-format) that is accessible at a `/metrics` endpoint. Applications that are already serving HTTP traffic can make this available directly with the help of [client libraries](https://prometheus.io/docs/instrumenting/clientlibs/), while other services may need a companion exporter which runs a web server and gathers data in a native way before exporting.

Examples of stand-alone exporters are the [node_exporter](https://github.com/prometheus/node_exporter) which queries hardware and OS level metrics from *nix systems. The [mysqld_exporter](https://github.com/prometheus/mysqld_exporter) is another example that runs alongside a MySQL database using the database connection to gather metrics about the database server itself.

**To instrument RAPIDS we care about exporting three sets of metrics:**

-   System metrics via the [node_exporter](https://github.com/prometheus/node_exporter).
-   GPU metrics via the [DCGM-exporter](https://github.com/NVIDIA/gpu-monitoring-tools#dcgm-exporter).
-   [Dask](https://dask.org/) metrics that are [natively exposed via the Dask dashboard's web server](https://docs.dask.org/en/latest/setup/prometheus.html).

## Service discovery

When running Prometheus at scale it is common to use service discovery to allow Prometheus to automatically discover metrics endpoints.

When running Prometheus on Kubernetes, for example, the [Kubernetes service discovery](https://prometheus.io/docs/prometheus/latest/configuration/configuration/#kubernetes_sd_config) will use the Kubernetes API to discover all running HTTP services and will attempt to call the `/metrics` endpoint on each service looking for metrics.

It is also possible to write custom service discovery by either using [DNS records](https://prometheus.io/docs/prometheus/latest/configuration/configuration/#dns_sd_config), a key/value store like [Consul](https://prometheus.io/docs/prometheus/latest/configuration/configuration/#consul_sd_config), or simply [a text file](https://prometheus.io/docs/prometheus/latest/configuration/configuration/#file_sd_config) that is periodically updated with endpoints.

For simplicity, we will manually configure Prometheus to monitor a single host running RAPIDS.

## Installing our components

For this example, we will run RAPIDS on a Ubuntu 20.04 workstation with two NVIDIA GPUs, the latest NVIDIA drivers, and [NVIDIA Docker](https://github.com/NVIDIA/nvidia-docker) installed.

### RAPIDS

To make deployment simple, here we will be using [Docker](https://www.docker.com/) and [Docker Compose](https://docs.docker.com/compose/). Let's start by creating a compose file for RAPIDS.

```yaml
version: "3.9"
  services:
    rapids:
      image: rapidsai/rapidsai:0.18-cuda11.0-runtime-ubuntu16.04-py3.8
        ports:
          - "8888:8888"  # Jupyter
          - "8786:8786"  # Dask communication
          - "8787:8787"  # Dask dashboard
        environment:
            JUPYTER_FG: "true"
        deploy:
          resources:
            reservations:
              devices:
              - capabilities: [gpu]
```

Here we are defining one container running the RAPIDS image, exposing all the necessary ports, setting Jupyter to run as our foreground process, and allowing access to all our GPUs.

Then we can get RAPIDS up and running.

```
docker-compose up -d
```

Now we should be able to access port `8888` in our browser to view [Jupyter Lab](https://jupyterlab.readthedocs.io/en/stable/index.html).

![A screenshot of Jupyter Lab running in the RAPIDS Docker image.](https://miro.medium.com/max/2882/0*8qg6foILhhdGDdQ-)

Next, let's start our Dask cluster. You can do this in a notebook or via the [Dask Jupyter Lab Extension](https://github.com/dask/dask-labextension). Let's click the Dask logo on the side and click `NEW`.

![A screenshot of a Dask cluster being created using the Jupyter Lab Dask Extension.](https://miro.medium.com/max/2882/0*fa_zXZloWCWFBtqE)

Now if we visit port `8787` in our browser, we will see the Dask dashboard.

![A screenshot of the Dask dashboard.](https://miro.medium.com/max/2882/0*coH2UYT9VxeUaMlr)

And if we visit the `/metrics` endpoint, we will see text format metrics that we can scrape with Prometheus.

![A screenshot of the Dask prometheus metrics endpoint.](https://miro.medium.com/max/2882/0*YS_lUvK6om_IasWI)

## Start Prometheus

Next, let's start Prometheus and have it scrape these Dask metrics.

In our current directory, we will create a new directory with `mkdir prometheus` and within that a config file called `prometheus.yml` with the following contents.

```yaml
global:
  scrape_interval: 15s

scrape_configs:
- job_name: rapids
  static_configs:
  - targets: ['10.51.100.43:8787']
```

The IP here is the IP of the workstation on our LAN, so update it to be whatever yours is.

Then let's add another service to our Docker Compose file.

```yaml
  prometheus:
    image: prom/prometheus:latest
    volumes:
      - ./prometheus/:/etc/prometheus/
    ports:
      - "9090:9090"
```

Then run `docker-compose up -d` again.

Now we can head to port `9090` on our system and try out some PromQL queries in Prometheus. For example, we can get the number of Dask workers in our cluster with `dask_scheduler_workers{job="rapids"}`. Our system has two GPUs so we can see two workers reported here.

![A screenshot of the prometheus dashboard showing a query for the number of Dask workers and a refult of two connected and idle workers.](https://miro.medium.com/max/2882/0*eY7QdSQRAvUlpxQW)

## Collecting more metrics

In addition to our Dask cluster metrics, we also want to collect system and GPU metrics. So let's add those exporters as services in our `docker-compose.yml` file.

```yaml
  node_exporter:
    image: quay.io/prometheus/node-exporter:latest
    command:
      - '--path.rootfs=/host'
    network_mode: host
    pid: host
    volumes:
      - '/:/host:ro,rslave'

  gpu_exporter:
    image: nvcr.io/nvidia/k8s/dcgm-exporter:2.0.13-2.1.2-ubuntu18.04
    ports:
      - "9400:9400"
    deploy:
      resources:
        reservations:
          devices:
          - capabilities: [gpu]
```

Then run `docker-compose up -d` again.

Now we need to update our Prometheus configuration to include these two exporters.

```yaml
global:
  scrape_interval: 15s

scrape_configs:
- job_name: rapids
  static_configs:
  - targets: ['10.51.100.43:8787']
  - targets: ['10.51.100.43:9100']
  - targets: ['10.51.100.43:9400']
```

And then we need to restart Prometheus with `docker-compose restart prometheus`.

Now if we head back to the Prometheus dashboard we can perform a query like `DCGM_FI_DEV_GPU_TEMP` to get our GPU temperatures.

![A screenshot of the prometheus dashboard showing a GPU temperature query and two results of 30 and 32 degrees celcius.](https://miro.medium.com/max/2882/0*LmNirUv-_EEk4XxE)

## Grafana dashboards

Now what we have all our metrics being collected by Prometheus let's install Grafana so we can make plots and dashboards.

We need to create a directory to store Grafana config with `mkdir grafana` and also give it ownership by the Grafana user `sudo chown -R 472 grafana`.

Then let's add one last section to our `docker-compose.yml`.

```yaml
  grafana:
    image: grafana/grafana:latest
    volumes:
      - ./grafana:/var/lib/grafana
    ports:
      - "3000:3000"
```

And start the service with `docker-compose up -d`.

Now we can visit port `3000`, log in with the credentials `admin:admin,` and run through the Grafana first time setup.

![A screenshot of the Grafana Dashboard home page with the first time setup flow visible.](https://miro.medium.com/max/2882/0*lH10btV10QqUbFXP)

We need to tell Grafana about Prometheus to click "Add your first data source" and choose a Prometheus source.

![A screenshot of the Grafana add sources page with prometheus selected.](https://miro.medium.com/max/2882/0*aiDDZ16HdAyRdPDh)

Then input the URL of the Prometheus server and click "Save & Test".

![A screenshot of the Grafana add Prometheus source page with the IP address filled in.](https://miro.medium.com/max/2882/0*CJL-HnQHSP40QOeo)

![A screenshot of the Grafana add Prometheus source page where the "save and test" button has been clicked and a notification saying "Data source is working" is displayed.](https://miro.medium.com/max/2882/0*2YZtpz3v5SolafL4)

Then we can head back to the home page by clicking the Grafana logo and click "Create your first Dashboard".

![A screenshot of blank new dashboard page in Grafana.](https://miro.medium.com/max/2882/0*o1mb4zXNnIjcVXrK)

This gives us a new dashboard with one empty panel, click "Add an empty panel".

Now we can enter a query and configure our plot. For this first example let's query the number of connected Dask workers and display it as a Stat plot. Once you're happy with it click Apply.

![A screenshot of a Grafana plot editor showing a stat plot with the query for number of Dask workers and the plot displays the number two.](https://miro.medium.com/max/2882/0*1pvI9FeSa9tQ93Bs)

We can then keep clicking the "Add Panel" button to create plots for all the metrics we want to have on our dashboard. Take some time to experiment here and see how we can visualize all the data collected by Prometheus.

![A screenshot of our new Grafana dashboard with the number of workers plot on it.](https://miro.medium.com/max/2882/0*lRd7mjDDSfI6dRY2)

**Some good resources when designing dashboards are:**

-   [Prometheus Official Best Practices](https://prometheus.io/docs/practices/consoles/)
-   [Grafana Official Best Practices](https://grafana.com/docs/grafana/latest/best-practices/best-practices-for-creating-dashboards/)
-   [Tips for Designing Grafana Dashboard by Percona](https://www.percona.com/blog/2019/11/22/designing-grafana-dashboards/)
-   [Creating the perfect Grafana dashboard by Logz.io](https://logz.io/blog/creating-the-perfect-grafana-dashboard/)
-   [Grafana dashboards from basic to advanced by Metric Fire](https://www.metricfire.com/blog/grafana-dashboards-from-basic-to-advanced/)

## Conclusion

In this post, we used Prometheus and Grafana to instrument our RAPIDS deployment and display useful metrics on a dashboard. This allows us to gain more insight into our workflows and how they are performing on our system.

You can find the full example config files in [this GitHub Gist](https://gist.github.com/jacobtomlinson/d66766c0773780c4b2e666d2da60199e) and an [interactive example dashboard on RainTank](https://snapshot.raintank.io/dashboard/snapshot/UeFwR3zMvzxwLePJbhYGwf8kJJraSnqc?orgId=2).

![A screenshot of an example RAPIDS Grafana dashboard showing metrics including: GPU temperature, GPU power usage, free disk space, number of Dask tasks in a variety of states, CPU utilization, GPU utilization, host memory and NVLink bandwidth.](https://miro.medium.com/max/3200/0*sNzYijX3Gkgo3axB)

Monitoring with Prometheus scales from a single node to multi-node clusters. While we only covered setting up monitoring on a single node in this post, we intend to cover multi-node Kubernetes deployments in the future.
