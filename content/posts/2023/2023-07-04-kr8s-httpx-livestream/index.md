---
title: "Livestream notes: Replacing aiohttp with httpx in kr8s"
date: 2023-07-04T00:00:00+00:00
draft: false
author: "Jacob Tomlinson"
categories:
  - blog
tags:
  - coding
  - python
  - kr8s
  - httpx
  - asyncio
  - trio
  - aiohttp
  - kubernetes
aliases:
  - /2023/july-4th-livestream
---

```info
This post will be updated with notes from the livestream throughout the day.
```

Today I will be streaming some open source code refactoring. [Come and join in on Twitch!](https://www.twitch.tv/constrainedcoding). Don't forget to say hi in the chat 😊.

## Overview

The main goal for the day is to close [kr8s-org/kr8s#77](https://github.com/kr8s-org/kr8s/issues/77) by adding support for [`trio`](https://trio.readthedocs.io/en/stable/) to [`kr8s`](https://github.com/kr8s-org/kr8s).

To do this we need to swap out [`aiohttp`](https://docs.aiohttp.org/en/stable/) for [`httpx`](https://www.python-httpx.org/). We could use any library that implements [`anyio`](https://anyio.readthedocs.io/en/3.x/) but `httpx` seems like the most popular choice.
