---
title: 'govspeak: An open source markup language'
date: 2015-07-22T00:00:00+00:00
draft: false
categories:
- blog
tags:
- collaboration
author: Jacob Tomlinson
canonical: http://archived.informaticslab.co.uk/collaboration/2015/07/22/govspeak-an-open-source-markup-language.html
thumbnail: jekyll
---

_Originally published on the [Met Office Informatics Lab blog](http://archived.informaticslab.co.uk/collaboration/2015/07/22/govspeak-an-open-source-markup-language.html) on July 22nd, 2015._

The Informatics Lab website is created with an application called [Jekyll][jekyll]. Recently I made an enhancement to it which I'm very excited about. It allows us to write our articles in a markup language called [Govspeak][govspeak], which is an extension to the excellent [markdown][markdown].

## What is markdown?

Markdown is a way of writing web content without having to write any code (HTML/CSS) by adding styling to your text with simple text symbols. It is commonly used for creating documentation for open source projects and for writing articles on blogs.

For example you can prefix a line with a hash and it will transform it into a heading, a double-hash for a sub-heading, a triple-hash for a sub-sub-heading, and so on.

```
# My heading
## My sub-heading
```

Or wrap some text with a asterisk or underscore and you make it _italic_, or a double-asterisk/underscore for **bold**.

```
Some words are _italic_ and some are **bold**.
```

You can find out more about markdown [here][markdown].

## What does Govspeak do that markdown doesn't?

Govspeak is an open source extended version of markdown created by the [Government Digital Service][gds] team who maintain [GOV.UK][govuk]. It was created to allow the content authors of GOV.UK to use some more advanced features such as [steps][govspeak-steps], [legislative lists][govspeak-leglists] and [abbreviations][govspeak-abbr]. These contructs are common in government texts.

## Why are we using Govspeak?

The original reason for using Govspeak was down to me wanting to use callouts in our [first Lab School article][lab-school]. Callouts are common on the web, they are just paragraphs or sentences which are highlighted with some kind of coloured border.  Callouts are not available in standard markdown and so I started looking for an alternative, which is when I found Govspeak.

Govspeak also has some useful features such as abbreviations. This means you can define a list of acronyms and their meaning at the end of an article. When the web page is displayed the acronym will have an underline and if you hover over it you will see a popup with the full text. This was added to Govspeak because acronyms are common in government text, but they are also very common in technical writing so it is very useful for us too.

## Contributing to Govspeak

I got very excited when I found Govspeak for two main reasons. The first is that it is open source, which means I can add enhancements to it and submit those contributions back to the author. The second is that it was developed by another government department. Working with people on a open source project always leaves you with a warm fuzzy feeling inside, but knowing you're working with distant colleagues is even nicer. It also gives credit to GDS's decision to open source the project as it has been extremely useful to us.

Sadly when I imported Govspeak into Jekyll there was a dependancy conflict which stopped me from being able to use it, but thanks to the nature of open source I fixed the issue, submitted a patch and improved Govspeak a little bit. So now anyone wanting to use it with Jekyll can do so.

## Conclusion

I wanted to write about this small change to our website because for me it highlights one of the main reasons why I like working on open source projects. Someone wrote a very useful piece of code and shared it with the world. I made a minor improvement to the code when using it for my own purposes, and now everyone else can benefit from that enhancement too.

[gds]: https://twitter.com/gdsteam
[govspeak]: https://github.com/alphagov/govspeak
[govspeak-abbr]: https://github.com/alphagov/govspeak#abbreviations
[govspeak-leglists]: https://github.com/alphagov/govspeak#legislative-lists
[govspeak-steps]: https://github.com/alphagov/govspeak#steps
[govuk]: http://gov.uk
[jekyll]: http://jekyllrb.com/
[lab-school]: http://www.informaticslab.co.uk/announcements/2015/06/22/introducing-lab-school.html
[markdown]: http://daringfireball.net/projects/markdown/
