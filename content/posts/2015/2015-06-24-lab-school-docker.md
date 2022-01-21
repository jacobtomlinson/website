---
title: 'Lab School: Docker'
date: 2015-06-24T00:00:00+00:00
draft: false
author: Jacob Tomlinson
categories:
- blog
tags:
- docker
canonical: https://archived.informaticslab.co.uk/lab-school/2015/06/24/lab-school-docker.html
canonical_title: the Informatics Lab Blog
thumbnail: docker
---

_Originally published on the [Met Office Informatics Lab blog](https://archived.informaticslab.co.uk/lab-school/2015/06/24/lab-school-docker.html) on June 24th, 2015._

Welcome to the first ever [Lab School][lab-school] session. This session aims to give you an overview of docker and how we are currently using it in the Lab.

## What is Docker?

When people talk about [docker][docker] they are most likely talking about the `docker engine`. Docker is in fact an organization which creates tools to assist with the creation and deployment of software in a portable manor. As well as creating the docker engine they also have projects which orchestrates and manages deployed applications, we will come across more of these later.

### Docker engine

The `docker engine` (which we will just refer to as `docker` from now on) is a piece of software which allows a sort of operating system level virtualization on linux based computers. This is called containerization and allows you to run applications within a container, which is a portable unit and can be moved between linux machines.

### Containers

The biggest misconception I have found when it comes to containers is that people think they are just like virtual machines. They aim to achieve the same outcomes as a virtual machine including portability, speed and flexibility but they do it in a different way.

All linux systems are fundamentally the same, they are all running the linux kernel. The differences are the tools which are layered on top of the kernel. If you think about it these tools are just files in a filesystem which are executed in a certain order and interact with the kernel to achieve different goals.

Docker works by allowing you to mount multiple filesystems on one machine and then run applications within the context of the different filesystems. For example you could create a ubuntu based container which is running on a Red Hat server. If you run the `bash` tool within the ubuntu container you can make use of all the standard ubuntu tools, for example `apt-get`. Those tools will interact with the kernel run by the Red Hat host and function as if you were on a ubuntu machine.

Docker makes use of linux kernel functionality such as `cgroups`, `selinux` and `chroot jails` to keep the processes separate and contained within their containers.

### The differences from a virtual machine

You may be thinking "this sounds a lot like virtualization" and you're right, however the container will not do many of the things you would expect a virtual machine to do automatically.

Containers run exactly one process, no more, no less. So if you execute a command within a ubuntu container that command will act as if it is running on a ubuntu machine. However if that command relies on operating system services to be running, such as cron, it will not be able to find them. The host operating system may be running these services but the container is not.

To work around this you simply need more containers. This is the beauty of containerization, you can run a database service on ubuntu, which is accessed by a Java application running on Red Hat, which is backed up by a cron job running on debian, all on a host server running SUSE. The important thing to realise is that in this example SUSE is the only "operating system" actually running and taking up resources, the others are just running single processes but giving those processes access to the tools of alternate operating systems.

_You can also specify particular versions of operating systems, so you could run a container based on ubuntu 13.10 on a ubuntu 14.04 server without worrying about compatibility issues._

Hopefully you can see now why this is exciting. It gives you a level playing field. You can develop an application on your desktop with 100% confidence it will behave exactly the same way on a production server.

## Exercises

You will now work through some exercises which will help you get to grips with docker and also experience how we are using them in the Lab.

### Requirements
The following exercises are designed to be run on an [AWS EC2][aws-ec2] instance. If you are taking part in the session live then please request an instance from one of the Lab core members, if you are taking part afterwards please sign up for an [AWS free account][aws-free] and create an EC2 micro instance based on the "Amazon Linux" image. [This guide][ec2-guide] should get you started, just follow it up to the "Configuring your account" section.

## Exercise 1: Installing docker
Before we do anything we need docker. To get started you'll need to ssh on to your instance. It's always good practice to run an update when you create an instance.

```
sudo yum update -y
```

Once this has finished we're ready to install docker. Thankfully it is available in the Amazon linux repositories.

```
sudo yum install docker -y
```

Hooray you now have docker! All you need to do now is set it to start on boot (we might reboot our instance later) and then start it for the first time.

```
sudo chkconfig docker on
sudo service docker start
```

To test docker you should be able to run:

```
sudo docker -v
```

which should print out the version and build.

```
docker version 1.6.2, build 7c8fca2/1.6.2
```

To avoid having to run docker commands using sudo we can add our `ec2-user` to the `docker` group.

```
sudo usermod -G docker ec2-user
```

_You will need to log out and back in again from your ssh session for this to take effect._

## Exercise 2: My first container

Now that you have docker we can go ahead and create our first container. Docker containers are just commands run within the context of an alternative file system. These alternative filesystems are called images. You can build your own images but for now we are going to use an off-the-shelf one. Docker provides a place to store images called the [Docker Hub][docker-hub]. If you create a container which doesn't have an image locally, but it finds one on the hub, it will download it for you.

Let's get started by creating a simple installation of an apache web server. We are going to download the apache image from the hub and create a container to run it.

```
docker run -p 80:80 httpd
```

In this example we want docker to use the default `httpd` image.

_Images are usually named following the convention of `author/image`, however we are using an official docker image which means we can omit the author._

We are also telling docker to link port `80` on the container to port `80` on the host. This means we can visit our instance in our web browser and we should see the "It works!" apache default page.

![Apache test page](https://i.imgur.com/Mn5vDbr.png)

If you see the same page in your browser then you've successfully created an apache container.

_You should also see some log output in your command line._

## Exercise 3: Extending an existing image

Creating an apache instance is all well and good but what if we actually want it to serve our own content? To achieve this we need to take the default apache image and add our own files to it.

We are going to create a new directory to contain this project.

```
mkdir apache-app
cd apache-app
```

Now we can create our `Dockerfile`. I like to use `vi` for text editing but you may be more comfortable with `nano`.

```
nano Dockerfile
```

Set the contents of your `Dockerfile` to look like this, but set yourself as the maintainer of course.

```
FROM httpd:2.4
MAINTAINER Jacob Tomlinson <jacob.tomlinson@informaticslab.co.uk>

COPY ./my-html/ /usr/local/apache2/htdocs/

EXPOSE 80

CMD ["httpd-foreground"]
```

Going through this line by line we see that we are inheriting this image `FROM` `httpd`. This means we want docker to download the httpd image as before but do some stuff to it afterwards. We are also explicitly stating the version of the image (which directly maps to the version of apache) we want rather than just going with the latest one.

_The `httpd` image is really just the result of another Dockerfile which you can [find on GitHub](https://github.com/docker-library/httpd/blob/63cd0ad57a12c76ff70d0f501f6c2f1580fa40f5/2.4/Dockerfile)._

We are setting ourselves as the `MAINTAINER` so when someone else comes along and wants to use our image they know who to bother with their problems.

Next we are going to `COPY` the contents of a directory called `my-html` into the image and place them in apache's default content location, which for 2.4 is `/usr/local/apache2/htdocs/`.

We want to `EXPOSE` port `80` to allow it to be accessed outside this container, think of it like a software firewall within docker.

Then we specify our command that we want to run when the container is started with `CMD`.

_The last two commands are actually already defined within the httpd image, I wanted to put them in here to show what is happening but to also show that you can redefine a `CMD`. Docker will run the last `CMD` to be defined which is useful if you want to override the functionality of an image._

You may notice that the `my-html` directory doesn't exist yet. Let's create it.

```
mkdir my-html
```

Then let's create an `index.html` file within it.

```
nano my-html/index.html
```

Set the content to something along the lines of:

```
<html>
  <head>
    <title>Hello world!</title>
  </head>
  <body>
    <h1>Hello world!</h1>
    <hr>
    <p>Docker is the best!</p>
  </body>
</html>
```

Now we can build our image.

```
docker build -t my-httpd .
```

See we've specified the image name with `-t`. This is the name we will use to run it.

```
docker run -p 80:80 my-httpd
```

Excellent, now when you navigate to your EC2 instance in your browser you should see  your lovely new index page.

![Custom apache test page](https://i.imgur.com/1PzBUyL.png)

## Exercise 4: Docker Compose

In the Lab we have progressed to the stage of needing to run multiple containers at once which make up a service. If we just used pure docker we would have to run each one individually and remember to specify the correct arguments (like `-p`) each time. We could create a bash file which contains all of these lines but there is a better way, [docker compose][docker-compose].

Compose is another tool provided by docker and it allows you to write down your structure of containers and all of their arguments in a [yaml][yaml] file. You can then simply run `docker-compose up` and your containers will be created.

To get started we need to install docker compose.

```
sudo pip install docker-compose
```

Again to test that it is installed correctly we run:

```
docker-compose -v

```

Which should print the version of docker-compose, cpython and openssl.

```
docker-compose version: 1.3.1
CPython version: 2.7.9
OpenSSL version: OpenSSL 1.0.1k-fips 8 Jan 2015
```

Let's start simply by creating a compose file for our apache application. To do this we just need to create a new file called `docker-compose.yml`

```
nano docker-compose.yml
```

In this file we can put:

```
apache:
  build: .
  ports:
   - "80:80"
```

This file defines a new service called `apache`, it tells docker that to build this service it will find the Dockerfile in the current directory `.` and that we want to again bind port 80 on the container to 80 on the host.

Now we can run our application.

```
docker-compose up
```

You should see similar output to when running docker manually but the log outputs will be prefixed with the container name. When docker compose creates a container it names it after the service followed by an underscore and the number one. This is because you can scale containers with compose easily and it will create additional containers with incrementing numbers.

## Exercise 5: Multiple containers in compose

Now that we are comfortable creating a container with compose let's create a second container and link them together. Instead of using our apache container we're going to create a [python django][django] application running with a `postgres` database.

Let's make a new project directory and switch to it.

```
mkdir ~/django-app
cd ~/django-app
```

We'll need a Dockerfile for our django app.

```
nano Dockerfile
```

With the following contents:

```
FROM python:2.7

ENV PYTHONUNBUFFERED 1

RUN mkdir /code

WORKDIR /code

COPY src/requirements.txt /code/

RUN pip install -r requirements.txt

COPY src/ /code/
```

Here we are using `python:2.7` as our base image. This is another official image which will ensure we have python installed and at version `2.7`.

We have our first use of `ENV`, this sets an environment variable within the container. In this case we want our python to be unbuffered which will help with our output in the docker console.

Next we `RUN` a command to create a directory for our django app to live in on our container. We then use `WORKDIR` to switch the current working directory within our container.

We `COPY` in a `requirements.txt` file which specifies our python dependancies. We'll create this file in a minute. Then we are using `pip` the python package manager to install those dependancies.

Finally we want to copy the contents of our `src` directory on the host to the `code` directory on the container.

Now we want to create our `requirements.txt` file for docker to use.

```
mkdir src
nano src/requirements.txt
```

Add the following lines:

```
Django
psycopg2
```

Next we are going to define our `docker-compose.yml` file.

```
nano docker-compose.yml
```

Add the contents:

```
db:
  image: postgres
web:
  build: .
  command: python manage.py runserver 0.0.0.0:8000
  volumes:
    - src/:/code/
  ports:
    - "80:8000"
  links:
    - db
```

Firstly we are creating a database service called `db`. We just want a plain old `postgres` server so we don't need a Dockerfile for it, we can reference the image name from the Docker Hub directly. As we will be accessing this service from another container we don't need to expose any ports to the outside world.

Then we create a `web` service for our django application. We are again going to be building the Dockerfile in the current directory, we are also specifying the command for the container to run. This can be done in the Dockerfile but you can also do it here.

As well as copying the contents of our `src` directory into the container on build we are going to mount it as a volume which means when the container makes changes to those files they will be persisted on the host even if the container is destroyed and rebuilt.

%This is actually a bad practice, see "data only containers" in the further reading section below for more information.%

We are linking our ports again but django runs on port `8000` by default so this time we are going to link port `80` on the host to `8000` on the container.

Finally we are going to tell our container to link with the db container. This means they will have network access between each other and it also sets up a local dns record (in `/etc/hosts`) on the web container so we can access the database at the hostname `db`.

Before we can start our web service django needs you to initialise its project and also tell it where the database is. We can do this with the docker compose `run` command, this runs the container but executes a different command to the one specified in the yaml file.

```
docker-compose run web django-admin.py startproject composeexample .
```

When running this docker will discover your images are not built, it will automatically download and build them for you and then move on to execute the command.

When this command finishes it should generate a new directory in your `src` directory called `composeexample` along with a `manage.py` file. Check that they are there.

```
ls src
```

The `composeexample` directory will contain a `settings.py` file. This is where you need to put the database configuration.

Open the file for editing, docker will have created these files as root so you'll need a `sudo` on this one.

```
sudo nano src/composeexample/settings.py
```

Then update the `DATABASES = ...` declaration to look like this:

```
DATABASES = {
    'default': {
        'ENGINE': 'django.db.backends.postgresql_psycopg2',
        'NAME': 'postgres',
        'USER': 'postgres',
        'HOST': 'db',
        'PORT': 5432,
    }
}
```

You can see we are just selecting `postgres` as our database engine and then pointing it at host `db` which, thanks to our nifty DNS settings, will point to the `db` container.

Now we can try and run the django command `syncdb` which will take your django models and update the database to reflect them.

```
docker-compose run web python manage.py syncdb
```

_If it asks you about a `superuser` just say no, we're not worried about what this command is doing as it will be lost next time we run docker-compose anyway. We're just testing the database connection._

Finally you can start the containers.

```
docker-compose up
```

Refresh your EC2 instance's page in your browser and you should now see the default django test page.

![Django test page](https://i.imgur.com/IFC3Yf9.png)

## Conclusion

Congratulations, you are now working with docker. Hopefully you can see the power provided by adding this little bit of overhead to your applications. We've only scratched the surface here so I suggest you read as much as you can about docker, starting with the links below.

#### Further reading
  * [Data only containers][data-only-containers] - Persistent data storage done right
  * [Swarm, machine and compose][docker-orchestration] - Building and clustering docker servers
  * [Kubernetes][kubernetes] - Docker orchestration platform by Google

#### References
  * Header image by [Kevin Talec][kevin-talek], licensed under [CC BY-SA 2.0][cc-by-sa-20].
  * Exercises 2 & 3 are based on the examples in the [Docker Hub httpd repository][docker-hub-httpd].
  * Exercise 5 is an expanded version of the [docker compose django guide][docker-compose-django].

[aws-ec2]: http://aws.amazon.com/ec2
[aws-free]: http://aws.amazon.com/free/
[cc-by-sa-20]: https://creativecommons.org/licenses/by-sa/2.0/
[data-only-containers]: http://container42.com/2013/12/16/persistent-volumes-with-docker-container-as-volume-pattern/
[django]: https://www.djangoproject.com/
[docker]: https://www.docker.com/
[docker-compose]: https://docs.docker.com/compose/
[docker-compose-django]: https://docs.docker.com/compose/django/
[docker-hub]: https://hub.docker.com/
[docker-hub-httpd]: https://registry.hub.docker.com/_/httpd/
[docker-orchestration]: https://blog.docker.com/2015/02/orchestrating-docker-with-machine-swarm-and-compose/
[ec2-guide]: http://www.crmarsh.com/aws/
[httpd-dockerfile]: https://github.com/docker-library/httpd/blob/63cd0ad57a12c76ff70d0f501f6c2f1580fa40f5/2.4/Dockerfile
[kevin-talek]: https://www.flickr.com/photos/kevtalec/
[kubernetes]: http://kubernetes.io/
[lab-school]: http://www.informaticslab.co.uk/announcements/2015/06/22/introducing-lab-school.html
[yaml]: http://yaml.org/
