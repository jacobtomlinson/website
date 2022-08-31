---
title: "Cleaning up conda environments"
date: 2019-08-23T00:00:00+01:00
draft: false
author: "Jacob Tomlinson"
categories:
  - blog
tags:
  - conda
  - python
  - hygine
thumbnail: conda
---

Often when I'm developing or debugging in Python I end up creating throw away conda environments. They will be to test some package installation or combination of packages and once I've finished I will probably never use them again.

Usually I accidentally leave these hanging around and end up with a cluttered `conda env list`.

Here's a quick one liner to remove any environments with the words `tmp`, `temp` or `test` in them.

```
$ conda env list | grep 'tmp\|temp\|test' | awk '{ print $1 }' | xargs -n1 -d'\n' conda env remove -n
```
