---
title: "Kubeflow and Kubernetes"
date: 2022-09-09T00:00:00+00:00
draft: false
author: "Jacob Tomlinson"
---

Happy Friday!

_Was this forwarded to you? [You can subscribe here!](https://jacobtomlinson.dev/newsletter)_

## 1. Blog: Accelerating ETL on KubeFlow with RAPIDS

[![Screenshot of JupyterLab with the Dask Lab extension open on the left and various Dask plots arranged on the screen](accelerating-etl-featured.png)](https://developer.nvidia.com/blog/accelerating-etl-on-kubeflow-with-rapids/)

Last week a blog post I wrote about running RAPIDS on Kubeflow was [published on the NVIDIA Developer Blog](https://developer.nvidia.com/blog/accelerating-etl-on-kubeflow-with-rapids/).

This post was particularly exciting because it marks the culmination of months of engineering work. At the start of 2022 I set out to ensure that RAPIDS integrates seamlessly with KubeFlow and as a result ended up completely overhauling how Dask deployments work on Kubernetes.

It has taken many hours of engineering effort to even make this post possible, so a huge thank you to everyone both within RAPIDS and also in the Dask community who came together to make this happen.

## 2. Code snippet: Expanding environment variables in arbitrary strings Python

Sometimes I want to allow a user to configure a template for a string and then [populate it from various sources](https://github.com/dask/dask-kubernetes/blob/5718f1f7eaf1f5220dda6cc0dc4e7162a54751a3/dask_kubernetes/classic/kubecluster.py#L591-L593).
We may want to generate a uuid or substitute in environment variables.

I often do this by leveraging Python's [str.format()](https://docs.python.org/3/library/stdtypes.html#str.format).
While f-strings are my preferred way of templating strings these days the older format style is great for creating a template string in one place and populating it in another.

We can then populate user provided strings with a whole array of useful things like a generated uuid or even the whole set of environment variables. If a kwarg to format isn't used in the template it will quietly ignore it.

```python
import os
import uuid

# Imagine the user configured this string somewhere
user_string = "Hello {USER} your uuid is {uuid}"

# Now let's populate the user string with all sorts of useful optional variables
filled_string = user_string.format(uuid=str(uuid.uuid4())[:10], **os.environ)

print(filled_string)  # for me prints "Hello jacob your uuid is a27fcecd-1"
```

## 3.
## 4.
## 5.

---

That's all folks! I hope you had a good week and I'll see you next Friday.
