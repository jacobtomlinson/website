---
author: Jacob Tomlinson
title: Guiding your contributor's agents to better behaviours
date: 2026-02-19T00:00:00+00:00
draft: false
categories:
  - blog
tags:
  - Open Source
  - AI
  - Coding Agents
  - Claude Code
  - GitHub
---

Many open source maintainers have noticed an uptick in low-effort AI generated PRs recently, myself included. The most frustrating of these is when someone prompts their agent to `"Fix <url to issue> and make a PR with the changes"`. Reviewing these PRs can be time consuming because diffs can be large and the contributors rarely respond to review feedback, they just prompt and move on.

Worst of all is when the discussion in the issue is incomplete or hasn't reached a conclusion. In these cases the agent writes code that isn't what the project needs or wants. Honestly, I don't know what is motivating people to do this, I don't know what benefit they get from spending their Claude quota in this way. But for every low-effort prompt a maintainer needs to spend 10-100x more time triaging and reviewing those PRs. 

_I wonder how much time was collectively wasted as a result of [crabby rathbun](https://bsky.app/profile/jacobtomlinson.dev/post/3meolufa43k2c)?_

Many projects are adding AI policies to help govern these things, for example see the [Pandas Automated contributions policy](https://pandas.pydata.org/docs/dev/development/contributing.html#automated-contributions-policy). This is great, having these things written down gives maintainers an easy document to refer to when closing out low-effort work. But in my opinion having these policies wont stop low-effort PRs from happening, they just give you a mechanism to close them in a socially acceptable way.

We are at a pivotal turning point in software engineering where writing the code has gone from the hard part to the easy part. In 2026 I can generate an almost unlimited amount of code. So the problem now becomes generating the "right" code. So much more of our time as developers is being spent in planning, design, architecture and discussion, with the code creation being a final automated step. I think this is great, we probably should've been putting more effort into these areas for years. AI is forcing us to do some soul searching, growth can be painful but we will all be better for it.

I'm more excited about building software than I have been for years thanks to modern AI tooling. But I don't want to spend my time reviewing low-effort AI PRs, I want to spend my time designing, discussing and planning so that I can generate high-effort AI PRs myself.

To help fight the low-effort PRs I've been experimenting with adding guidance to the projects I work on to instruct agents not to work on things that aren't clearly marked as ready for implementation. Agent memory files like `CLAUDE.md` or `AGENTS.md` are a great place to put general guidance, but you can also create specific [skills](https://platform.claude.com/docs/en/agents-and-tools/agent-skills/overview) that will be triggered when the agent is prompted to do a certain thing.

For example in [`sphinx-llm`](https://github.com/NVIDIA/sphinx-llm) I have a `resolve-issue` skill which gets triggered when you ask the agent to work on a specific issue. The skill has some specific guidance around not working on things that aren't ready. Here's an abridged version of the skill to highlight the point I'm making, but you can view the [full skill](https://github.com/NVIDIA/sphinx-llm/blob/463eb0ef6b1aeb7fff7eec78d76ee9d9e7cca78e/.claude/skills/resolve-issue/SKILL.md) if you're interested.

```markdown
---
name: resolve-issue
description: |
  Triggers when you are asked to resolve a specific GitHub issue for the sphinx-llm project. Guides you through a test-driven workflow to resolve the issue, ensuring it is ready for work and following project conventions.
---

# Resolve GitHub Issue

This skill guides you through resolving a GitHub issue for the sphinx-llm
project using a test-driven development workflow.

## Instructions

When the user asks you to resolve a GitHub issue, follow these steps:

### Check for `ready` Label

**IMPORTANT:** Before proceeding with implementation, check if the issue has
the `ready` label.

- **If the `ready` label is present:** Proceed with implementation
- **If the `ready` label is missing:** Stop and inform the user:
  > "This issue does not have the `ready` label, which indicates it likely
  > requires more discussion before implementation. Please ensure the issue has
  > been discussed and approved by maintainers before proceeding."
```

My workflow in many projects these days is to spend 90% of my time crafting a good issue that explains what I want to happen, then simply prompting the agent to go ahead and implement it. Skills and instructions in the project guide and constrain the agent to implement things in a way that is appropriate for that specific project. The code generation part isn't that different from the low-effort prompters, but they are picking up random issues that may not be well written or are still actively being discussed. Garbage in, garbage out.

Today I wrote an issue about a small feature I wanted to add in `sphinx-llm` and I triggered the Claude skill to resolve it. I was pleased to see that it ran into my own safeguard here and Claude stopped and told me this issue isn't marked as `ready`.

```
❯ /resolve-issue 84                       

⏺ Let me start by fetching the issue details.
⏺ Bash(gh issue view 84 --repo NVIDIA/sphinx-llm)
  ⎿  title:     Make llms-full.txt configurable
     state:     OPEN
     author:    jacobtomlinson
     … +11 lines (ctrl+o to expand)

⏺ This issue does not have the ready label, which indicates it likely requires more discussion before implementation. Please ensure the issue has been discussed and approved by maintainers before proceeding.
```

Other times I've done this I have noticed it reason through this and say things like "You are the project maintainer, you opened the issue and you're telling me to solve it, so it's probably fine to keep going" but overall I would expect this simple instruction to at least make low-effort contributors stop and think before continuing.

This has gotten me thinking more about how we can include more prompts in our projects that could guide contributors through better behaviours. Perhaps instead of just saying "Please ensure the issue has been discussed and approved by maintainers before proceeding" we could add instructions which would cause the user to apply some critical thinking to the problem and engage in the discussion.

Maybe that's too hopeful and people will just bounce off an issue if their agent is hesitant to resolve it and asks them follow up questions. Either way it stops the low-effort PRs.
