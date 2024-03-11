---
title: "A beginner's guide to managing Kubernetes resources in Python with kr8s"
date: 2024-03-11T00:00:00+00:00
draft: false
author: "Jacob Tomlinson"
categories:
  - blog
tags:
  - python
  - kubernetes
  - tutorial
  - kr8s
---

Managing Kubernetes resources with Python has never been easier thanks to the [`kr8s` Kubernetes client for Python](https://github.com/kr8s-org/kr8s).

A core goal of `kr8s` is to provide a simple and complete Python client for Kubernetes to allow folks to get up and running quickly. If you've used `kubectl` then you'll find the `kr8s` syntax familiar. Whether you want to write scripts, automate repetitive tasks, build applications that interact with the Kubernetes API or build operators, `kr8s` is a great place to start.

## Setup

To get started you'll need a Kubernetes cluster and some credentials. `kr8s` looks in the same places for your credentials as `kubectl` so if you can run `kubectl` commands you're good to go.

```console
$ kubectl get nodes        
NAME                        STATUS   ROLES           AGE   VERSION
kind-control-plane          Ready    control-plane   59d   v1.29.0
```

In our Python environment we can install the `kr8s` package using `pip`.

```bash
pip install kr8s
```

Now let's verify that we can interact with the Kubernetes API and make the equivalent call to `kubectl get nodes`.

```python
import kr8s

for node in kr8s.get("nodes"):
    print(node.name)
```

If you see the same node names printed out then everything is set up!

## Creating resources

Let's start by creating some Kubernetes resources. There are many different ways to do this with `kr8s`, so let's explore the most common ones.

First we can generate common resources from a few bits of key information, this is similar to the `kubectl run` command.

```python
from kr8s.objects import Pod

# Generate a simple Pod spec using the `nginx` container image
pod = Pod.gen(name="webserver", image="nginx:latest", ports=[80])  
# Create the Pod
pod.create()  
```

We could also define the Pod spec explicitly as a Python dictionary.

```python
from kr8s.objects import Pod

pod = Pod({
        "apiVersion": "v1",
        "kind": "Pod",
        "metadata": {
            "name": "webserver",
        },
        "spec": {
            "containers": [{
                "name": "webserver", 
                "image": "nginx:latest",
                "ports": [{
                    "containerPort": 80, 
                    "protocol": "TCP",
                }]
            }]
        },
    })

pod.create()
```

Or we could store the resource in a YAML file and use `kr8s` to load and create the object, similar to `kubectl create -f`.

```yaml
# nginx-pod.yaml
apiVersion: v1
kind: Pod
metadata:
  name: webserver
spec:
  containers:
  - name: webserver
    image: nginx:latest
    ports:
    - containerPort: 80
```

```python
from kr8s.objects import objects_from_files

# We do this in a loop because there may be more than one resource in a file
for resource in objects_from_files("nginx-pod.yaml"):
    resource.create()
```

To learn more about creating resources see the [`kr8s` example documentation](https://docs.kr8s.org/en/stable/examples/creating_resources.html).

## Listing resources

We've already seen how we can list resources using `kr8s.get()` but often we need to do some filtering or querying of objects.

You can use the same selectors that `kubectl` supports for filtering by labels or fields. For example we can list only running Pods.

```python
import kr8s

for pod in kr8s.get("pods", field_selector="status.phase=Running"):
    print(pod.name)
```

The resource objects returned by `kr8s` allow you to access the data from the resource either using dict notation `pod["metadata"]` or using dot notation `pod.metadata`. Having our resources available in Python allows you to use the full power of Python to work with these objects. For example if we wanted to sort our Pods by their restart count we can use the built-in `sort` list operation and a `lambda` to select which field to sort by.

```python
import kr8s

# Here we are listing Pods in all namespaces
pods = kr8s.get("pods", namespace=kr8s.ALL)  
pods.sort(
    key=lambda pod: pod.status.containerStatuses[0].restartCount, 
    reverse=True
)

for pod in pods:
    print(pod.name, pod.status.containerStatuses[0].restartCount)
```

We can also get a reference to an object directly if we know it's name.

```python
from kr8s.objects import Pod

pod = Pod.get("webserver")  # The nginx Pod we created earlier
```

See the [listing resources documentation](https://docs.kr8s.org/en/stable/examples/listing_resources.html#list-ready-pods) for more examples.

## Patching and updating objects

Resource objects in `kr8s` can update the remote objects in the Kubernetes cluster with the `patch()` method. For example we could add labels to the Pod we created earlier.

```python
from kr8s.objects import Pod

pod = Pod.get("webserver")

# Update the labels
pod.patch({"metadata": {"labels": {"foo": "bar"}}})

print(pod.metadata.labels)  # prints '{'foo': 'bar'}'
```

[JSON 6902](https://jsonpatch.com/) style patching is also supported which allows you to make more targeted updates.

```python
pod.patch(
    [{"op": "add", "path": "/metadata/labels/patched", "value": "true"}],
    type="json",
)
```

Many `kr8s` objects also have utility methods to simplify making common changes to resources. For example you can add labels with `.label()`.

```python
pod.label({"fizz": "buzz"})
```

Here are some more common examples:

```python
from kr8s.objects import Deployment

# Scale a deployment
deploy = Deployment.get("metrics-server", namespace="kube-system")
deploy.scale(1)

# Cordon a node
node = Node("k8s-node-1")
node.cordon()
```

## Interacting with running Pods

The Pod resource also has additional operations you may want to perform such as viewing logs, running commands or forwarding ports from your local machine.

```python
# Print the logs from our nginx Pod
for line in pod.logs():
    print(line)
```

```python
# Exec a command in our nginx Pod and print the stdout
command = pod.exec(["cat", "/etc/os-release"])
print(command.stdout.decode())
```

```python
# Forward port 80 of nginx to 8080 on our local machine
pod.portforward(80, local_port=8080).run_forever()
# Open http://localhost:8080/ to view the nginx welcome page
```

For more examples or running commands or setting up port forwarding in the background check out the [Pod operations example documentation](https://docs.kr8s.org/en/stable/examples/pod_operations.html).

## Deleting resources

Deleting resources works the same way as creating them. All we need to know about a resource to delete it is it's name (and namespace if not default).

```python
from kr8s.objects import Pod

# Get our nginx Pod
pod = Pod.get("webserver")
# Delete it
pod.delete()  
```

## Conclusion

In this guide we've only scratched the surface of what is possible with interacting with Kubernetes in Python. Hopefully touching on the basics of creating and manipulating resources gives some ideas for things you could build with `kr8s` and the Kubernetes API.

For more inspiration head over to the [guides section of the `kr8s` documentation](https://docs.kr8s.org/en/stable/guides/index.html) for examples of building end-to-end projects like operators.
