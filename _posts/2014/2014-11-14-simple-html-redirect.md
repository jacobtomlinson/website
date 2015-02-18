---
title: Simple HTML Redirect
author: Jacob Tomlinson
layout: post
permalink: /2014/11/14/simple-html-redirect/
category: Web Development
thumbnail: google-developers
tags:
  - html
  - web development
  - redirect
  - code snippet
---

I often find myself in need of a quick html redirect page. Most of the time I
use the [example from Stack Overflow][1] but it involves changing the url in 3 places.

I've decided to create my own example with a little PHP in it to make it simpler
to use.

```html
<?php
// Change this url
$url = "http://www.example.com";
?>
<!DOCTYPE HTML>
<html lang="en-US">
    <head>
        <meta charset="UTF-8">
        <meta http-equiv="refresh" content="1;url=<?php echo $url ?>">
        <script type="text/javascript">
            window.location.href = "<?php echo $url ?>"
        </script>
        <title>Page Redirection</title>
    </head>
    <body>
        <!-- Note: don't tell people to `click` the link, just tell them that it is a link. -->
        If you are not redirected automatically, follow the <a href='<?php echo $url ?>'>link</a>.
    </body>
</html>
```

or if you don't have access to PHP then just use the original one.

```html
<!DOCTYPE HTML>
<html lang="en-US">
    <head>
        <meta charset="UTF-8">
        <meta http-equiv="refresh" content="1;url=http://example.com">
        <script type="text/javascript">
            window.location.href = "http://example.com"
        </script>
        <title>Page Redirection</title>
    </head>
    <body>
        <!-- Note: don't tell people to `click` the link, just tell them that it is a link. -->
        If you are not redirected automatically, follow the <a href='http://example.com'>link</a>.
    </body>
</html>
```
[1]: http://stackoverflow.com/a/5411601/1003288
