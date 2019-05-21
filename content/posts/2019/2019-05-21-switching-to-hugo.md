---
title: "Switching to Hugo"
date: 2019-05-21T14:21:28+01:00
draft: false
author: "Jacob Tomlinson"
categories:
  - blog
tags:
  - hugo
  - static sites
  - blog
thumbnail: hugo
---

It has been nearly two years since I published a new blog post on this website. That doesn't mean I haven't been writing things. It's just that much of my content has been posted on other platforms. I've decided recently to gather everything together and make this website the canonical source for the things I produce. This includes blog posts, talks, videos and more.

## Building an old Jekyll blog

Unsurprisingly when I came to update this blog I attempted to build it locally using Jekyll and nothing worked. This isn't the first time I've had problems building a Jekyll blog, and I put this mostly down to the fact that I never use Ruby unless I'm using Jekyll. Whenever I update my OS I get a new Ruby version, things break and I don't have a proper Ruby development environment set up. I shouldn't have to have one to write a blog anyway.

## Hugo

[Hugo](https://gohugo.io/) has been on my radar for a while now. It is another static site generator but this one is written in Go. The main benefit of this is that Go applications compile to a single binary, which means there is no development environment to maintain. Wonderful!

There are many other benefits of using Hugo, but a simple build environment is the thing I care most about. Therefore I decided I was going to port this blog over to Hugo.

## Porting

There are two main things which need porting over; the content, and the theme.

### Content

Markdown is markdown. Porting my Jekyll posts over to Hugo posts was as simple as moving them from `_posts` into `content/posts`.

There are a couple of differences around the way that Hugo builds pages. The date is taken from the frontmatter instead of the file name. I actually prefer this as I often take a few days to write and publish a post and tweaking some frontmatter feels more natural then renaming a file.

I also went through my posts and replaced the `url` frontmatter with `aliases` instead. This was a conscious decision to embrace Hugo's url formatting but also to make sure my old links didn't break.

### Theme

Migrating the theme was a little more complicated. This is partly because the theme I created for my blog called [Carte Noire](https://github.com/jacobtomlinson/carte-noire) was built before Jekyll separated themes out into a better file structure.

This meant that my theme code and content were quite intertwined and getting updated from the main theme was tricky. Things have progressed in the Jekyll world, but I hadn't kept pace. So now seemed like a good time to separate things once and for all. Hugo has an excellent theme system where they live in a `themes` folder and you are encouraged to include them as git submodules.

To create my submodule I've created [Carte Noire Hugo](https://github.com/jacobtomlinson/carte-noire-hugo) which is a direct port of my Jekyll theme into the Hugo structure. Much of the syntax was extremely similar. Jekyll uses the Liquid language and Hugo uses Go's built in templating. It wasn't quite a copy and paste exercise but it was mostly replacing things like `{{ post.title}}` with `{{ .Title }}`.

### Cleaning up

Switching away from Jekyll has helped me massively clean up this repository. My `.travis.yml` is simpler, I can delete my `Rakefile` and `Gemfile` and sort everything into what feels like a much more sensible heirarchy.

## Next steps

Now that I've migrated across and hopefully created myself a more maintainable blog I intend to scour the web for content that I have created and retrospectively syndicate and link to it from here.
