---
author: Jacob Tomlinson
date: 2017-06-15T00:00:00+00:00
categories:
  - Git
tags:
  - git
  - github
thumbnail: git
title: Generate git release notes automatically
aliases:
  - /2017/06/15/git-list-commits-since-last-release/
---


It is common practice for release notes to consist of a list of the Pull Requests which have been merged since the last release. Some projects divide these into categories, for example breaking changes, enhancements and bug fixes. If you are a project maintainer you may want to be able to generate this automatically.

### Squash and merge

The first step to making your life easier is to use the [squash and merge feature](https://github.com/blog/2141-squash-your-commits) in GitHub. This results in each PR being squashed into a single commit, with the commit message being the title of the PR followed the PR number.

Using this feature consistently results in the commit history of your master branch simply being a list of pull requests.

### Extract the history

To get this list in an easy to copy and paste format ensure you are on your master branch and it is up to date with tags.

```shell
$ git checkout master
$ git pull origin master --tags
```

Now you can use some `git log` and `cut` wizardry to print out the commit messages since the last tag.

```shell
$ git log $(git describe --tags --abbrev=0)..HEAD --oneline | cut -d' ' -f 2-
Catch network errors (#184)
Change default log location (#183)
Allow ssl on web server (#182)
Expand module path. (#180)
Add restart and stop methods to opsdroid. (#179)
Update docs badge (#177)
Added symlink to readme (#176)
Removed index and getting started in favour of the main README (#175)
Add ubuntu 16.04 instructions to README (#173)
```

This will print you a neat list of changes which you can then move around and categorise or simply copy and paste into your release notes.
