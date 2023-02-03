---
title: "Quick and dirty way to pre-pull container images on Kubernetes"
date: 2023-02-02T00:00:00+00:00
draft: false
author: "Jacob Tomlinson"
categories:
  - blog
tags:
  - kubernetes
  - containers
---

Sometimes when I give live demos with Kubernetes clusters I want to make sure that the container images I'm going to use are already pulled onto all of the nodes in my cluster. The last thing I want is for a `Pod` to be created to then sit in a `Pending` state while an image is pulled, especially given how large containers can be in the Data Science space.

I could run my demo through once ahead of time and clean up any resources, but if my cluster has multiple nodes I can't guarantee the same Pods will land on the same nodes.

To avoid this I use a little `DaemonSet` that will ensure all images are pulled onto every node.

```yaml
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: prepuller
spec:
  selector:
    matchLabels:
      name: prepuller
  template:
    metadata:
      labels:
        name: prepuller
    spec:
      # Configure an init container for each image you want to pull
      initContainers:
        - name: prepuller-1
          # Set the image you want to pull
          image: ORG/IMAGE:TAG
          # Use a known command that will exit successfully immediately
          # Any no-op command will do but YMMV with scratch based containers
          command: ["sh", "-c", "'true'"]

        # - name: prepuller-2
        #   image: ...
        #   command: ["sh", "-c", "'true'"]

        # etc...

      # Use the pause container to ensure the Pod goes into a `Running` phase
      # but doesn't take up resource on the cluster
      containers:
        - name: pause
          image: gcr.io/google_containers/pause
```
