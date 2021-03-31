---
title: "Talk: Intro to distributed computing on GPUs with Dask in Python"
date: 2021-03-31T00:00:00+00:00
draft: false
author: "Jacob Tomlinson"
categories:
  - talks
tags:
  - Python
  - RAPIDS
  - GPUs
  - Dask
  - BlazingSQL
  - BlazingDB
  - Talk
  - Public Speaking
  - Slides
thumbnail: dask
---

This week I presented a talk on using [Dask](https://dask.org/) with [RAPIDS](https://rapids.ai/) as part of a [BlazingSQL](https://blazingsql.com/) webinar series.

## Abstract

[RAPIDS](https://rapids.ai/) is an end to end data science stack built entirely for CUDA GPUs. Faster analytics, at scale, for lower total cost of ownership. [Dask](https://dask.org/) natively scales Python and the RAPIDS ecosystem stack onto multiple servers and GPUs, supporting unprecedented scale. [BlazingSQL](https://blazingsql.com/) is a distributed SQL engine built in Python. It performs incredibly fast SQL queries on the RAPIDS DataFrame and ensures optimal usage of GPU primitives.

BlazingSQL is built using the same libraries underpinning the RAPIDS ecosystem. As RAPIDS improves, so does BlazingSQL.

This workshop runs through the basics of using Dask alongside RAPIDS to perform out-of-core distributed GPU computations in Python. We cover high level APIs such as [DataFrames](https://docs.dask.org/en/latest/dataframe.html) and [Arrays](https://docs.dask.org/en/latest/array.html) and then dive under the covers to explore [delayed functions](https://docs.dask.org/en/latest/delayed.html) and [distributed futures](https://docs.dask.org/en/latest/futures.html).

## Video

<iframe width="100%" height="500" src="https://www.youtube.com/embed/py1YPs6s6so\" frameborder="0" allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>

## Notebook

<script src="https://gist.github.com/jacobtomlinson/6f16abb716f50f81a6687bd67efd2f61.js"></script>
