---
title: What is semalt and why are they in my analytics?
author: Jacob Tomlinson
layout: post
permalink: /2014/09/03/what-is-semalt-and-why-are-they-in-my-analytics/
category: Web Development
thumbnail: google-analytics
tags:
  - google analytics
  - analytics
  - semalt
  - spam
---

You may be among the many people having their Google Analytics stats skewed by
fake referrals from semalt.

I've been noticing for a while that I get the occasional referral from semalt
on this blog. I was a bit curious at the time but as it was such a small percentage
of my traffic I didn't worry about it.

However when reviewing the analytics of another website I created I noticed that
the majority of its traffic was coming from semalt. The website is very low traffic
and so the amount of traffic from semalt was skewing the analytics dramatically.

I investigated further and it seems that semalt are some kind of analytics, SEO,
webmaster type company and you can pay them for services similar to Google
Analytics and Webmaster Tools. However they do seem slightly strange, particularly
because some services such as McAfee SiteAdvisor advises you not to visit them and
my work enterprise firewall blocks them completely under the category 'malware'.

If you do view their website (although I recommend you don't) they appear to be
a Ukranian startup of 11 people. They even have a dedicated page with information
about their web crawler. They state that it is "a software algorthym" which "simulates
real user behaviour" to collect their analytics. They go on to say that they do not
think their crawlers should cause any problems and that the "World Wide Web is
overfilled with crawlers".

I imagine that because their crawler simulates user behaviour Google Analytics has
a hard time filtering it out in the same way that it filters out other crawlers
like yahoo and bing.

They provide a facility to remove your domain from their list but if McAfee and others
are correct about this being a malware site then you may not want to type your
information in there.

It is possible to filter them out in Google Analytics, instructions can be found
[here][1], but a sure fire way to block them is to add the following to your
`.htacess` file:

```
RewriteEngine on
RewriteCond %{HTTP_REFERER} semalt\.com [NC]
RewriteRule .* -- [F]
```

[1]: https://productforums.google.com/forum/#!topic/analytics/ePCUyPkDVvs
