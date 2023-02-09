---
title: "Deploying multi-gpu workloads on Kubernetes in Python"
date: 2023-02-02T00:00:00+00:00
draft: false
author: "Jacob Tomlinson"
event:
  name: PyData DC Meetup
  link: https://www.meetup.com/pydatadc/events/290894158/
  type: Talk
  location: Online
length: 30
abstract: true
video: https://www.youtube.com/watch?v=jkw8kFVEKb0
slides: https://speakerdeck.com/jacobtomlinson/deploying-multi-gpu-workloads-on-kubernetes-in-python
---

The RAPIDS suite of open source software libraries gives you the freedom to execute end-to-end data science and analytics pipelines entirely on GPUs with minimal code changes and no new tools to learn.

Dask provides advanced parallelism for Python by breaking functions into a task graph that can be evaluated by a task scheduler that has many workers.

By using Dask to scale out RAPIDS workloads on Kubernetes you can accelerate your workloads across many GPUs on many machines. In this talk we will discuss how to install and configure Dask on your Kubernetes cluster and use it to run accelerated GPU workloads on your cluster.
