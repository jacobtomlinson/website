---
author: Jacob Tomlinson
date: 2016-03-31T00:00:00+00:00
categories:
  - Linux
tags:
  - Intro to Docker
  - Tutorial
thumbnail: docker
title: Interactive Docker containers
aliases:
  - /2016/03/31/interactive-docker-containers/
---


## Interactive shell

In this post we are going to create a container we can interact with. We can then have a poke around inside the container and see what it is and how it works.

Open your terminal and run:

```
docker run --rm -ti ubuntu /bin/bash
```

This will download the ubuntu image from Docker Hub and start a new container form it. There are a couple of new arguments here that we haven't seen yet. `-t` tells docker to assign a pseudo tty and `-i` keeps `stdin` open to allow us to interact with the shell. We have also added the command `/bin/bash` to the end, this overrides the default command set in the image.

Once the container runs you should get something similar to:

```
root@4e8a1c232723:/#
```

The hostname of your container will be different but you should see that you are running a bash session as root within your container.

Seeing a prompt like this feels very much like being inside a virtual machine. However as this is a container you are sharing the kernel with the host system and any other containers you run simultaneously.

## Filesystem

If you type `ls` it will list the root filesystem of the container, this is the filesystem which came inside the Docker image and your bash process is contained inside it.

```
# ls
bin   dev  home  lib64  mnt  proc  run   srv  tmp  var
boot  etc  lib   media  opt  root  sbin  sys  usr
```

You can move around the filesystem, edit files and execute commands. However if you create or modify a file and then exit from the container those changes will not be persisted. This is because each instance of a container is an ephemeral copy of the base image. Using a technology called copy-on-write Docker doesn't actually have to copy the base image for each container. It stores all changes separately to the base image and merges them together cleverly when the containers accesses the filesystem. When the container exits these changes are discarded.

It is possible to persist data between containers using volumes, these will be covered in a later post.

You can also find the container images themselves stored under `/var/lib/docker` on the host system. They are just directories like any other but Docker bind-mounts these to create the layered copy-on-write filesystem and uses `chroot` to contain the processes within the image.

## Process namespace

If you list the processes running on the system you will only see two, the bash session and the `ps` itself.

```
# ps aux
USER       PID %CPU %MEM    VSZ   RSS START   TIME COMMAND
root         1  0.0  0.3  18180  3340 14:14   0:00 /bin/bash
root        20  0.0  0.2  15572  2080 14:45   0:00 ps aux
```

This is different to a virtual machine in may ways. For starters your bash session has `PID 1`. On all linux systems `PID 1` is the init system, that may be `init.d`, `systemd`, `upstart` or something else. An init system handles the starting of all the other processes required to run your operating system, this includes networking, cron, disk mounts, graphical user interfaces and everything else you use. However in a container we already have a functioning operating system, the host. The container can therefore do one thing and one thing only without worrying about running an operating system, so having a single process with `PID 1` is logical.

If you were to run the `ps aux` command on the linux host you would see all the running processes, including those inside the container.

```
$ ps aux
USER       PID %CPU %MEM    VSZ   RSS START   TIME COMMAND
root         1  0.1  0.1   9764  1736 13:49   0:05 init
...
root      1522  0.0  0.3  18180  3340 14:14   0:00 /bin/bash
...
docker    2028  0.0  0.1  13016  1740 14:45   0:00 ps aux
```

Here we can see the true `PID 1` of the system, we can also see the same `/bin/bash` process from our container but with `PID 1522` instead of `1`. The reason our container thinks the process is `PID 1` is due to kernel namespaces. The process is in both the main namespace, with `PID 1522`, and in the container namespace with `PID 1`. This is the containment part of containers, when we run `ps` within the container it is restricted to the container namespace and can only see the other processes inside the container.

## Conclusion

Containers appear to be a whole different system, but we've demonstrated here that they are just normal linux processes running on a normal linux machine. The only difference is they are kept separate from the other things running on the system and have access to a subset of resources. The exciting thing is each container can have a different subset of resources. They can behave like an entirely separate system with different tools and dependancies even if they conflict with the packages on the host system or in other containers.

[docker-hub]: https://hub.docker.com/
[install-docker]: https://docs.docker.com/engine/installation/
[what-is-docker]: /linux/2016/03/22/what-is-docker/
