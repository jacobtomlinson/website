---
title: Deploying opsdroid using ZEIT
date: 2017-09-16T00:00:00+00:00
draft: false
categories:
- blog
tags:
- ospdroid
- chatbots
- deployment
- zeit
author: Jacob Tomlinson
canonical: https://medium.com/opsdroid/deploying-opsdroid-using-zeit-38296888a651
canonical_title: the Opsdroid Blog
---

[ZEIT](https://zeit.co/) is a great platform for deploying your [opsdroid](https://opsdroid.github.io/) instance. Particularly because it is free for light use, which many opsdroid deployments will be.

ZEIT is a platform for deploying node.js or Docker based applications. It provides you with a GUI or command line tool to quickly spin up your projects on their platform. Paid users can enjoy additional features like custom domain names and guaranteed uptime, but for personal applications the free tier is more than suitable.

This article will focus on how to get an opsdroid instance up and running using the ZEIT command line tool called `now`.

## Installing

To get started you need to sign up for a [free ZEIT account](https://zeit.co/login) and install the command line tool.

```console
$ npm install -g now
```

Once you've done this you need to run the command `now login` which will ask you for your email address and send you an email which will walk you through the authentication.

## Configuration

The `now` tool gives you the ability to deploy any `Dockerfile` onto the ZEIT platform. It will build the image automatically for you (you don't need to have docker installed), run the container and return you a url which is proxied to the exposed port on the container.

To get started you need to create your own personal config, so let's start with a new directory for your ZEIT config to live in and a `Dockerfile` within.

_Hint: You can look at the opsdroid [demo config](https://github.com/opsdroid/zeit-demo) for inspiration!_

```console
$ mkdir my-zeit-config
$ cd my-zeit-config
$ touch Dockerfile
```

Now open the `Dockerfile` in your favourite editor and enter the following configuration:

```dockerfile
FROM opsdroid/opsdroid:v0.9.0
EXPOSE 8080
COPY configuration.yaml /root/.opsdroid/configuration.yaml
```

You can see this is a pretty basic `Dockerfile` We are taking the official opsdroid image, exposing port `8080` (the default port for the opsdroid api) and copying a custom configuration file into the image.

We also need to create our custom configuration file. A good starting point for this if you don't have one already is the example config which is bundled with opsdroid. Let's `curl` this config from the GitHub repository into our directory.

```console
$ curl https://raw.githubusercontent.com/opsdroid/opsdroid/v0.9.0/opsdroid/configuration/example_configuration.yaml > configuration.yaml
```

We also need to change the `host` in the web block from `127.0.0.1` to `0.0.0.0` so that the api can be accessed from outside the container.

```yaml
# Web server
web:
  host: '0.0.0.0'
```

## Deployment

Now we have our config in place we can deploy our container to ZEIT. Simply run:

```console
$ now
```

It will prompt you to check you are happy that as you are on the free tier your config and logs will be made available for the world to see (don't panic we'll address this later). Type `y` to continue and watch you container build and deploy!

![](https://i.imgur.com/9q2ssUNh.png)

Your opsdroid container will now be running and `now` will have automatically copied the deployment url to your clipboard. Give it a couple of seconds to get up and running and then test the api with the `curl` command.

```console
$ curl -L <paste your ZEIT url>
{"message": "Welcome to the opsdroid API"}
```

That's it! You now have your very own opsdroid running for free on ZEIT. You can have a look at the config and logs on the ZEIT website if you want to see information about your deployment.

## Domain aliases

When you create a deployment it will give you a new url every time, this isn't ideal if you're running a connector which requires a consistent webhook endpoint like Facebook Messenger. Luckily ZEIT supports aliases so you can give your deployments a consistent name.

```console
$ now alias <paste your ZEIT url> <some new name>
```

This will create a new domain called `<some new name>.now.sh` and alias your deployment to it. Then when you do further deployments you just run the command again with the new deployment url and it will update the alias.

## Secrets and privacy

You might be wondering about keeping sensitive information such as connector API keys private, and the fact that anyone can read your logs.

To help manage your secrets ZEIT allows you to securely store secrets using the `now` tool and it can expose them to your application as environment variables which you can use in your opsdroid config.

```console
$ now secret add my-secret "My super secret information"
```

Then you can either pass that secret to your deployment with the `now -e` [option](https://zeit.co/docs/features/env-and-secrets), or you can create a file called `now.json` and configure them in there.

```json
# Example now.json
{
  "env": {
    "MY_SECRET": "@my-secret"
  }
}
```

You can then use that variable in your opsdroid config.

```yaml
connectors:
  - name: someconnector
    api-key: $MY_SECRET
```

Also if you're worried about conversations leaking out in your logs you can simply turn off console logging in the opsdroid config.

```yaml
logging:
  console: false
```

Or if you still wish to get some logs for debugging you may just want to set the level higher.

```yaml
logging:
  console: true
  level: error
```

Wrap up
=======

That about covers it! You can see how you can quickly and easily deploy your own opsdroid instance on ZEIT for free with minimal effort and maintenance.

If you follow this guide and set up your own instance why not share your experience, you could [earn stickers](https://medium.com/opsdroid/stickers-for-contributors-a0a1f9c30ec1)!
