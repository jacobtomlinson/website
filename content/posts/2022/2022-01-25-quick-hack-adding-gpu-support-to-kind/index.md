---
title: "Quick hack: Adding GPU support to kind"
date: 2022-01-25T00:00:00+00:00
draft: false
author: "Jacob Tomlinson"
categories:
  - blog
tags:
  - Community
  - Open Source
  - Kubernetes
  - Kind
  - Golang
---

I needed GPU support in [kind](https://kind.sigs.k8s.io/), so I added it. I'm also prone to yak shaving so it's quick, dirty and not going upstream.

When developing tools for [Kubernetes](https://kubernetes.io/) I like to use kind which runs a whole cluster inside a single [Docker](https://www.docker.com/) container. I especially like using it via [pytest-kind](https://pypi.org/project/pytest-kind/) which makes running Python unit tests against a Kubernetes cluster a breeze.

Today as of kind 0.11.1 there is no support for passing GPUs through to the Kubernetes cluster and attempts made in [kubernetes-sigs/kind#1886](https://github.com/kubernetes-sigs/kind/pull/1886) were rejected. It seems there is a desire to add this support to kind in the future, but disagreements on how to implement it. Sadly I don't have time to dive into that and try and implement a robust solutions that would be accepted by the kind maintainers, so I decided to quickly hack together a version that I could use right away.

You can find [my fork of kind here with a Pull Request that adds GPU support](https://github.com/jacobtomlinson/kind/pull/1).

My PR adds a `gpus` config option to kind nodes which passes the `--gpus=all` flag to Docker. So all you need to use it is the NVIDIA drivers, Docker and the [NVIDIA runtime](https://github.com/NVIDIA/nvidia-docker). If you can run this quick test you're all set.

```console
$ docker run --rm --gpus=all nvcr.io/nvidia/k8s/cuda-sample:vectoradd-cuda10.2
[Vector addition of 50000 elements]
Copy input data from the host memory to the CUDA device
CUDA kernel launch with 196 blocks of 256 threads
Copy output data from the CUDA device to the host memory
Test PASSED
Done
```

## Installing my fork

If you have [golang](https://go.dev/) installed you can pull down my fork and build yourself a `kind` binary with GPU support. Alternatively if you're on 64 bit linux you can grab [this binary](https://github.com/jacobtomlinson/kind/files/7935041/kind.zip) I already built.

```console
$ git clone https://github.com/jacobtomlinson/kind.git

$ cd kind

$ git branch gpu && git pull origin gpu

$ make install

$ kind --version
kind version 0.12.0-alpha+bf63502ed1dfde
```

## Creating a GPU kind cluster

Create a kind cluster template YAML and specify `gpus: true` in the `control-plane` node's config.

```yaml
# kind-gpu.yaml
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
name: gpu-test
nodes:
  - role: control-plane
    gpus: true
```

Then create the cluster.

```console
$ kind create cluster --config kind-gpu.yaml
```

I like to use the awesome [kubectx](https://github.com/ahmetb/kubectx) command to switch my context over to my new kind cluster.

```console
$ kubectx kind-gpu-test
```

## Installing NVIDIA plugins

In order for us to be able to schedule GPUs in our cluster we need the NVIDIA plugins installed. We can
do this via the NVIDIA Operator. Let's install that with [helm](https://helm.sh/).

As our host machine already has NVIDIA drivers installed we need to disable the driver install step.

```console
$ helm repo add nvidia https://nvidia.github.io/gpu-operator \
  && helm repo update

$ helm install --wait --generate-name \
  -n gpu-operator --create-namespace \
  nvidia/gpu-operator \
  --set driver.enabled=false
```

## Testing

Now that we have things set up let's test that we can schedule a pod with a GPU, we can use the vector add example from earlier.

```yaml
# gpu-pod.yaml
apiVersion: v1
kind: Pod
metadata:
  name: vectoradd
spec:
  restartPolicy: OnFailure
  containers:
  - name: vectoradd
    image: nvcr.io/nvidia/k8s/cuda-sample:vectoradd-cuda10.2
    resources:
      limits:
        nvidia.com/gpu: 1
```

```console
$ kubectl apply -f gpu-pod.yaml
```

If we run `kubectl get pods -w` we should see our pod go from `Pending` to `ContainerCreating` to `Running` to `Completed`.

Then if we check the logs we should see the same output as before.

```console
$ kubectl logs vectoradd
[Vector addition of 50000 elements]
Copy input data from the host memory to the CUDA device
CUDA kernel launch with 196 blocks of 256 threads
Copy output data from the CUDA device to the host memory
Test PASSED
Done
```

## What next?

I built a quick version of `kind` with GPU support so that I can quickly test the thing I'm actually meant to be working on. Ideally I would like to help add
this feature into upstream kind for everyone to use, but I don't have the time right now. Maybe I'll find some time another day.

Sadly this fork of kind will slowly go out of date unless this feature is added upstream. These are the consequences of forking a project, you own the maintenance, and I do not intend on maintaining this.

You're welcome to use my forked version, but here be dragons.
