---
title: 'Generate git release notes automatically'
author: Jacob Tomlinson
layout: post
category: Git
thumbnail: git
tags:
  - git
  - github
---

It is common practice for release notes to consist of a list of the Pull Requests which have been merged since the last release. Some project divide these into categories, for example breaking changes, enhancements and bug fixes. If you are a project maintainer you may want to be able to generate this automatically. 

### Squash and merge

The first step to making your life easier is to use the squash and merge feature in GitHub. This results in each PR being squashed into a single commit, with the commit message being the title of the PR followed the PR number.

Using this feature consistently results in the commit history of your master branch simply being a list of pull requests.

### Extract the history

To get this list in an easy to copy and paste format make sure you ar eon your master branch and it is up to date with tags.

```
git pull origin master --tags
```

Now you can use some `git log` and `cut` wizardry to print out the commit messages since the last tag.

```
git log $(git describe --tags --abbrev=0)..HEAD --oneline | cut -d' ' -f 2-
```

This will print you a neat list of changes which you can then move around and categorise or simply copy and paste into your release notes.
