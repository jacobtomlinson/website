---
title: Sort top command by cpu usage and set to default in OS X
author: Jacob Tomlinson
layout: post
permalink: /2013/06/04/sort-top-command-by-cpu-usage-and-set-to-default-in-os-x/
thumbnail: command-line
has_been_twittered:
  - yes
category: Apple
tags:
  - alias
  - bash
  - os x
  - terminal
  - top
---
As I come from a linux background but seem to spend more and more of my free time using OS X I keep noticing little differences in the way the command line works on a mac. One difference which as been bugging me recently is that way that the *top* command orders itself. I&#8217;m used to having it ordered by highest processor usage at the top on linux, which I find the most useful as generally when I run top I&#8217;m looking to see what is chewing up my cpu. However when you run top in OS X it orders by pid, so the newest processes are at the top.

To change this you can give some extra parameters to set the ordering. So to order by cpu on a mac you give the command


```
top -o cpu
```


which is great but I don&#8217;t really want to be typing this in every time, I want it to be the default.

To do this you can create an alias in your &#8216;.profile&#8217; file. This file should be in your home directory but if you don&#8217;t have one you can just create one with the command


```
touch ~/.profile
```


In there you just need to add the line


```
alias top="top -o cpu"
```


this means when you run the command `top` it will call `top -o cpu` instead. To apply these changes you can either type `. .profile` or just close the terminal and open it again. You can run the default top command by sticking a *\* before the command which will ignore the alias, more details on setting an alias <a title="Overriding Aliases" href="http://en.wikipedia.org/wiki/Alias_(command)#Overriding_aliases" target="_blank">here</a>.

There are of course other ways around this, one being to use the more advanced and powerful <a title="Man page for htop command" href="http://linux.die.net/man/1/htop" target="_blank">htop</a> command which can be installed with <a title="Macports" href="http://www.macports.org/" target="_blank">macports</a> or <a title="Homebrew" href="http://mxcl.github.io/homebrew/" target="_blank">brew</a>. But for a quick little workaround this does the trick for me.
