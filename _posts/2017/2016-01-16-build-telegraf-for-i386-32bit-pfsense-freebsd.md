---
title: 'Building Telegraf for 32bit FreeBSD'
author: Jacob Tomlinson
layout: post
category: Monitoring
thumbnail: influxdata
tags:
  - telegraf
  - freebsd
  - influxdata
  - 32bit
  - i386
  - docker
---

Set the version we want to use
```
export VERSION=1.1.2
```

Get the source, checkout the version we want and build for FreeBSD i386
```
go get github.com/influxdata/telegraf
cd /go/src/github.com/influxdata/telegraf
git checkout ${VERSION}
env GOOS=freebsd GOARCH=386 make
```

Change to the output volume
```
cd /output
```

Get the 64bit package which includes the log rotation, startup script, etc and extract it
```
wget https://dl.influxdata.com/telegraf/releases/telegraf-${VERSION}_freebsd_amd64.tar.gz
tar -xvzf telegraf-${VERSION}_freebsd_amd64.tar.gz
rm telegraf-${VERSION}_freebsd_amd64.tar.gz
```

Replace the 64bit binary with the 32bit one and archive it again
```
mv /go/bin/freebsd_386/telegraf telegraf/usr/bin/telegraf
tar -cvzf telegraf-${VERSION}_freebsd_i386.tar.gz telegraf
rm -rf telegraf
```
