---
title: Why is there no space in the MySQL password parameter?
author: Jacob Tomlinson
layout: post
permalink: /2013/05/31/why-is-there-no-space-in-the-mysql-password-parameter/
has_been_twittered:
  - yes
categories:
  - Linux
tags:
  - CLI
  - Command Line
  - Database
  - MySQL
  - terminal
---
After troubleshooting a MySQL issue with a colleague we began discussing a &#8220;feature&#8221; of the MySQL command line which insists that you don&#8217;t put a space in the password parameter when using the short parameter. We both felt that it was rather inconsistent to allow the usage of *-h hostname* or -u username but insist on *-ppassword* instead of *-p password*. You can of course use the full parameter *&#8211;password=password* but as most people use the shorthand commands it just seems slightly unintuitive. After doing a bit of reading it appears that this is due to the value being optional. If you don&#8217;t specify a password in the command it will prompt you for one, therefore if you had the option to include a space it wouldn&#8217;t be able to tell if you&#8217;ve specified a password or the next parameter. This still seems a little kludgy to me but I guess there is some reasoning behind it.

Here is the official reasoning from the MySQL website.

> For a long option that takes a value, separate the option name and the value by an “*=*” sign. For a short option that takes a value, the option value can immediately follow the option letter, or there can be a space between:* -hlocalhost* and *-h localhost* are equivalent. An exception to this rule is the option for specifying your MySQL password. This option can be given in long form as *[&#8211;password=pass_val][1]* or as *[&#8211;password][1]*. In the latter case (with no password value given), the program prompts you for the password. The password option also may be given in short form as *-ppass_val* or as *-p*. However, for the short form, if the password value is given, it must follow the option letter with *no intervening space*. The reason for this is that if a space follows the option letter, the program has no way to tell whether a following argument is supposed to be the password value or some other kind of argument.
> 
> &nbsp;
> 
> *Source - *<http://dev.mysql.com/doc/refman/5.5/en/command-line-options.html>

 [1]: http://dev.mysql.com/doc/refman/5.5/en/connecting.html#option_general_password