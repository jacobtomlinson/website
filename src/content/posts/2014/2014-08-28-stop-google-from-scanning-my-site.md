---
author: Jacob Tomlinson
date: 2014-08-28T00:00:00+00:00
categories:
  - Web
  - Development
permalink: /2014/08/28/stop-google-from-scanning-my-site/
tags:
  - google
  - bots
  - robots.txt
thumbnail: google-developers
title: How to stop Google from scanning my site
url: /2014/08/28/stop-google-from-scanning-my-site/
---


Sometimes there may be occasions where you don't want Google (and other search engines)
to scan some or all of your website.

This may be because you don't want them to scan your admin control panel, or perhaps
you've set up a testing version of a website that you don't want most people to see.

This is where robots.txt comes in useful. This is a file which you place in the
root of your website that gives instructions to robots like Google's crawler on
what they can and can't look at.

Not all robots check this however so don't expect this to work as a way to stop
bad bots from scanning your website, but it's useful for the well behaved ones
like Google and Bing.

If you want to your whole website from being scanned create your robots.txt file
and put these two lines in it:

```
User-agent: *
Disallow: /
```

Or if you want to just block specific urls then enter something like this:

```
User-agent: *
Disallow: /admin
Disallow: /logs
```

Of course your admin pages should be secured properly and you really shouldn't
have logs available to the internet, but this is just an example after all.
