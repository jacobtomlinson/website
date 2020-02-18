---
title: "Talk: RAPIDS - Open GPU Data Science"
date: 2020-02-18T00:00:00+00:00
draft: false
author: "Jacob Tomlinson"
categories:
  - talks
tags:
  - Python
  - RAPIDS
  - GPUs
  - Talk
  - Public Speaking
  - Slides
  - PyData Cardiff
thumbnail: gpu
---

Last week I was fortunate to speak at the [PyData Cardiff meetup](https://www.meetup.com/PyData-Cardiff-Meetup/).

I presented an overview of [RAPIDS](https://rapids.ai/), a suite of open source software libraries which give you the freedom to execute end-to-end data science and analytics pipelines entirely on GPUs. Much can change between releases of RAPIDS so the latest overview of libraries, benchmarks, and updates is consolidated in a [release deck](https://docs.rapids.ai/overview). The deck I presented at PyData Cardiff was a variation of the 0.11 deck which I remixed to appeal to the audience of the meetup, which was mainly data scientists, individual researchers and students.

## Abstract

The RAPIDS suite of open source software libraries (https://rapids.ai/) allow you to run data science and analytics pipelines entirely on GPUs, but following familiar Python APIs including Numpy, Pandas and SciKit Learn.

RAPIDS relies on NVIDIA® CUDA® primitives for low-level compute optimization, but exposes that GPU parallelism and high-bandwidth memory speed through user-friendly Python interfaces.

RAPIDS also focuses on common data preparation tasks for analytics and data science. This includes a familiar DataFrame API that integrates with a variety of machine learning algorithms for end-to-end pipeline accelerations without paying typical serialization costs. RAPIDS also includes support for multi-node, multi-GPU deployments, enabling vastly accelerated processing and training on much larger dataset sizes.

## Video

<iframe width="100%" height="500" src="https://www.youtube.com/embed/mL4hMoOzIiQ?start=456" frameborder="0" allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>

## Slides

<script async class="speakerdeck-embed" data-id="59f0a5c182e04f7b876d61ed988672ec" data-ratio="1.77777777777778" src="//speakerdeck.com/assets/embed.js"></script>
