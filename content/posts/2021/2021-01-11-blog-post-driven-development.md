---
title: "Blog post driven development"
date: 2021-01-11T00:00:00+00:00
draft: false
author: "Jacob Tomlinson"
categories:
  - blog
tags:
  - work
  - BPDD
thumbnail: journal
---

**Hot take**: Open source software developers are not very good at telling folks about new functionality.

So we need _blog post driven development_.

## What?!

That statement is probably unfair, but I want to highlight that open source projects often do not enjoy the benefit of dedicated documentation and marketing teams that exist within large organisations.

When you work on a piece of software in your spare time, whether personally or professionally, you probably don't put enough effort into telling people about the changes you are making.

## How open source projects announce functionality

Features within open source projects tend to follow these written steps:

- Open an issue, "I need the project to do X".
- Raise a Pull Request, "adding X".
  - This often contains documentation on "how to do X".
- Changelog/release notes is updated, "x was added".

Some projects may go one step further and put out a release blog post, but this is often the changelog with bows on.

Once this is done the feature is out there in the world for everyone to discover and use. I want to make the argument here that this is not enough.

## Abstract features are not enough

Adding a feature to a project is great, I do it all the time. In fact I do it so often that I regularly forget that I've done it. More than once have I raised a pull request for something that I've already completed, merged and released six months earlier.

If I can't keep track of all the small features in the projects that I work on how can I expect the users and other contributors to do the same?

The reason why incremental improvements to a piece of software are hard to track is because on their own they have little value, but as part of a larger narrative they hang together to empower some new workflow.

## Narrative use in agile software development

In [agile software development](https://en.wikipedia.org/wiki/Agile_software_development) it is common to work with [user stories](https://en.wikipedia.org/wiki/User_story) as your unit of individual tasks. A story is a way of describing some piece of required functionality in the context of the person using it.

> As a [persona], I [want to], [so that].

You describe the user ([persona](https://en.wikipedia.org/wiki/Persona_(user_experience))), the functionality they require and the value that brings to them and the larger organisation.

Stories can also be grouped together into [epics](https://www.atlassian.com/agile/project-management/epics) which commonly represents the work performed by a team, but can be used to represent a complete workflow performed by a persona.

Many open source projects follow the Agile methodology, but only loosely. It is not very common (in my experience anyway) to come across user stories on GitHub. But work is often captured in issues as individual tasks, with larger goals captured with projects and milestones.



