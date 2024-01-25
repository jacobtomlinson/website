---
title: "Running Dask workloads on multiple cluster backends with zero code changes using dask-ctl"
date: 2024-01-25T00:00:00+00:00
draft: false
author: "Jacob Tomlinson"
categories:
  - blog
tags:
  - python
  - dask
  - dask-ctl
  - distributed computing
---

Sometimes you want to write some code using Dask which can then be run against multiple different cluster backends. For example for local testing you might want to use `LocalCLuster`, but in production use `KubeCluster`. Or perhaps you want to easily switch between an on premise HPC with `SLURMRunner` or the cloud with `Coiled`.

You could write this logic into your code via a series of `if` statements, but mixing cluster lifecycle code and application code is a poor separation of concerns.

Instead you should use `dask-ctl` to specify everything as a YAML config file. 

## Example problem

To get started let's create an example Dask script that maps a function over a data structure in parallel.

```python
from dask.distributed import LocalCluster, Client

cluster = LocalCluster()
client = Client(cluster)

def double(x):
    """Double an integer."""
    return x * 2

data = range(10)  # Start with a list of number 0-99
futures = client.map(double, data)  # Map our function over our data
results = client.gather(futures)  # Gather the results
print(results)
# [0, 2, 4, 6, 8, 10, 12, 14, 16, 18]
```

When we run this code we create a local Dask cluster, connect a client to it and then submit some work. 

However, if we want to use a different cluster backend in different deployment scenarios we have to modify our application to handle creating something other than `LocalCluster`, and we probably need to introduce some config system to choose which cluster to use. 

## Creating clusters with `dask-ctl`

To avoid this we can use the `dask-ctl.lifecycle` submodule to create our clusters for us.

```python {hl_lines=[1, 4]}
import dask_ctl.lifecycle
from dask.distributed import Client

cluster = dask_ctl.lifecycle.create_cluster(local_fallback=True)
client = Client(cluster)

def double(x):
    """Double an integer."""
    return x * 2

data = range(10)  # Start with a list of number 0-99
futures = client.map(double, data)  # Map our function over our data
results = client.gather(futures)  # Gather the results
print(results)
# [0, 2, 4, 6, 8, 10, 12, 14, 16, 18]
```

When we call `dask_ctl.create_cluster()` it looks for a cluster spec file located in the same directory named `dask-cluster.yaml`. 
You can override the default path to the cluster spec file by passing it as a keyword-argument.

```python
create_cluster("/path/to/my-cluster.yaml", local_fallback=True)
```

```info
This can also be set with the `DASK_CTL__CLUSTER_SPEC` environment variable.
```

Also by setting `local_fallback=True` if no file is found then it will create a `LocalCluster` using the default configuration.

This way our code behaves as expected when ran locally, but if we deployed it to a production environment we could also supply a `dask-cluster.yaml` file that describes the type of cluster we want to use. 

Here's an example config file that would launch a cluster using [dask-kubernetes](https://kubernetes.dask.org/).

```yaml
# dask-cluster.yaml
version: 1
module: "dask_kubernetes.operator"
class: "KubeCluster"
kwargs:
  name: "foo"
  namespace: "bar"
  image: "ghcr.io/dask/dask:latest"
  n_workers: 20
```

## Switching backends quickly

You may also want to keep multiple different backend configs in your project.

```text
├── cluster-a.yaml  # Dask cluster spec for Kubernetes
├── cluster-b.yaml  # Dask cluster spec for HPC
└── myscript.py
```

Then you can switch to them quickly with an environment variable.

```console
$ DASK_CTL__CLUSTER_SPEC="cluster-a.yaml" python myscript.py
Running on Kubernetes backend
$ DASK_CTL__CLUSTER_SPEC="cluster-b.yaml" python myscript.py
Running on HPC backend
```
