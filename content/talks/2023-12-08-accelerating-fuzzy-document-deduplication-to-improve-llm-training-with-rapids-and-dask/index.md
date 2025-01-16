---
title: "Accelerating fuzzy document deduplication to improve LLM training with RAPIDS and Dask"
date: 2023-12-08T00:00:00+00:00
draft: false
author: "Jacob Tomlinson"
event:
  name: PyData Global 2023
  link: https://global2023.pydata.org/cfp/talk/3GLSEV/
  type: Talk
  location: Online
length: 25
abstract: true
slides: https://speakerdeck.com/jacobtomlinson/accelerating-fuzzy-document-deduplication-to-improve-llm-training-with-rapids-and-dask
video: https://www.youtube.com/watch?v=mB0cxP_4b24&list=PLGVZCDnMOq0poULd1C4oUdPbPkTe4poJx&index=18
---

Training Large Language Models (LLMs) requires a vast amount of input data, and the higher the quality of that data the better the model will be at producing useful natural language. NVIDIA NeMo Data Curator is a toolkit built with RAPIDS and Dask for extracting, cleaning, filtering and deduplicating training data for LLMs.

In this session, we will zoom in on one element of LLM pretraining and explore how we can scale out fuzzy deduplication of many terabytes of documents. We can run a distributed Jaccard similarity workload by deploying a RAPIDS accelerated Dask cluster on Kubernetes to remove duplicate documents from our training set.

