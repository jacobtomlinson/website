---
title: "Migrating from Docker Hub autobuilds to Github Actions"
date: 2021-06-09T00:00:00+00:00
draft: true
author: "Jacob Tomlinson"
categories:
  - blog
tags:
  - Docker
  - Docker Hub
  - CI/CD
  - GitHub Actions
---

This week [Docker Hub announced that they are removing the autobuilds feature from the free user tier](https://www.docker.com/blog/changes-to-docker-hub-autobuilds/) due to abuse from crypto miners.

Here's a guide on migrating your Docker builds over to GitHub Actions.

```yaml
on:
  push:
    branches:
    - master
  pull_request:

jobs:
  build:
    name: 'Build'
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@v2

      - name: Login to DockerHub
        uses: docker/login-action@v1
        if: github.repository == 'opsdroid/opsdroid' && github.event_name == 'push' && startsWith(github.ref, 'refs/tags')
        with:
          username: ${{ secrets.DOCKERHUB_USERNAME }}
          password: ${{ secrets.DOCKERHUB_TOKEN }}

      - name: Build and push
        uses: docker/build-push-action@v2
        with:
          push: github.repository == 'dask/dask-docker' && github.event_name == 'push' && startsWith(github.ref, 'refs/tags')
          platforms: linux/amd64
          tags: ${{ steps.tags.outputs.tags }}
```