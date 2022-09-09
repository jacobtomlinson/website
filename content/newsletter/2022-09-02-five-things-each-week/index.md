---
title: "Issue 1: Five things each week"
date: 2022-09-02T00:00:00+00:00
draft: false
author: "Jacob Tomlinson"
---

Thank you for reading my very first newsletter! Everyone and their dog seems to have a newsletter these days, so I thought why not me.

It seems fitting to use my first newsletter to give a bit of background on why I am starting a newsletter.

## 1. Why a newsletter?

Let's start with a goal, and it's a simple one:

> I want to increase [my Twitter following](https://twitter.com/_jacobtomlinson).

As a software engineer working both remotely and in open source it's really important for me to stay connected to my colleagues, collaborators and the open source community. One important tool that I use for this is Twitter. Not only can I amplify the things that are important to me but I can also have conversations with users, contributors and other maintainers. It is a very useful tool.

I plan to direct newsletter readers to my Twitter and amplify my newsletter issues via tweets in the hope that it creates a virtuous cycle and I get more followers over time.

## 2. What goes in my newsletter?

I regularly tweet about what I am working on, but I also maintain a list the things I work on each day that I use for reference during the various standups I attend. Sometimes I note something down that I want to share with a wider audience but it's too much for a tweet and not enough for a blog post. Sometimes I want to share something like a podcast episode that I have enjoyed, but with a little more context than a tweet allows. I also want to amplify things I have worked on like blog posts or stream collabs.

Having a newsletter seems like a great place for this in-between-ish content to go.

## 3. When will I send out my newsletter?

The plan is to send out my newsletter every Friday afternoon. Given that the intent is to provide a bit of a retrospective of the week this seems fitting.

I am also in Europe and many of the folks I interact with in open source are in North America so by sending it out during my afternoon it gives the majority of my readers some Friday lunchtime reading material.

## 4. What technology/platform am I using for my newsletter?

My [personal website](https://jacobtomlinson.dev/) is the home to [my blog](https://jacobtomlinson.dev/posts/) but also intends to be an archive of my professional footprint on the internet like keeping track all of the [conference talks I've given](https://jacobtomlinson.dev/talks/).

So I definitely want my newsletters to be archived onto my website.

I explored a few newsletter tools like MailChimp, ButtonDown and Revue but didn't want to have to write on their platforms and manually archive the content to my blog. Instead I've opted to write the newsletters in the same way I write blog posts, as markdown files in a Git repo that gets built with [Hugo](https://gohugo.io/).

I'm using [Mailgun's API](https://www.mailgun.com/) to handle sending out the emails and a little automation via [Netlify Functions](https://www.netlify.com/products/functions/) and [GitHub Actions](https://github.com/features/actions) to add/remove subscribers and to send out the newsletter from the built HTML files.

I intend to write a full blog post about this.

## 5. How do you subscribe?

To sign up you can head to the [newsletter page](https://jacobtomlinson.dev/newsletter/) on my website and submit your email address via the box at the top. You'll get an email asking you to click a link to confirm to ensure that everyone on the mailing list definitely wants to be on it.

You can also head back to the same form and use the "Unsubscribe" button on the right if you change your mind, or click the unsubscribe link in the email. This isn't a lifelong commitment, it's just for as long as you find it useful 😁.

Thanks for reading!

![Lenny the dog looking confused](lenny.jpg "Lenny hopes you had a good week! (I did say everyone and their dog)")
