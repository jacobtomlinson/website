---
title: SSH without a password on OS X with ssh-copy-id
author: Jacob Tomlinson
layout: post
permalink: /2013/01/24/ssh-copy-id-os-x/
category: Apple
thumbnail: command-line
tags:
  - os x
  - ssh
  - ssh-copy-id
  - terminal
---

### What is ssh-copy-id?
ssh-copy-id is a script that uses ssh to log into a remote machine (presumably using a login password,
so password authentication should be enabled, unless you&#8217;ve done some clever use of multiple
identities) It also changes the permissions of the remote user&#8217;s home, `~/.ssh`, and
`~/.ssh/authorized_keys` to remove group writability (which would otherwise prevent you from logging in,
if the remote `sshd` has `StrictModes` set in its configuration).

If the `-i` option is given then the
identity file (defaults to `~/.ssh/id_rsa.pub`) is used, regardless of whether there are any keys in your
`ssh-agent`. Otherwise, if this: `ssh-add -L` provides any output, it uses that in preference to the
identity file. If the `-i` option is used, or the `ssh-add` produced no output, then it uses the
contents of the identity file. Once it has one or more fingerprints (by whatever means) it uses ssh
to append them to `~/.ssh/authorized_keys` on the remote machine (creating the file, and directory,
  if necessary).

### Installation method 1: Copy the script
So to simply install the script on your OS X machine simply type the commands

```bash
sudo curl https://raw.githubusercontent.com/beautifulcode/ssh-copy-id-for-OSX/master/ssh-copy-id.sh -o /usr/local/bin/ssh-copy-id
sudo chmod +x /usr/local/bin/ssh-copy-id
```

Once you have done this you should be able to use ssh-copy-id as you can on linux.

### Installation method 2: Homebrew
You also have the option to install ssh-copy-id with homebrew, if you don&#8217;t know what homebrew is it&#8217;s basically a package installer for adding command line programs which are common in the unix community but aren&#8217;t included with OS X.

First you will need to install homebrew if you don&#8217;t have it already, you can find some very simple (a one line command at time or writing) instructions on the <a title="Homebrew" href="http://brew.sh" target="_blank">homebrew website</a>.

Then to install ssh-copy-id you simply run the command

```bash
sudo brew install ssh-copy-id
```

**Usage**  
For those unfamiliar with using ssh-copy-id or other ssh tools here is a quick guide on how to set up passwordless SSH logins.

First you need to generate a key, this can be done with the command

```bash
ssh-keygen -t rsa
```

This will prompt you for a location and name which you can leave as default by just hitting enter. Make a note of the location and name as you will need it in a minute.

Next you need to push that key onto your remote server. To do this you need to run the following command

```bash
ssh-copy-id -i [path to rsa file] user@machine
```

So the `-i` sets which rsa file we are going to use, so we just put the path to the file we just generated with ssh-keygen. Then you put the login for the machine you want to SSH to. It will run the command and then once it has finished you should be able to SSH to your machine and it will not ask you for a password.
