---
title: "How to set environment variables on your Dask workers"
date: 2022-05-05T00:00:00+00:00
draft: false
author: "Jacob Tomlinson"
categories:
  - blog
tags:
  - Python
  - Dask
  - Snippet
---

When working with Dask clusters you often need the remote worker environment to match you local environment. This generally means having the same packages and data available.

However sometimes you might have things like API keys set locally as environment variables. You might have AWS keys set locally which allows you to do things like read data from S3 with Pandas.

```python
>>> import pandas as pd
>>> df = pd.read_csv("s3://foo/bar.csv")
```

However when doing the same thing with Dask Dataframe the data loading from S3 will be done on the workers which will not have the same environment variables set.

```python
>>> import dask.dataframe as dd
>>> from dask.distributed import Client
>>> import os

>>> client = Client("...")  # Connect to your Dask cluster

>>> ddf = dd.read_csv("s3://foo/bar.csv")
>>> df = ddf.compute()
💥  # Will result in a permission error even though your keys are set locally
```

Some Dask cluster deployment tooling allows you to configure environment variables at cluster creation time.

```python
>>> from dask_cloudprovider.azure import AzureVMCluster
>>> from dask.distributed import Client

>>> cluster = AzureVMCluster(..., env_vars={"FOO": "BAR"})
>>> client = Client(cluster)
```

However this isn't always available in every Dask deployment method.

Instead you can use your `Client` to run an out of band function on all the workers to set the environment variables.

```python

>>> from dask.distributed import Client
>>> import os

>>> client = Client("...")  # Connect to your Dask cluster

>>> def set_env():
...    os.environ["FOO"] = "BAR"
...

>>> client.run(set_env)
{'tcp://127.0.0.1:51521': None, 'tcp://127.0.0.1:51522': None, 'tcp://127.0.0.1:51523': None, 'tcp://127.0.0.1:51524': None}
```

Now we can check that our environment variables were applied correctly by querying them the same way.

```python
>>> client.run(lambda: os.environ["FOO"])
{'tcp://127.0.0.1:51521': 'BAR', 'tcp://127.0.0.1:51522': 'BAR', 'tcp://127.0.0.1:51523': 'BAR', 'tcp://127.0.0.1:51524': 'BAR'}
```
