---
author: Jacob Tomlinson
date: 2016-03-22T00:00:00+00:00
categories:
  - Linux
tags:
  - Intro to Docker
  - Tutorial
thumbnail: docker
title: What is Docker?
aliases:
  - /2016/03/22/what-is-docker/
---


Welcome to the first post in 'Intro to Docker', a series I'm writing to get beginners started with Docker. Each post in this series will be no more than a 5 minute read and will cover exactly one topic surrounding Docker.

## What is Docker?

Firstly Docker is a company. They build tools for creating and managing linux containers (more on those in a minute). The main tool they produce, also commonly referred to simply as Docker, is the [Docker Engine][docker-engine].

The Docker Engine is a service which runs on your linux computer and provides an API for building, running and managing containers. When you install the engine you get a command line application, also called `docker`, which you use to interact with the service API.

Docker also produce other tools to use alongside the Docker Engine. [Kitematic][kitematic] provides a graphical interface to the Engine API. [Compose][docker-compose] simplifies the creation of multiple containers to form a stack. [Swarm][docker-swarm] allows you to run containers across multiple servers transparently using a single API endpoint. And [Machine][docker-machine] is a tool for quickly and easily creating new docker hosts using platforms such as AWS, GCE or even VirtualBox.

## What are containers?

A container is a linux process, or tree of processes, which are running in isolation from the other processes on the system. These processes are isolated using a linux kernel feature called [namespaces][kernel-namespaces].

When a container is created it doesn't have access to anything else on the system other than the linux kernel. Therefore in order for the process to function correctly we need to give it access to a few system resources, primarily the filesystem and the network stack.

Docker can provide a filesystem to each container. This effectively allows you to have a different linux distribution in each container, as all of the system tools and libraries that make up a distribution reside on the filesystem. It also allows you to do clever things with the networking such as multiple containers listening on the same port.

It is easy to confuse containers with virtual machines when you start using them. When you run a container you have something which appears to be a complete independent operating system, like a VM. However your container is just a process on the host operating system, and is accessing system resources via the kernel. This is different to a VM which is running its own kernel and is accessing the host system resources through virtual hardware interfaces. A VM thinks it is speaking to hardware, but in fact it is speaking to the hypervisor software.

## Why use Docker?

A more suitable question would probably be "why use containers?", as most of the benefits I am about to explain are true of any container system. However Docker is currently the most mature and widely used of all the container systems. Here are a few of the many reasons why containers are useful:

### Portability

Containers are static, isolated and repeatable. You can build a container on your laptop and then move it to a server without any noticeable effect. You can also be certain that whether you start 1 or 1,000 containers if they are based on the same image then they will be exactly the same.

### Size
They can also be very small depending on the image you use. Does you application only need a few system libraries? Then you don't need the entire ubuntu filesystem. Using a base image such as [Alpine][alpine] allows you to have containers that are less than 30MB instead of 700MB.

### Dependency isolation

Imagine you want to run two applications which have a conflicting dependency, normally this takes a lot of fiddling to get running, or in some cases can even be impossible. Running them in containers allows you to provide them with different filesystems, each with a different version of the dependency.

### Creation speed

Containers start quickly thanks to a filesystem feature called copy-on-write. This means that if you start 10 containers using the same image you only need one copy of the image on disk. Any changes you make inside your container are stored separately to the base image. If you wish to start another container it runs almost instantly as it doesn't have to make a copy of the filesystem.

## When to not use Docker?

As always you should remain pragmatic when it comes to running containers. There are downsides which may make containers unsuitable for what you are trying to do. Just because it seems like the majority of people are jumping on the container bandwagon doesn't mean you have to try and squeeze everything from your Java application to your Oracle database into a container. Here are some things to be aware of:

### Persistence

Containers are designed to be ephemeral. It is possible to persist data when you stop and start a container, but it doesn't come naturally. This means that containers are excellent for stateless scalable services such as a NodeJS application, however they are not so great for running databases. It is definitely possible to run databases in containers but it you can end up restricting and managing your containers so much that you may as well not use them.

### Complexity

Be aware of the [K.I.S.S principal][kiss]. Complexity can be bad and should usually be avoided. If it is possible to achieve something using a simpler implementation then you should opt for that. Containers can add complexity, and they can reduce complexity. Ensure you understand the problem you are trying to solve and that containers will help you solve it before jumping in with both feet.

## Conclusion

Docker is an excellent tool for producing portable and repeatable linux services. Used in the right situation it can provide you with a flexible, scalable and secure infrastructure. It can also provide you with a consistent and repeatable development environment and a smooth deployment workflow.

The rest of this series will walk through using Docker starting with the very basics!

[alpine]: https://hub.docker.com/_/alpine/
[docker-compose]: https://www.docker.com/products/docker-compose
[docker-engine]: https://docs.docker.com/engine/
[docker-machine]: https://www.docker.com/products/docker-machine
[docker-swarm]: https://www.docker.com/products/docker-swarm
[kernel-namespaces]: http://man7.org/linux/man-pages/man7/namespaces.7.html
[kiss]: https://en.wikipedia.org/wiki/KISS_principle
[kitematic]: https://www.docker.com/products/docker-kitematic
