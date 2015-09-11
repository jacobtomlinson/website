---
title: 'Amazon S3: s3cmd put ([Errno 32] Broken pipe)'
author: Jacob Tomlinson
layout: post
permalink: /2014/07/31/amazon-s3-s3cmd-put-errno-32-broken-pipe/
category: Linux
thumbnail: aws
tags:
  - aws
  - error
  - linux
  - s3
  - terminal
---
Recently I decided to use Amazon&#8217;s S3 as another location to store some of my server backups. However I found when testing that I was unable to upload my backup tarballs to S3. I ended up with the following errors.

```
jacob@server:backups$ s3cmd put backup-20140731.tar.gz s3://test-bucket/backup-20140731.tar.gz
backup-20140731.tar.gz -&gt; s3://test-bucket/backup-20140731.tar.gz  [1 of 1]
    253952 of 250487870     0% in    1s   159.37 kB/s  failed
WARNING: Upload failed: /backup-20140731.tar.gz ([Errno 104] Connection reset by peer)
WARNING: Retrying on lower speed (throttle=0.00)
WARNING: Waiting 3 sec...
backup-20140731.tar.gz -&gt; s3://test-bucket/backup-20140731.tar.gz  [1 of 1]
    569344 of 250487870     0% in    1s   357.37 kB/s  failed
WARNING: Upload failed: /backup-20140731.tar.gz ([Errno 104] Connection reset by peer)
WARNING: Retrying on lower speed (throttle=0.01)
WARNING: Waiting 6 sec...
backup-20140731.tar.gz -&gt; s3://test-bucket/backup-20140731.tar.gz  [1 of 1]
    405504 of 250487870     0% in    1s   282.39 kB/s  failed
WARNING: Upload failed: /backup-20140731.tar.gz ([Errno 104] Connection reset by peer)
WARNING: Retrying on lower speed (throttle=0.05)
WARNING: Waiting 9 sec...
backup-20140731.tar.gz -&gt; s3://test-bucket/backup-20140731.tar.gz  [1 of 1]
     98304 of 250487870     0% in    1s    64.06 kB/s  failed
WARNING: Upload failed: /backup-20140731.tar.gz ([Errno 32] Broken pipe)
WARNING: Retrying on lower speed (throttle=0.25)
WARNING: Waiting 12 sec...
backup-20140731.tar.gz -&gt; s3://test-bucket/backup-20140731.tar.gz  [1 of 1]
     28672 of 250487870     0% in    1s    14.69 kB/s  failed
WARNING: Upload failed: /backup-20140731.tar.gz ([Errno 32] Broken pipe)
WARNING: Retrying on lower speed (throttle=1.25)
WARNING: Waiting 15 sec...
backup-20140731.tar.gz -&gt; s3://test-bucket/backup-20140731.tar.gz  [1 of 1]
     16384 of 250487870     0% in    4s     3.84 kB/s  failed
ERROR: Upload of 'backup-20140731.tar.gz' failed too many times. Skipping that file.
```

The file I&#8217;m uploading isn&#8217;t huge (238MB). So it wasn&#8217;t the known limit of trying to <a title="Amazon S3 docs - 5GB put limit" href="http://docs.aws.amazon.com/AmazonS3/latest/dev/UploadingObjects.html" target="_blank">upload files which are greater than 5GB</a>.

```
jacob@server:backups$ ll
total 238M
-rw------- 1 jacob 238M Jul 31 06:00 backup-20140731.tar.gz
```

It seems that this is a common issue when uploading files to a new bucket. Something to do with the DNS propagation means it takes a little while before your bucket can successfully accept files. More info <a title="Google Code thread mentioning DNS issues" href="https://code.google.com/p/s3ql/issues/detail?id=363#c13" target="_blank">here</a>.

After waiting a few hours running the same command completes successfully.
