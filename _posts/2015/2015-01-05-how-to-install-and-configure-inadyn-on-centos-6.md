---
title: How to install and configure inadyn on CentOS 6
author: Jacob Tomlinson
layout: post
permalink: /2015/01/05/how-to-install-and-configure-inadyn-on-centos-6/
category: Centos
thumbnail: centos
tags:
- centos
- dyndns
- inadyn
- chkconfig
excerpt: "Inadyn is a command line utility for periodically checking and updating your ip
address with DynDNS."
---

### Introduction

Inadyn is a command line utility for periodically checking and updating your ip
address with DynDNS.

The guide will walk you through installing and configuring inadyn as a service
on Centos 6.

All commands in this guide will either need to be run as the root user or with
sudo.

### Requirements

If you're running CentOS x86_64 you'll need to install the i686 version of glibc
for inadyn to work.

```
yum install glibc.i686
```

## Install Binary

First we need to download and install the binary. DynDNS makes the utility, along
with the source code, available as a zip file which you can download from their
website.

Let's download and extract this zip file into /tmp.

```
cd /tmp
wget http://cdn.dyndns.com/inadyn.zip
unzip inadyn.zip
```

Next we'll copy the binary into /usr/bin and make it executable.

```
cp inadyn/bin/linux/inadyn /usr/bin/inadyn
chmod 755 /usr/bin/inadyn
```

## Create config file

When running inadyn you can specify your options on the command line but as we
will be running it as a service we're going to put all of our options in a
config file.

Inadyn will check this config file automatically if you run it without options
so we'll open it for editing.

```
vi /etc/inadyn.conf
```

Enter the following options and update your username, password and alias.

```
# Basic configuration file for inadyn
#
# /etc/inadyn.conf
update_period_sec 600 # Check for a new IP every 600 seconds
username <username>
password <password>
dyndns_system dyndns@dyndns.org
alias <your dyndns domain>
background
```

As we've stored our password in plain text let's ensure that only those who need
to see it can.

```
chmod 640 /etc/inadyn.conf
```

## Create init script

Next we'll create an init script for inadyn so it can start automatically on
boot.

```
vi /etc/init.d/inadyn
```

This is a modified version of the script provided by trendyserial [here][1].

I've updated it to make it chkconfig friendly for CentOS and also added a status
option so we can check if it is running. Copy it into your file.

```
#!/bin/bash
#
# inadyn				Startup script for the DynDNS update service
#
# chkconfig: 5 85 15
# description: inadyn updates DynDNS with the current ip address	\
#	             of the server.
# processname: inadyn
# config: /etc/inadyn.conf
# pidfile: /tmp/inadyn.pid
#
case "$1" in
  start)
    if [ -f /tmp/inadyn.pid ]; then
      PID=$(cat /tmp/inadyn.pid)
      kill -0 ${PID} &>/dev/null
      if [ $? = 0 ]; then
        echo "inadyn is already running."
      else
        /usr/bin/inadyn
        pidof inadyn > /tmp/inadyn.pid
        PID=$(cat /tmp/inadyn.pid)
        kill -0 ${PID} &>/dev/null
        if [ $? = 0 ]; then
          echo "inadyn started succesfully."
        else
          echo "Error starting inadyn"
        fi
      fi
    else
      /usr/bin/inadyn
      pidof inadyn > /tmp/inadyn.pid
      PID=$(cat /tmp/inadyn.pid)
      kill -0 ${PID} &>/dev/null
      if [ $? = 0 ]; then
        echo "inadyn started succesfully."
      else
        echo "Error starting inadyn"
      fi
    fi
  ;;

  stop)
    if [ -f /tmp/inadyn.pid ];then
      PID=$(cat /tmp/inadyn.pid)
      kill -0 ${PID} &>/dev/null
      if [ $? = 0 ]; then
        /bin/kill ${PID}
        kill -0 ${PID} &>/dev/null
        if [ $? = 1 ]; then
          echo "inadyn stopped succesfully."
        else
          echo "Error stopping inadyn"
        fi
      else
        echo "inadyn is already stopped."
      fi
    else
      echo "inadyn is already stopped."
    fi
  ;;

  status)
    PID=$(cat /tmp/inadyn.pid)
    if ps -p $PID > /dev/null
    then
      echo "inadyn is running"
    else
      echo "inadyn is not running"
    fi
  ;;

  reload|restart)
    $0 stop
    $0 start
  ;;

  *)
    echo "Usage: $0 start|stop|restart|reload"
    exit 1
esac
exit 0
```

We also need to make the script executable so chkconfig can run it.

```
chmod +x /etc/init.d/inadyn
```

## Add to chkconfig

We need to make chkconfig aware of the script and turn it on so that it
starts on boot.

```
chkconfig --add inadyn
chkconfig inadyn on
```

We can check to make sure inadyn has been added to chkconfig

```
chkconfig --list inadyn
inadyn         	0:off	1:off	2:on	3:on	4:on	5:on	6:off
```

## Start service

Then we need to start the service for the first time.

```
service inadyn start
```

You should get some output which looks like this

```
inadyn started succesfully.
```

If you want to see what inadyn is doing you can see its log messages in
/var/log/messages.

If you check now you should see some messages similar to these

```
Jan  5 22:00:22 gateway INADYN[15973]: INADYN: Started 'INADYN version 1.96.2' - dynamic DNS updater.
Jan  5 22:00:22 gateway INADYN[15973]: I:INADYN: IP address for alias 'mydomain.dyndns.org' needs update to 'x.x.x.x'
Jan  5 22:00:23 gateway INADYN[15973]: I:INADYN: Alias 'mydomain.dyndns.org' to IP 'x.x.x.x' updated successful.
```

## Uninstallation

To remove inadyn from your system for any reason simply run the following commands

```
service inadyn stop
chkconfig --del inadyn
rm /etc/init.d/inadyn /etc/inadyn.conf /usr/bin/inadyn
```

[1]: http://www.linuxquestions.org/questions/linux-software-2/how-do-i-execute-inadyn-automatically-during-boot-541367/#post4518378
