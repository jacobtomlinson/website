---
author: Jacob Tomlinson
title: Foundation Software
date: 2026-11-18T00:00:00+00:00
draft: false
categories:
  - blog
tags:
  - AI
  - Coding
  - Personalisation
---

People are not just vibe coding all sorts of side projects from scratch, they are forking the existing tools they use and are building on top of that code too.

I've found myself doing this a bunch recently. For a long time I've run a home server where I self host lots of different applications. Sometimes I notice a gap in that project, whether it's a bug, a missing feature, whatever. In the past for bugs and general purpose features I would probably fix them and make a PR upstream. However if I wanted something personal, niche or unusual I just wouldn't bother. The upstream project almost certainly wouldn't accept that PR and the effort required to maintain a fork against upstream was too big to be able to justify. But with agents it's easy not only to generate the feature I want but most importantly I can have a recurring job that rebases my changes onto new releases or just cherry picks fixes from the main branch.

There's something powerful about feeling ownership over the software you use. This feels like the true vision of open source.

It's clear that 2026 has been a weird year for open source. On one hand there is [more code on GitHub than ever](https://github.blog/news-insights/octoverse/), and the constant stream of open AI models coming out of labs around the world has empowered many developers. On the other hand many projects have closed their PRs and taken an anti-AI stance, while other communities have seen a huge drop in interest because it's easier to just vibe your own.

I find it interesting to think about the future of open source software. Despite the rumours it's certainly not dead, but it will be different. Libraries are more important than ever because agents can use them as building blocks to quickly flesh out their codebases (but they may be [cribbed, vendored or used as reference](../open-source-libraries-are-evolving-into-reference-implementations/) more). While at the same time simple projects like habit trackers or meal planners are being vibe coded almost daily because they are simple and often very personal pieces of software. Lots of Reddit communities I am a part of don't allow submitting vibe coded projects because there are so many and there is so much duplication.

I don't want to spend my time rebuilding things that already exist. Of course I could build a document manager from scratch, but I would much rather use [paperless-ngx](https://github.com/paperless-ngx/paperless-ngx) with some customisation to make it behave the way I would like and leave my agents to keep my fork up to date with the community. As a result instead of looking for the best solution to a problem I have I've started looking for the best _foundation software_ for me to build my own solution on top of.

I want to use projects which have:
- Good agent instructions in their `AGENTS.md`
- Helpful skills for running tests, verifying new changes, etc
- Good modular design and layout so my agents can navigate it quickly
- Clean records of how things work and the reasoning why in developer documentation (for example [ADRs](https://github.com/mattpocock/skills/blob/main/skills/engineering/domain-modeling/ADR-FORMAT.md))

The better these things are the more likely my agent is going to [generate clean code](../guiding-your-contributors-agents-to-better-behaviours/) when it makes changes instead of slop.

As we shift in this direction as a community it begs the question "what about upstreaming contributions?". I have hundreds of commits of code sitting in my forks which I will never upstream. In some ways I kinda feel sad about that, I like giving back and I like making software better for everyone. But at the same time I know that people won't want my niche and specific changes.

Instead I think we need to flip the paradigm. We should publish these forks and share them, put them online for everyone to see, but with no expectation of merging. There is huge value in sharing our weird and wonderful tweaks. But instead of opening a PR and requesting someone to consider merging it the maintainers can just explore all the forks to find trends. As a maintainer of software projects I would love to be able to poke around and explore all the forks that people make. I would be able to send my agents off into the various worktrees hunting for patterns. Are there things that get modified frequently in the same way? If so we should consider making that same change upstream. Are there things that are frequently changed but in a myriad of different ways? If so we should consider making that component even easier to customise.

In her talk [All Code Is Slop (Except Mine)](https://www.youtube.com/watch?v=O_zcc6Xl12g) Lauren Tan compares building software with agents to tending a garden. Sometimes you need to just let things grow and then cut them back. You need to put the frames and structures in place to help guide and propagate the plants you care about. I enjoy Lauren's analogies (her [Michelin Kitchen](https://www.youtube.com/watch?v=MN9dGgmLyso) one is great too) but the garden one got me thinking. Lots of people have gardens. They can be very personal places where you explore and experiment and make things just the way you like. We don't often tend a communal garden, we usually let our local authority manage our parks and other green spaces, and that's ok. It's ok to grow your own garden and enjoy it yourself, just like it's ok to build your own software for your own use. When we have a garden we often show it to our friends and families. Avid gardeners may also visit stately homes or other places with show gardens to learn and get ideas. We should do more sharing without the expectation of adoption. Sharing and discussing still builds community.

I feel like we are entering a golden age of building software. The barrier to entry for building and customising things has gotten a lot lower. But I don't want to spend my time rebuilding things that already exist. I want to spend my time tailoring my software to my needs and building new things and I would much prefer to have a solid foundation to build on top of.
