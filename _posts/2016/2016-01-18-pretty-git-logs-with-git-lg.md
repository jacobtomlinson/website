---
title: 'Pretty git logs with `git lg`'
author: Jacob Tomlinson
layout: post
category: Quick Tip
thumbnail: code-fork
tags:
  - git
  - log
excerpt: 'Make your git logs easier to look at and track branches.'
---

This is a repost of a [Stack Overflow answer][stack-overflow-answer], mainly to preserve it for myself. [Slipp Thompson][stack-overflow-slipp-thompson] posted some really nice aliases for showing branch topology in the git command line.

To use it add the following to your `~/.gitconfig`.

```
[alias]
    lg = !"git lg1"
    lg1 = !"git lg1-specific --all"
    lg2 = !"git lg2-specific --all"
    lg3 = !"git lg3-specific --all"

    lg1-specific = log --graph --abbrev-commit --decorate --format=format:'%C(bold blue)%h%C(reset) - %C(bold green)(%ar)%C(reset) %C(white)%s%C(reset) %C(dim white)- %an%C(reset)%C(bold yellow)%d%C(reset)'
    lg2-specific = log --graph --abbrev-commit --decorate --format=format:'%C(bold blue)%h%C(reset) - %C(bold cyan)%aD%C(reset) %C(bold green)(%ar)%C(reset)%C(bold yellow)%d%C(reset)%n''          %C(white)%s%C(reset) %C(dim white)- %an%C(reset)'
    lg3-specific = log --graph --abbrev-commit --decorate --format=format:'%C(bold blue)%h%C(reset) - %C(bold cyan)%aD%C(reset) %C(bold green)(%ar)%C(reset) %C(bold cyan)(committed: %cD)%C(reset) %C(bold yellow)%d%C(reset)%n''          %C(white)%s%C(reset)%n''          %C(dim white)- %an <%ae> %C(reset) %C(dim white)(committer: %cn <%ce>)%C(reset)'
```

Examples of how each command looks.

#### git lg
![git lg](http://i.stack.imgur.com/gkcfL.png)


#### git lg2
![git lg2](http://i.stack.imgur.com/7fWK9.png)


#### git lg3
![git lg3](http://i.stack.imgur.com/mgrEp.png)

[stack-overflow-slipp-thompson]: http://stackoverflow.com/users/177525/slipp-d-thompson
[stack-overflow-answer]: http://stackoverflow.com/a/34467298/1003288
