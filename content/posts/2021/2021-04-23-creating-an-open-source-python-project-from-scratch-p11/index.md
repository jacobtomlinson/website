---
title: "Communicating with your open source community"
series: ["Creating an open source Python project from scratch"]
date: 2021-04-23T00:00:00+00:00
draft: false
author: "Jacob Tomlinson"
categories:
  - blog
tags:
  - Python
  - GitHub
  - Tutorial
thumbnail: python
---

Once your open source Python project has users and a community you will likely want to communicate with them in an official capacity. Perhaps you want to tell them about a new release, show a use case where someone is using your tool or solicit feedback on an upcoming feature.

In this article we will talk about some common ways that maintainers communicate with their community.

## Broadcasting

When interacting with your community via GitHub issues or on Stack Overflow this will mostly be one-to-one communication.

But for cases where you need to make an announcement and tell folks about something you will need a one-to-many platform to broadcast to your community.

### Release notes

Often the most common thing you want to communicate to your community is that there has been a new release and that it has cool new features and bug fixes in it. So the first place to tell everyone about your release is in the release notes themselves.

In [Part 8](/posts/2021/automating-releases-of-python-packages-with-github-actions/) of this series we automated the release procedure for our `is-number` project. All we have to do to get a new release into the world is create a git tag and push it to GitHub.

```
$ export RELEASE=0.0.2
$ git commit --allow-empty -m "Release $RELEASE"
$ git tag -a $RELEASE -m "Version $RELEASE"
$ git push --tags
```

But if we head to [our releases page on GitHub](https://github.com/jacobtomlinson/is-number/releases) we just see a list of the tags we created.

![Screenshot of the is-number releases page on GitHub](https://i.imgur.com/qf7xtWP.png)

Instead of creating our tag from the command line we could also click the "Draft a new release" button on that page and create it from there. We can fill in our tag name, release name and add some release notes. We can also head over to our [commit history](https://github.com/jacobtomlinson/is-number/commits/master) to see what changed since our last version to help us write our notes.

![Screenshot of the draft a new release page with 0.0.3 information filled in](https://i.imgur.com/pV5tl12.png)

Now if we click publish release our tag will be created and our release automation will trigger as expected, but our releases page will have way more useful information for our users.

![Screenshot of is-number releases page with 0.0.3 showing more information than the other releases](https://i.imgur.com/SALP8tt.png)

### Twitter

Now that we have a new release out we need to tell our users. Not everyone will want to subscribe and watch our project on GitHub, so we need to go and find them where they already are. And for the tech community these days that means [Twitter](https://twitter.com).

#### Signing up

When you sign up for Twitter you'll be asked for a name, email and date of birth.

- For the **name** we could say `is-number`, `Is Number` or maybe `Python: Is Number?`. Something that allows folks to clearly understand what the account is representing.
- For the **email address** you could either sign up for a project specific Gmail account, or if you already use Gmail you can use the `+` feature where you can append extra words to the end of your own email address. For example if your email address is `alice@example.com` you could use `alice+is-number@example.com` for this project. Gmail ignores everything after the `+` so it'll still end up in your inbox, but it also adds a label with everything after the `+` to help you sort and filter your messages. It also allows you to create a unique email for Twitter even if you already signed up with your own account.
- For the **date of birth** set this to whatever you want, just make sure it is more than 12 years ago. It's tempting to set this to the date you made your first commit, or some other project milestone. But you have to be 12 years old to use Twitter, and setting your account younger than this can result in it being banned. I typically set the date and month to match the first commit, but set the year to something like 1970 or my own birth year.

Once we've set a password we will need to choose a profile picture. To keep things simple I recommend just taking a screenshot of your GitHub page, or just writing the project name in a text document and taking a screenshot of that. We don't need to worry about anything fancy like a logo at this point, although we will discuss branding in a future post. Just make sure the image clearly shows what the project is.

#### Tweeting

Once you have your account set up we need to start tweeting. I find the best way to manage this for project accounts is to come up with a list of situations where you want to tweet, then you can refer to that list later as various events happen in your project.

Here's a quick list of situations you probably should tweet:

- Tweet on new releases
- Tweet links if someone writes a blog post about your project
- Tweet links to active discussions on GitHub is you want more feedback
- Retweet someone if they tweet something positive about your project
- Quote tweet announcements from projects you depend on and mention what features you use

Be sure to always add context to your tweets.

Saying "Version 0.0.3 is out now!" doesn't contain much useful information. Instead try to say "Version 0.0.3 is out now which contains features X, Y and Z! Check it out now (link to release notes)".

If a project you depend on tweets about a new release you should quote tweet it and say something like "Excited to see version x.x.x of project A, we use feature X in `is-number`!".

#### Growing followers

Trying to grow your following on Twitter seems like the thing to do, but be mindful of your goals here. Twitter accounts for projects can typically be used in a couple of ways:

- Make announcements to your existing users and community
- Advertise your project and grow your community

So far we have only discussed the first one. If this is your only goal then don't worry too much about growing your audience. Add badges to your README, documentation and even release notes to point your users to the Twitter account. Those that want to subscribe to your announcements will follow you, and that's probably enough.

If you are also going to try to use Twitter to grow your audience then there is a [bunch of advice](https://business.twitter.com/en/blog/how-to-increase-twitter-followers.html) out on the internet for that.

### Blog

The last thing you may want to do to reach your audience is to start a project blog. Popular platforms for project blogs are [GitHub Pages](https://pages.github.com/), [Medium](https://medium.com/), [Tumblr](https://www.tumblr.com/) and [Blogger](https://www.blogger.com/).

Blogs can be useful for sharing information which doesn't fit into any of your other project information sources. Good topics to blog about are:

- **Case studies of your project**. Often these work best when co-authored with a user who is actively using your project.
- **Recaps of talks and events**. Did you talk about your project at a conference? Then share the abstract and a link to the YouTube video in a blog post.
- **Tutorials targeted at specific user groups** can be great, but first consider whether that content should be part of your documentation instead.
- **Deep dives into the guts of your project**. Why does it work a certain way? How does some bit of functionality work under the hood?
- **Comparisons with other tools**. While useful for helping users make decisions about using your tool these can be hard because it is [easy to be biased](https://matthewrocklin.com/blog/work/2017/03/09/biased-benchmarks).
- **Recapping an event or period of time**. A review of the last 12 months can make a great annual post. Or recapping all the contributions made during [Hacktoberfest](https://hacktoberfest.digitalocean.com/) can be a nice summary of the event.

When writing blog posts for your project be mindful about whether that content would be better placed somewhere else. If you are announcing some new feature should that just be in the release notes? If you are writing a guide on using your project should it be documentation?

Blog posts work best when they discuss specific scenarios or user groups, or when they cover high level topics and meta issues.

If you do decide to have a project blog don't forget to tweet about each new post!

## Summary

In this post we have covered:

- Writing release notes
- Creating a project Twitter account
- When to tweet as your project
- Creating project blogs
- What content makes a good blog post

In future posts we will cover:

- Communicating with your community
- Creating a contributor community
- Handling future maintenance