---
title: "Most stale bots are anti-user and anti-contributor, but they don't have to be"
date: 2024-12-05T00:00:00+00:00
draft: false
author: "Jacob Tomlinson"
categories:
  - blog
tags:
  - github
  - open source
  - projects
---

If you've been around open source projects on GitHub you may have encountered a project with a [stale bot](https://github.com/actions/stale). 

Here's how a common stale bot interaction goes; You've found a problem and you open an issue, but nobody responds. Then 30 days later you get a notification from a bot saying "Beep boop, there hasn't been any activity here for a while. I'm going to mark this as stale". Then a month after that you get another notification from the bot saying "Closing this issue due to inactivity". 

In the meantime you've either worked around the bug or pivoted entirely to avoid it. As a user this experience sucks.

![](./kubernetes-stalebot.png "An issue from Kubernetes that was closed without human involvement ([kubernetes/kubernetes#8268](https://github.com/kubernetes/kubernetes/issues/8268))")

This can be even more painful for a contributor when a stale bot closes a Pull Request. The amount of effort that goes into making a PR can be substantial, and to have it closed out by a bot because nobody has looked at it can be very frustrating. It's likely that person will never try and contribute to your project again.

## A maintainer's perspective

As a maintainer of various open source projects I can tell you that it's very tempting to add a stale bot to a project when you're feeling overwhelmed. If your project is popular it may have hundreds or even thousands of open issues and PRs and often it's just impossible to keep up. Most open source projects are run by volunteers and there are never enough folks offering to do triage.

When you do have some time for triaging you typically start with the most recent things because you're more likely to get engagement back from the person who opened the issue. You would be amazed how many people open an issue to report a bug but never come back again. Often issues don't immediately have all the information that a maintainer needs to resolve it. It's common to ask follow up questions only to be ghosted by the person who opened the issue, even if you reply within a few hours. The longer you leave it before asking a follow up question the more likely you are to be ghosted.

![](./kr8s-ghosted-issue.png "I tend to politely close issues if I get ghosted for a couple of weeks ([kr8s-org/kr8s#406](https://github.com/kr8s-org/kr8s/issues/406))")

This means that the issues that don't get triaged immediately are unlikely to get triaged at all. They slip through the cracks and just get older and older. As a maintainer it's tricky to triage an issue that was created six months ago and figure out if it's still relevant. The person who opened it is unlikely to respond and if there isn't enough information to know what needs fixing then the only option is to close it. These are the cases where a stale bot can be appealing.

## Next action responsibility

I want to make the argument that stale bot automations can be valuable for maintainers to reduce their workload and keep a project in a good state. But if they are done badly, and the majority are, then they are toxic for your project's user and contributor community.

The core problem here comes down to figuring out who's responsibility is it to take the next action. When a user opens an issue it is the responsibility of the maintainer to triage the issue. When the maintainer asks the user a follow up question it becomes the user's responsibility to respond with more information.

Open source collaboration is all about resolving and closing issues, but closing things prematurely is a negative community behaviour. Issues go through an initial phase of validation where we try and answer questions like: Is it clear what is broken? Do we know what needs changing to resolve the issue? Is the issue in a state where a contributor could take it and action it? Until we have the answer to these questions it is the responsibility of the maintainer and the user who opened the issue to work together to answer them. If the user disengages before the issue is validated then the only course of action is to close it. This is a good opportunity for stale bot automation.

However, once an issue becomes a clear actionable task it should never be closed due to inactivity. It should sit on the backlog until someone picks it up. Of course if nobody actions the issue for a long time the project may move on and the issue may become irrelevant or be superseded, but a human needs to make that call, not a bot.

Similarly when a contributor opens a Pull Request it is the maintainer's responsibility to review the PR. If the maintainer reviews and asks for changes it becomes the contributor's responsibility to either challenge the feedback or make the requested changes. The contributor and reviewer will go back and forth and iterate until it is merged or rejected. Overall it's the maintainer's responsibility to keep pushing a PR towards that final state. The PR should only be automatically closed if the contributor hasn't responded to feedback for a long time. 

When the author of a PR goes quiet the maintainer can do two things. They can take it over and keep making changes themselves until it's in a margeable state, or they can close it. Will McGugan wrote a great blog post titled ["Pull Requests are cake or puppies"](https://textual.textualize.io/blog/2023/07/29/pull-requests-are-cake-or-puppies/) where he categorises PRs as either cake, a simple and clear benefit for the project, or a puppy, an improvement that comes with ongoing maintenance needs. If the stale PR is cake then the maintainer is more likely to push it over the finish line and get it merged. But if it's a puppy then the contributor needs to advocate for it to be merged and if they lose motivation it will probably be closed.

Broadly speaking a stale bot should never mark something as stale or close it if the responsibility is on the maintainer to take an action. However, if the responsibility is on the user/contributor then I think it's ok to automate closing these things out.

## An example workflow

GitHub labels are a great way to build workflows to ensure that maintainers can easily see whether the issue needs maintainer input or user input. Each project has their own way of doing things, but I thought it might be useful to outline how I like to handle things.

When an issue is opened it should automatically be given a {{% label "black" "#c2e0c6" %}}needs triage{{% /label %}} label. This shows that the maintainer needs to take action.

If the maintainer asks a question they should replace the label with a {{% label "white" "#3C5E2A" %}}needs info{{% /label %}} label. Or if they request a [minimal reproducible example](https://matthewrocklin.com/minimal-bug-reports.html) they could add a {{% label "white" "#D14B1A" %}}needs reproducible example{{% /label %}} label. It's best to automate things so that when the user comments on the issue these get automatically removed and potentially replaced with a {{% label "white" "#B64491" %}}needs review{{% /label %}} or just the {{% label "black" "#c2e0c6" %}}needs triage{{% /label %}} label again to signal to maintainers that they need to do something again.

Once the user and maintainer have communicated enough to turn the issue into a clear actionable task then the maintainer should add a label to communicate what skill level is required to action the issue. For this I like to use {{% label "white" "#7057ff" %}}good first issue{{% /label %}} {{% label "white" "#1d76db" %}}good second issue{{% /label %}} and {{% label "black" "#E97E74" %}}good expert issue{{% /label %}}. 

```info 
"Good Issues" are tasks that are clearly desribed and actionable.

- {{% label "white" "#7057ff" %}}good first issue{{% /label %}} is simple enough for a newbie contributor to pick up. 
- {{% label "white" "#1d76db" %}}good second issue{{% /label %}} requires some familiarity with GitHub and the project. 
- {{% label "black" "#E97E74" %}}good expert issue{{% /label %}} requires someone with deep understanding of the codebase. 
```

Only issues that have {{% label "white" "#3C5E2A" %}}needs info{{% /label %}} or {{% label "white" "#D14B1A" %}}needs reproducible example{{% /label %}} should be considered by the stale bot. And once an issue has any of the "Good Issue" labels on then all automations should end.

This same process can be applied to Pull Requests but using the review status. A new PR should be reviewed by a maintainer. If they need changes they can use the review feature to request changes. Only PRs with a "request changes" review should be considered by the stale bot. Users can request another review which dismisses the stale bot.

## Proceed cautiously

Automations that close stale issues and PRs can help keep a project's backlog under control. But closing things because you don't have time to get to them leaves a bad taste in people's mouths, especially if you never even look at the issue and let a bot handle the closing for you.

If you do choose to add a stale bot to your project be sure to use labels so that only issues with an outstanding user action get closed.
