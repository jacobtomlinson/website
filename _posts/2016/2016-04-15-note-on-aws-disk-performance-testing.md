---
title: 'A note on AWS disk performance testing'
author: Jacob Tomlinson
layout: post
category: AWS
thumbnail: aws
tags:
  - AWS
  - EBS
  - Performance
  - Disk
---

Here is an interetsing note on testing the disk performance of your AWS instances. Before you can acurately test the performance of your EBS disk you need to read all the sectors of the disk at least once.

### Why?

Imagine that I have created a 1TB EBS volume on AWS and filled it with sensitive user data. I no longer require this volume so I am going to delete it. 

Now imagine that you, a different user, has come along and requested a 50GB EBS volume from AWS. By chance you have been allocated a section of blocks which were within the range of my now deleted volume. This means you could potentially read some of my sensitive data from your raw disk.

Of course AWS don't want to allow this to happen to their customers so they erase the disk and remove the sensitive data. The interesting thing is that they don't do this when you remove the volume as that would impact the performance of the EBS service. Instead they erase each block when the new user attempts to read/write to it. 

This means that whenever you provision a new EBS volume there is a performance impact whenever you access a block for the first time, as AWS has to erase it before performing your action.

### Solution

Most of the time this isn't an issue for users. However if you are building an application which needs high performance disk access from the start, or you want to accurately test the disk performance of your instance, you should manually reset all the blocks yourself.

You can do this by reading your whole disk into `/dev/null` using the tool `dd`. 

```
sudo dd if=/dev/xvda of=/dev/null bs=16K
```

Just replace `/dev/xvda` with the disk you wish to reset. This action makes no changes to the disk itself, it just resets the underlying EBS storage.
