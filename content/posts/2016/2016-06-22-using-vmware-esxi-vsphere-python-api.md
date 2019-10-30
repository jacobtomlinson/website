---
author: Jacob Tomlinson
date: 2016-06-22T00:00:00+00:00
categories:
  - VMware
tags:
  - vmware
  - python
  - api
thumbnail: vmware
title: Getting started with VMwares ESXi/vSphere API in Python
aliases:
  - /vmware/2016/06/22/using-vmware-esxi-vsphere-python-api/
---


In 2013 VMware dropped their Python library for accessing the API for ESXi/vSphere on GitHub. This is great, however it isn't the easiest library in the world to use. This quick guide will show you how to connect to an ESXi host or vSphere cluster and get some info about a virtual machine.

## Installation

First you must install the [pyvmomi](https://github.com/vmware/pyvmomi/) library. You can either grab the newest release from pypi or the latest source straight from GitHub using [pip](https://pypi.python.org/pypi/pip).

```shell
# Install newest version from pypi
pip install pyvmomi

# Install latest source from GitHub
pip install git+https://github.com/vmware/pyvmomi.git
```

Now in Python you will be able to import the library.

```python
import pyVim
```

## Connecting to a host or cluster

To connect to your host or cluster you must identify the API end point, this will usually be port 443 on the ip of your host or vSphere appliance.

```python
from pyVim import connect

my_cluster = connect.Connect("192.168.2.171", 443, "username", "password")
```

If you are running your host without SSL configured correctly, in a home lab for example, you will get the error `[SSL: CERTIFICATE_VERIFY_FAILED]`. To work around this ensure you have installed the latest source from GitHub, or a version greater than `v6.0.0.2016.4`, and use the `connect.ConnectNoSSL()` method instead.

You should also disconnect your session when you are done with it.

```python
connect.Disconnect(my_cluster)
```

## Querying a VM

Once you are connected to your cluster you will be able to find your VMs, get their information, issue commands and more. It is a very powerful api so be careful.

To get a VM object or a list of objects you can use the [searchIndex](https://github.com/vmware/pyvmomi/blob/master/docs/vim/SearchIndex.rst) class. The class had methods to search for VMs by UUID, DNS name, IP address or datastore path.

```python
from pyVim import connect

# Connect to a cluster or host
my_cluster = connect.Connect("192.168.2.171", 443, "username", "password")

# Get a searchIndex object
searcher = my_cluster.content.searchIndex

# Find a VM
vm = searcher.FindByIp(ip="192.168.70.54", vmSearch=True)

# Print out vm name
print vm.config.name

# Disconnect from cluster or host
connect.Disconnect(my_cluster)
```

In the above example we connected to a host or cluster, found a VM and printed out it's name. We could easily have started, stopped or even destroyed it.

## Further reading

Once you have run through this guide I'm sure you'll be itching to explore the power of this API. You can find a full list of all the objects included in the library [here](https://github.com/vmware/pyvmomi/tree/master/docs).
