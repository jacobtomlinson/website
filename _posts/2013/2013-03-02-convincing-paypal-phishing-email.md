---
title: Convincing Paypal Phishing Email
author: Jacob Tomlinson
layout: post
permalink: /2013/03/02/convincing-paypal-phishing-email/
category: News
thumbnail: paypal
tags:
  - Hacking
  - Paypal
  - Phishing
---
So I woke up yesterday morning to find a receipt from Paypal for $149.49 on my iPhone. I haven&#8217;t bought anything for $149.49 so right away I was worried that my Paypal account had been broken into. The first thing I did was to log into my Paypal account on my laptop to check to see more information about the transaction. When logging in there was no trace of this transaction.

At that point I became suspicious that this was a phishing email and wasn&#8217;t actually real. When looking a bit closer into the email I noticed a few things which gave it away as fake. The links don&#8217;t go to Paypal, they go off to some random website. My name isn&#8217;t on the email, Paypal always address you by your full name, something like &#8220;Dear Jacob Tomlinson&#8221; but this email just says &#8220;Hello&#8221;.

For a further comparison I found an old email from Paypal and had a look at the differences. Here they are for reference.

![Paypal Phishing Email](http://i.imgur.com/cNkWXkc.png)

Paypal Phishing Email

![Real Paypal Email](http://i.imgur.com/IXfXkdU.png)

Real Paypal Email

I decided to do a bit of digging into the origin of this email so I copied the link but chopped the script name off the end. This took me to http://www.thorpe-hall.co.uk/templates/beez/. Which appears to be a login to the phishing control panel.

![Phishing Login Page](http://i.imgur.com/2jp4bYI.png)

Phishing Login Page

I then went to the root domain that this control panel is hosted on http://www.thorpe-hall.co.uk. This appears to just be a perfectly innocent website for a manor house in Yorkshire. So my guess would be that they have been compromised and had some malicious code installed on their website without them knowing. I attempted to contact them about this but didn&#8217;t manage to speak to anybody and have since sent an email to them.

![Website hijacked and used as phishing site](http://i.imgur.com/b4QXsin.png)

Website hijacked and used as phishing site

I assume that if I had clicked on any of the links in the email I would have seen a Paypal login page (I may investigate this from a secure computer if I get a chance) and when I logged in the hacker would have been sent my login details for Paypal and would have been able to withdraw money from my accounts.

So make sure if you get any emails from Paypal for things you haven&#8217;t bought that you don&#8217;t click any links and go straight to Paypal and get in touch with them.

**Update: 11/03/2013**  
Since writing this article I have had response from Thorpe Hall who thanked me for notifying them of the issue and asked if I could provide the screenshots and so I sent them back an email with a couple of links to the images in this blog post. Hopefully this will allow them to get the issue resolved quickly.
