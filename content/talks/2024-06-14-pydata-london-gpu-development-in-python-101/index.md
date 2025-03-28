---
title: "GPU Development in Python 101 "
date: 2024-06-14T00:00:00+00:00
draft: false
author: "Jacob Tomlinson"
event:
  name: PyData London 2024
  link: https://london2024.pydata.org/cfp/talk/LVBAQL/
  type: Tutorial
  location: London, UK
length: 180
video: https://www.youtube.com/watch?v=sMpqsoIvrio
resources: https://github.com/jacobtomlinson/gpu-python-tutorial
abstract: true
featured: true
thumbnail: thumb.jpg
excerpt: Writing GPU code in Python is easier today than ever, this tutorial covers how you can get started with accelerating your code.
---

{{< youtube sMpqsoIvrio >}}

## Abstract

Since joining NVIDIA I’ve gotten to grips with the fundamentals of writing accelerated code in Python. I was amazed to discover that I didn’t need to learn C++ and I didn’t need new development tools. Writing GPU code in Python is easier today than ever, and in this tutorial, I will share what I’ve learned and how you can get started with accelerating your code.

In this tutorial we will cover:

- What is a GPU and why is it different to a CPU?
- An overview of the CUDA development model.
- Numba: A high performance compiler for Python.
- Writing your first GPU code in Python.
- Managing memory.
- Understanding what your GPU is doing with pyNVML (memory usage, utilization, etc).
- RAPIDS: A suite of GPU accelerated data science libraries.
- Working with Pandas dataframes on the GPU.
- Working with Numpy style arrays on the GPU.
- Performing some scikit-learn style machine learning on the GPU.

Attendees will be expected to have a general knowledge of Python and programming concepts, but no GPU experience will be necessary. The key takeaway for attendees will be the knowledge that they don’t have to do much differently to get their code running on a GPU.
