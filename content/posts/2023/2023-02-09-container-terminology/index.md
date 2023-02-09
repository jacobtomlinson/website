---
title: "Being intentional with container terminology"
date: 2023-02-09T00:00:00+00:00
draft: false
author: "Jacob Tomlinson"
categories:
  - blog
tags:
  - docker
  - containers
  - oci
  - writing
---

When writing and speaking about linux container technologies I'm trying to be more intentional with the words I use, which means often avoiding the word _docker_. My goal is to communicate clearly to both experts and novices alike.

## Why

A large point of confusion I see when reading about containers is between the generic term **containers** and the brand name **docker**.

My perspective is that [Docker](https://www.docker.com/) is an important company that hugely enhanced and popularized linux container technology. They have tools such as the [Docker Runtime](https://www.docker.com/products/container-runtime/) and [Docker CLI](https://docs.docker.com/engine/reference/commandline/cli/) which today are the most popular way for folks to interact with containers in their day-to-day work. But over time they are slowly fading into the background as just one brand within a broad ecosystem and other tools are taking their place.

Containers are an important component in many other systems such as Kubernetes, [which don't use Docker tools at all](https://kubernetes.io/blog/2022/02/17/dockershim-faq/). They build on the container standards which are governed by the [Open Container Initiative (OCI)](https://opencontainers.org/). While Docker has contributed heavily to those standards there are also many popular alternatives to the Docker tools which implement exactly the same standards and often have identical user interfaces and experiences such as [Podman](https://podman.io/) and [nerdctl](https://github.com/containerd/nerdctl).

Brands like "Hoover" and "Kleenex" have become generic terms for their products and "Docker" is on the same path. So _docker_ is often used as a substitute for the word _container_. I feel like this adds ambiguity and confusion when writing because folks might be unsure whether you are specifically talking about Docker technologies or just talking about generic linux container technology.

## Terminology

To avoid this confusion here are my personal guidelines on container terminology that I'm going to refer back to in the future when writing.

### Brands

Docker is a company, `docker` is a command line tool.

- ✅ **Do** use the proper noun `Docker` when talking about the company or products like the `Docker Runtime`.
- ✅ **Do** use the lowercase name when referring to the CLI tool `docker`.

### Containers

A container is a linux process (or many processes) running in an isolated environment with its own filesystem, network namespace, user namespace and more.

- ✅ **Do** refer to them simply as a `container` or `containers`.
- ✅ **Do** suggest that people can use `docker` to run a `container`.
- ❌ **Do not** call them `docker containers` (or `podman containers` or anything else).
- ❌ **Do not** pluralize them as `dockers` ([looking at you Unraid community](https://www.reddit.com/r/unRAID/comments/626f46/comment/dfl8uvd/?utm_source=share&utm_medium=web2x&context=3)).

### Container images

A container image is a lightweight, standalone, executable package of software that includes everything needed to run an application: code, runtime, system tools, system libraries and settings.

- ✅ **Do** refer to them as `container images` or simply `images`.
- ❌ **Do not** call them `docker images`.

### Container registries

A registry is a standardized API which can store and serve container images. It may or may not have a frontend UI like Docker Hub.

- ✅ **Do** refer to registries as `container registries` or simply `registries`.
- ❌ **Do not** call them `docker registries`.
- ✅ **Do** suggest that folks can push a `container image` to a `registry`.
- ❌ **Do not** suggest that folks push their `docker image` to `docker hub`.
- ✅ **Do** acknowledge that [Docker Hub](https://hub.docker.com/) exists and is the default registry in the `docker` CLI.
- ❌ **Do not** simply tell folks to push images to Docker Hub, also mention a few additional options such as [Amazon ECR](https://aws.amazon.com/ecr/) or [Quay](https://quay.io/).

### Other Docker products

Docker make many other products that enhance the container experience such as [Docker Compose](https://docs.docker.com/compose/) for defining many containers in one configuration and [Docker Swarm](https://docs.docker.com/engine/swarm/) for orchestrating containers on a cluster of machines.

- ✅ **Do** title case product names like `Docker Compose`.
- ✅ **Do** only use lower case `docker-compose` when referring to the CLI tool.
- ❌ **Do not** simply refer to `Docker Compose` as `Compose`.
