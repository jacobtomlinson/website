---
title: 'Google Charts IE7 IE8 Issue: Date formatting problem'
author: Jacob Tomlinson
layout: post
permalink: /2013/11/28/google-charts-ie7-ie8-issue-date-formatting-problem/
category: Web Development
thumbnail: google-developers
tags:
  - Google Charts
  - IE7
  - IE8
  - JavaScript
---
Just a quick post about an issue I&#8217;ve had with Google Charts on IE7/8.

When viewing my page in Firefox or Chrome my graph displayed as expected.

![Correct graph](http://i.imgur.com/KWqjM9b.png)

However when I tried to view it in IE8 or IE7 the date axis just showed 1st of Jan 1970 and no chart.

![Broken graph](http://i.imgur.com/xiMYNlw.png)

After a bit of digging it appears that this is to do with the JavaScript function I used to convert a string to a date object. All of my dates are stored in a json file with the format YYYY-MM-DD along with an integer value, to convert the date string into a JS date object I used the following code

```javascript
var data = new google.visualization.DataTable();
var dataArray = [];

$.each(json.item, function(i, item){
  dataArray.push([new Date(item.Date), parseInt(item.Value)]);
});

data.addRows(dataArray);
```

the important part of this being the `new Date(item.Date)` which just reads in the date string and converts it to a Date object.

However IE8/7 seems to be unable to parse a string in the format YYYY-MM-DD and must have it in the format YYYY/MM/DD. I attempted this by adding a .replace(&#8220;-&#8221;,&#8221;/&#8221;) to the end of the variable but then it caused it to stop displaying in Firefox.

To counter this I made use of some browser detection I had implemented at the top of the page

```html
<!--[if lt IE 9]> <html lang="en" class="legacyie"> <![endif]-->
<!--[if gt IE 8]><!--> <html lang="en" class=""> <!--<![endif]-->
```

This basically adds a class to the html container to notify that the version of IE is older than version 9 (although there is no chance this will work in IE6 but that doesn&#8217;t matter as we&#8217;re only interested in 7 and 8).

Then I modified my JavaScript above to look like this

```javascript
var data = new google.visualization.DataTable();
var dataArray = [];

$.each(json.item, function(i, item){

  // Create temporary variable to store the date
  parsedDate = item.Date;

  // If the 'legacyie' class has been set perform the string replace
  if( $("html").hasClass("legacyie") ) { parsedDate = parsedDate.replace("-", "/"); };

  dataArray.push([new Date(parsedDate), parseInt(item.Value)]);
});

data.addRows(dataArray);
```

Now my graph displays as it always has in Firefox/Chrome but also not works in IE7/8

![Fixed Graph](http://i.imgur.com/6KSxIwn.png)
