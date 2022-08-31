---
title: "How to merge Kubernetes kubectl config files"
date: 2019-10-15T00:00:00+01:00
draft: false
author: "Jacob Tomlinson"
categories:
  - kubernetes
tags:
  - kubernetes
  - cli
  - tips
thumbnail: kubernetes
---

Sometimes when working with a new Kubernetes cluster you will be given a config file to use when
authenticating with the cluster. This file should be placed at `~/.kube/config`. However you may
already have an existing config file at that location and you need to merge them together.

Here is a quick command you can run to merge your two config files.

```console
# Make a copy of your existing config
$ cp ~/.kube/config ~/.kube/config.bak

# Merge the two config files together into a new config file
$ KUBECONFIG=~/.kube/config:/path/to/new/config kubectl config view --flatten > /tmp/config

# Replace your old config with the new merged config
$ mv /tmp/config ~/.kube/config

# (optional) Delete the backup once you confirm everything worked ok
$ rm ~/.kube/config.bak
```

Here is all of that (except the cleanup) as a one-liner.

```console
$ cp ~/.kube/config ~/.kube/config.bak && KUBECONFIG=~/.kube/config:/path/to/new/config kubectl config view --flatten > /tmp/config && mv /tmp/config ~/.kube/config
```