---
title: "How to enable SSH on Binder"
date: 2020-06-17T00:00:00+00:00
draft: false
author: "Jacob Tomlinson"
categories:
  - blog
tags:
  - Binder
  - SSH
thumbnail: mybinder
---

__⚠️ This post is no longer valid.__

Running SSH on Binder has not been possible since late 2020. Due to abuse from botnets Binder will now kill sessions running `sshd`.

<div style="text-align: center;">

![This is why we can't have nice things](https://i.imgur.com/mY4jZkS.png)

</div>

---

_Original post for archive purposes_

While preparing material for a [Dask tutorial](https://github.com/jacobtomlinson/dask-video-tutorial-2020) I wanted to be able to ssh to `localhost` on [Binder](https://mybinder.org/). This was to allow me to demonstrate the SSH features in Dask, while remaining in my Binder instance.

I can imagine other situations where being able to SSH to `localhost` on Binder would be useful so thought I would write up some instructions.

## Unprivileged SSH Daemon

When you run a binder repo you end up as the user `jovyan` inside a Docker container built from your git repo using [repo2docker](https://github.com/jupyter/repo2docker). By default you will not have `sudo` available and will not be able to run SSH regularly as root.

We can however run the SSH daemon as an unprivileged user with some configuration modifications.

## Installing openssh-server

Before we do anything we need to ensure that `openssh-server` is installed. You can do this by creating an `apt.txt` file within your `binder` directory with the name of the package in.

```
openssh-server
```

## Bootstrapping SSH

Next we need to run some commands when our binder server starts. We can do this by creating a bash script at `binder/start`.

```bash
#!/bin/bash

# Bootstrap commands go here

exec "$@"
```

This is a minimal start script which will ensure Jupyter starts successfully. In the next sections we will go through each command that you need to add to the script, but don't worry if you get lost while following along as we have the complete script at the end.

### Set up SSH client config

The first thing we want to do is create a `.ssh` directory for our `jovyan` user as we won't already have one.

```bash
mkdir ~/.ssh && chmod 700 ~/.ssh
```

Then we need to generate an SSH key for our `jovyan` user to use.

```bash
ssh-keygen -t rsa -f .ssh/id_rsa -N ''
```

We then need to also allow incoming SSH connections to authenticate with that key. This allows us to SSH to ourself at `localhost` without entering a password.

```bash
cat .ssh/id_rsa.pub > .ssh/authorized_keys
chmod 600 .ssh/authorized_keys
```

### Set up SSH server config

An SSH server requires a key pair that it uses to prove that it is who it says it is. It avoids man-in-the-middle attacks where someone pretends to be an SSH server you are expecting to log into and captures your username and password. When you connect to an SSH server for the first time it asks you to accept the host key, and then in future it will check that the host key matches. If malicious actor tries to MITM you then this check will fail.

Let's generate a host key for our server to use.

```bash
ssh-keygen -t rsa -f .ssh/ssh_host_rsa_key -N ''
```

Now we need to create our configuration for our SSH server. The default config lives at `/etc/ssh/sshd_config`, however we do not have permission to modify that as the `jovyan` user, so let's create our own at `binder/sshd_config` with the following contents.

```
Port 2222
UsePrivilegeSeparation no
HostKey /home/jovyan/.ssh/ssh_host_rsa_key
UsePAM no
PidFile /home/jovyan/.ssh/sshd.pid

X11Forwarding yes
PrintMotd no
ChallengeResponseAuthentication no
AcceptEnv LANG LC_*
Subsystem	sftp	/usr/lib/openssh/sftp-server
```

Let's talk through this line by line.

First we set the SSH port to `2222`. This is because as an unprivileged user we can't start services on the common low port of `22`. Instead we must set it to something over `1024`, so we are choosing `2222`.

Next we set `UsePrivilegeSeparation` to `no`. Privilege separation creates subprocesses with the user that is logging in, however you need to be `root` for this to work so we are disabling it.

Then we specify the path to the `HostKey` that we just generated.

Next we also set `UsePAM` to `no` because we would need to be `root` to do this. This does mean that our SSH server will not be able to use PAM modules for authentication, but as we just want to SSH to localhost this is fine.

Lastly we set the `PidFile` to `/home/jovyan/.ssh/sshd.pid`. When we start the SSH daemon it will create a pid file, and usually this is under `/run` which we do not have write access to as an unprivileged user.

The final options here are the defaults from `/etc/ssh/sshd_config` which we need to include, but as we are not changing them I will not explain what they do.

I think we should also probably put this file under `.ssh/` in order to keep everything together, so let's copy it there with our `start` script.

```bash
cp binder/sshd_config .ssh/sshd_config
```

### Start the SSH daemon

Now we are ready to start the SSH daemon process. We will call it from our `start` script but detach it to run in the background with `&`.

```bash
/usr/sbin/sshd -f .ssh/sshd_config -D &
```

### User experience additions

This is enough to get an SSH server running on our Binder instance, however there are some things that we could do to make the experience for the user even better. Specifically I want a user to be able to run `ssh localhost` as soon as they connect to their binder and it "just work".

Right now our SSH daemon is running on port `2222` because we were unable to listen on low ports. This means users would need to run `ssh -p 2222 localhost` to be able to connect. However we could update our SSH client configuration to store this instead. Your SSH client configuration lives at `.ssh/config` and we should add something like this.

```
Host localhost
  Port 2222
```

As this file does not exist and we are creating things from our `start` script we can add the following line to create it for us.

```bash
printf "Host localhost\n  Port 2222\n" > .ssh/config
```

Now a user can run `ssh localhost` and it will automatically connect to port `2222`. However the user will then be asked to accept the host key, because this is the first time we have connected to the server. We could be nice and set this up for them too, so that they aren't prompted to accept this. Add this to your `start` script.

```bash
sleep 5  # Give the SSH server a chance to start
ssh-keyscan -p 2222 -H localhost >> ~/.ssh/known_hosts
```

This connects to `localhost` and requests the public part of the host key and immediately saves it to our `known_hosts` file. Sadly `ssh-keyscan` doesn't take your `~/.ssh/config` into account so we still need to specify the port, and it needs the SSH server to be running so we need to give it a few seconds.

The last thing we should do is ensure the `PATH` variable is set correctly. When Jupyter starts a terminal it ensures the `PATH` includes things like `conda`, but our SSH connection will not do this. So we should add the additional paths to our `.bashrc` file.

```bash
printf '\n\nexport PATH="/srv/conda/envs/notebook/bin:/srv/conda/condabin:/home/jovyan/.local/bin:/home/jovyan/.local/bin:/srv/conda/envs/notebook/bin:/srv/conda/bin:/srv/npm/bin:$PATH"\n' >> .bashrc
```

There we are, if you test this out our new user should be able to start up their Binder and immediately run `ssh localhost`.

You are welcome to try out the [Dask Tutorial binder](https://mybinder.org/v2/gh/jacobtomlinson/dask-video-tutorial-2020/master?urlpath=lab) that I was creating when writing this if you want to see it in action.

## Startup script

Now that we have all the pieces our final `binder/start` script should look like this.

```bash
#!/bin/bash

# Make SSH directory
mkdir ~/.ssh && chmod 700 ~/.ssh

# Generate user SSH key and authorize it for ssh to localhost
ssh-keygen -t rsa -f .ssh/id_rsa -N ''
cat .ssh/id_rsa.pub > .ssh/authorized_keys
chmod 600 .ssh/authorized_keys

# Generate host key
ssh-keygen -t rsa -f .ssh/ssh_host_rsa_key -N ''

# Put SSH daemon config in place
cp binder/sshd_config .ssh/sshd_config

# Start SSH daemon
/usr/sbin/sshd -f .ssh/sshd_config -D &

# Configure client to use port 2222 when ssh to localhost
printf "Host localhost\n  Port 2222\n" > .ssh/config

# Approve host key in client when server starts
sleep 5  # Give ssh a chance to start
ssh-keyscan -p 2222 -H localhost >> ~/.ssh/known_hosts

# Fix PATH so that you can use conda, etc
printf '\n\nexport PATH="/srv/conda/envs/notebook/bin:/srv/conda/condabin:/home/jovyan/.local/bin:/home/jovyan/.local/bin:/srv/conda/envs/notebook/bin:/srv/conda/bin:/srv/npm/bin:$PATH"\n' >> .bashrc

exec "$@"
```

We also have an `binder/apt.txt` file which looks like this.

```
openssh-server
```

And we have a `binder/sshd_config` which looks like this.

```
Port 2222
UsePrivilegeSeparation no
HostKey /home/jovyan/.ssh/ssh_host_rsa_key
UsePAM no
PidFile /home/jovyan/.ssh/sshd.pid

X11Forwarding yes
PrintMotd no
ChallengeResponseAuthentication no
AcceptEnv LANG LC_*
Subsystem	sftp	/usr/lib/openssh/sftp-server
```

## External access

One last thing we may want to think about is external access. Now that we have an SSH server running in our Binder session theoretically it would be possible to SSH to that server from anywhere right?

Well half right. Yes we do have an SSH server running on port `2222` and we have an SSH key which would allow us access. However our Binder instance is a Docker container running inside a Kubernetes cluster in the cloud with no external network access other than HTTPS access to Jupyter.

You could look at using a service like [serveo](https://serveo.net) or [ngrok](https://ngrok.com/) to forward the port out. But if you're looking to remotely access your session you may prefer to use a free service like [Teleconsole](https://www.teleconsole.com/).

## Conclusion

I wanted to be able to SSH to `localhost` from a Binder session for a Dask Tutorial so that I could show off the [`SSHCluster` cluster manager](https://docs.dask.org/en/latest/setup/ssh.html).

I hope that if you also want to be able to SSH to `localhost` in your Binder sessions that this was useful for you too.
