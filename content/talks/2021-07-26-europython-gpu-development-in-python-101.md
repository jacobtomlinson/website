---
title: "GPU development with Python 101"
date: 2021-07-26T00:00:00+00:00
draft: false
author: "Jacob Tomlinson"
event:
  name: EuroPython 2021
  link: https://ep2021.europython.eu/talks/6M8oWP2-gpu-development-with-python-101/
  type: Tutorial
  location: Online
length: 180
video: null
abstract: true
resources: https://github.com/jacobtomlinson/gpu-python-tutorial
---

Writing GPU code in Python is easier today than ever!

I joined NVIDIA in 2019 and I was brand new to GPU development. In that time, I’ve gotten to grips with the fundamentals of writing accelerated code in Python. I was amazed to discover that I didn’t need to learn C++ and I didn’t need new development tools. Writing GPU code in Python is easier today than ever, and in this tutorial, I will share what I’ve learned and how you can get started with accelerating your code.

The first half of this training will take the form of a traditional tutorial. We will work through various materials and examples to get you started with GPU development in Python using open source libraries.

This material will cover:
- What is a GPU and why is it different to a CPU?
- An overview of the CUDA development model.
- Numba: A high performance compiler for Python.
- Writing your first GPU code in Python.
- Managing memory.
- Understanding what your GPU is doing with pyNVML (memory usage, utilization, etc).
- RAPIDS: A suite of GPU accelerated data science libraries.
- Working with Numpy style arrays on the GPU.
- Working with Pandas style dataframes on the GPU.
- Performing some scikit-learn style machine learning on the GPU.

In the second half, I will be joined by a colleague Graham Markall who is experienced in Python GPU development and we will work through a variety of programming challenges together to demonstrate how to think about developing for a GPU when tackling your day-to-day problems.

Tentative schedule:
- Intro to GPUs (20 mins)
- Writing low level GPU code in Python with Numba (30 mins)
- Inspecting your GPU usage with pyNVML (10 mins)
- Writing high level GPU code in Python with RAPIDS (30 mins)
- Collaborative solving of programming challenges on the GPU in Python (90 mins)

Attendees will be expected to have a general knowledge of Python and programming concepts, but no GPU experience will be necessary. Our key takeaway for attendees will be the knowledge that they don’t have to do much differently to get their code running on a GPU.