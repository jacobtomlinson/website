---
author: Jacob Tomlinson
title: Using multiple config files with kubectl and other Kubernetes tools
date: 2025-02-27T00:00:00+00:00
draft: false
categories:
  - blog
tags:
  - quick tips
  - tips
  - kubectl
  - kubernetes
  - kubeconfig
---

If you want to point tools like [`kubectl`](https://kubernetes.io/docs/reference/kubectl/) to a config file other than `~/.kube/config` you can set the [environment variable `KUBECONFIG`](https://kubernetes.io/docs/concepts/configuration/organize-cluster-access-kubeconfig/). But did you know that `KUBECONFIG` behaves sort of like a path, and `kubectl` will load all the config files it finds?

Let's say my `~/.kube/config` file has config for one cluster.

```console
$ kubectl config get-contexts
CURRENT   NAME                 CLUSTER              AUTHINFO             NAMESPACE
*         my-awesome-cluster   my-awesome-cluster   my-awesome-cluster
```

If I am provided with another config file for another cluster it might not be clear where to put it. I could just put it in my `~/.kube/` directory and call it `~/.kube/other.config` and point my `KUBECONFIG` to it.

```console
$ KUBECONFIG=~/.kube/other.config kubectl config get-contexts
CURRENT   NAME                 CLUSTER              AUTHINFO             NAMESPACE
*         some-other-cluster   some-other-cluster   some-other-cluster 
```

```info
I could [merge it with my existing config](https://jacobtomlinson.dev/posts/2019/how-to-merge-kubernetes-kubectl-config-files/). But what if this new config is going to change from time to time and I may need to redownload it occasionally. I don't want to have to pull apart my config or delete all the entries from it just to merge a new one in.
```

I can also update my `KUBECONFIG` environment variable to be a colon `:` separated list and then `kubectl` will load both.

```console
$ KUBECONFIG=~/.kube/config:~/.kube/other.config kubectl config get-contexts
CURRENT   NAME                 CLUSTER              AUTHINFO             NAMESPACE
*         my-awesome-cluster   my-awesome-cluster   my-awesome-cluster   
          some-other-cluster   some-other-cluster   some-other-cluster
```

If you have lots of config files that you use regularly you could update your `.bashrc`/`.zshrc` to do this all the time.

```bash
# .bashrc
# ...

# Explicitly set the default location
export KUBECONFIG="$HOME/.kube/config"

# Add more configs in a way that feels like PATH=/foo/bin:$PATH
export KUBECONFIG="$KUBECONFIG:$HOME/.kube/other.config"
```

```info
When specifying multiple configs any write operations like `kubectl config use-context foo` will be applied to the first config in the list. So it's best that this stays set to `~/.kube/config` unless you explicitly want to modify a different config.
```

For more information see the [Kubernetes documentation](https://kubernetes.io/docs/concepts/configuration/organize-cluster-access-kubeconfig/).
