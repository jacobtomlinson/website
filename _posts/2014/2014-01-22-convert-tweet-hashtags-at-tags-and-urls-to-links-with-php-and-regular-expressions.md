---
title: Convert tweet hashtags, at-tags and urls to links with PHP and Regular Expressions
author: Jacob Tomlinson
layout: post
permalink: /2014/01/22/convert-tweet-hashtags-at-tags-and-urls-to-links-with-php-and-regular-expressions/
thumbnail: twitter
category: Web Development
tags:
  - Guide
  - PHP
  - Regular Expressions
  - Tutorial
  - Tweet
  - Twitter API
---
Now of course if you're using the <a title="Twitter API" href="https://dev.twitter.com/" target="_blank">Twitter API</a> you can use <a title="Twitter Entities" href="https://dev.twitter.com/docs/entities" target="_blank">Twitter entities</a> but in this tutorial we're going to use <a title="Regular Expressions" href="http://en.wikipedia.org/wiki/Regular_expression" target="_blank">regular expressions</a>.

### What are Regular Expressions?

So you may be wondering what regular expressions or regex are. Basically they are a very powerful search and replace feature implemented in almost every modern programming language for manipulating strings. Now we're not going to get into how regular expressions work here but if you want to know more then see this<a title="Regular Expressions Learning Game" href="http://regexone.com/" target="_blank"> useful tool</a>.

### In PHP

So to make use of regular expressions in PHP you will use the function <a title="preg_replace() docs" href="http://uk3.php.net/preg_replace" target="_blank">preg_replace()</a>, this function must be given 3 parameters; a search string written in regex, a replacement string written in regex and then the string you are manipulating.

So what we're going to do is to search for hashtags, at-tags and urls in turn and replace each with an HTML anchor tag pointing to the correct url.

```php
$tweet = "@george check out http://www.google.co.uk #google";

//Convert urls to <a> links
$tweet = preg_replace("/([\w]+\:\/\/[\w-?&;#~=\.\/\@]+[\w\/])/", "<a target=\"_blank\" href=\"$1\">$1</a>", $tweet);

//Convert hashtags to twitter searches in <a> links
$tweet = preg_replace("/#([A-Za-z0-9\/\.]*)/", "<a target=\"_new\" href=\"http://twitter.com/search?q=$1\">#$1</a>", $tweet);

//Convert attags to twitter profiles in &lt;a&gt; links
$tweet = preg_replace("/@([A-Za-z0-9\/\.]*)/", "<a href=\"http://www.twitter.com/$1\">@$1</a>", $tweet);

echo $tweet;
```

Which gives the output

Now this is all well and good but you don't want to be implementing this code every time you want to "linkify" a tweet so lets wrap it up in a function which you can put at the top of your code or in an included module.

```php
function linkify_tweet($tweet) {

  //Convert urls to <a> links
  $tweet = preg_replace("/([\w]+\:\/\/[\w-?&;#~=\.\/\@]+[\w\/])/", "<a target=\"_blank\" href=\"$1\">$1</a>", $tweet);

  //Convert hashtags to twitter searches in <a> links
  $tweet = preg_replace("/#([A-Za-z0-9\/\.]*)/", "<a target=\"_new\" href=\"http://twitter.com/search?q=$1\">#$1</a>", $tweet);

  //Convert attags to twitter profiles in <a> links
  $tweet = preg_replace("/@([A-Za-z0-9\/\.]*)/", "<a href=\"http://www.twitter.com/$1\">@$1</a>", $tweet);

  return $tweet;

}
```

So now we can simply &#8220;linkify&#8221; our tweets with the following code.

```php
$tweet = "@george check out http://www.google.co.uk #google";

echo linkify_tweet($tweet);
```
