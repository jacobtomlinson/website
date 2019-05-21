---
author: Jacob Tomlinson
date: 2016-03-29T00:00:00+00:00
categories:
  - Linux
tags:
  - Intro to Docker
  - Tutorial
thumbnail: docker
title: Running a Docker container
url: /2016/03/29/running-a-docker-container/
---


## Installing Docker

Installing Docker on your machine is required but beyond the scope of this series. Getting Docker up and running is an ever evolving and improving process and anything put here will go stale reasonably quickly. As Docker uses linux kernel features you will need a running linux operating system. Therefore installing Docker on linux is easy, however installing it on Windows and Mac involve running a lightweight linux virtual machine.

For up to date instructions on how to get the Docker Engine running on your machine visit the [official documentation][install-docker].

## Checking your install

This series will focus on the command line interface for Docker. Before we get started you need to make sure you can access the Docker service. If you followed one of the guides above it should leave you with a working command line application.

To test this go to your terminal and type `docker info`. If you get output similar to the following you're good to go. Otherwise go back to your install guides and make sure you have everything set up correctly.

```
$ docker info
Containers: 1
Images: 76
Server Version: 1.9.1
Storage Driver: aufs
 Root Dir: /mnt/sda1/var/lib/docker/aufs
 Backing Filesystem: extfs
 Dirs: 78
 Dirperm1 Supported: true
Execution Driver: native-0.2
Logging Driver: json-file
Kernel Version: 4.1.13-boot2docker
Operating System: Boot2Docker 1.9.1 (TCL 6.4.1); master : cef800b - Fri Nov 20 19:33:59 UTC 2015
CPUs: 1
Total Memory: 996.2 MiB
Name: default
ID: OOYR:ZKOI:OJSC:QBSS:2DUE:MGPC:7KDU:2J24:QZVM:CS4W:7Z6F:3UFQ
Debug mode (server): true
 File Descriptors: 13
 Goroutines: 18
 System Time: 2016-03-29T13:49:59.080635418Z
 EventsListeners: 0
 Init SHA1:
 Init Path: /usr/local/bin/docker
 Docker Root Dir: /mnt/sda1/var/lib/docker
Username: jacobtomlinson
Registry: https://index.docker.io/v1/
Labels:
 provider=virtualbox
```

## Running your first container

As mentioned in the [What is docker?][what-is-docker] post a container is simply a process running in an alternate kernel namespace with a separate filesystem and network interface. Docker manages all of this for us, all we need is the filesystem and the command we want to run inside the container.

Docker also handles this for us in the form of images. An image is effectively a tar archive with a bit of associated meta data. The Docker cli tool can create, push and pull these images from centralized repositories. More on that later.

For now we will just pull and run an image from the default repository called [Docker Hub][docker-hub].

## Hello world

Run the following in your terminal

```
docker run --rm hello-world
```

This command will give two lots of output. First is the image being pulled from the hub.

```
Unable to find image 'hello-world:latest' locally
latest: Pulling from library/hello-world
b901d36b6f2f: Pull complete
0a6ba66e537a: Pull complete
Digest: sha256:8be990ef2aeb16dbcb9271ddfe2610fa6658d13f6dfb8bc72074cc1ca36966a7
Status: Downloaded newer image for hello-world:latest
```

Next is the output from the container itself.

```
Hello from Docker.
This message shows that your installation appears to be working correctly.

To generate this message, Docker took the following steps:
 1. The Docker client contacted the Docker daemon.
 2. The Docker daemon pulled the "hello-world" image from the Docker Hub.
 3. The Docker daemon created a new container from that image which runs the
    executable that produces the output you are currently reading.
 4. The Docker daemon streamed that output to the Docker client, which sent it
    to your terminal.

To try something more ambitious, you can run an Ubuntu container with:
 $ docker run -it ubuntu bash

Share images, automate workflows, and more with a free Docker Hub account:
 https://hub.docker.com

For more examples and ideas, visit:
 https://docs.docker.com/userguide/
```

You may also notice we added a `--rm` to the command, this means the container will remove itself when it finishes executing. This isn't always necessary but while we are learning it makes sense to mop up after ourselves.

## Conclusion

Congratulations! You have Docker all set up and have run your first container. In the next post we will explore a little deeper into what a container actually is.

[docker-hub]: https://hub.docker.com/
[install-docker]: https://docs.docker.com/engine/installation/
[what-is-docker]: /linux/2016/03/22/what-is-docker/
