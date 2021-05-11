---
title: "Intro to distributed computing on GPUs with Dask in Python"
date: 2021-03-31T00:00:00+00:00
draft: false
author: "Jacob Tomlinson"
event:
  name: BlazingSQL Webinars
  link: https://www.youtube.com/channel/UCjxMdJuZ6d-c-X6xiWpN1eg
  type: Webinar
  location: Online
length: 60
video: https://www.youtube.com/embed/py1YPs6s6so
abstract: true
resources: https://gist.github.com/jacobtomlinson/6f16abb716f50f81a6687bd67efd2f61
---

[RAPIDS](https://rapids.ai/) is an end to end data science stack built entirely for CUDA GPUs. Faster analytics, at scale, for lower total cost of ownership. [Dask](https://dask.org/) natively scales Python and the RAPIDS ecosystem stack onto multiple servers and GPUs, supporting unprecedented scale. [BlazingSQL](https://blazingsql.com/) is a distributed SQL engine built in Python. It performs incredibly fast SQL queries on the RAPIDS DataFrame and ensures optimal usage of GPU primitives.

BlazingSQL is built using the same libraries underpinning the RAPIDS ecosystem. As RAPIDS improves, so does BlazingSQL.

This workshop runs through the basics of using Dask alongside RAPIDS to perform out-of-core distributed GPU computations in Python. We cover high level APIs such as [DataFrames](https://docs.dask.org/en/latest/dataframe.html) and [Arrays](https://docs.dask.org/en/latest/array.html) and then dive under the covers to explore [delayed functions](https://docs.dask.org/en/latest/delayed.html) and [distributed futures](https://docs.dask.org/en/latest/futures.html).
