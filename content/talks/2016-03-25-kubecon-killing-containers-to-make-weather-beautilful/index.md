---
title: "Killing Containers to Make Weather Beautiful"
date: 2016-03-25T00:00:00+00:00
draft: false
author: "Jacob Tomlinson"
event:
  name: Kubecon EU
  link: https://kubeconeurope2016.sched.com/event/6BWQ
  type: Talk
  location: London, UK
length: 30
abstract: true
video: https://www.youtube.com/watch?v=i_gxAFA9_WY
slides: https://speakerdeck.com/jacobtomlinson/kubecon-eu-2016-killing-containers-to-make-weather-beautiful
---

The Met Office Informatics Lab includes scientists, developers and designers. We build prototypes exploring new technologies to make environmental data useful. Here we describe a recent project to process multi-dimensional weather data to create a fully interactive 4D browser application. We used long-running containers to serve data and web pages and short-running processes to ingest and compress the data. Forecast data is issued every three hours so our data ingestion goes through regular and predictable bursts (i.e. perfect for autoscaling).

We built a Kubernetes cluster in an AWS group which auto-scales based on load. We used replication controllers to process the data. Every three hours ingestion jobs are added to a queue and the number of ingestion containers are set in proportion to the queue length. Each worker completes exactly one ingestion job from the queue and then exits, at which point Kubernetes creates a new one to process the next message. This has allowed us to remove the lifespan logic from the containers and keep them light, fast and massively scalable. We are now in the process of using this in our production systems.