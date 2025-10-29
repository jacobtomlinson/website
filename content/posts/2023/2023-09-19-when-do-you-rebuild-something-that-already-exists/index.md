---
title: "When to rebuild things that already exist"
date: 2023-09-19T00:00:00+00:00
draft: true
author: "Jacob Tomlinson"
categories:
  - blog
tags:
  - software development
  - projects
  - talk
---

This year [I built a library that already exists](https://jacobtomlinson.dev/posts/2023/introducing-kr8s-a-new-kubernetes-client-library-for-python-inspired-by-kubectl/). The existing solutions [didn’t quite meet my needs](https://jacobtomlinson.dev/posts/2023/comparison-of-kr8s-vs-other-python-libraries-for-kubernetes/), I wanted something that ticked all of my boxes. When I was thinking about building something new people referred me to [xkcd #927](https://xkcd.com/927/). **But I did it anyway.**

![](xkcd-927-standards.png "Fortunately, the charging one has been solved now that we've all standardized on mini-USB. Or is it micro-USB? Shit.")

Don't get me wrong, I totally agree with the sentiment of the comic. It takes [a mandate from a supernational policital and economic union](https://www.theverge.com/2023/9/16/23872237/apple-iphone-15-usb-c-switch-guide) to get everyone to converge on even the most obvious and sensible standards, and I doubt the European Union cares about Python libraries quite as much as phone chargers. But there is definitely a time and a place to redo work that someone else has already done.

## The duality of permenance in software

Writing software is weird. In some ways it feels like building sandcastles that will get washed away the next time the tide comes in. Software is continuously superseded by new software.

> The often-made analogy between constructing a bridge or house and constructing a software system is therefore flawed. **Building software is perhaps closer to constructing a sand castle at the beach** or playing a game of sticks. Building sand castles instead of concrete structures requires different skills and tools.
>
> _An excerpt from "Code Craft" by Gerard J. Holzmann published in [Computing Edge Magazine in May 2017](https://ieeecs-media.computer.org/assets/pdf/ce-may17-final.pdf)_

But this doesn't hapeen to 100% of it 100% of the time. I guarantee that the device you are reading this on will execute code today [that was written decades ago](https://github.com/bminor/glibc/blame/master/io/chmod.c). So in many ways software that meets the bar of being "good enough" becomes permanent. 

Do you think the stone masons who built this Roman wall in Exeter had imagined it would still be in use thousands of years later to separate a shopping centre from some offices? And if they knew do you think they would've done anything differently?

![](exeter-roman-wall.jpg "Photo provided under Creative Commons by Sophia Feltham - https://www.geograph.org.uk/photo/7482577")

So as software engineers who experience this duality at an accelerated rate we have to be comfortable with both the idea that what we build today will likely be discarded soon and replaced with something better, while at the same time feeling the responsibility that there is a chance that this code may end up still being used decades down the line.

## Reasons to start again

Choosing whether to start again boils down to the things that exist today not meeting your needs, but having no way to affect change on the existing solutions to bridge the gap. This inability or lack of desire to change things can be techinical, legal, political, commercial or even personal.

In the world of proprietary software this is pretty cut and dry. The software that exists is closed-source and without joining or buying the company that built them you simply cannot make changes. So if you want to build the thing that you need your only option is to start again.

However with open-source software this gets a little more complicated. I love open-source and I would always encourage you to try and contribute to a project to fix bugs and fill in gaps before starting over. But even though the majority of projects accept contributions there are many reasons why your contribution will not be accepted. The project could just be unmaintained, lots of software is maintained by volunteers and for a variety of reasons they may just go away and stop reviewing code and making new releases. 

You might need to change things substantially enough that it degrades the experience of other users and so maintainers reject your Pull Request. You might just want to take the project in a direction that the maintainers aren't happy with. In these cases there is another option that is open to you, forking.

## Alternatives to starting again

Forking a project is a major undertaking. You gain control over the future of the fork, but you also inherit the baggage and technical debt of the existing codebase. You are likely less experienced with the project than the maintainers of the original repository and so you will probably make slower progress. [Many forks run out of steam if the main repo is still active and they are being developed in parallel](https://www.infoworld.com/article/3219689/nodejs-forks-again-this-time-over-a-political-dispute.html), often only forks of unmaintained repos seem to gain any traction. 

[paperless-ngx](https://github.com/paperless-ngx/paperless-ngx) is one such example where the original project became stale and was forked, then that fork ran out of steam and was forked again. This fork of a fork has a vibrant community, but is effectively the continuation of the original project, not an alternative path.

The other option you have is wrapping. Instead of changing the library you want to use you could introduce a layer on top that exposes the functionality you need. You build this layer using the existing library and your own code to fill in the gaps. This allows you to have full agency over the code you are writing while also decoupling the implementation from your code allowing you or others to reuse your wrapper in other projects. 

The downside to this is that you are now depending on this upstream implementation and you will likely run into situations where upstream changes break your wrapper, this is even more likely if you start digging into internals and using private methods to fill in gaps (maybe this is more of a Python problem). You are also beholden to bugs in the upstream library and you may not have the agency to fix them quickly even if users of your wrapper are reporting them to you.

## Licensing

You might want to start again for licensing reasons. If the software you want to change has a license like the [GPLv3](https://www.gnu.org/licenses/gpl-3.0.en.html) you may want whatever you build to have a more permissive license in order to keep the door to future commercialisation open. Equally, if the software you want to contribute to has a license like the [Apache 2.0](https://www.apache.org/licenses/LICENSE-2.0) you might want to start again with a GPLv3 license to keep the door to commercialisation firmly closed. In both cases you need to [write new code under your desired license](https://en.wikipedia.org/wiki/Clean_room_design).

## Not invented here

Many projects are rebuilt due to [Not Invented Here (NIH) syndrome](https://en.wikipedia.org/wiki/Not_invented_here) which is used as a pejorative term and suggests it is an antipattern. 

Often NIH is synonymous with ego, control and a misplaced feeling that you can do exactly the same thing but better, faster, cheaper, etc. This is definitely something to be very mindful of but I would encourage you to try not to dismiss NIH as a negative immediately and instead consider that there are many valid reasons why **you** should rebuild something that someone else has already built.

I find the best way to rule out the antipattern of NIH is to ask yourself "am I building exactly the same thing but better? or am I building a novel and different thing?". If we are honest with outselves every time we start from scratch there is some amount of NIH going on, just make sure it's not the driving factor in your decision making.

## Revolution over evolution

If the existig solutions don't quite meet your requirements is that because there is a gap, or is it because the design fundamentally doesn't solve problem? Starting again is a perfect opportunity to do things very differently and think about things with a different mindset. 

Instead of contributing to an existing project or even forking it in order to bend it to your will is there another approach that more neatly solves your needs?

Revolution also doesn't need to be about a paradigm shift that allows you to do more, it also can be a simplification that allows you to do the same with less.

Code is a liability, it needs to be maintained and we should all strive to write less code. But we should also stive to depend on less code too, so if we can replace our dependencies with something more lightweight and maintainable then our code can also become more maintainable.

## Learning

The other totally valid reason to rebuild something is to simply learn how it works. Building something from scratch to imitate something that already exists is a fantastic way to learn. 

I've recently been enjoying [videos on YouTube by @hyperplexed](https://www.youtube.com/@Hyperplexed) where they look at award winning websites and rebuild some of the elements to understand how they are created. 

The things they build are toy examples to explain how something works, but what if that exmple takes on a life of its own? In the video [Unravelling the Magic behind Polyrhythms](https://www.youtube.com/watch?v=Kt3DavtVGVE) @hyperplexed reproduces some visual animations from the [@project_jdm channel](https://www.youtube.com/@project_jdm) to understand how they were created, but then the end result was interesting enough that they published their own [1hr Relaxing Polyrhythmic Spirals](https://www.youtube.com/watch?v=gvGfgCcfAEg) video which stands on its own as a nice piece of audiovisual art.

I love this idea of building a toy demo simply as an act to learn how something works, then improving on it with your new found understanding to the point where it takes on a useful life of its own.
