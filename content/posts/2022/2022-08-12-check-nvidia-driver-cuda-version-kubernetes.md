---
title: "How to check your NVIDIA driver and CUDA version in Kubernetes"
date: 2022-08-12T00:00:00+00:00
draft: false
author: "Jacob Tomlinson"
categories:
  - blog
tags:
  - Kubernetes
  - NVIDIA driver
  - CUDA
---

When using GPUs with Kubernetes it can be important to know which driver and CUDA versions are installed on the nodes.

## Method 1: Node labels

If you are using the [NVIDIA GPU Operator](https://docs.nvidia.com/datacenter/cloud-native/gpu-operator/overview.html) to install your drivers the feature discovery service will automatically add labels to your nodes with this version info.

```console
$ kubectl describe node -A | grep nvidia.com/cuda | uniq
                    nvidia.com/cuda.driver.major=495
                    nvidia.com/cuda.driver.minor=46
                    nvidia.com/cuda.driver.rev=
                    nvidia.com/cuda.runtime.major=11
                    nvidia.com/cuda.runtime.minor=5
```

Here we see the driver version is `495.46` and the CUDA version is `11.5`.

These labels can be extremely useful if you have many nodes in your cluster with different driver/CUDA versions and you want to restrict your `Pods` to only run with specific versions.

## Method 2: Create a pod with `nvidia-smi`

To find this information on your own machine you usually use `nvidia-smi`, so to do this on Kubernetes you can create a pod that runs `nvidia-smi` and check the logs to see its output.

```console
$ cat << EOF | kubectl create -f -
apiVersion: v1
kind: Pod
metadata:
  name: nvidia-version-check
spec:
  restartPolicy: OnFailure
  containers:
  - name: nvidia-version-check
    image: "nvidia/cuda:11.0.3-base-ubuntu20.04"
    command: ["nvidia-smi"]
    resources:
      limits:
         nvidia.com/gpu: "1"
EOF
pod/nvidia-version-check created
```

Now we can view the logs.

```console
$ kubectl logs nvidia-version-check
+-----------------------------------------------------------------------------+
| NVIDIA-SMI 495.46       Driver Version: 495.46       CUDA Version: 11.5     |
|-------------------------------+----------------------+----------------------+
| GPU  Name        Persistence-M| Bus-Id        Disp.A | Volatile Uncorr. ECC |
| Fan  Temp  Perf  Pwr:Usage/Cap|         Memory-Usage | GPU-Util  Compute M. |
|                               |                      |               MIG M. |
|===============================+======================+======================|
|   0  Quadro RTX 8000     Off  | 00000000:15:00.0 Off |                  Off |
| 33%   29C    P8    16W / 260W |      5MiB / 48601MiB |      0%      Default |
|                               |                      |                  N/A |
+-------------------------------+----------------------+----------------------+

+-----------------------------------------------------------------------------+
| Processes:                                                                  |
|  GPU   GI   CI        PID   Type   Process name                  GPU Memory |
|        ID   ID                                                   Usage      |
|=============================================================================|
+-----------------------------------------------------------------------------+
```

Here we see the driver version is `495.46` and the CUDA version is `11.5`.

Last we should clear up our container again.

```console
$ kubectl delete pod nvidia-version-check
pod "nvidia-version-check" deleted
```

## Method 3: Call `nvidia-smi` with a `kubectl run` one-liner

This will only work if all of your nodes have NVIDIA drivers because we can't specify GPU resources and the `Pod` could land on any node, but if they are you can call `nvidia-smi` with a one-liner.

```console
$ kubectl run nvidia-smi --restart=Never --rm -i --tty --image nvidia/cuda:11.0.3-base-ubuntu20.04 -- nvidia-smi
+-----------------------------------------------------------------------------+
| NVIDIA-SMI 495.46       Driver Version: 495.46       CUDA Version: 11.5     |
|-------------------------------+----------------------+----------------------+
| GPU  Name        Persistence-M| Bus-Id        Disp.A | Volatile Uncorr. ECC |
| Fan  Temp  Perf  Pwr:Usage/Cap|         Memory-Usage | GPU-Util  Compute M. |
|                               |                      |               MIG M. |
|===============================+======================+======================|
|   0  Quadro RTX 8000     Off  | 00000000:15:00.0 Off |                  Off |
| 33%   30C    P8    16W / 260W |      5MiB / 48601MiB |      0%      Default |
|                               |                      |                  N/A |
+-------------------------------+----------------------+----------------------+
|   1  Quadro RTX 8000     Off  | 00000000:2D:00.0 Off |                  Off |
| 33%   31C    P8    13W / 260W |     14MiB / 48593MiB |      0%      Default |
|                               |                      |                  N/A |
+-------------------------------+----------------------+----------------------+

+-----------------------------------------------------------------------------+
| Processes:                                                                  |
|  GPU   GI   CI        PID   Type   Process name                  GPU Memory |
|        ID   ID                                                   Usage      |
|=============================================================================|
+-----------------------------------------------------------------------------+
pod "nvidia-smi" deleted
```
