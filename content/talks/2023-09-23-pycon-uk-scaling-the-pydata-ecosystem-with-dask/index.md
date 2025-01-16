---
title: "Scaling the PyData Ecosystem with Dask"
date: 2023-09-23T00:00:00+00:00
draft: false
author: "Jacob Tomlinson"
event:
  name: PyCon UK 2023
  link: https://pretalx.com/pyconuk-2023/talk/EYTZY3/
  type: Talk
  location: Cardiff, UK
length: 90
abstract: true
---

Dask is a flexible library for parallel computing in Python. Dask provides high-level interfaces to extend the PyData ecosystem to larger-than-memory or distributed environments, as well as lower-level interfaces to customise workflows. No previous experience is required, though knowledge of Python, NumPy and pandas is preferred.

This tutorial is aimed at people who are new to parallel and distributed computing, and to using PyData libraries at scale. The primary focus is on Dask’s commonly used features and building an intuition for distributed workflows. However, since Dask is a powerful tool with a lot more capabilities, we also include an overview of the breadth of APIs and deployment options provided by Dask.

Tutorial outline :

1. Overview [10 minutes]

Quick showcase of using Dask for a larger-than-memory dataset.
Give a brief overview of what Dask is, why it's needed, and where it fits into the PyData ecosystem.
How to start a Dask client and access the dashboard.

2. Diagnostic tools for understanding performance [10 mins]

Need for visualization tools in distributed computing.
Quick overview of visualizing task graphs and widgets for local computations.
Deep dive into Dask Dashboard for distributed computations.

3. Parallel DataFrames using Dask DataFrame [30 mins ]

Introduction to Dask DataFrames. Discuss how Dask DataFrame mirrors pandas.
Common pandas operations with Dask (reading, groupby, apply, map_partitons).

4. Parallel Arrays using Dask Array [ 30 mins ]

Introduction to Dask Arrays. Discuss how Dask collections mirror their single-threaded counterparts (e.g. Dask Array mirror NumPy).
Visualizing task graphs and using widgets to understand the computation in detail.
Best Practices around choosing chunks.

5. Brief introduction to Delayed, Distributed, and Deployments [ 10 minutes ]

When to choose low-level collections.
What are the Delayed and Futures APIs?
Dask’s LocalCluster and other deployment options.
Additional resources to learn more about the topics in this section.

