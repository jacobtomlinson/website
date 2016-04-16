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

Here is an interesting note on testing the disk performance of your AWS instances. Before you can acurately test the performance of your EBS disk you need to read all the sectors of the disk at least once.

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

Just replace `/dev/xvda` with the disk you wish to reset. This action makes no changes to the data already on the disk as it is only reading. It just triggers all the resets to the unused blocks on the underlying EBS storage.

### Example

Here is an example of resetting an 8GB boot disk with 24 IOPS on a `t2.micro` instance.

```
$ sudo dd if=/dev/xvda of=/dev/null bs=16K
524288+0 records in
524288+0 records out
8589934592 bytes (8.6 GB) copied, 189.241 s, 45.4 MB/s
$ sudo dd if=/dev/xvda of=/dev/null bs=16K
524288+0 records in
524288+0 records out
8589934592 bytes (8.6 GB) copied, 133.22 s, 64.5 MB/s
$ sudo dd if=/dev/xvda of=/dev/null bs=16K
524288+0 records in
524288+0 records out
8589934592 bytes (8.6 GB) copied, 133.226 s, 64.5 MB/s
```

Note that the second and third reads of the disk are 42% faster than the initial read. This is because the first read triggered all the resets and the subsequent reads are now directly accessing the disk.
