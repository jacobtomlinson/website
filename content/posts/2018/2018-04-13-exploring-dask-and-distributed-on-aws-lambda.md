---
title: Exploring Dask and Distributed on AWS Lambda
date: 2018-04-13T00:00:00+00:00
draft: false
categories:
- blog
tags:
- dask
- distributed
- AWS
- AWS Lambda
author: Jacob Tomlinson
canonical: https://medium.com/informatics-lab/exploring-dask-and-distributed-on-aws-lambda-55d81d9641d
---

I spent some time this week exploring whether it would be possible to run [Dask](https://dask.pydata.org/en/latest/) and [Distributed](https://distributed.readthedocs.io/en/latest/) on a function as a service platform like [AWS Lambda](https://aws.amazon.com/lambda/).

My hypothesis was that by leveraging even high levels of service abstraction on public cloud platforms we would be able to make Dask and Distributed even more scalable and responsive.

![](https://i.imgur.com/5SgfoeIh.png)

Introduction
------------

Dask is a python library for writing parallel Python code, it allows you to write code which will make full use of all the CPU cores on your computer. It also optimizes calculations by creating a task graph which is only executed when you try to access the results of the calculation, but if you perform further operations such as sub-setting before accessing the result it removes the unnecessary sections of the graph.

Distributed is a library which allows you to run your Dask calculations on a cluster of computers, allowing your code to scale to very large systems. It provides memory sharing, clustering, task management and life cycle management of workers which are connected to a scheduler. Dask graphs are submitted to the scheduler and calculated on the workers.

Currently as part of the [Pangeo](https://github.com/pangeo-data) project we are running Dask and Distributed on [Kubernetes](https://kubernetes.io/). Running it on a platform as a service system like Kubernetes allows us to abstract away from packing workers onto linux servers and think about running them on a homogeneous cluster instead. However the next step up the abstraction ladder is function as a service which would mean we don't even have to worry about cluster capacity or scaling.

My plan
-------

The Distributed scheduler has the ability to add and remove workers depending on its current task queue, with the intention of the queue being emptied in five seconds. So in this case a person running some Dask powered Python code would submit tasks to the scheduler, it would invoke enough Lambda functions to complete all the tasks in the queue in five seconds and then the functions would exit.

To start this work I began looking at how to easily deploy Python code onto Lambda, with the goal of deploying the Dask worker code as a function. Using the [boto3 library](https://boto3.readthedocs.io/en/latest/) I was able to create new Lambda functions and invoke them directly from my Python session.

Generally a Lambda function has a single purpose and you invoke it through an API and it returns a result. However in this case we want to invoke a function which starts a generic worker that requests work from a central scheduler, the scheduler then passes it a task which it completes. It continues requesting tasks until none are left and then it exits.

Reality
-------

Given the short amount of time I gave myself to look at this I was unable to get a working demo, and I feel that even if I persevered it would be severely limited for a number of reasons.

When running a worker you pass it information about where to find the scheduler and how the scheduler can find it. As lambda functions are designed to be ephemeral, short lived and highly scalable there are limitations on how to access it. In theory a lambda function doesn't have an IP address and isn't able to bind ports and so it is not possible for the worker to tell the scheduler how to address it. Ports could be bound onto a remote server using ssh remote port forwarding or similar techniques but this would introduce a bottleneck into the system.

Workers store state in memory. Often this is only for the duration of a calculation but eventually at the end of the calculation there will be an answer which has to be held somewhere. Currently Distributed will remove all but one worker at the end of a calculation and that remaining worker will store the result. However as Lambda functions have a maximum lifespan of five minutes this isn't practical for storing results. Therefore we would need a small pool of persistent workers running on a platform other than Lambda which can scale out using Lambda when necessary.

The worker needs to have all of the same Python libraries available to it, and at the same versions, as the Python process which submitted the tasks to the scheduler. This is because the graph is run by pickling functions and sending them to the workers to be executed, and so the function can't be un-pickled if the libraries used to make it are not there. According to the AWS [documentation](https://docs.aws.amazon.com/lambda/latest/dg/limits.html#w135aac55b9c19) when creating a lambda function you can only pass 50MB of compressed code which can not expand to more than 250MB. The [Docker image](https://github.com/informatics-lab/singleuser-notebook) we are using on our Kubernetes cluster to power our workers is over 2GB. We would not need to put everything from the image into our Lambda function but it will certainly be more than 50MB.

Future
------

I have some thoughts on how we can tackle these challenges, but it is by no means a small undertaking.

We could create bespoke Lambda functions on the fly using the Dask graph. For each unique operation that is in the graph we could use the Lambda API to create a new function which performs that task. We could also calculate which dependencies are needed and pass them in the code payload. Then the Distributed scheduler could invoke those directly to calculate those tasks, rather than having generic workers consuming from a queue.

We could separate the state from the calculations. Using a service like [ElastiCache](https://aws.amazon.com/elasticache/) the workers could store their memory somewhere else, this would allow workers to come and go without losing memory.

Conclusion
----------

I feel like this work is heading in the right direction and that the future of distributed computation will make use of highly abstracted function execution services like Lambda. However with the tools we have today there are too many issues to get this working without having major limitations.

But just because this isn't a reality today doesn't mean it can't be in the near future...
