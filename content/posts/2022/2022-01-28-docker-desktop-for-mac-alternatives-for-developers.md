---
title: "Docker Desktop for Mac alternatives for developers"
date: 2022-01-28T00:00:00+00:00
draft: false
author: "Jacob Tomlinson"
categories:
  - blog
tags:
  - Docker
  - Kubernetes
  - Kind
  - Podman
  - Minikube
  - Docker Desktop
---

In a couple of days [Docker will begin charging employees of companies with >250 employees to use Docker Desktop](https://www.docker.com/blog/updating-product-subscriptions/). I have no problem with paying for software that brings me value, but you wouldn't believe how complex it can be for large companies to sign employees up to subscription services. Paperwork everywhere! To avoid this I'm evaluating alternatives for Docker Desktop to use on my MacBook.

## What is Docker Desktop for Mac?

[Docker Desktop for Mac](https://hub.docker.com/editions/community/docker-ce-desktop-mac) is an application ([also available for Windows](https://hub.docker.com/editions/community/docker-ce-desktop-windows)) that runs Docker inside a virtual machine transparently. All you need to do is install Docker Desktop and start the service and you can run Docker containers from the command line on your Mac. It neatly makes your local filesystem available to containers via volume mounts and maps all ports back to your Mac. It feels just like running Docker on Linux. Recently it has also supported running [Kubernetes](https://kubernetes.io/) on your Mac, although I've always found it finnicky and used alternatives like [kind](https://kind.sigs.k8s.io/).

It also has a GUI for accessing your containers, images and settings.

I've been a user [since the beta](https://www.docker.com/blog/docker-for-mac-windows-beta/) and it was called "Docker for Mac". In all honesty I miss the old UI.

![Docker for Mac from 2016](https://i.imgur.com/A5eT7c7.png)

## My Requirements

Here are the core features that Docker Desktop provides that I use and care about that I am looking for a replacement for.

- Use the `docker` commands in the same way you would on Linux.
- Mount volumes from my Mac into the container with the `-v` flag.
- Expose ports on my Mac with the `-p` flag.
- Create Kubernetes clusters inside containers with [kind](https://kind.sigs.k8s.io/).
- Use Docker on a poor network connection (like on a train).
- Configure the resources that Docker is using on my Mac (CPU, memory, storage).
- It should be easy to install/use!

Things that Docker Desktop provides that I don't need.

- A graphical UI.
- Kubernetes, because I prefer to use kind.

_I like kind because I can create multiple Kubernetes clusters with different versions and configurations, this is great when building and testing tools for Kubernetes. This may be less relevant to you if you consider yourself a Kubernetes user rather than a developer. I also like how portable it is and how it integrates nicely with test suites like [pytest-kind](https://pypi.org/project/pytest-kind/)._

To test out my alternatives I am going to install them and do the following things:

- Run an Nginx container that serves an HTML file from my MacBook on port 8080.
- Run my Nginx container again with the WiFi turned off.
- Create a `kind` cluster.
- Modify the hardware allocation that Docker has access to.

I'll score each option out of 5, a point for volumes, ports, works offline, runs kind and can easily modify the hardware allocation.

For my testing I'm going to create a directory in my home directory called `test` on my Mac with an `index.html` file in that looks like this:

```html
<!DOCTYPE html>
<html>
<body>
<h1>It works!</h1>
<p>Docker volumes and ports working as expected.</p>
</body>
</html>
```

## Docker Desktop alternatives

I'm going to work through this non-exhaustive list of alternatives that I found by googling around.

- [Colima](https://github.com/abiosoft/colima)
- [Podman](https://podman.io/)
- [Minikube](https://minikube.sigs.k8s.io/docs/)
- [Remote Docker on another machine](https://www.cloudsavvyit.com/11185/how-and-why-to-use-a-remote-docker-host/)
- [Running a virtual machine myself](https://medium.com/carvago-development/my-docker-on-macos-part-1-setup-ubuntu-virtual-machine-both-intel-and-apple-silicon-cpu-5d886af0ebba)

Did I miss any other options? [Let me know](https://twitter.com/_jacobtomlinson)!

### Colima

[Colima](https://github.com/abiosoft/colima) is built on [Lima](https://github.com/lima-vm/lima) which is a virtual machine tool for MacOS with automatic file sharing and port forwarding. Using Lima feels a lot like using WSL on Windows.

Colima builds on that foundation to run a VM with Docker installed and it also configures your local docker context for you.

#### Installing

```console
$ brew install docker colima

$ colima start
INFO[0000] starting colima
INFO[0000] creating and starting ...                     context=vm
INFO[0030] provisioning ...                              context=docker
INFO[0031] restarting VM to complete setup ...           context=docker
INFO[0031] stopping ...                                  context=vm
INFO[0037] starting ...                                  context=vm
INFO[0058] starting ...                                  context=docker
INFO[0063] waiting for startup to complete ...           context=docker
INFO[0063] done
```

That's it! Colima is now running Docker in a Lima VM and it created a new `docker context` and switched to it so we can use the `docker` CLI straight away.

```console
$ docker ps
CONTAINER ID   IMAGE     COMMAND   CREATED   STATUS    PORTS     NAMES

$ docker context list
NAME                TYPE                DESCRIPTION                               DOCKER ENDPOINT                                    KUBERNETES ENDPOINT                ORCHESTRATOR
colima *            moby                colima                                    unix:///Users/jtomlinson/.colima/docker.sock
default             moby                Current DOCKER_HOST based configuration   unix:///var/run/docker.sock                        https://127.0.0.1:6443 (default)   swarm
desktop-linux       moby                                                          unix:///Users/jtomlinson/.docker/run/docker.sock
```

#### Testing

Let's start our `nginx` container.

```console
$ docker run --rm -p 8080:80 -v $HOME/test:/usr/share/nginx/html:ro nginx
```

Now I can head to my browser and check that it works.

![Web browser showing the test page](https://i.imgur.com/EOXMg93.png)

Hooray! That worked, my volume mounted successfully and the port was made available on my laptop. I also successfully disabled my wifi and ran the command again and got the same result.

Next let's create a `kind` cluster.

```console
$ kind create cluster --name test
Creating cluster "test" ...
 ✓ Ensuring node image (kindest/node:v1.21.1) 🖼
 ✓ Preparing nodes 📦
 ✓ Writing configuration 📜
⢆⡱ Starting control-plane 🕹️
[kubelet-check] It seems like the kubelet isn't running or healthy.
[kubelet-check] The HTTP call equal to 'curl -sSL http://localhost:10248/healthz' failed with error: Get "http://localhost:10248/healthz": dial tcp [::1]:10248: connect: connection refused.
```

Hrm that's disappointing. It looks like `kind` failed to start the `kubelet` inside the container. I've run this multiple times in various ways and get the same result each time. This command works perfectly when using the `default` Docker context that comes with Docker Desktop.

This feels like something that could be resolved, but is a bit of a roadblock when getting started. Here's the [full error](https://gist.github.com/jacobtomlinson/76c410118beff174ef9d564694b9712d) if you're interested in debugging this.

Let's move on and try to modify the resources available to the Lima VM. It seems that [by default the VM has 2 CPU cores, 2GiB of memory and 60Gib of storage](https://github.com/abiosoft/colima#customizing-the-vm). We can modify the CPU and memory by stopping and starting Colima.

```console
$ colima stop
INFO[0000] stopping colima
INFO[0000] stopping ...                                  context=docker
INFO[0001] stopping ...                                  context=vm
INFO[0005] done
$ colima start --cpu 4 --memory 8
INFO[0000] stopping colima
INFO[0000] stopping ...                                  context=docker
INFO[0001] stopping ...                                  context=vm
INFO[0006] done
INFO[0000] using docker runtime
INFO[0000] starting colima
INFO[0000] starting ...                                  context=vm
INFO[0020] provisioning ...                              context=docker
INFO[0021] starting ...                                  context=docker
INFO[0026] waiting for startup to complete ...           context=docker
INFO[0026] done
```

However if we want to modify the storage allocation we need to delete the VM and recreate it, which is straight forward but means we will lose our container and images.

```console
$ colima stop
INFO[0000] stopping colima
INFO[0000] stopping ...                                  context=docker
INFO[0001] stopping ...                                  context=vm
INFO[0005] done

$ colima delete
are you sure you want to delete colima and all settings? [y/N] y
INFO[0001] deleting colima
INFO[0001] deleting ...                                  context=vm
INFO[0001] done

$ colima start --cpu 4 --memory 8 --disk 100
INFO[0000] starting colima
INFO[0000] creating and starting ...                     context=vm
INFO[0030] provisioning ...                              context=docker
INFO[0031] restarting VM to complete setup ...           context=docker
INFO[0031] stopping ...                                  context=vm
INFO[0037] starting ...                                  context=vm
INFO[0057] starting ...                                  context=docker
INFO[0063] waiting for startup to complete ...           context=docker
INFO[0063] done
```

#### Pros/Cons

Colima seems pretty nice. It was easy to install and use and you can dynamically change CPU and memory allocations on the fly. However it does seem to have some issues with running more complex containers that will stop me from using it for now, which is a shame because it seemed pretty perfect.

It gets 4/5.

### Podman

Next up is [Podman](https://podman.io/), a Docker alternative from RedHat. Podman follows the same command line API as Docker so it makes for a nice drop in replacement.

#### Installing

```console
$ brew install podman

$ podman machine init

$ podman machine start
INFO[0000] waiting for clients...
INFO[0000] listening tcp://127.0.0.1:7777
INFO[0000] new connection from  to /var/folders/0l/fmwbqvqn1tq96xf20rlz6xmm0000gp/T/podman/qemu_podman-machine-default.sock
Waiting for VM ...
Machine "podman-machine-default" started successfully
```

Now we have podman running we can use the `podman` cli in exactly the same way as we use docker.

```console
$ podman ps
CONTAINER ID  IMAGE       COMMAND     CREATED     STATUS      PORTS       NAMES
```

Optionally you could even `alias docker=podman`.

#### Testing

Let's start our `nginx` container.

```console
$ docker run --rm -p 8080:80 -v $HOME/test:/usr/share/nginx/html:ro nginx
Error: statfs /Users/jtomlinson/test: no such file or director
```

Oh no, it looks like our volume mounts don't work here. This is because podman doesn't make files from my mac available inside the podman VM. It looks like [there are plans to make this happen](https://github.com/containers/podman/issues/8016), but they haven't been implemented yet.

Let's try running the command without the volume mount.

```console
$ docker run --rm -p 8080:80 nginx
```

This works so port mappings seem to be fine.

![Default nginx page](https://i.imgur.com/KKSDMYT.png)

Disabling my wifi and running the command again works, so network access is not required. However I did notice that hitting ctrl+c didn't stop the container and I had to manually run `podman rm -f <name>` to stop it. Not a problem but definitely different to `docker`.

Now let's launch a Kubernetes cluster with `kind`. We need to make a few modifications to our `podman` vm following the [kind documentation here](https://kind.sigs.k8s.io/docs/user/rootless/).

```console
$ podman machine ssh

$ sudo bash

$ mkdir /etc/systemd/system/user@.service.d/

$ cat <<EOF >/etc/systemd/system/user@.service.d/delegate.conf
[Service]
Delegate=yes
EOF

$ cat <<EOF >/etc/modules-load.d/iptables.conf
ip6_tables
ip6table_nat
ip_tables
iptable_nat
```

Then logout and restart the VM.

```console
$ podman machine stop

$ podman machine start
```

We can run our `kind create` command now but need to set an environment first to enable the experimental `podman` support in `kind`.

```console
$ KIND_EXPERIMENTAL_PROVIDER=podman kind create cluster --name test
using podman due to KIND_EXPERIMENTAL_PROVIDER
enabling experimental podman provider
Cgroup controller detection is not implemented for Podman. If you see cgroup-related errors, you might need to set systemd property "Delegate=yes", see https://kind.sigs.k8s.io/docs/user/rootless/
Creating cluster "test" ...
 ✓ Ensuring node image (kindest/node:v1.21.1) 🖼
 ✓ Preparing nodes 📦
 ✓ Writing configuration 📜
 ✓ Starting control-plane 🕹️
 ✓ Installing CNI 🔌
 ✓ Installing StorageClass 💾
Set kubectl context to "kind-test"
You can now use your cluster with:

kubectl cluster-info --context kind-test

Have a question, bug, or feature request? Let us know! https://kind.sigs.k8s.io/#community 🙂
```

This looks good, however when I run the `kubectl cluster-info` command above I get the following error.

```console
$ kubectl cluster-info --context kind-test
Alias tip: k cluster-info --context kind-test
2022/01/28 14:20:31 tcpproxy: for incoming conn 127.0.0.1:61571, error dialing "192.168.127.2:61448": connect tcp 192.168.127.2:61448: connection was refused
2022/01/28 14:20:32 tcpproxy: for incoming conn 127.0.0.1:61573, error dialing "192.168.127.2:61448": connect tcp 192.168.127.2:61448: connection was refused
2022/01/28 14:20:33 tcpproxy: for incoming conn 127.0.0.1:61575, error dialing "192.168.127.2:61448": connect tcp 192.168.127.2:61448: connection was refused
```

Something is definitely not right with the network connectivity into the podman VM.

Let's try our last task of changing the VM resources. With podman these can only be set during `podman machine init` so we need to delete everything and start from the beginning.

```console
$ podman machine stop

$ podman machine rm

$ podman machine init --cpus 2 --memory 2048 --disk-size 20

$ podman machine start
```

#### Pros/Cons

Podman is heading in a great direction as a deamonless alternative to Docker. But it is exactly that, not Docker. I've not only noticed that there are differences between the behavior of certain commands but also that it needs to be treated specially by `kind`. Given that I work on tools and documentation targeting folks using Docker it would make my like easiest if I am running Docker too.

It managed to map ports, works offline and can reallocate hardware, but I'm going to subtract a point for not being Docker.

2/5.

### Minikube

Now let's look at Minikube. Minikube is a tool to help you set up and run Kubernetes an any operating system for local development. I've already mentioned that I want to use `kind` for my Kubernetes work, but Minikube creates a VM with Docker running and runs it's Kubernetes on that, so we can also use that Docker directly.

#### Installing

```console
$ brew install minikube

$ minikube start --driver=hyperkit
😄  minikube v1.25.1 on Darwin 12.1
✨  Using the hyperkit driver based on user configuration
👍  Starting control plane node minikube in cluster minikube
🔥  Creating hyperkit VM (CPUs=2, Memory=4000MB, Disk=20000MB) ...
🐳  Preparing Kubernetes v1.23.1 on Docker 20.10.12 ...
    ▪ kubelet.housekeeping-interval=5m
    ▪ Generating certificates and keys ...
    ▪ Booting up control plane ...
    ▪ Configuring RBAC rules ...
🔎  Verifying Kubernetes components...
    ▪ Using image gcr.io/k8s-minikube/storage-provisioner:v5
🌟  Enabled addons: storage-provisioner, default-storageclass
🏄  Done! kubectl is now configured to use "minikube" cluster and "default" namespace by default

$ eval $(minikube -p minikube docker-env)
```

Here we start the Minikube VM and use the `docker-env` command to set up our local environment to connect to the Docker in the VM.  We would need to do this every time we open a shell or add this to our `.bashrc` which isn't very convenient. We could manually add a Docker context to use instead. First we need to check the values set by the `minikube docker-env` command and then create a context.

```console
$ minikube -p minikube docker-env
export DOCKER_TLS_VERIFY="1"
export DOCKER_HOST="tcp://192.168.64.3:2376"
export DOCKER_CERT_PATH="/Users/jtomlinson/.minikube/certs"
export MINIKUBE_ACTIVE_DOCKERD="minikube"

# To point your shell to minikube's docker-daemon, run:
# eval $(minikube -p minikube docker-env)

$ docker context create minikube --description "minikube" --docker "host=tcp://192.168.64.2:2376,ca=/Users/jtomlinson/.minikube/certs/ca.pem,cert=/Users/jtomlinson/.minikube/certs/cert.pem,key=/Users/jtomlinson/.minikube/certs/key.pem"

$ docker context use minikube
```

Either way you should now be able to access Docker.

```console
$ docker ps
CONTAINER ID   IMAGE                  COMMAND                  CREATED          STATUS          PORTS     NAMES
19fee5ef6441   6e38f40d628d           "/storage-provisioner"   2 seconds ago    Up 2 seconds              k8s_storage-provisioner_storage-provisioner_kube-system_6bbd1d44-058e-451f-8502-6472c209c3f7_1
8e7032248f94   a4ca41631cc7           "/coredns -conf /etc…"   32 seconds ago   Up 33 seconds             k8s_coredns_coredns-64897985d-b4thw_kube-system_13403e07-ec08-45b5-b378-19da17114f93_0
71de010480fd   b46c42588d51           "/usr/local/bin/kube…"   33 seconds ago   Up 33 seconds             k8s_kube-proxy_kube-proxy-ndq8n_kube-system_3e529d6e-3c8a-4899-a8d5-a10690afed77_0
9ef2a86e03cc   k8s.gcr.io/pause:3.6   "/pause"                 33 seconds ago   Up 33 seconds             k8s_POD_coredns-64897985d-b4thw_kube-system_13403e07-ec08-45b5-b378-19da17114f93_0
6e0513aab64b   k8s.gcr.io/pause:3.6   "/pause"                 33 seconds ago   Up 34 seconds             k8s_POD_kube-proxy-ndq8n_kube-system_3e529d6e-3c8a-4899-a8d5-a10690afed77_0
807bc12d32a5   k8s.gcr.io/pause:3.6   "/pause"                 34 seconds ago   Up 34 seconds             k8s_POD_storage-provisioner_kube-system_6bbd1d44-058e-451f-8502-6472c209c3f7_0
90bd9e84c02c   71d575efe628           "kube-scheduler --au…"   54 seconds ago   Up 54 seconds             k8s_kube-scheduler_kube-scheduler-minikube_kube-system_b8bdc344ff0000e961009344b94de59c_0
c797c4aef3aa   25f8c7f3da61           "etcd --advertise-cl…"   54 seconds ago   Up 54 seconds             k8s_etcd_etcd-minikube_kube-system_784bb9912f58da5bb41429cf74502a65_0
dfb754a80566   f51846a4fd28           "kube-controller-man…"   54 seconds ago   Up 54 seconds             k8s_kube-controller-manager_kube-controller-manager-minikube_kube-system_d3f0dbc1c3a23fddbc9f30b9e08c775e_0
0ec76bb6c0ab   b6d7abedde39           "kube-apiserver --ad…"   54 seconds ago   Up 54 seconds             k8s_kube-apiserver_kube-apiserver-minikube_kube-system_e96765452e79009354d0106c7a1d1e66_0
197ad29faec6   k8s.gcr.io/pause:3.6   "/pause"                 55 seconds ago   Up 54 seconds             k8s_POD_kube-scheduler-minikube_kube-system_b8bdc344ff0000e961009344b94de59c_0
168bd574907d   k8s.gcr.io/pause:3.6   "/pause"                 55 seconds ago   Up 55 seconds             k8s_POD_kube-controller-manager-minikube_kube-system_d3f0dbc1c3a23fddbc9f30b9e08c775e_0
5ee93e0e24a7   k8s.gcr.io/pause:3.6   "/pause"                 55 seconds ago   Up 55 seconds             k8s_POD_kube-apiserver-minikube_kube-system_e96765452e79009354d0106c7a1d1e66_0
ab47d13d0ded   k8s.gcr.io/pause:3.6   "/pause"                 55 seconds ago   Up 55 seconds             k8s_POD_etcd-minikube_kube-system_784bb9912f58da5bb41429cf74502a65_0
```

Note we can see all the containers running in the Minikube Kubernetes cluster too.

#### Testing

Let's try our test container again.

Unfortunately volumes and ports are not made available by Minikube automatically, so we need to open a couple of terminals.

In the first we need to mount our home directory into the Minikube VM.

```console
$ minikube mount $HOME:$HOME
📁  Mounting host path /Users/jtomlinson into VM as /Users/jtomlinson ...
    ▪ Mount type:
    ▪ User ID:      docker
    ▪ Group ID:     docker
    ▪ Version:      9p2000.L
    ▪ Message Size: 262144
    ▪ Options:      map[]
    ▪ Bind Address: 192.168.64.1:62386
🚀  Userspace file server: ufs starting
✅  Successfully mounted /Users/jtomlinson to /Users/jtomlinson

📌  NOTE: This process must stay alive for the mount to be accessible ...
```

Then in the second we need to tunnel traffic into our VM.

```console
$ minikube tunnel
Password:
Status:
        machine: minikube
        pid: 25571
        route: 10.96.0.0/12 -> 192.168.64.3
        minikube: Running
        services: []
    errors:
                minikube: no errors
                router: no errors
                loadbalancer emulator: no errors
```

Now we can run our container.

```console
$ docker run --rm -p 8080:80 -v $HOME/test:/usr/share/nginx/html:ro nginx
```

We should now be able to access our container, but not at `localhost`. Minikube routes traffic through a virtual IP that we can see in the `minikube tunnel` output. So in this case we need to visit `http://192.168.64.3:8080`.

![Our test page showing at the virtual IP](https://i.imgur.com/ptHTHQ4.png)

This also works without WiFi.

Next let's try running `kind`.

Again the cluster seems to create successfully but running the `kubectl` command fails.

```console
$ kubectl cluster-info --context kind-test

To further debug and diagnose cluster problems, use 'kubectl cluster-info dump'.
The connection to the server 192.168.64.3:62939 was refused - did you specify the right host or port?
```

I even tried editing my `~/.kube/config` file and changed the address from `127.0.0.1` to the one provided by the `minikube tunnel` command, but no luck.

Finally let's try changing the resources. Again we need to stop and delete the VM and start again from scratch.

```console
$ minikube stop

$ minikube delete

$ minikube start --driver=hyperkit --memory 8192 --cpus 2 --disk-size=100g
```

#### Pros/cons

Minikube seems like a useful tool for running Kubernetes clusters in VMs on your machine. The fact that it runs Docker inside means that technically you can connect to it and leverage it directly. But as we've seen that experience is less than ideal.

I'll give it half a point for ports and volumes because they work, but painfully.

2/5.

### Run Docker on a virtual machine or remote machine

We could also run Linux ourselves either in a virtual machine or remotely on another system. However both of these options would mean visiting a different IP to access your containers and not being able to mount files from your Mac directly. Using a remote machine also means that you would lose access if you didn't have a decent network connection.

This means that a VM would only get 2/5 and a remove machine only 1/5.

In this case you might as well develop directly on that Linux machine over SSH, it'll make your life easier.

## Conclusion

As of today, three days before Docker Desktop becomes a paid app, it remains unparalleled as the simplest and easiest Docker solution for Mac.

[Colima](https://github.com/abiosoft/colima) came very close with ease of installation and tight binding to the Mac filesystem and ports, however running more complex things like `kind` are not quite working today. I'm optimistic that it will be the best option for me to replace Docker Desktop, but just needs a little polish.
