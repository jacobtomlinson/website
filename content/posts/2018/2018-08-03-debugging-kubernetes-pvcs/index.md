---
title: Debugging Kubernetes PVCs
date: 2018-08-03T00:00:00+00:00
draft: false
categories:
- blog
tags:
- Kubernetes
- Persistent Volume Claim
- Debugging
author: Jacob Tomlinson
canonical: https://itnext.io/debugging-kubernetes-pvcs-a150f5efbe95
canonical_title: IT Next
---

Sometimes I find that something goes wrong in a container and some data stored in a [persistent volume](https://kubernetes.io/docs/concepts/storage/persistent-volumes/) gets corrupted. This may result in me having to get my hands dirty and have a poke around in the filesystem myself.

Recently I tried to do this with a [Prometheus](https://github.com/prometheus/prometheus) container and found that there was no shell or environment inside the container (à la best practice!). This meant I needed to attach the [PVC](https://kubernetes.io/docs/concepts/storage/persistent-volumes/#persistentvolumeclaims) to a different pod with an environment I could use to debug.

## Detach the volume

The container was stuck in a restart loop because of the corrupted data. So the first step was to scale the deployment to zero.

```console
$ kubectl scale deployment my-deployment --replicas=0
deployment.extensions "my-deployment" scaled
```

## Create a debugging pod

Now I need to inspect the deployment to find out which PVC I want to explore.

```console
$ kubectl describe deployment my-deployment | grep ClaimName
ClaimName:  my-claim
```

Then I want to create a new pod spec which mounts the same PVC but using a different docker image. In this example I'll use busybox as I just need a basic shell, but you could use any debugging tools image here.

```yaml
# my-pvc-debugger.yaml
kind: Pod
apiVersion: v1
metadata:
  name: volume-debugger
spec:
  volumes:
    - name: volume-to-debug
      persistentVolumeClaim:
       claimName: <CLAIM NAME GOES HERE>
  containers:
    - name: debugger
      image: busybox
      command: ['sleep', '3600']
      volumeMounts:
        - mountPath: "/data"
          name: volume-to-debug
```

I then create this pod and run a shell inside it.

```console
$ kubectl create -f /path/to/my-pvc-debugger.yaml
pod "volume-debugger" created
$ kubectl exec -it volume-debugger sh
/ #
```

Now that I'm inside the container I can explore the volume which is mounted at `/data` and fix the issue.

## Scaling back

Once I'm happy with the volume I can exit my shell within the container and delete the debugger pod.

```console
/ # logout
$ kubectl delete -f /path/to/my-pvc-debugger.yaml
```

Finally I can scale my deployment back up.

```console
$ kubectl scale deployment my-deployment --replicas=1
deployment.extensions "my-deployment" scaled
```

## Conclusion

In a perfect world we should never have to get hands on with our volumes, but occasionally bugs cause if to have to go and clean things up. This example shows a quick way to hop into a volume for a container which does not have any user environment.
