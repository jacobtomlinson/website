---
title: "How to check out the default git branch"
date: 2020-09-09T00:00:00+00:00
draft: false
author: "Jacob Tomlinson"
categories:
  - blog
tags:
  - git
  - Open source
  - code
thumbnail: git
---

Many open source projects are taking steps to update terminology to be more inclusive. The largest of these changes has been renaming the "trunk" branch of git repositories from `master` to `main`.

This is great and I fully support the transition, however this has resulted in some minor annoyances when working on various projects. I use git aliases to checkout out the trunk branch (`gcm`) and to also pull from upstream repositories (`glum`) as I generally work on forks. I'm now in a position where some projects still use `master` as their default branch, some use `main`, some repos like documentation or websites may use `gh-pages`, and some use other conventions all together.

To resolve this I'm updated my `gcm` and `glum` aliases to use a new `gdb` alias which queries the upstream repository for it's default branch name. As my workflow uses forks 99% of the time I have defaulted the remote to `upstream` but if you use another workflow you may wish to change this to `origin` or something else.

```bash
function gdb () {
	REMOTE=${1:-upstream}
	git fetch $REMOTE
	git remote show $REMOTE | grep "HEAD branch" | sed 's/.*: //'
}


unalias gcm
function gcm () {
       git checkout $(gdb $1)
}

unalias glum
function glum () {
	git pull upstream $(gdb upstream)
}
```

These functions replace my existing aliases provided by [oh-my-zsh](https://github.com/ohmyzsh/ohmyzsh/tree/master/plugins/git).

### gcm

`gcm` checks out the default branch name from the upstream repository, but also takes an argument to pass in an alternative remote name.

```console
$ gcm  # Checks out the default branch name from the 'upstream' remote
$ gcm origin # Checks out the default branch name from the 'origin' remote
```

### glum

`glum` pulls the default branch from the `upstream` remote. As the `u` in `glum` stands for `upstream` this one does not take an argument.

```console
$ glum  # Pull the default branch name from the upstream remote
```