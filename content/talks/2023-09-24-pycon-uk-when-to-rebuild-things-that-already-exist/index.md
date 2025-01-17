---
title: "When to rebuild things that already exist"
date: 2023-09-24T00:00:00+00:00
draft: false
author: "Jacob Tomlinson"
event:
  name: PyCon UK 2023
  link: https://pretalx.com/pyconuk-2023/talk/UGV7TJ/
  type: Talk
  location: Cardiff, UK
length: 25
abstract: true
video: https://www.youtube.com/watch?v=f6VtdBLAKn4
slides: https://speakerdeck.com/jacobtomlinson/when-to-rebuild-things-that-already-exist
---

This year I built a library that already exists. The existing solutions didn’t quite meet my needs, I wanted something that ticked all of my boxes. When thinking about building something new people referred me to xkcd #927. But I did it anyway.

For the last 6 years I’ve maintained dask-kubernetes, a Python library for deploying Dask clusters on Kubernetes. In that time I’ve tried nearly every Python Kubernetes client library on PyPI. In fact dask-kubernetes today uses over five different libraries and tools to interact with the Kubernetes API. Each one has different strengths and weaknesses, features and bugs. To satisfy all of the needs of Dask Kubernetes there is no one library that can do it alone.

Should I continue to build wrappers and shims in dask-kubernetes to homogenize the various dependencies? Should I contribute to an existing one to fill in the blanks? Or can I build one library to rule them all?

Earlier this year I decided to build exactly the library I needed. Not a perfect universal library to supersede everything, not a wrapper for everything that exists. Just the library I need to solve my problems, to reduce complexity in my projects and to help me learn the things I need to know to maintain these projects into the future.

In this talk I will dig into my perspective on when to wrap a dependency, when to contribute to a dependency and when to build a new dependency from scratch.
