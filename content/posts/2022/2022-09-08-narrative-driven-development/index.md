---
title: "Narrative driven development"
date: 2022-09-08T00:00:00+00:00
draft: true
author: "Jacob Tomlinson"
categories:
  - blog
tags:
  - work
  - BPDD
thumbnail: journal
---

In July I published a blog post on [using Dask on KubeFlow with the Dask Kubernetes Operator](https://jacobtomlinson.dev/posts/2022/using-dask-on-kubeflow-with-the-dask-kubernetes-operator/). I originally outlined that post in January before the Dask Operator even existed as part of my planning for that work.

I like to call this process _Narrative Driven Development (NDD)_ and I do it all the time. Before starting a large piece of technical work I put some thought into how it will be communicated to the user community. The idea is similar to [Documentation Driven Development](https://dev.to/this-is-learning/a-better-way-to-code-documentation-driven-development-1kem) and definitely isn't limited to blog posts, while I often do _Blog Post Driven Development_ it could also be _Conference Talk Driven Development_ or even _Twitter Thread Driven Development_.

The intent is to think about the problem we are trying to solve for our users and consider how we would like to communicate this to our users when we are done. Having a narrative that describes the value of what we are building helps define what success looks like, but also really makes you think about API design and UX ahead of time. 

One benefit to _Conference Talk Driven Development_ is that it enforces an artificial deadline on the work too, which I sometimes find useful. There is a date on the calendar where that narrative has to be delivered.

This narrative approach is somewhat inspired by [Amazon's approach to product design they call "Working Backwards"](https://www.productplan.com/glossary/working-backward-amazon-method/). You start with the end user and draft the press release for your new product, then you decide if that story is compelling enough and worth the engineering effort required to achieve it, if so you build a project roadmap, allocate some resources and start your development cycle. 

## Narrative first development in open source

In open source software development it is much harder to enforce processes like Amazon's, so NDD is intended to be lightweight and done by individuals. It also helps open source developers communicate our work, which is something we are typically not good at.

### How open source projects announce functionality

Features within open source projects tend to follow these written steps:

- Someone identifies a need, either from their own use of a project or from being active within the community.
- They open an issue, "I need the project to do X".
- Someone raises a Pull Request, "adding X".
  - This often contains documentation on "how to do X".
- On the next release the changelog/release notes is updated, "x was added".

Some projects may go one step further and put out a release blog post, but this is often the changelog with bows on.

Once this is done the feature is out there in the world for everyone to discover and use. I want to make the argument here that this is not enough.

### Abstract features are not enough

Adding a feature to a project is great, I do it all the time. In fact I do it so often that I regularly forget that I've done it. More than once I have raised a pull request for something that I've already completed, merged and released six months earlier.

If I can't keep track of all the small features in the projects that I work on how can I expect users and other contributors to do the same?

The reason why incremental improvements to a piece of software are hard to track is because on their own they have little value, but as part of a larger narrative they hang together to empower some new workflow. 

For example I also recently published [Accelerating your ETL on Kubeflow with RAPIDS](https://developer.nvidia.com/blog/accelerating-etl-on-kubeflow-with-rapids/) on the NVIDIA Developer blog. In order to construct that narrative I had to add a bunch of features to `dask-kubernetes`, the RAPIDS Docker images and we even built the Dask Operator and [Container Canary](https://github.com/NVIDIA/container-canary) from the ground up. Each of those tasks on their own were hard to justify, but as part of a larger narrative they were clearly necessary and easier to articulate.

### Narrative use in agile software development

This kind of workflow is common in other types of development. In [agile software development](https://en.wikipedia.org/wiki/Agile_software_development) folks work with [user stories](https://en.wikipedia.org/wiki/User_story) as their unit of individual tasks. A story is a way of describing some piece of required functionality in the context of the person using it.

> As a [persona], I [want to], [so that].

You describe the user ([persona](https://en.wikipedia.org/wiki/Persona_(user_experience))), the functionality they require and the value that brings to them and the larger organisation.

Stories can also be grouped together into [epics](https://www.atlassian.com/agile/project-management/epics) which commonly represents the work performed by a team, but can be used to represent a complete workflow performed by a persona.

Many open source projects follow the Agile methodology, but only loosely. It is not very common, in my experience, to come across user stories on GitHub. But work is often captured in issues as individual tasks, with larger goals captured with projects and milestones.

NDD helps capture the same user focused values as user stories but with less effort and more flexibility.

### Getting open source developers to do marketing

By drafting up a blog post before implementing a feature you also reduce the barrier to actually publishing a blog post at the end to promote your new feature. Clicking merge on a PR isn't the end of a piece of development work, as much as folks like to think it is. That code still needs to be bundled in a release and then communicated to the community. Often it needs to be communicated multiple times in different ways and via different channels. 

By having a blog post draft up your sleeve you can polish it up with minimal effort and publish. Then you or others can releatedly come back to the post to crib out snippets for tweets, changelog entries, newsletters, etc. This up front effort hugely reduces the marketing effort that is needed when you've moved on and are thinking about the next thing you're building.

> If there is no blog post about something it didn't happen!

## What does Narrative Driven Development actually look like?

We've already seen a couple of examples of features and blog posts that were the result of narrative driven development, but let's talk through what the process actually looks like.

I work on various projects in the Python open source community including Dask, RAPIDS and Opsdroid. I've been working on these projects for years at this point, so I feel like I have a good model in my head around what exists today, how users use those tools and what is missing. My backlog of feature ideas is long and when I think of something I usually start by adding an item to a wishlist on the whiteboard in my home office.

I call it a wishlist because there are too many things on there and not enough time to do them, so it is a list of all the things I wish could exist. In order to make one of those wishes a reality I start by constructing a narrative of why that feature would be useful.

### Start with a high level issue

It is common in open source development, especially on GitHub, to start a new feature by opening an issue. This communicates what you want to work on and solicits feedback or help from developers on the project. This is the beginning of your narrative, but it is targeted at developers and not users. That's ok, but it's valuable to acknowledge who your audience is. Some power users may read issues and give the occasional thumbs up, but this is not the majority of folks who will benefit from the feature.

I still like to try and pitch the issue from a user perspective though. Explain what the feature is, how it will be implemented, how it will benefit the user and ideally include a proposed code snippet.

Generally in the open source projects I work on we use lazy concensus. So issues like this may get little or no input from other developers unless someone strongly wants it or strongly objects. Silence is taken as a sign that you're ok to move forwards.

### Building the user narrative

Next I start my blog post draft. Sometimes I do this in a Google Doc, or as a draft post on my blog or some other blog I contribute to. Whatever has the least friction.

My blog posts always start out as a couple of lists of bullet points. The first list is three things I want folks to take away from the post, for example the post you are reading now started with the following list:

- Encourage more folks to do narrative driven development
- Encourage more developers to write blog posts about the features they build
- Encourage more OSS developers to quantify what they are doing in terms of user value

I usually delete this list again at the end but find it useful to refer back to while I'm writing to ensure I'm communicating the message I started out to do and that my scope is realistic. I also do this for conference talks and any other outward communication.

This bullet point list will basically encompass the MVP for whatever you are building.

The second bullet point list is usually some structure. What sections would my post need in order to communicate my key takeaways. Once I have a rough outline I convert this list into headings and write TODO in the body of each section.

The first section is almost always some kind of executive summary style announcement.

> "You can now do X and achieve Y! Here's a cool code snippet or screenshot to demonstrate, keep reading to find out how and why."

Depending on the feature I may write a bunch of explaination on what the new feature is. For example with the Dask Operator I wanted to communicate that you can create Dask clusters with `kubectl` or the Python `KubeCluster` object, you can disconnect and reconnect from a running cluster and it'll work seamlessly with Kubeflow which generally uses Istio for network encryption and service discovery.

The rest of the TODOs will stay there until the engineering work has been done, but I have the bones of what I want to be able to communicate to users. This should probably not have taken longer than an hour or so.

### But it's all lies

The funny thing is at this point everything you've written in the blog post draft will be untrue. 

None of those things exist, the code examples don't work and it is all a fantasy. But by phrasing things as if they have already happened it really puts you in the mindset of "if I could say anything here what would I say?". I often spend a bunch of time playing around with the code snippets trying to decide what the most readable thing is that I could include here, which then hugely informs the API design once I start coding.

It also means that when I'm doing the actual engineering work I'm thinking about how the decisions I make will affect the blog post. Will my decision make a statement true? Will it make a code snippet valid? Will I need to change the blog post to match a decision, and if so is that good or bad for the user? 

This blog post is a draft so it should evolve over time, but you should only be changing it for the better.

### Raising more issues

Now that you have this narrative in place it should be clear what steps you need to take in order to make the things you've written true. This will often spawn many tasks that I like to also write up as issues. For large projects like the Dask Operator I'll also create a milestone to group those issue under, the milestone will be complete once reality matches what I've written in the blog post.

This also helps me spread the development load when I'm in the fortunate position of having folks to help out. The bulk of the development effort for the Dask Operator was shared between myself and an NVIDIA intern. Having a narrative written out and broken down into tasks meant that this particular 4 month internship was one of the most productive internships I've ever seen (although it helped that Matt was one of the best interns I've ever had the fortune of hiring). It was easy to communicate the vision of the work and justify all decisions because everything was being done in aid of making an exciting sounding blog post become real. It also meant that he could be more autonimous and make decisions because he could rationalise independently about whether something would match the narrative.

### Finishing the blog post

Once many PRs have been merged, issues have been closed and the milestone is complete you should now have a draft blog post that is now telling the truth. Once the the next release cycle has happened and users can actually install our changes I'll return to the blog post to finish it up.

Any remaining TODOs can be filled in relatively easily. You're not trying to stare at a blank page and come up with the best way to announce this new thing. You're taking a story that you've been thinking about and getting excited about and polishing it into something you can share online.

### Next steps

There are always more features. Your blog post may get comments from users asking for things or you may have had ideas during development for how you want to extend this next. For me this is when I start thinking about the next blog post. What story do I want to tell next? What workflow to I want to make true? Which features will that require?

For the Dask Operator work I want 