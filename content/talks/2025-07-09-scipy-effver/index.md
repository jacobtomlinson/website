---
title: "EffVer - Version your code by the effort required to upgrade"
date: 2025-07-T00:00:00+00:00
draft: false
author: "Jacob Tomlinson"
event:
  name: SciPy
  link: https://www.scipy2025.scipy.org/schedule
  type: Talk
  location: Tacoma, WA, USA
length: 25
abstract: true
slides: https://speakerdeck.com/jacobtomlinson/effver-version-your-code-by-the-effort-required-to-upgrade
video: https://www.youtube.com/watch?v=dx5K2Sytk4c
---

Version numbers are hard to get right. Maintainers want to communicate to users what the impact of adopting a new version will be, but poor communication can lead to a lot of frustration. There are a few popular version schemes in use today including Semantic Versioning (SemVer) and Calendar Versioning (CalVer). However, projects in the Python community often don’t strictly conform to these standards which leads to confusion.

In this talk we will discuss Intended Effort Versioning (EffVer), a new scheme that captures the reality of what many Python projects do today. This formalisation has been officially adopted by projects including Jupyter Hub, Matplotlib, JAX and many more.

It’s very common for projects in the Python ecosystem to try to follow Semantic Versioning (SemVer), a scheme that adds specific meanings to each version segment and guarantees backward compatibility in all but major releases. In practice many projects violate the semantics laid out in SemVer because often reality is more complex than the scheme allows for.

Some projects go another route and use Calendar Versioning (CalVer), a scheme that intentionally gives no meaning in the segments other than the date the software was released. This information causes challenges in other ways, such as not communicating how much effort it will be to adopt the new version or misleading users to believe that releases of different projects from the same time period will be compatible.

EffVer is forward and backward compatible with SemVer, making it easy to adopt for projects already loosely following SemVer. However, instead of communicating specific semantics about what a release contains, it communicates the magnitude of the effort required to adopt those changes. This is already what many Python packages do today, but by formalising things it makes it easy for maintainers to reason about what the next version of a project should be, and allows users to more clearly understand how much effort they need to spend to adopt a new version.
