---
title: "Building a user community for your open source project"
series: ["Creating an open source Python project from scratch"]
date: 2021-04-16T00:00:00+00:00
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

Now that our open source Python project exists and users can install it we will want to turn our attention to sustainability, reach and ongoing maintenance. By putting it out there and gaining users you are opening yourself up to questions, bug reports and feature requests.

In this post we will discuss growing a community around your project for supporting your users.

## Users

When a user discovers your project the main resources they will have are your `README.rst` file and your documentation. If you've followed through this blog post series you should have covered enough information here that they can install and use your project. But what happens when they run into trouble?

Generally user queries fall into one of the following categories:

- **Bug**: The documentation says X but I can't get it to work. Is something broken?
- **Feature request**: I really wish your project did Y.
- **Usage question**: How do I do Z?
- **Design question**: Why does your project do P instead of Q like "insert competitor project here"?
- **Vague or long winded question**: The documentation says you can do A but I have a system which is unique and only allows B which is almost the same as A but things don't seem to be working and I can't try A exactly but if I try C then I kinda get the same result but not quite, etc, etc...

The first instinct for users who are familiar with GitHub may be to go and raise an issue. Issues are typically used for reporting bugs or requesting new features, but it's common for novice users to use them as a central place for all these types of queries.

As a community manager we should consider appropriate platforms to field each of these query types, but also confirm to the [principle of least astonishment](https://en.wikipedia.org/wiki/Principle_of_least_astonishment) and support users wherever they ask for support.

### Issue templates

Let's start with issue templates as this is the first place folks tend to go, so let's meet them there. Issue templates are a GitHub feature which allows you to include template files within your project that will pre-populate the issue form when users go to fill it in.

Let's create a bug report template in our `is-number` project. We do this by creating an `ISSUE_TEMPLATE` directory within the `.github` directory and then creating markdown templates with YAML frontmatter.

```console
$ mkdir .github/ISSUE_TEMPLATE
```

Now let's create our new file called `.github/ISSUE_TEMPLATE/bug.md`.

````markdown
---
name: Bug Report
about: Report a bug
labels:
  - bug
---

**What happened**:

**What you expected to happen**:

**Minimal Complete Verifiable Example**:

<!-- See http://matthewrocklin.com/blog/work/2018/02/28/minimal-bug-reports or https://stackoverflow.com/help/mcve for an example -->

```python
# Put your MCVE code here
```

**Anything else we need to know?**:

**Environment**:

- is-number version:
- Python version:
- Operating System:
- Install method (conda, pip, source):
````

The first section of our template enclosed in the `---` dashed is our YAML frontmatter. This is some data used by GitHub to describe our template. In this data we have the template a name and description and also list any issue labels which we want to have attached to issues created via this template.

Notice we added some instructions to the user within a `<!-- -->` HTML comment. We did this because user's will probably not remove anything from the template, only add, and we do not want these instructions to be visible in our final issue. If a user submits an issue containing HTML comments they will not be rendered by GitHub.

Then I'm going to checkout a new branch, commit this, push it up, raise a pull request and merge it. Alternatively you can just commit and push it to your main branch, but I strongly recommend following the same pull request process that you would expect contributors to follow.

```console
$ git checkout -b add-issue-template
$ git add .github/ISSUE_TEMPLATE/bug.md
$ git commit -m "Add issue template"
$ git push --set-upstream origin add-issue-template
```

![Screenshot of Add issue template pull request](https://i.imgur.com/jkqkRhP.png)

Now if we head to "Issues" and create a new issue we are presented with our template.

![Screenshot of issue template selection page](https://i.imgur.com/JCtG6vu.png)

![Screenshot of Add issue template](https://i.imgur.com/5ZZPgRa.png)

You can create more than one template, so let's add another one for feature requests called `.github/ISSUE_TEMPLATE/feature.md`.

````markdown
---
name: Feature Request
about: Request a new feature to be added
labels:
  - enhancement
---

<!-- Please do a quick search of existing issues to make sure that this has not been asked before. -->
````

This one is a little less prescriptive, but we've included a comment asking the user to check that this hasn't already been requested. Now if I branch, commit, push, and merge this change I can see my second option on the "create issue" screen.

![GitHub create issue page with bug and feature options](https://i.imgur.com/a9D2oE5.png)

You can also specify external links to be included in this multiple choice if your project uses other services for some query types. This is great because we can catch all our users who head over to raise an issue and direct them to the appropriate place.

### Q&A

For folks who want to ask a general "How do I do X?" question you could definitely add another issue template for this, and there is nothing wrong with doing that, but it's more common for technical folks to ask these questions on a dedicated Q&A website. [StackOverflow](https://stackoverflow.com/) is the most popular of these services, but GitHub also recently introduced a Q&A feature in their new [GitHub Discussions](https://docs.github.com/en/discussions) feature.

For the purposes of this post let's choose StackOverflow as our preferred location for `is-number` Q&A. It doesn't really matter what you choose, but it is helpful to just choose one and direct your community there.

Anyone can create a new question on StackOverflow about our `is-number` package. It would make sense for them to tag the question with things like `python`, and maybe if our project grows popular enough [StackOverflow will add a dedicated `is-number` tag too](https://stackoverflow.com/help/privileges/create-tags). Having a tag is useful because you can [subscribe to email notifications for specific tags](https://stackexchange.com/filters/413765/my-filter).

To ensure our community asks questions here we will want to add an option to our GitHub Issue template list which instead of creating an issue directs folks over to StackOverflow. We can do this by creating a `config.yml` file within our `.github/ISSUE_TEMPLATE` directory.

```yaml
blank_issues_enabled: false
contact_links:
  - name: General Question
    url: https://stackoverflow.com/questions/tagged/python
    about: "If you have a question like *How do I use is-number to check if a list is all numbers?* then please ask on Stack Overflow using the #Python tag."
```

Now if I branch, commit, push and merge this I'll see a third option on the "create new issue" page.

![GitHub create issue page with bug, feature and StackOverflow options](https://i.imgur.com/WKqcIIK.png)

### Forums

Some projects create forums for their community. These can also be great places to answer Q&A and also have longer design discussions about the project. Some folks have a section for use cases where the community can discuss and share how they are using the project.

Common choices for a forum are [Discourse](https://www.discourse.org/about), [Google Groups](https://groups.google.com/a/nvidia.com/forum/#!homeredir) and [GitHub Discussions](https://docs.github.com/en/discussions).

Just be sure to add a link to the `config.yml` file to direct folks over there.

### Chat

The last type of community platform that you commonly see in open source is chat rooms. These are great for real time communication with your community and can be a quick way to dig through a vague user question to figure out what they are really asking and then direct them to the appropriate place.

There are downsides to chat too. Often content can get buried in there which will not be surfaced by a search engine. If a user asks a question on StackOverflow and gets a good answer it is easy for others to find that page in the future. If the same happens in chat it's common for the answer to only be seen by a few people and it's common for the same question to be asked repeatedly by different folks.

I enjoy using chat to coordinate in real time with contributors to the project and hash things out faster than via GitHub issues but I'm also quick to direct folks to more appropriate channels.

If someone asks a Q&A style question I'll often say something like "Great question! Would you mind asking on StackOverflow so that others can benefit from it? Feel free to ping me the link in chat once you've done it so I can go in and answer it".

Also if I have a discussion with a contributor or another maintainer about a specific issue it's good to head to the issue after and summarize the conversation for others to see.

Chat is a tradeoff between convenience and speed with interruption and impermanence. If you do want to add chat to your community make sure you also link it from the issue template list in `config.yml`.

Popular chat services include [Slack](https://slack.com/), [Matrix](https://matrix.org/), [Gitter](https://gitter.im/), [Discord](https://discord.com/) and [IRC](https://en.wikipedia.org/wiki/Internet_Relay_Chat).

For `is-number` let's use Gitter. Head over to Gitter and log in with your GitHub Account. Then head to the [create community from repo](https://gitter.im/home/explore#createcommunity) page and find your GitHub repo.

![Screenshot of Gitter repo selection page](https://i.imgur.com/Bct6sQg.png)

Then create your room.

![Screenshot of is-number Gitter chat](https://i.imgur.com/glJ6EPS.png)

Folks can now head to [https://gitter.im/is-number/community](https://gitter.im/is-number/community), log in with their GitHub account and chat about `is-number`.

We should also add a link to our issue templates.

```markdown
  - name: Chat
    url: https://gitter.im/is-number/community
    about: "Chat about anything else with the community."
```

Then branch, commit, push and merge it to see the result.

![GitHub create issue page with bug, feature, StackOverflow and Gitter options](https://i.imgur.com/WjKqyAy.png)

### FAQ/Troubleshooting

Once your community starts to grow and folks are asking and answering questions you may want to add an FAQ or Troubleshooting section to the docs. Perhaps take the content from popular StackOverflow questions about your project, or questions which are frequently asked in chat, and write up some documentation on it. This is especially helpful for common gotchas which are not necessarily something that can be fixed in the project, but users often bump into.

Let's add an FAQ page called `docs/faq.rst` to `is-number`.

```rst
FAQ
===

Here's a list of commonly asked questions. If you can't find what you're looking for here you may
want to consider asking a question on `StackOverflow <https://stackoverflow.com/questions/tagged/python>`_.

Why should I use ``is-number``?
-------------------------------

You probably shouldn't! This project was built as part of a `blog post series on creating an open source Python project from scratch
<https://jacobtomlinson.dev/series/creating-an-open-source-python-project-from-scratch/>`_.

It contains a trivially small amount of code and is poking a bit of tongue-in-cheek fun at the JavaScript community
who are known for publishing very small packages on NPM and then having `half the internet break if one package is taken down
<https://www.theregister.com/2016/03/23/npm_left_pad_chaos/>`_.

If you want to check if something is a number in your Python code you should probably just write your own utility function
like this instead of adding another dependency.

.. code-block:: python

    def is_number(in_value):
        try:
            float(in_value)
            return True
        except (ValueError, TypeError):
            return False
```

We also need to make sure we link it in the `toctree` in `index.rst`.

```rst
...

.. toctree::
   :maxdepth: 2
   :caption: Contents:

   api
   faq
```

Now if I branch, commit, push and merge this code I can head over to my documentation website and see my new FAQ section.

![Screenshot of the is-number FAQ page](https://i.imgur.com/ksE6XnP.png)

### Creating a joined up experience

Note that from the FAQ page I've also directed folks to StackOverflow for anything else that isn't covered there. One of the most important things about creating a cohesive community based on multiple platforms is making it easy for folks to get to the right place.

We've made sure that we link all of our community resources from our "create an issue" page on GitHub. But what happens if a user with an FAQ style question ends up in chat? Or if someone with an API design question asks on StackOverflow?

For platforms that support it we should try to do the same as we have on GitHub by listing the types of queries and linking to the appropriate place to ask.

Slack has a [bot which can welcome new users to a channel](https://slack.com/intl/en-gb/slack-tips/automatically-onboard-new-channel-members). This message could share these links for you. Discourse also has a similar [welcome message feature](https://meta.discourse.org/t/changing-welcome-message/70407).

For platforms like StackOverflow there isn't an equivalent feature, so here you'll probably have to keep an eye on things and direct folks yourself. To make this easier on yourself you could point folks to the "new issue" page on GitHub which already has these links, or you could add some kind of list in your documentation.

For our `is-number` project let's add some more badges to our `README.rst` and documentation `index.rst` with links to our communities.

In the badges section of each file add some more community badges.

```rst
.. image:: https://img.shields.io/badge/FAQ-documentation-blue.svg
   :target: https://is-number.readthedocs.io/en/latest/faq.html
   :alt: Community FAQ
.. image:: https://img.shields.io/badge/Q&A-StackOverflow-orange.svg
   :target: https://stackoverflow.com/questions/tagged/python
   :alt: Q&A StackOverflow
.. image:: https://img.shields.io/badge/chat-gitter-green.svg
   :target: https://gitter.im/is-number/community
   :alt: Gitter chat
```

Then if we branch, commit, push and merge this change we can see our new badges on GitHub and our documentation.

![Screenshot of both GitHub and our documentation showing the new badges](https://i.imgur.com/Pqo4rDl.png)

## Summary

In this post we have covered:

- Creating issue templates
- Choosing a Q&A platform
- Forums
- Chat
- Adding an FAQ to the docs
- Joining up all our new community resources

In future posts we will cover:

- Communicating with your community
- Creating a contributor community
- Handling future maintenance
