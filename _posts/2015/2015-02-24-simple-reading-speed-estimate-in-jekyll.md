---
title: Simple reading speed estimate in Jekyll
author: Jacob Tomlinson
layout: post
category: Jekyll
thumbnail: jekyll
tags:
- jekyll
- reading speed
excerpt: "I wanted to share a nice code snippet I found for estimating the
reading speed of an article in jekyll."
---

When browsing through off-the-shelf Jekyll themes recently I stumbled across
one I really like called [Pixyll][1]. There are lots of things I like
about the theme but one thing in particular is the reading speed estimate
at the top of each article. Not only is it a nice feature but the code is simple and
concise too!

The code was first [contributed][4] to the project by GitHub user [WrinklyNinja][2] and I decided I
liked it so much I would add it to my own blog.

This simple one line of code (which I included immediately after my frontmatter)
takes the number of words inside the content of
your article and divides by 180 words per minute. This is lower than the
generally accepted 200-250 words per minute of the average adult but I think
having it a bit lower takes into account thoroughly reading through code snippets
and looking at images.

{% raw %}
```
{% assign minutes = content | number_of_words | divided_by: 180 %}
```

The next line simply rounds up from 0 if the article will take less than one minute
to read.

```
{% if minutes == 0 %}{% assign minutes = 1 %}{% endif %}
```

Then to display the minute count on the page simply use the variable in whatever
sentence you like.

```
A {{ minutes }} minute read
```

Taking it a bit further, Rassol made a [nice addition][6] which allows you to override the number
of minutes in the pages YAML Frontmatter. However he decided to add the logic
to the display code , I felt it would probably be better to add it to the calculation
because there is no point in calculating it if it has already been specified.

[My changes][5] wrap around the first two lines to look like this.

```
{% if page.minutes %}
  {% assign minutes = page.minutes %}
{% else %}
  {% assign minutes = content | number_of_words | divided_by: 180 %}
  {% if minutes == 0 %}{% assign minutes = 1 %}{% endif %}
{% endif %}
```

And you can simply specify the number of minutes in the frontmatter.

```
minutes: 8
```
{% endraw %}

Adding a reading length to your article lets people know how long
they need to spend on your content and make a quick decision about whether
they have enough time to do so.


[1]: https://github.com/johnotander/pixyll
[2]: https://github.com/WrinklyNinja
[4]: https://github.com/johnotander/pixyll/commit/0979ada039a7a36b3383c0739b1e7d7edb1f34f9
[5]: https://github.com/johnotander/pixyll/commit/0742c98c9e379de607ff840595ceb29583e139a7
[6]: https://github.com/johnotander/pixyll/commit/16fc4ea2c83eadd40781f914f8654567926e6b7e
