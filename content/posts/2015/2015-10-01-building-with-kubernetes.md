---
title: Building with Kubernetes
date: 2015-10-01T00:00:00+00:00
draft: false
tags:
- infrastructure
categories:
- blog
author: Jacob Tomlinson
thumbnail: kubernetes
canonical: http://archived.informaticslab.co.uk/infrastructure/2015/10/01/building-with-kubernetes.html
canonical_title: the Informatics Lab Blog
---

_Originally published on the [Met Office Informatics Lab blog](http://archived.informaticslab.co.uk/infrastructure/2015/10/01/building-with-kubernetes.html) on October 1st, 2015._

For our 3D visualisation project we wanted to build a data processing service using [Docker containers][lab-school-docker]. We quickly found that once you are running more than a couple of containers you need a way to manage them. After looking into the different tools available we decided to give [Kubernetes][kubernetes] a go, this is what we learned.

## What is Kubernetes?

[Kubernetes][kubernetes] is an open source Docker orchestration tool created by Google. It oversees your cluster of Docker hosts and starts, stops and moves your containers around for maximum efficiency and resilience.

Not only does it scale and migrate your containers around but it also runs proxying services to ensure they can always talk to each other and the outside world. It can even manage the underlying infrastructure depending on the platform you select.

## Managing a cluster

Building a cluster is fairly straight forward if you use the [simple one-line commands][kube-getting-started] available. These will get you up and running on a platform of your choice. You specify what type of servers you would like it to use ([EC2 m4.large instances][aws-ec2-m4] for example) and how many docker hosts you want to have.

Building a cluster using the AWS tools gives you an auto-scaling group for your docker hosts, therefore if a host is removed a new one is created and automatically added to the cluster. This is fantastic because if something odd happens to a host you can just terminate it. Kubernetes will migrate the containers to a different host and AWS will replace it with a shiny new one. This is also very useful for security patching your hosts as you can just replace them one by one and the replacements will install the latest security updates as part of their build process. As you replace hosts Kubernetes will seamlessly migrate your containers around the cluster which may temporarily decrease your capacity but your service will be fine and your users will be none the wiser.

However once your cluster is up and running this way it isn't totally clear how to manage it yourself. The only downside I've found with using Kubernetes is that troubleshooting and debugging problems related to the cluster itself is very non-trivial. If something goes wrong with the provisioning of the hosts themselves then it becomes difficult to remedy due to the opaque nature of the "getting started" scripts. This has lead me to rebuilding my entire cluster more than once to try and solve a small problem.

It is worth bearing in mind that although this project is incredibly powerful it is still very young and not as straight forward as it will be after a few more version increments. There are also other projects out there like [OpenShift Origin][openshift-origin] or [Mesosphere][mesosphere] which are wrapping Kubernetes up in a more user friendly and enterprise ready way.

## Deploying containers

The really exciting thing about Kubernetes is being able to define your application in code. This works by creating a collection of YAML manifests which describe what containers you want to have, where to get the images from, and how they should connect together. You never tell Kubernetes to create anything, you just tell it that 'x' number of this thing should exist and it figures out the best way to make that happen.

In Kubernetes you have the concept of a pod, this is a collection of containers which are treated as one entity. For the most part you can think of one of your containers as one pod, as you will most likely only put one container in the pod. Kubernetes will also add a few service level containers to the pod to handle the lifespan and proxying of your container. These service level containers are hidden from you unless you really want to see them, but it is important to know they exist to understand why you need a pod.

If you were to deploy a pod on a Docker host and that host goes down the pod goes with it. To counter this Kubernetes has a construct which allows you to define how many of each pod you want to have running at any given time. Therefore if you tell Kubernetes you want one pod it will start it up for you and then if the pod dies for any reason Kubernetes will replace it on a different host. This also allows you to easily scale your application up and down by changing the required quantity of each pod.

To expose your applications you create service manifests which define how these containers should be opened to the internet. You can think of these as a cross between a firewall and a load balancer. If you are using a service which has software defined load balancers, like Google Compute Engine (GCE) or Amazon Web Services (AWS), Kubernetes will automatically create one for each service and configure them accordingly.

## Thinking differently

Containerization has already started making us think differently about how to architect our applications with concepts like micro services, statelessness and scalability. These concepts allow you to implement principals such as ["fail fast, fail cheap"][fail-fast-fail-cheap] and ["treat your servers like cattle, not pets"][cattle-not-pets].

Kubernetes pushes these concepts even further by forcing you to apply them from the outset. It will start and stop your containers of its own accord to be most efficient. This means that if your containers are not stateless then your application will not be happy about being killed and replaced. Designing your applications with this in mind is great because if your containers are killed for reasons outside of Kubernetes' control (e.g a docker host being manually terminated for security patching) it will just recreate the container on a different host as it would have done if it were moving the container itself. Therefore you end up with a system which is resilient and self healing by default.

You also stop thinking about your services as long running processes. So if you have a background process which takes a long time to run and Kubernetes kills it half way through that would be inconvenient. However if you design it in a way where it does a little work, exits, gets restarted, does a little more, exits, and so on then if your process is killed or moved you don't mind.

An example of where we have implemented this is in [converting weather data files into images][data-encoding] where we split the data into small chunks that we store in a message queue. The messages can be processed in parallel by a number of containers which remove the job from the queue and exit when they have completed it. Therefore if a container fails we don't care because the failed job will timeout and appear back on the queue and the next time a processing container is started it will take over that job. This is possible because Kubernetes will always ensure there are a set number of processing containers running, regardless of whether they exited successfully or not. We can also scale the number of containers based on the length of the queue so if a lot of data comes in at once (which is common in weather forecasting) then the infrastructure will scale up to process it quickly, but then scale back down when they are not required to reduce costs.

## Conclusion

The current state of Kubernetes is hard to deploy. You can use the quick start scripts but you'll probably end up in a mess sooner or later, and unless you are an experienced system administrator you'll find building a cluster from scratch an uphill struggle. Kubernetes is designed for internal use at Google, therefore it works wonders for large workloads with dedicated support teams.

That being said, in my eyes Kubernetes is the start of the next generation of IT infrastructure. It abstracts the clustering and resilience logic out of your applications, allowing them to focus on their job and not on managing their own lifespans. This is a dream for developers who want to build small, fast moving applications components and release them easily without worrying about servers and infrastructure.


[aws-ec2-m4]: https://aws.amazon.com/ec2/instance-types/#M3
[cattle-not-pets]: http://www.lauradhamilton.com/servers-pets-versus-cattle
[data-encoding]: http://www.informaticslab.co.uk/technical/2015/10/05/data-encoding.html
[fail-fast-fail-cheap]: http://blog.flux7.com/3-strategies-to-fail-fast-fail-cheap-get-agile
[kubernetes]: http://kubernetes.io/
[kube-getting-started]: http://kubernetes.io/docs/user-guide/walkthrough/
[lab-school-docker]: http://www.informaticslab.co.uk/lab-school/2015/06/24/lab-school-docker.html
[mesosphere]: https://mesosphere.com/
[openshift-origin]: http://www.openshift.org/
