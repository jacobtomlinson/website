---
title: "The art of wrangling your GPU Python environments "
date: 2024-12-04T00:00:00+00:00
draft: false
author: "Jacob Tomlinson"
event:
  name: PyData Global 2024
  link: https://global2024.pydata.org/cfp/talk/B9GZWJ/
  type: Talk
  location: Online
length: 30
abstract: true
---

Debugging software itself is a hard task, but debugging GPU software environments can be even more challenging. Understanding the intricate interactions between hardware, drivers, CUDA, C++ dependencies, and Python libraries can be far more complex.

In this talk we will dig into how these different layers interact and how you can address some of the common pitfalls that folks run into when configuring GPU Python environments. We will also introduce a new tool, RAPIDS Doctor, that aims to take the challenge out of ensuring your software environments are in good shape. RAPIDS Doctor checks and diagnoses environmental health issues straight from the command line, ensuring that your setup is fully functional and optimized for performance.

Projects like RAPIDS, a rapidly growing suite of GPU Accelerated ML & Data Science libraries, along with communities like Pytorch, Tensorflow and others are continuously looking to simplify the setup required to leverage GP-Us in your PyData workflows.

Many users seek to install and use RAPIDS but are unclear of certain system requirements that it depends on. To install RAPIDS you generally need a GPU, NVIDIA Drivers, CUDA Toolkit, and RAPIDS packages (and compatible dependencies). While most of the software can be installed via conda/pip, the drivers must be installed outside of your Python environment and consistent with GPU requirements.

RAPIDS Doctor is a new command line tool that will have capabilities to check for multiple system dependencies. As users are frequently installing RAPIDS in a variety of cloud environments, this is particularly useful in getting a quick rundown on incompatibilities that may cause issues down the line. Additionally, RAPIDS Doctor also prescribes a treatment for diagnosed issues, such as quick fix suggestions in your terminal or even autofixes.

In this talk we will demonstrate the suite of use-cases of Rapids Doctor and the diversity of health checks that it has expertise in. Whether you're a seasoned developer or just starting with Python software development with GPUs, this tool streamlines the setup process and enhances your productivity, allowing you to focus on your data science and machine learning projects without the headaches of environmental troubleshooting.
