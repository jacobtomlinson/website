---
author: Jacob Tomlinson
title: Revisiting old contributions for Hacktoberfest
date: 2025-10-01T00:00:00+00:00
draft: false
categories:
  - blog
tags:
  - Open Source
  - Hacktoberfest
---

It's [hacktoberfest](https://hacktoberfest.com/) time again!

For the last few years hacktoberferst has been opt-in for project maintainers to avoid the bombardment of spam PRs across GitHub. To find participating projects you can view the [hacktoberfest tag on github](https://github.com/topics/hacktoberfest).

Usually I end up contributing to projects I own/maintain and this fills up my quota pretty quickly. This year I want to try and contribute more to other people's projects. However I don't necessarily want to spin up from scratch as a new contributor to some random project, those folks are probably already dealing with a bunch of first-time contributors.

My plans is to revisit old projects I've contributed to in the past. I've made thousands of PRs on GitHub over the years, so I'm going to go back to a repo I've worked on before and be a second-time contributor instead.

Here's a quick bash command that uses the [`gh` GitHub CLI tool](https://cli.github.com/) to find participating repos you've contributed to before.

```bash
gh search prs --author="@me" --limit 1000 | \
awk '{print $1}' | sort -u | \
while IFS= read -r repo; do \
  gh api "repos/$repo/topics" --jq '.names[]' | grep -q hacktoberfest && echo "$repo" ;\
done
```

The GitHub search limits you to 1000 PRs at a time with no pagination beyond that, so if you've made more than 1000 PRs you might need to do it in a few time range batches.

```bash
gh search prs --author="@me" --limit 1000 "created:2010-01-01..2019-12-31" | ...
```
