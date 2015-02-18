---
title: Has Stack Overflow been hijacked?
author: Jacob Tomlinson
layout: post
permalink: /2013/02/27/has-stack-overflow-been-hacked/
category: News
thumbnail: stack-overflow
tags:
  - Google
  - Hijack
  - Stack Overflow
---
Something funny seems to be happening with Google and Stack Overflow.

I first noticed this last night when on my Macbook at home. I went to www.google.co.uk, typed in "stackoverflow" and was presented by the usual page. However I noticed that the url displyed under the link on the search results said www.doioig.gov. When clicking the link it took me to www.doioig.gov instead of www.stackoverflow.com. I thought to myself that this was probably just a temporary issue and went to the correct url myself.

But when I got to work this morning and logged onto my RHEL box and tried to get onto Stack Overflow again I went to www.google.co.uk (well used firefox's search bar) and typed "stackoverflow" and again I have been presented with the result but linking to www.doioig.gov.

![Google Search for Stack Overflow](http://i.imgur.com/NwpQL7P.png)

It seems that www.doioig.gov is the page for "Office of Inspector General" and appears to be a very harmless US government website.

![DOI Website](http://i.imgur.com/C9tvnt2.png)

I did a quick search for "stackoverflow doioig" (as you can see in the screenshots) to see if anyone else had seen this issue but there is nothing in the search that appears relevant at time of writing.

So I am curious to know if this is an accidental hijacking of stackoverflow's result on google, it has the correct page description and everything which is odd, or whether something more interesting is happening at google.

**Update 11/03/2013**  
It appears the the answer has been found here <http://meta.stackoverflow.com/questions/169405/google-indexing-issue-for-keyword-stackoverflow>.  
It seems that the website www.doioig.gov has been replaced with a new site and so the developer put in some redirect code. They must've not know how to do this redirect and looked it up on stackoverflow. They copied the code from here <http://stackoverflow.com/revisions/5411601/3> and updated the url to the new website. However they only updated one url. The code from stackoverflow has an html redirect using the refresh tag and also a javascript redirect using window.location. They updated the url in the window.location but left the example stackoverflow url in the html refresh. So anybody with javascript enabled in their browser was taken to the new website but if you had javascript disabled (as googlebot does) it would have redirected you to stackoverflow.com. This website must have has a high page rank and lots of influence and therefore google updated their listing to show doioig.gov as stackoverflow's new url. Google have since rectified this and the search now shows up the correct website.

 [1]: http://www.jacobtomlinson.co.uk/wp-content/uploads/2013/02/stackoverflowdoioig.png
 [2]: http://www.jacobtomlinson.co.uk/wp-content/uploads/2013/02/doioig.png
