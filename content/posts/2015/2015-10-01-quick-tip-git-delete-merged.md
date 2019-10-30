---
author: Jacob Tomlinson
date: 2015-10-01T00:00:00+00:00
categories:
  - Quick
  - Tip
excerpt: Ever find yourself with lots of branches you have already merged?
tags:
  - quick tips
  - git
  - delete merged branches
thumbnail: git
title: Quick Tip - git delete merged branches
aliases:
  - /2015/10/01/quick-tip-git-delete-merged/
---


Here's a quick line to run in your terminal to delete all local git branches which have already been merged into master.

```
git branch --merged | grep -v "\*" | grep -v master | grep -v dev | xargs -n 1 git branch -d
```

This is a nice safe command which will only remove branches which have already been merged.
