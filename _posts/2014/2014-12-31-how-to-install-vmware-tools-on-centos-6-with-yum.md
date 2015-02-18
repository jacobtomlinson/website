---
title: How to install VMware Tools on Centos 6 with yum
author: Jacob Tomlinson
layout: post
permalink: /2014/12/31/how-to-install-vmware-tools-on-centos-6-with-yum/
category: Centos
thumbnail: vmware
tags:
- centos 6
- VMware
- yum
---

### Introduction

Often when managing a large number of systems you want to manage all software
installs the same way. So when it comes to VMware Tools you may not want to
follow the [official instructions][1] but instead install using yum, especially
if you're automating a large number of headless systems.

VMware makes their tools available via a web hosted yum repository, however
this means you must also update the tools using yum. You server will show the
tools status in vCenter as "Running (3rd-party/Independant)" and you will not be
able to specify the "Update on boot" option in your vm settings.

## Add repository

First we must add VMware's repository. Create a new repository file with your
text editor (we'll use vi for now)

```
vi /etc/yum.repos.d/vmware.repo
```

and add the following content

```
[vmware-tools]
name=VMware Tools for Red Hat Enterprise Linux $releasever - $basearch
baseurl=http://packages.vmware.com/tools/esx/latest/rhel6/$basearch
enabled=1
gpgcheck=1
gpgkey=http://packages.vmware.com/tools/keys/VMWARE-PACKAGING-GPG-RSA-KEY.pub
```

## Install Package

Once you've done this you can simply install the headless version of the tools
(no X) with this simple command. You'll need to accept the installation of the
GPG key.

```
yum install vmware-tools-esx-nox
```

## Updating VMware Tools

From now on if you need to update VMware Tools you can simply run

```
yum update vmware-tools-esx-nox
```

or as part of a regular yum update.

## Removing VMware Tools

If you change your mind and want to remove VMware Tools or switch back to the
tar and perl scripts method of installing VMware Tools simply run the following
command to remove them.

```
yum remove vmware-tools-esx-nox
```

[1]: https://www.vmware.com/support/ws55/doc/ws_newguest_tools_linux.html#wp1127177
