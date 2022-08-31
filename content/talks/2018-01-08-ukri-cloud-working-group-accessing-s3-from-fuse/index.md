---
title: "Accessing S3 from FUSE"
date: 2018-01-08T00:00:00+00:00
draft: false
author: "Jacob Tomlinson"
event:
  name: UKRI Cloud Working Group
  link: https://cloud.ac.uk/workshops/jan2018/
  location: London, UK
  type: Talk
length: null
abstract: true
video: null
slides: null
---

My data lives in an object store, but my tools expect a POSIX file path, what do I do? In the Informatics Lab we have been experimenting with using FUSE filesystems on a distributed compute cluster to provide parallel access to files on AWS S3. We have also created a library called pysssix which provides a slim and intentionally minimal access to S3 for fast data access.