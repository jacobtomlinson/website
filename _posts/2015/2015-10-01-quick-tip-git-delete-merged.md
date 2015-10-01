---
title: "Quick Tip - git delete merged branches"
author: Jacob Tomlinson
layout: post
category: Quick Tip
thumbnail: code-fork
tags:
- quick tips
- git
- delete merged branches
excerpt: 'Ever find yourself with lots of branches you have already merged?'
---

Here's a quick line to run in your terminal to delete all local git branches which have already been merged into master.

```
git branch --merged | grep -v "\*" | grep -v master | grep -v dev | xargs -n 1 git branch -d
```

This is a nice safe command which will only remove branches which have already been merged.
