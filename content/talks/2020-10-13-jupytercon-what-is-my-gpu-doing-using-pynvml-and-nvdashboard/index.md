---
title: "What is my GPU doing? Using PyNVML and NVDashboard to access GPU metrics"
date: 2020-10-13T00:00:00+00:00
draft: true
author: "Jacob Tomlinson"
event:
  name: JupyterCon
  link: https://cfp.jupytercon.com/2020/schedule/presentation/221/what-is-my-gpu-doing-using-pynvml-and-the-nvdashboard-jupyter-lab-extension-to-access-gpu-metrics/
  location: Online
  type: Talk
length: 25
abstract: true
video: https://www.youtube.com/watch?v=5t1HIR2xBTc
slides: null
---

With the growing number of GPU accelerated Python packages it is more important than ever to understand what your GPU devices are doing. You need to know how they are performing and how well they are being utilized. This talk will demonstrate a couple of open-source Python packages to help you access utilization information programmatically and view the data in realtime dashboards.

This talk will walk you through collecting and viewing various utilization information from NVIDIA GPU devices with PyNVML and NVDashboard. Understanding what your GPUs are doing and how well they are being utilized can be tricky. Until now this information is generally accessed via the command line but with these libraries you can access this information in your browser, or in your code.

PyNVML is an open-source Python wrapper for the NVIDIA Management Library (NVML), which is a C-based API for monitoring and managing various states of NVIDIA GPU devices. NVML is directly used by the better-known NVIDIA System Management Interface (nvidia-smi).

NVDashboard is an open-source browser-based dashboard written in Python using Bokeh. It displays data collected with PyNVML in various time series and utilization plots. There is also a Jupyter Lab extension which enables you to view these dashboards in your Jupyter session right next to your notebooks.

We will run through various examples of using these libraries, from installing NVDashboard and getting it running in your Jupyter session to building your own simple diagnostics tools with PyNVML. It doesn’t matter whether you are new to GPU development in Python or an experienced user, by the end of this session you will leave with a good understanding of getting detailed information about what your hardware is up to.
