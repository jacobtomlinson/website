---
title: "High Throughput Computing with Dask: Intro Tutorial"
date: 2021-01-21T00:00:00+00:00
draft: false
author: "Jacob Tomlinson"
event:
  name: CECAM Dask on HPC Seminar Series
  link: https://www.cecam.org/workshop-details/1022
  location: Online
  type: Tutorial
length: 120
abstract: true
video: https://www.youtube.com/watch?v=Tl8rO-baKuY
slides: null
resources: https://github.com/jacobtomlinson/dask-video-tutorial-2020
---

High-throughput (task-based) computing is a flexible approach to parallelization. It involves splitting a problem into loosely-coupled tasks. A scheduler then orchestrates the parallel execution of those tasks, allowing programs to adaptively scale their resource usage. Individual tasks may themselves be parallelized using MPI or OpenMP, and the high-throughput approach can therefore enable new levels of scalability.

Dask is a powerful Python tool for task-based computing. The Dask library was originally developed to provide parallel and out-of-core versions of common data analysis routines from data analysis packages such as NumPy and Pandas. However, the flexibility and usefulness of the underlying scheduler has led to extensions that enable users to write custom task-based algorithms, and to execute those algorithms on high-performance computing (HPC) resources.

This workshop will be a series of virtual seminars/tutorials on tools in the Dask HPC ecosystem. All sessions will be held online, with a live Zoom for some registered participants and a live YouTube stream for the public.
