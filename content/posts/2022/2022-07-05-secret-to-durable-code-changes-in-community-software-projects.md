---
title: "The secret to making code contributions that stand the test of time"
date: 2022-07-05T00:00:00+00:00
draft: false
author: "Jacob Tomlinson"
categories:
  - blog
tags:
  - Open Source
  - SoftwareEngineering
  - Community
---

When you contribute code to collaborative projects, whether they are open source community projects or large internal projects inside organisations, the feeling of having your code running inside a large application can be very rewarding.

I spend a lot of my time contributing code to [Dask](https://www.dask.org/) which is used by many folks around the world. It is exciting knowing that code I've written in Dask is being [run on supercomputers at NASA](https://www.nas.nasa.gov/hecc/support/kb/using-dask-at-nas_648.html) and powers popular services like [recommendations at Spotify](https://www.nvidia.com/en-us/on-demand/session/gtcspring21-s32017/).

```warning
There are [many ways to contribute](https://twitter.com/choldgraf/status/1544046930375966720?s=20&t=bOt29HEwPB1Sfq2xughMWw) to software projects other than directly writing code, but in this post I want to focus on code contributions specifically and the feeling of excitement I get when I think about actual lines on Python that I have typed being executed in interesting and important places.
```

The other side of this feeling is a strange one. Code gets replaced and improved over time, especially in active community projects, and that's a good thing. But it means that if I stopped contributing code now the number of lines of my code that run in cool places will dwindle and eventually drop back to zero.

I've definitely reached the stage in my career where [I want there to be less of my code in the world](https://twitter.com/futureimperfect/status/1311519750418247681?s=20&t=c4kGyGHtMbJqVd8A8F_QQw) and I get excited when I see my code replaced with more optimal, readable, maintainable contributions. But there is also a sadness to seeing my code being removed.

I think this sadness stems from associating my lines of code with my impact on the codebase, but I want to challenge this because it isn't how things work. My goal is usually to change the way software works to make it more robust, more flexible and ultimately more useful. I want to add new features and fix bugs so that it delivers more value to the users. So if my actual lines of code are going to be replaced over time how do I ensure that the replacement code still does what I intended our code to do and builds upon my goals?

The secret to making changes to software that stand the test of time is to make **tests and documentation** your main vehicle for change.

## Capturing your intent with tests

When you change code in a project you are applying your intensions directly. Fixing bugs is a nice way to picture this, if a function does something that it shouldn't you can fix the bug by changing the code in the function. However when someone else comes along and changes the function again or even does a refactor that removes the function altogether it supersedes your bug fix and the code you wrote is gone. They could also accidentally reintroduce the bug that you fixed which effectively reverts your contribution.

To guard against regressions like this it is common to also write a unit test that verifies that the bug is gone. Every contribution that comes after yours must pass the tests so even if the function is modified or removed the bug cannot come back.

Therefore the actual contribution you made to the project is the test. The fix you made to the function was just a temporary implementation detail. Your intent was to stop the bug from happening, and by writing a test you ensured that it doesn't.

[Test Driven Development](https://en.wikipedia.org/wiki/Test-driven_development) (TDD) is already a popular workflow for many developers where you start by writing the tests before you implement the actual code. I see TDD as a developer defining their intentions ahead of time and then updating the world to meet those expectations.

In my experience tests also get updated over time, especially if they are brittle or too specific. But I rarely see tests being removed or even modified away from their core intentions. So if you want a piece of software to do something you should write a test to check that it does the that thing. You could even write whole test suite that lays out a full application and hand it to another developer to implement and you can be sure that the application will do what you wanted it to do even though you have written no code in the actual implementation.

Your intentions are safe when they are verified by unit tests.

## Communicating your intent with documentation

At a high level, documentation explains to your users how to use your software, but if you look a little deeper the documentation is effectively a contract that you create with your community. The documentation describes what your software does and why that is useful. Tests can then be used to validate that the claims made in the documentation are true, but ultimately the docs are the public communication of your intentions. This is how we can build reliable software even with many developers making frequent changes. As long as the software does what users expect, based on what the documentation describes, then they will perceive it as being reliable.

If your intention is to implement a new feature in a project then it is critical that you document that feature so that users can discover and use it.

> If a feature is implemented but no one can read about it in the documentation, does it even exist?

By describing a feature in a narrative with code examples you are laying out how you intend something will work and how it can be used. This documentation then becomes the measure for whether something is implemented correctly, if the code does not behave like the documentation describes then there is a bug and something is broken.

Similar to our test example if you implement a new feature and write some documentation that describes it, then the documentation is the impactful thing that you have contributed. You've extended the contract with the user and laid out what should exist. The code you write that implements the feature is necessary because you need to make sure the documentation you wrote is true, but that is all that it is.

## The What, Why and How of contributions

As a developer if I make a change to a piece of code it is because I have a specific intention in mind. I want to fix a bug, implement a new feature, improve performance or improve maintainability.

When I make changes I try to think about "why" am I changing it, can I describe the problem I am solving in writing with code examples? Then I think about "what" does my change need to do, can I describe the desired behavior with a test? Then lastly I think about "how" can I implement it with some code.

We see this methodology often within the open source community but I feel like it comes around naturally with experience rather than discussed explicitly. So let's discuss it now.

Many projects require you to raise an issue before starting on a pull request, this is because you need to capture the "why" ahead of time. The issue can be a description of a bug, a design proposal for a feature or a discussion around performance or maintenance. I also find it helpful to think about the issue as the foundation that will eventually become the documentation and write it in a way that communicates to the community what my intentions are rather than communicating to a developer what my intentions are.

Many developers also use TDD, I am personally a huge fan. So I will start a new branch and implement a test that measures whether the code meets my intentions. I haven't done any implementation yet so this will be a failing test that highlights a bug or uses a new feature the way I intend users to use it.

Then I write the actual implementation that makes the test pass and the documentation true. This is often the thing folks will do first, but doing it in this order really highlights that the actual code is less important than communicating and validating your intended change.

Once things pass I'll create a Pull Request. This is the next opportunity to communicate my intent after opening the issue. I try to write PRs in the style of documentation, it describes to a user what the PR is solving and how to use it. I may also include information on the implementation decisions that are targeted at a reviewer but I commonly see folks doing this on GitHub by doing a first review of their PR themselves and including this info in comments.

I try to keep the PR up to date as the review goes on and once things are generally accepted I convert the PR text into a documentation page and add that to the PR.

## What happens after the merge

Once the PR is merged there is now modified code, some tests and some documentation in the project. There will be bugs in this, because there always are. Other folks will have other ideas for improving things, because they always do. So over time people will come along and make more changes to this part of the code. But they will do so in a way that aims to get all tests passing and doesn't invalidate documentation.

Of course someone could disagree with you and undo your intentions. But if they fundamentally want to change what you intended to happen in your PR they will need to consider how to communicate that to users and they will need to remove/modify tests. The result of this is that they will also consider the "why" of your intention and hopefully witll incorporate that into the "why" of their own intention.

This isn't a silver bullet to ensure any change you make lives on forever. Features do get removed and things get deleted over time. But by ensuring your intentions are documented and tested you are ensuring the project builds upon your contributions even if your implementation code is ultimately replaced.

My favorite example of this is the [Met Office integration in Home Assistant](https://www.home-assistant.io/integrations/metoffice/). I contributed this feature a long time ago to pull open weather data from my (then) employer into my home automation setup. If you have a [look at the integration code](https://github.com/home-assistant/core/blame/dev/homeassistant/components/metoffice/weather.py) today and do a git blame I don't think you'll find many lines written by me. But the feature still exists and I can still use it, and that is because of me and intention. It's just way better now than when I implemented it.
