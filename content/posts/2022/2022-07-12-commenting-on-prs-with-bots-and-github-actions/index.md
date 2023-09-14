---
title: "Commenting on Pull Requests with GitHub Actions"
date: 2022-07-12T00:00:00+00:00
draft: false
author: "Jacob Tomlinson"
categories:
  - blog
tags:
  - Open Source
  - GitHub Actions
  - Bots
  - Automation
---

When someone opens a Pull Request (PR) on your GitHub project it can be helpful for a bot to comment on the PR. You might want to thank the user for the contribution, provide some useful information such as [giving a binder link where folks can try out the PR](https://github.com/dask/dask-tutorial/pull/260#issuecomment-1181821332), or providing more verbose output from some tests or other checks.

Let's work through setting this up and dig into the caveats and limitations of doing so with GitHub Actions.

## Tokens and security

Before we can make comments on a PR from within our GitHub Actions we need to think about how we will authenticate.

GitHub Actions often provides us with a [`GITHUB_TOKEN` secret](https://docs.github.com/en/actions/security-guides/automatic-token-authentication) that can be used to allow your action to interact with GitHub. This is commonly used to push artifacts on releases and is often used in Actions that run on the `main` branch.

When a contributor forks your repo and raises a Pull Request from that fork the token they get will be read only. So we can't use it to make comments.

We could also create a [Personal Access Token](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token) to a specific account or a bot account. We could create a [repository secret](https://docs.github.com/en/actions/security-guides/encrypted-secrets) to store the token and then use this in our Action. However we also face the same restriction with PRs from forks, we don't want users to have our secrets in their workflows because they could change their code to use the token in a bad way or even print the token so that they can have it themselves.

To work around this the Action that we are going to write will use the [`pull_request_target` event](https://docs.github.com/en/actions/using-workflows/events-that-trigger-workflows#pull_request_target) to trigger our action. In comparison to the common `pull_request` event the `pull_request_target` event will trigger every time a commit it pushed to a PR, but it will run the action against the target branch (usually `main`) instead of the PR branch. This way we know the code run by the action is trusted, and GitHub allows secrets and the `GITHUB_TOKEN` in these workflows.

The downside to this is that it makes commenting the results of a regular workflow more challenging because they will usually be using the `pull_request` event instead. So bear that in mind when making comments.

## Making a comment

Let's make a workflow that comments on all new Pull Requests with a welcome message to thank the contributor.

```yaml
name: Thank Contributor
on: pull_request_target

jobs:
  thank_contributor:
    runs-on: ubuntu-latest
    steps:
      - name: Comment PR
        uses: thollander/actions-comment-pull-request@v1
        with:
          message: |
            Thank you for your contribution. We aim to review it within 48 hours!
          comment_includes: "Thank you"
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```

This small action will trigger on every commit of a Pull Request. It uses the `thollander/actions-comment-pull-request@v1` action to make the comment with a message thanking the user.

We have also specified `comment_includes` here which means if more commits are pushed the action will look for an existing comment containing that string and if one exists it will update it instead of making a new comment. This is great for ensuring a comment only happens once, or updating some output from the workflow if you have other steps. If you want your comment to happen on every new commit then remove this option.

We also specify the `GITHUB_TOKEN` but if you are using a PAT stored in a secret you would configure that here.

In order to trigger this action we need to merge it into the `main` branch. Then on new PRs our comment will be applied.

![Bot comment showing up on new PR](https://i.imgur.com/6z7b5bD.png)

Just remember that if you open PRs to update this action in the future it will trigger the version on main, not the one in the PR until it has been merged. Keep your secrets safe folks!
