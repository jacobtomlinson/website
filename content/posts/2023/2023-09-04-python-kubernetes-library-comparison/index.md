---
title: "Comparison of kr8s vs other Python libraries for Kubernetes"
date: 2023-09-04T00:00:00+00:00
draft: false
author: "Jacob Tomlinson"
categories:
  - blog
tags:
  - kubernetes
  - python
  - kr8s
  - lightkube
  - pykube-ng
  - kubernetes-asyncio
---

I've been [working on `kr8s`](https://jacobtomlinson.dev/posts/2023/introducing-kr8s-a-new-kubernetes-client-library-for-python-inspired-by-kubectl/) for a while now and one of my core goals is to build a Python library for Kubernetes that is the most simple, readable and produces the most maintainable code. It should enable folks to [write dumb code](https://matthewrocklin.com/blog/work/2018/01/27/write-dumb-code) when working with Kubernetes.

In this post I want to compare how `kr8s` is doing in achieving that goal by comparing it against the other Kubernetes libraries out there. This isn't intended to be a demonstration of why `kr8s` is "better" than the other libraries, I'm just trying to think out loud, [build in public](https://www.willmcgugan.com/blog/tech/post/promoting-your-open-source-project-or-how-to-get-your-first-1k-github-stars/#building-in-public) and take a step back and see whether `kr8s` is hitting it's target. 

I'm going to draw on code examples from the docs of all of the libraries I am comparing against along with examples from [`dask-kubernetes`](https://kubernetes.dask.org/en/latest/), which is the primary place I am using `kr8s` in my own work. I'll then rewrite each example with each library in a style that aims to satisfy the goals `kr8s` aspires to. 

I'll be subjectively comparing them by the following goals:

- Code should feel familiar to folks who know `kubectl`.
- Simple things should be simple.
- Boilerplate code should be kept to a minimum.
- Complex things such as using multiple API clients should be possible, but opt-in.

Here are some non-goals of `kr8s`, which may be strengths of the other libraries I am comparing against, but will not be considered when comparing them:

- Mimic the Kubernetes API.
- Resource schemas and models should be strictly enforced.

```warning
Its important to note that this is probably not a _fair_ comparison because the other libraries have different strengths and goals. If you are using this post to choose between them then it is down to you to exercise your own judgement on which strengths are important to you.
```

## The other libraries

Here's a table of the libraries I'll be comparing against and their self-described strengths.

| Name                  | Sync | Asyncio | Strengths                            |
| --------------------- | ---- | ------- | ----------------------------------- |
| [`kubernetes`](https://github.com/kubernetes-client/python)          |  ✅  |  ❌      | • Official library. <br/>• Exactly the same features / API objects in both client-python and the Kubernetes version due to auto generation. |
| [`kubernetes-asyncio`](https://github.com/tomplus/kubernetes_asyncio)  |  ❌  |  ✅      | • Semi-official asyncio version of `kubernetes`. |
| [`pykube-ng`](https://pykube.readthedocs.io/en/latest/)           |  ✅  |  ❌      | • Lightweight client. <br/>• HTTP interface using requests using kubeconfig for authentication. <br/>• Python native querying of Kubernetes API objects |
| [`lightkube`](https://lightkube.readthedocs.io/en/stable/)           |  ✅  |  ✅      | • Extensive type hints to avoid common mistakes and to support autocompletion. <br/>• Models and resources generated from the swagger specifications using standard dataclasses. <br/>• Support for installing a specific version of the kubernetes models (1.15 to 1.27). <br/>• Lazy instantiation of inner models. |

_Note that `kr8s` and `lightkube` support both sync and async usage, so I will include examples using both APIs._

```info
If you are the author of one of the libraries I am comparing against I want to take a moment to say thank you for all of your hard work. Your library has inspired much of the work here. 

My goal with `kr8s` is to simply add another option with a different emphasis. If you feel like I've misrepresented your work or if any of these comparisons are unfair or could be done better then please reach out to me so I can improve them.
```

## Setup

To run these comparisons I have done the following steps.

Create a new Kubernetes cluster with [`kind`](https://kind.sigs.k8s.io/).

```console
$ kind create cluster 
```

Create a new Python environment with [`conda`](https://github.com/conda-forge/miniforge).

```console
$ conda create -n k8s-comparison python=3.10 ipython
$ conda activate k8s-comparison
```

Install the packages.

```console
$ pip install kubernetes kubernetes-asyncio pykube-ng lightkube kr8s
```

## Comparisons

For each comparison I will take a code example from one of the libraries and rewrite it for each library. I'll aim to keep with the style the example was written in. I may make some modifications to the original example to ensure it can be copy and pasted into a python terminal to be reproduced or to just tweak it in the interest of fairness.

Let's get started!

### List nodes

**Goal:** Print out all of the node names in the cluster.

**Source:** [`lightkube` docs homepage example](https://lightkube.readthedocs.io/en/stable/).

To do this with `kubectl` we can get the nodes and output the names.

```console
$ kubectl get nodes -o name
```

{{< tabs >}}

{{< tab tabName="`kr8s`" >}}
```python
import kr8s

for node in kr8s.get("nodes"):
    print(node.name)
```
{{< /tab >}}

{{< tab tabName="`kubernetes`" >}}
```python
from kubernetes import client, config

config.load_kube_config()

v1 = client.CoreV1Api()
for node in v1.list_node().items:
    print(node.metadata.name)
```
{{< /tab >}}

{{< tab tabName="`pykube-ng`" >}}
```python
import pykube

api = pykube.HTTPClient(pykube.KubeConfig.from_file())
for node in pykube.Node.objects(api):
    print(node.name)
```
{{< /tab >}}

{{< tab tabName="`lightkube`">}}
```python
from lightkube import Client
from lightkube.resources.core_v1 import Node

client = Client()
for node in client.list(Node):
    print(node.metadata.name)
```
{{< /tab >}}
{{< /tabs >}}

{{< tabs >}}

{{< tab tabName="`kr8s (async)`" >}}
```python
import kr8s.asyncio

for node in await kr8s.asyncio.get("nodes"):
    print(node.name)
```

```info
Maybe this would feel better as an `async for`? I might change this in `kr8s`.
```
{{< /tab >}}

{{< tab tabName="`kubernetes-asyncio`" >}}
```python
from kubernetes_asyncio import client, config
from kubernetes_asyncio.client.api_client import ApiClient

await config.load_kube_config()

async with ApiClient() as api:
    v1 = client.CoreV1Api(api)
    nodes = await v1.list_node()
    for node in nodes.items:
        print(node.metadata.name)
```
{{< /tab >}}

{{< tab tabName="`lightkube (async)`">}}
```python
from lightkube import AsyncClient
from lightkube.resources.core_v1 import Node

client = AsyncClient()
async for node in client.list(Node):
    print(node.metadata.name)
```
{{< /tab >}}
{{< /tabs >}}

### List all Pods

**Goal:** List all Pods in all namespaces and print their IP, namespace and name.

**Source:** [kubernetes-asyncio README](https://github.com/tomplus/kubernetes_asyncio#example)

To do this with `kubectl` we can 

```console
$ kubectl get pods -A --no-headers -o custom-columns="IP:.status.podIP,NAMESPACE:.metadata.namespace,NAME:.metadata.name"

```

{{< tabs >}}

{{< tab tabName="`kr8s`" >}}
```python
import kr8s

for pod in kr8s.get("pods", namespace=kr8s.ALL):
    print(pod.status.podIP, pod.metadata.namespace, pod.metadata.name)
```
{{< /tab >}}

{{< tab tabName="`kubernetes`" >}}
```python
from kubernetes import client, config

config.load_kube_config()
api_client = client.CoreV1Api()

for pod in api_client.list_pod_for_all_namespaces().items:
    print(pod.status.pod_ip, pod.metadata.namespace, pod.metadata.name)
```
{{< /tab >}}

{{< tab tabName="`pykube-ng`" >}}
```python
import pykube

api = pykube.HTTPClient(pykube.KubeConfig.from_file())
for pod in pykube.Pod.objects(api).filter(namespace=pykube.all):
    print(pod.obj["status"]["podIP"], pod.namespace, pod.name)
```
{{< /tab >}}

{{< tab tabName="`lightkube`">}}
```python
from lightkube import Client, ALL_NS
from lightkube.resources.core_v1 import Pod

client = Client()
for i in client.list(Pod, namespace=ALL_NS):
    print(i.status.podIP, i.metadata.namespace, i.metadata.name)
```
{{< /tab >}}
{{< /tabs >}}

{{< tabs >}}

{{< tab tabName="`kr8s (async)`" >}}
```python
import kr8s

for pod in await kr8s.asyncio.get("pods", namespace=kr8s.ALL):
      print(pod.status.podIP, pod.metadata.namespace, pod.metadata.name)
```
{{< /tab >}}

{{< tab tabName="`kubernetes-asyncio`" >}}
```python
from kubernetes_asyncio import client, config
from kubernetes_asyncio.client.api_client import ApiClient

await config.load_kube_config()

async with ApiClient() as api:
    v1 = client.CoreV1Api(api)
    ret = await v1.list_pod_for_all_namespaces()
    for i in ret.items:
        print(i.status.pod_ip, i.metadata.namespace, i.metadata.name)
```

```info
I removed the event loop creation boilerplate to make this more readable.
```
{{< /tab >}}

{{< tab tabName="`lightkube (async)`">}}
```python
from lightkube import AsyncClient, ALL_NS
from lightkube.resources.core_v1 import Pod

client = AsyncClient()
async for i in client.list(Pod, namespace=ALL_NS):
    print(i.status.podIP, i.metadata.namespace, i.metadata.name)
```
{{< /tab >}}
{{< /tabs >}}

### Get ready Pods

**Goal:** Get a list of Pod resources that have the `Ready=True` condition.

**Source:** [`pykube-ng` docs README usage example](https://pykube.readthedocs.io/en/latest/#usage)

To do this with `kubectl` we can use `kubectl get pods` with a `jsonpath` output to filter on the ready condition.

```console
$ kubectl -n kube-system get pods -o jsonpath='{range .items[*]}{.status.containerStatuses[*].ready.true}{.metadata.name}{"\n"}{end}'
```

{{< tabs >}}

{{< tab tabName="`kr8s`" >}}
```python
import kr8s

for pod in kr8s.get("pods", namespace="kube-system"):
    if pod.ready():
        print(pod.name)
```
{{< /tab >}}

{{< tab tabName="`kubernetes`" >}}
```python
from kubernetes import client, config

config.load_kube_config()
api_client = client.CoreV1Api()

def check_ready(pod):
    ready = [c.status for c in pod.status.conditions if c.type == "Ready"]
    return ready and all(ready)

pods = api_client.list_namespaced_pod("kube-system").items
for pod in filter(check_ready, pods):
    print(pod.metadata.name)
```
{{< /tab >}}

{{< tab tabName="`pykube-ng`" >}}
```python
import pykube

api = pykube.HTTPClient(pykube.KubeConfig.from_file())
for pod in pykube.Pod.objects(api).filter(namespace="kube-system"):
    if pod.ready:
        print(pod.name)
```
{{< /tab >}}

{{< tab tabName="`lightkube`">}}
```python
from lightkube import Client
from lightkube.resources.core_v1 import Pod

def check_ready(pod):
    ready = [c.status for c in pod.status.conditions if c.type == "Ready"]
    return ready and all(ready)

client = Client()
for pod in client.list(Pod, namespace="kube-system"):
    if check_ready(pod):
        print(pod.metadata.name)
```
{{< /tab >}}
{{< /tabs >}}

{{< tabs >}}

{{< tab tabName="`kr8s (async)`" >}}
```python
import kr8s

for pod in await kr8s.asyncio.get("pods", namespace="kube-system"):
    if await pod.ready():
        print(pod.name)
```
{{< /tab >}}

{{< tab tabName="`kubernetes_asyncio`" >}}
```python
from kubernetes_asyncio import client, config
from kubernetes_asyncio.client.api_client import ApiClient

await config.load_kube_config()

def check_ready(pod):
    ready = [c.status for c in pod.status.conditions if c.type == "Ready"]
    return ready and all(ready)

async with ApiClient() as api:
    v1 = client.CoreV1Api(api) 
    pods = await v1.list_namespaced_pod("kube-system")
    for pod in filter(check_ready, pods.items):
        print(pod.metadata.name)
```
{{< /tab >}}

{{< tab tabName="`lightkube (async)`">}}
```python
from lightkube import AsyncClient
from lightkube.resources.core_v1 import Pod

def check_ready(pod):
    ready = [c.status for c in pod.status.conditions if c.type == "Ready"]
    return ready and all(ready)

client = AsyncClient()
async for pod in client.list(Pod, namespace="kube-system"):
    if check_ready(pod):
        print(pod.metadata.name)
```
{{< /tab >}}
{{< /tabs >}}

### Get pods by label selector

**Goal:** Starting from a dictionary containing a label selector get all Pods from all namespaces matching that label.

**Source:** [`pykube-ng` docs ](https://pykube.readthedocs.io/en/latest/howtos/find-pods-by-label.html).

To do this with `kubectl` we can use the `-A` flag for all namespaces and a label selector.

```console
$ kubectl get po -A -l "component=kube-scheduler"
```

{{< tabs >}}

{{< tab tabName="`kr8s`" >}}
```python
import kr8s

selector = {'component': 'kube-scheduler'}

for pod in kr8s.get("pods", namespace=kr8s.ALL, label_selector=selector):
    print(pod.namespace, pod.name)
```
{{< /tab >}}

{{< tab tabName="`kubernetes`" >}}
```python
from kubernetes import client, config

selector = {'component': 'kube-scheduler'}
selector_str = ",".join([f"{key}={value}" for key, value in selector.items()])

config.load_kube_config()

v1 = client.CoreV1Api()
for pods in v1.list_pod_for_all_namespaces(label_selector=selector_str, ).items:
    print(pod.metadata.namespace, pod.metadata.name)
```
{{< /tab >}}

{{< tab tabName="`pykube-ng`" >}}
```python
import pykube

selector = {'component': 'kube-scheduler'}

api = pykube.HTTPClient(pykube.KubeConfig.from_file())
for pod in pykube.Pod.objects(api).filter(namespace=pykube.all, selector=selector):
    print(pod.namespace, pod.name)
```
{{< /tab >}}

{{< tab tabName="`lightkube`">}}
```python
from lightkube import Client, ALL_NS
from lightkube.resources.core_v1 import Pod

selector = {'component': 'kube-scheduler'}

client = Client()
for pod in client.list(Pod, labels=selector, namespace=ALL_NS):
    print(pod.metadata.namespace, pod.metadata.name)
```
{{< /tab >}}
{{< /tabs >}}

{{< tabs >}}

{{< tab tabName="`kr8s (async)`" >}}
```python
import kr8s

selector = {'component': 'kube-scheduler'}

for pod in await kr8s.asyncio.get("pods", namespace=kr8s.ALL, label_selector=selector):
    print(pod.namespace, pod.name)
```
{{< /tab >}}

{{< tab tabName="`kubernetes_asyncio`" >}}
```python
from kubernetes_asyncio import client, config
from kubernetes_asyncio.client.api_client import ApiClient

selector = {'component': 'kube-scheduler'}
selector_str = ",".join([f"{key}={value}" for key, value in selector.items()])

await config.load_kube_config()

async with ApiClient() as api:
    v1 = client.CoreV1Api(api)
    pods = await v1.list_pod_for_all_namespaces(label_selector=selector_str)
    for pod in pods.items:
        print(pod.metadata.namespace, pod.metadata.name)
```
{{< /tab >}}

{{< tab tabName="`lightkube (async)`">}}
```python
from lightkube import AsyncClient, ALL_NS
from lightkube.resources.core_v1 import Pod

selector = {'component': 'kube-scheduler'}

client = AsyncClient()
async for pod in client.list(Pod, labels=selector, namespace=ALL_NS):
    print(pod.metadata.namespace, pod.metadata.name)
```
{{< /tab >}}
{{< /tabs >}}

### Simple labelling operator

**Goal:** Write an operator controller that periodically reconciles all deployments and adds a label to any with a certain annotation.

**Source:** [`pykube-ng` docs on how to write an operator](https://pykube.readthedocs.io/en/latest/howtos/write-an-operator.html).

You probably wouldn't do this with `kubectl` but if you wanted to you would do something like running a loop and listing all deployments with the annotation, piping that into another `kubectl` command to add the label and then sleeping.

```console
$ while true; do kubectl get deploy -o=jsonpath='{.items[?(@.spec.template.metadata.annotations.pykube-test-operator)].metadata.name}' | xargs kubectl label deploy - foo=bar; sleep 15; done

```

{{< tabs >}}

{{< tab tabName="`kr8s`" >}}
```python
import kr8s, time

while True:
    for deploy in kr8s.get("deployments", namespace=kr8s.ALL):
        if 'pykube-test-operator' in deploy.annotations:
            deploy.label(foo="bar")
    time.sleep(15)
```
{{< /tab >}}

{{< tab tabName="`kubernetes`" >}}
```python
from kubernetes import client, config
import time

config.load_kube_config()
appsv1 = client.AppsV1Api()

while True:
    for deploy in appsv1.list_deployment_for_all_namespaces().items:
        if 'pykube-test-operator' in deploy.metadata.annotations:
            appsv1.patch_namespaced_deployment(
                name=deploy.metadata.name,
                namespace=deploy.metadata.namespace,
                body={"metadata": {"labels": {"foo": "bar"}}}
            )
    time.sleep(15)
```
{{< /tab >}}

{{< tab tabName="`pykube-ng`" >}}
```python
import pykube, time

api = pykube.HTTPClient(pykube.KubeConfig.from_env())

while True:
    for deploy in pykube.Deployment.objects(api, namespace=pykube.all):
        if 'pykube-test-operator' in deploy.annotations:
            deploy.labels['foo'] = 'bar'
            deploy.update()
    time.sleep(15)
```

```info
I tweaked the original and moved the api client creation outside of the loop. It felt unecessary to do this every time.
```
{{< /tab >}}

{{< tab tabName="`lightkube`">}}
```python
from lightkube import Client, ALL_NS
from lightkube.resources.apps_v1 import Deployment
import time

client = Client()

while True:
    for deploy in client.list(Deployment, namespace=ALL_NS):
        if 'pykube-test-operator' in deploy.metadata.annotations:
            client.patch(
                Deployment, 
                name=deploy.metadata.name,  
                namespace=deploy.metadata.namespace,
                obj={"metadata": {"labels": {"foo": "bar"}}}
            )
    time.sleep(15)
```
{{< /tab >}}
{{< /tabs >}}

{{< tabs >}}

{{< tab tabName="`kr8s (async)`" >}}
```python
import kr8s, asyncio

while True:
    for deploy in await kr8s.asyncio.get("deployments", namespace=kr8s.ALL):
        if 'pykube-test-operator' in deploy.annotations:
            await deploy.label(foo="bar")
    await asyncio.sleep(15)
```
{{< /tab >}}

{{< tab tabName="`kubernetes_asyncio`" >}}
```python
from kubernetes_asyncio import client, config
from kubernetes_asyncio.client.api_client import ApiClient
import asyncio

await config.load_kube_config()

async with ApiClient() as api:
    appsv1 = client.AppsV1Api(api)
    while True:
        deployments = await appsv1.list_deployment_for_all_namespaces()
        for deploy in deployments.items:
            if 'pykube-test-operator' in deploy.metadata.annotations:
                await appsv1.patch_namespaced_deployment(
                    name=deploy.metadata.name,
                    namespace=deploy.metadata.namespace,
                    body={"metadata": {"labels": {"foo": "bar"}}}
                )
        await asyncio.sleep(15)
```
{{< /tab >}}

{{< tab tabName="`lightkube (async)`">}}
```python
from lightkube import AsyncClient, ALL_NS
from lightkube.resources.apps_v1 import Deployment
import asyncio

client = AsyncClient()

while True:
    async for deploy in client.list(Deployment, namespace=ALL_NS):
        if 'pykube-test-operator' in deploy.metadata.annotations:
            await client.patch(
                Deployment, 
                name=deploy.metadata.name,  
                namespace=deploy.metadata.namespace,
                obj={"metadata": {"labels": {"foo": "bar"}}}
            )
    await asyncio.sleep(15)
```
{{< /tab >}}
{{< /tabs >}}

### Scale a deployment

**Goal:** Scale the deployment `metrics-server` in the namespace `kube-system` to `1` replica.

**Source:** [LightKube Quickstart documentation](https://lightkube.readthedocs.io/en/stable/)

To do this with `kubectl` we can do:

```console
$ kubectl scale -n kube-system deployment metrics-server --replicas=1
```

{{< tabs >}}

{{< tab tabName="`kr8s`" >}}
```python
from kr8s.objects import Deployment

deploy = Deployment.get("metrics-server", namespace="kube-system")
deploy.scale(1)
```
{{< /tab >}}

{{< tab tabName="`kubernetes`" >}}
```python
from kubernetes import client, config

config.load_kube_config()
appsv1 = client.AppsV1Api()

appsv1.patch_namespaced_deployment_scale(
    "metrics-server", "kube-system", {'spec': {'replicas': 1}}
)
```
{{< /tab >}}

{{< tab tabName="`pykube-ng`" >}}
```python
from pykube import Deployment, HTTPClient, KubeConfig

api = HTTPClient(KubeConfig.from_file())
deploy = Deployment.objects(api).get(name="metrics-server", namespace="kube-system")
deploy.scale(1)
```
{{< /tab >}}

{{< tab tabName="`lightkube`">}}
```python
from lightkube import Client
from lightkube.resources.apps_v1 import Deployment
from lightkube.models.meta_v1 import ObjectMeta
from lightkube.models.autoscaling_v1 import ScaleSpec

client = Client()
obj = Deployment.Scale(
    metadata=ObjectMeta(name='metrics-server', namespace='kube-system'),
    spec=ScaleSpec(replicas=1)
)
client.replace(obj, 'metrics-server', namespace='kube-system')
```
{{< /tab >}}
{{< /tabs >}}

{{< tabs >}}

{{< tab tabName="`kr8s (async)`" >}}
```python
from kr8s.asyncio.objects import Deployment

deploy = await Deployment.get("metrics-server", namespace="kube-system")
await deploy.scale(1)
```
{{< /tab >}}

{{< tab tabName="`kubernetes_asyncio`" >}}
```python
from kubernetes_asyncio import client, config
from kubernetes_asyncio.client.api_client import ApiClient

await config.load_kube_config()

async with ApiClient() as api:
    appsv1 = client.AppsV1Api(api)

    await appsv1.patch_namespaced_deployment_scale(
        "metrics-server", "kube-system", {'spec': {'replicas': 1}}
    )
```
{{< /tab >}}

{{< tab tabName="`lightkube (async)`">}}
```python
from lightkube import AsynClient
from lightkube.resources.apps_v1 import Deployment
from lightkube.models.meta_v1 import ObjectMeta
from lightkube.models.autoscaling_v1 import ScaleSpec

client = AsyncClient()
obj = Deployment.Scale(
    metadata=ObjectMeta(name='metrics-server', namespace='kube-system'),
    spec=ScaleSpec(replicas=1)
)
await client.replace(obj, 'metrics-server', namespace='kube-system')
```
{{< /tab >}}
{{< /tabs >}}

## Conclusion

In a nutshell, after digging into the various Kubernetes Python libraries, it's safe to say that `kr8s` is on the right track. My main goal here was not to pick a winner but to see if `kr8s` is hitting its marks of being simple, readable, and maintainable. And it looks like it is!

Throughout the comparison, a few key things stood out:

It feels like `kubectl`, `kr8s` manages to feel familiar, which is great. If you know `kubectl`, you'll find `kr8s` approachable.

Keeping things simple is at the core of `kr8s` and it does a good job of keeping characters to a minimum while remaining readable.

I've also found it interesting to see just how verbose `kubernetes_asyncio` is. It's the library we were primarily using in `dask-kubernetes` and fuelled the need to find something more readable and maintainable. When I assessed the alternatives before building `kr8s` I disregarded `lightkube` as the README has a warning about being experimental and not ready for use, but actually it seems in a pretty good state today. It is definitely more verbose than `kr8s` but there are clear and obvious benefits such as type safety that come with those extra characters.

In the world of Kubernetes Python libraries, it's all about options. `kr8s` is a strong contender for those who want a simple, readable, and maintainable Kubernetes experience. But there are other compelling options out there too. Cheers to open source and the freedom to choose the tools that work best for you! 🚀
