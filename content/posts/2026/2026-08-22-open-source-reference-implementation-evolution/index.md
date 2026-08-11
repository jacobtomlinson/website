---
author: Jacob Tomlinson
title: Open Source libraries are evolving into reference implementations
date: "2026-08-11T00:00:00+00:00"
draft: false
categories:
  - blog
tags:
  - Open Source
  - Coding Agents
  - AI
  - Community
  - Software Culture
---

I've spent the majority of my professional career working on software libraries, chunks of reusable code that other people stack together to make _their_ code.

Libraries are valuable because they contain common code that gets reused a lot, code that has been heavily tuned for performance, or code that captures a niche domain that would be cumbersome to rewrite. Libraries exist to save people time. Open source libraries also add benefit by having many people contributing on a common codebase fixing bugs, improving performance and increasing correctness. 

**Library code is "better" than code I would write myself.**

Over the last few years the software industry has changed:
- Coding agents have exploded in popularity and many people now direct agents instead of writing code directly.
- Coding agents can generally one-shot well defined implementations with robust testing.
- Agents can hill-climb towards performance, correctness, simplicity, maintainability or whatever other goal you want to define using [autoresearch](https://github.com/karpathy/autoresearch).
- Software [supply chain attacks](https://www.sophos.com/en-us/blog/axios-npm-package-compromised-to-deploy-malware) have increased so much that people are more hesitant to use libraries than ever.

As a result I keep asking myself whether we even need to depend on libraries any more? Instead of pulling in thousands of lines of library code in order to use a few hundred in my application, and risk some [upstream poisoning adding malware to my application](https://en.wikipedia.org/wiki/XZ_Utils_backdoor), why not just get an agent to build the bits I need from scratch every time?

LLMs are also very good at translating from one language to another. It doesn't really matter to them what language you want to write something in. They are by nature transformers and so transforming an implementation from one language to another is part of their core capabilities. Therefore, every library ever made is valuable as a reference for building something in any other programming language.

In the past if I've been working on a Python application I've been limited to the Python libraries that are available to me. But now I can easily crib stuff from libraries, build bindings or wholesale port things from any language into my project.

I care less about having well maintained chunks of code that I can depend on directly. I care more about having well thought out reference implementations that my agents can use for reference, inspiration and benchmarking.

**Is library code "better" than what my agent can write for me?**

## An example

One example from a personal project that stands out was when I was tinkering with my smart home setup recently. I use [Immich](https://github.com/immich-app/immich) running on my NAS to self-host my photos. Immich is open source and has a face matching feature that allows you to view your photos grouped by person. I also have a smart doorbell with a camera which sends me a picture notification when someone pressed the bell. 

I had the idea that when someone presses the bell I could first send the image through my Immich face detection, and if it matches a known person it would include their name in the notification. I love that all that processing can happen locally.

> 🔔 John Smith is at the door!

I use [n8n](https://n8n.io) for some of my more complex home automations, which is a node based editor that stores workflows in a JSON format with small snippets of JavaScript. I had no idea how to use Immich to get names from a picture, but given that it's open source I knew it must be possible to figure out. It groups faces using an embedding, and you can attach people's names to their face, so somewhere in the code/data of my Immich server is a function that can take an image and generate an embedding and a lookup table of my friend's and relative's names and face embeddings.

It turned out that Immich doesn't have an API where you can post an image and it tells you who is in the image. However, I did some searching and found a [python library on GitHub](https://github.com/cmprmsd/immich-face-api) which takes a photo, calls the same embedding API that Immich calls, then calls the [pgvector](https://github.com/pgvector/pgvector) database that backs Immich directly to look up the names. 

I gave the URL of this library to Claude, along with the MCP of my n8n server and told it to make an n8n workflow that does the same thing, it should accept a webhook from [Home Assistant](https://www.home-assistant.io) that I can trigger with the doorbell, and it should finish with calling the Home Assistant notify service. It thought for a couple of minutes and then created a workflow in n8n JSON and a little bit of JavaScript to handle the database calls. It worked first time.

The python library I found on GitHub provided me with a reusable chunk of code, but instead of depending on it directly I ported it's logic into the system/langauge I happen to be using for a couple of pennies in tokens. That library acted as a reference implementation which my agent used to build out just the bits I need.

```info
Interestingly both the reference implementation I found as well as Immich itself are GPL-family licensed projects, and Claude definitely read the source of both when making the n8n workflow. Therefore, if I wanted to distribute my doorbell workflow beyond my own personal use it would also need to be published under a GPL license too as a derivative work.
```

## More open questions

- If fewer people directly depend on libraries what will [motivate people to maintain them](https://www.jeffgeerling.com/blog/2022/burden-open-source-maintainer/)?
- If I port a library to a different language, then spend some time/tokens making it more performant how do I contribute that back in a meaningful way?
- How do we track dependency licenses for software projects at work when dependencies are indirect reference implementations?
- How does licensing need to evolve to keep up? Projects like [chardet are testing whether you can relicense an AI rewrite](https://github.com/chardet/chardet/issues/327).
