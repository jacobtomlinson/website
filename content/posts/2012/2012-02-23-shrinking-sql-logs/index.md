---
author: Jacob Tomlinson
date: 2012-02-23T00:00:00+00:00
categories:
  - Web
  - Development
tags:
  - Databases
  - SQL
  - Storage
thumbnail: database
title: Shrinking SQL logs
aliases:
  - /2012/02/23/shrinking-sql-logs/
---


While reading some articles about minimizing the size of data while leaving the data fully searchable I came across a lot of info about finding common data and replacing it with pointers to a dictionary or store of common strings. After thinking about it for a while I thought I would just write a little post about a way I try and minimize the size of databases in projects that I work on. Now I&#8217;m going to start with an example of a very bad database I wrote a while back and how I used relational tables to shrink the size considerably. Now if you know anything about databases you should find all of this quite obvious but if you don&#8217;t then it may be worth taking note of.

While working on a blog based system a while back the client wanted me to record every page view and use the information to work out whether people are returning to the site or just viewing it once. I did suggest google analytics but they wanted me to implement my own thing (as they say the client is always right). So what i did was to create a table which recorded their IP address, their User Agent, a timestamp and the URI of the page they visited. I did this really simply in a table with 5 fields (one for ID plus the 4 pieces of data) and then wrote a simple PHP page to spew that information out in  a more readable way. The client was happy with this and I left it at that.

However a few thousand hits later I noticed that the table had grown a fair amount quite quickly as I hadn&#8217;t optimised it due to getting distracted with another section of the site. Now after having a look at the data I noticed a fair amount of it was being repeated. People were viewing more than one page so IPs and User Agents were repeated over and over and also different people were using the same browser/OS configuration and so had the same User Agents. This meant it would be quite easy to just replace the data with pointers to tables with shared information and due to foreign keys being smaller than the data itself would shrink the size of the database.

To do this I just created 3 extra tables for IP, User Agent and URI. There was no point in making one for the timestamps as it would be quite rare that lots of people would look at a post at exactly the same time but I guess if this was applied to a much larger system there would be more chance of that happening so it may be worth doing. These 3 tables just contained an ID field and a value field and then the existing fields in the stats table could just be replaced with the ID of the value required. Now you may think that this is just a lot of shuffling around for not much benefit but by doing this I managed to reduce the total data size by ~80%.

Now if you think in the original table that you would need VARCHAR and TEXT fields to store the data which can be considerably larger than key fields then it begins to shrink down. Lets assume that the client may write 100 blog posts each of which may averagely be viewed 100 times, some more some less but 100 seems a good average for the client I was working with. That would give you 10,000 stat rows all together. Using a bit of maths and assuming the average IP is 11 bytes and the average User Agent is 100 bytes plus the other fields and headers ends up with that table being around 1.3MB.

Now if you assume that each visitor may view 10% of the pages, obviously some may read them all and some may just read one, and if you also estimate that for every 100 views 5% of them share the same User Agent, this will show you that for every 10,000 stats there would only be 1000 unique IPs and 500 unique User Agents (I appreciate depending on the type of website your on this may be widly inaccurate but for the project I was on this is pretty close to the truth). So if you have replaced 10,000 User Agent fields which are 100 bytes (totalling ~975KB) with 10,000 4 byte pointers  and a table containing 500 User Agents which are 100 bytes and 500 IDs which are 4 bytes (totalling ~92KB) you can quickly see how the size is shrinking. Now if you apply this to the IP and URI fields too you end up shrinking this 1.3MB table down to 0.25MB. Now these are only small values I know and with disk space being so cheap these days you may find youself thinking &#8220;why does saving 1MB matter?&#8221; but if you think that on a server you may end up with lots of these projects which are doing far better than this one getting far more views that can end up ramping up the usage and by storing the data this way you will save hundreds of megabytes and reduce the frequency with which you have to archive out portions of your database. Plus it&#8217;s always good practice to try and keep memory usage to a minimum on any system.
