---
title: "Oversubscribing GPUs in Kubernetes"
date: 2023-02-07T00:00:00+00:00
draft: false
author: "Jacob Tomlinson"
categories:
  - blog
tags:
  - kubernetes
  - gpus
  - scheduling
---

Sometimes I want to oversubscribe the GPUs in my Kubernetes cluster. This is especially useful when I'm developing but could also be useful in light workloads where you have ample GPU memory and don't mind the occasional failure.

```warning
If you are thinking about doing this in production there are probably better options.
```

Newer NVIDIA Data Center GPUs support [MIG](https://docs.nvidia.com/datacenter/cloud-native/kubernetes/mig-k8s.html) which allows partitioning a GPU at the hardware level.

You might also find [MPS](https://towardsdatascience.com/how-to-increase-gpu-utilization-in-kubernetes-with-nvidia-mps-e680d20c3181) interesting which allows multiple processes to share the GPU at a CUDA level. However it seems Kubernetes support for this today is [still a work in progress](https://developer.nvidia.com/blog/improving-gpu-utilization-in-kubernetes/).

In this post we are going to use [**time slicing**](https://docs.nvidia.com/datacenter/cloud-native/gpu-operator/gpu-sharing.html) to share our GPUs between Pods. This works by running many CUDA processes on the same GPU and giving them equal time slices of computation.

```warning
With time slicing there is no memory or fault-isolation between processes. So processes should play nicely or they will all get OOM errors.
```

Ok enough warning you that this might not be a good idea. Let's do it!

## Prep

First we need a working Kubernetes cluster with some GPUs.

### Kubernetes Cluster

 I'm going to run one on my workstation with [my patched version of kind](../../2022/quick-hack-adding-gpu-support-to-kind/), but you get yours however you prefer.

```console
$ cat << EOF > kind-gpu.yaml
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
name: gpu-cluster
nodes:
  - role: control-plane
    gpus: true
EOF
```

```console
$ kind create cluster --config kind-gpu.yaml
Creating cluster "gpu-cluster" ...
 ✓ Ensuring node image (kindest/node:v1.23.1) 🖼
 ✓ Preparing nodes 📦
 ✓ Writing configuration 📜
 ✓ Starting control-plane 🕹️
 ✓ Installing CNI 🔌
 ✓ Installing StorageClass 💾
Set kubectl context to "kind-gpu-cluster"
You can now use your cluster with:

kubectl cluster-info --context kind-gpu-cluster

Thanks for using kind! 😊
```

```console
$ kubectl get nodes
NAME                        STATUS   ROLES                  AGE   VERSION
gpu-cluster-control-plane   Ready    control-plane,master   74s   v1.23.1
```

### NVIDIA Operator

Next we need to install the NVIDIA Operator which can install NVIDIA drivers and the Kubernetes GPU device plugin.

```info
As I'm using `kind` on my workstation which already has NVIDIA drivers installed I'll set a flag to skip that.
```

```console
$ helm install --repo https://helm.ngc.nvidia.com/nvidia gpu-operator \
  --wait --generate-name \
  --create-namespace -n gpu-operator \
  --set driver.enabled=false
NAME: gpu-operator-1675786339
LAST DEPLOYED: Tue Feb  7 16:12:27 2023
NAMESPACE: gpu-operator
STATUS: deployed
REVISION: 1
TEST SUITE: None
```

Once the operator has successfully installed you should see all of the Pods it created in a `Running` phase except for the two validator Pods which should be `Completed`.

```console
$ kubectl get pods --namespace gpu-operator
NAME                                                              READY   STATUS      RESTARTS   AGE
gpu-feature-discovery-t6pwg                                       1/1     Running     0          2m1s
gpu-operator-1675786339-node-feature-discovery-master-bc94sqsn7   1/1     Running     0          2m33s
gpu-operator-1675786339-node-feature-discovery-worker-l9rc5       1/1     Running     0          2m33s
gpu-operator-5cf698664-kswsm                                      1/1     Running     0          2m33s
nvidia-container-toolkit-daemonset-z2nst                          1/1     Running     0          2m1s
nvidia-cuda-validator-4cgbx                                       0/1     Completed   0          80s
nvidia-dcgm-exporter-9jkdt                                        1/1     Running     0          2m1s
nvidia-device-plugin-daemonset-lgd7l                              1/1     Running     0          2m1s
nvidia-device-plugin-validator-lz9qm                              0/1     Completed   0          59s
nvidia-operator-validator-vn9w4                                   1/1     Running     0          2m1s
```

Ok we should be good to run some GPU workloads.

## Some workload

Now that we have a GPU capable Kubernetes cluster let's throw too much work at it. My workstation has two GPUs in it, so let's create a deployment that needs four GPUs.

```console
$ cat << EOF | kubectl create -f -
apiVersion: apps/v1
kind: Deployment
metadata:
  name: gpu-workload
spec:
  replicas: 4
  selector:
    matchLabels:
      app: gpu-workload
  template:
    metadata:
      labels:
        app: gpu-workload
    spec:
      containers:
      - name: pause
        image: gcr.io/google_containers/pause
        resources:
          limits:
            nvidia.com/gpu: 1
EOF
deployment.apps/gpu-workload created
```

Now if we have a look at our Pods we should see two `Running` and two `Pending` because we hit our GPU capacity.

```console
$ kubectl get pods
NAME                            READY   STATUS    RESTARTS   AGE
gpu-workload-5cf65846c5-72vz2   1/1     Running   0          13s
gpu-workload-5cf65846c5-ngtzf   1/1     Running   0          13s
gpu-workload-5cf65846c5-g6brs   0/1     Pending   0          13s
gpu-workload-5cf65846c5-qjhtc   0/1     Pending   0          13s
```

## Enabling time-slicing

To enable time-slicing we are going to follow the [guide in the official docs](https://docs.nvidia.com/datacenter/cloud-native/gpu-operator/gpu-sharing.html).

First we need to create a configuration for the NVIDIA Device Plugin to use. Here we create a new config profile called `default` which allows one GPU to be sliced into four.

```console
$ cat << EOF | kubectl create -f -
apiVersion: v1
kind: ConfigMap
metadata:
  name: time-slicing-config
  namespace: gpu-operator
data:
    default: |-
        version: v1
        sharing:
          timeSlicing:
            resources:
            - name: nvidia.com/gpu
              replicas: 4
EOF
configmap/time-slicing-config created
```

Then we need to patch the NVIDIA Operator to tell it to use this config.

```console
$ kubectl patch clusterpolicy/cluster-policy \
   -n gpu-operator --type merge \
   -p '{"spec": {"devicePlugin": {"config": {"name": "time-slicing-config"}}}}'
```

Then we need to tell all of our nodes to use the default profile (although you can [configure this on a per-node basis too](https://docs.nvidia.com/datacenter/cloud-native/gpu-operator/gpu-sharing.html#applying-a-time-slicing-configuration-per-node)).

```console
$ kubectl patch clusterpolicy/cluster-policy \
   -n gpu-operator --type merge \
   -p '{"spec": {"devicePlugin": {"config": {"name": "time-slicing-config", "default": "default"}}}}'
```

Now if we give the controller a minute to restart and check our Pods again we should see everything in a `Running` phase.

```console
$ kubectl get pods
NAME                            READY   STATUS    RESTARTS   AGE
gpu-workload-5cf65846c5-72vz2   1/1     Running   0          10m
gpu-workload-5cf65846c5-g6brs   1/1     Running   0          10m
gpu-workload-5cf65846c5-ngtzf   1/1     Running   0          10m
gpu-workload-5cf65846c5-qjhtc   1/1     Running   0          10m
```

We can also check the node and see that it now appears that we have eight GPUs instead of the actual two because they are being time sliced into four pieces.

```console
$ kubectl describe node gpu-cluster-control-plane| grep -E "(nvidia.com/gpu:|Capacity:|Allocatable:)"
Capacity:
  nvidia.com/gpu:     8
Allocatable:
  nvidia.com/gpu:     8
```

## Closing

I work a bunch on various Kubernetes deployment tools, so I often want to have many Pods running in my cluster that all use GPUs. Being able to oversubscribe the GPUs in my workstation means I can try more stuff out locally. I'm not actually going to run any large workloads on this cluster, if I did I may run into memory errors and other problems from GPU sharing.

But this solves my problem!
