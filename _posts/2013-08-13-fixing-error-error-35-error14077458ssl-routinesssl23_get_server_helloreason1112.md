---
title: 'Fixing &#8220;ERROR: Error 35: error:14077458:SSL routines:SSL23_GET_SERVER_HELLO: reason(1112)&#8221;'
author: Jacob Tomlinson
layout: post
permalink: /2013/08/13/fixing-error-error-35-error14077458ssl-routinesssl23_get_server_helloreason1112/
thumbnail: command-line
has_been_twittered:
  - failed
twitter_failure_code:
  - 410
twitter_failure_reason:
  -
category: Apple
tags:
  - apple
  - curl
  - error
  - https
  - mac
  - munki
  - os x
  - ssl
  - sslv3
  - tlsv1
---
So the other day I came across the following error when using the munki configuration tool for mac.

```
ERROR: Error 35: error:14077458:SSL routines:SSL23\_GET\_SERVER_HELLO:reason(1112)
```

After a bit of digging I found that this is not a munki specific issue but rather an error from the curl command.

Munki uses curl to download manifests (in the form of plist files) from a munki server. This error was happening because when curl tried to communicate with the server over https it was being rejected.

Whenever you have issues communicating with a server over https and you&#8217;ve configured the server you should always check to make sure that the name in the certificate matches the ServerName in the apache configuration. This was my first mistake so I updated my `ssl.conf` file in `conf.d` to include the fully qualified domain name.

Now the error itself is curl saying that it&#8217;s trying to say hello to the server and initiate a ghandshake but the server is rejecting the protocol its using. The default protocol setting for https in apache looks like this


```
#   SSL Protocol support:
# List the enable protocol levels with which clients will be able to
# connect.  Disable SSLv2 access by default:
SSLProtocol all -SSLv2
```


So this is basically saying allow all protocols except SSLv2 (this is because there are known flaws in SSLv2). Now if you have a look at what curl is doing in the man page it is actually using TLSv1 as its default protocol. To get this working you have 2 options.

1. If you are able to see where curl is being called (or if you&#8217;re calling curl yourself) then you can add the `-ssl` flag like so

2. If you're managing the server you can allow it to handshake using the TLSv1 protocol by modifying the SSLProtocol in your apache config (mine is found at `/etc/https/conf.d/ssl.conf`)




This tells the server to add support for TLSv1.
