---
title: Collaborative article corrections in Jekyll
author: Jacob Tomlinson
layout: post
category: Jekyll
thumbnail: jekyll
tags:
- jekyll
- github
- collaboration
excerpt: 'How to avoid people saying "I really enjoyed your article on X, but your spelled pedantic wrong!".'
---

Don't you find it really useful when you publish an article on your blog and then someone comes up to you and says:

> I really enjoyed your article but you spelled pedantic wrong!

...well not any more!

There are a thousand great things about using [jekyll][jekyll] on [GitHub][github], and the one I'm going to highlight now is that anyone can edit your content. Most people who host their blog on [GitHub Pages][github-pages] using jekyll do so in a public repository, therefore anyone can fork it, make a change and open a pull request back in. Sadly for most readers this isn't really convenient, if someone is reading your article and they spot a typo it is highly unlikely they will get the sudden urge to look up your repository and fix it for you.

But thanks to jekyll's `page.path` value you can easily create a link from your page to the raw page in your repository. `page.path` is the location of your original markdown or textile file, not the location your page ends up after jekyll has done its magic, therefore you can use it to link to the file in your reporitory.

You may spot at the end of this article a nice link which allows you to 'suggest an edit' to this article. It is simply a link to the markdown file for this article on the master branch of my GitHub repository. From there it is easy to click the edit button, make your change, select 'open pull request', add a sentence on what you changed and wait for it to be merged in.

The code for the link looks something like this
{% raw %}

```
<a href="{{site.github_repository}}/blob/master/{{page.path}}">
  link
</a>

```

{% endraw %}

I've already defined `github_repository` in my `_config.yml` with the url of my blog repo so I can just link to the master branch blob followed by the actual path of this file within the repository.

Now all you need to do is include this link in an easy to find place on your page and if you're really lucky then instead of telling you about your mistakes your readers will fix them for you.

[github]: https://github.com/
[github-pages]: https://pages.github.com/
[jekyll]: http://jekyllrb.com/
