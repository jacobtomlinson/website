---
title: "cudf.pandas: the Zero Code Change GPU Accelerator for pandas"
date: 2024-02-07T00:00:00+00:00
draft: false
author: "Jacob Tomlinson"
event:
  name: PyData Exeter Meetup Feb 2023
  link: https://www.meetup.com/pydata-exeter/events/298439130/?eventOrigin=group_events_list
  type: Talk
  location: Exeter, UK
length: 20
abstract: true
slides: https://speakerdeck.com/jacobtomlinson/cudf-dot-pandas-the-zero-code-change-gpu-accelerator-for-pandas
---

Pandas is flexible, but often slow when processing gigabytes of data. Many frameworks promise higher performance, but they often support only a subset of the Pandas API, require significant code change, and struggle to interact with or accelerate third-party code that you can’t change. RAPIDS cuDF enables Pandas users to accelerate their existing workflows and third-party code with zero code change required. You can continue using Pandas on CPUs for small-scale local development and testing, and then enable “Pandas accelerator mode” to run on GPUs when you want better performance. Accelerator mode uses the GPU for supported operations and the CPU otherwise, dispatching under the hood as needed, and is compatible with most third-party libraries.
