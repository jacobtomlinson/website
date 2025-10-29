---
title: "Reimplementing kubectl for fun"
date: 2023-06-12T00:00:00+00:00
draft: true
author: "Jacob Tomlinson"
categories:
  - blog
tags:
  - project
  - kubernetes
  - python
  - kr8s
  - kubectl-ng
---

To resolve this I'm also tinkering with another side-project called [`kubectl-ng`](https://github.com/kr8s-org/kr8s/tree/main/examples/kubectl-ng). This is a reimplementation of `kubectl` in Python using `kr8s`, [`rich`](https://rich.readthedocs.io/en/stable/introduction.html) and [`typer`](https://typer.tiangolo.com/).

`kubectl-ng` is much more of a toy project and isn't really ready yet for others to use, and maybe never will be. By reimplementing `kubectl` I can understand more deeply what is abstracted away, what API calls `kubectl` makes under the hood and how that is exposed to the user. I can also find new things to push down to the `kr8s` level. For example the Kubernetes API has no concept of an upsert, so neither do the existing client libraries, but `kubectl` exposes this functionality via `kubectl apply`. 

By implementing `kubectl-ng apply` I can see that all it does is call `.patch()`, catches the 404 error in the case there the resource is missing and calls `.create()` instead. This is a great bit of functionality that I can push up into the `kr8s` API so that folks could call `Pod({...}).apply()` in the Python API.

This is also an opportunity for me to improve on `kubectl`. I regularly use [`k9s`](https://k9scli.io/) to interact with my Kubernetes clusters but sometimes I find k9s too heavy and kubectl too light. For example I might want to watch some Pods to see how they are behaving. I really like the way that k9s shows changes in information over time (using colour and the `Δ` symbol), but I would prefer to call `kubectl get pods -w`. However I really dislike how `kubectl` shows changes in information because it just prints duplicate lines with updated information. My plan for `kubectl-ng` is to implement `kubectl-ng get pods -w` to use `rich` to make information beautiful and more useful for interactive use.
