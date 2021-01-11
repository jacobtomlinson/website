---
title: "How to interactively debug GitHub Actions with netcat"
date: 2021-01-08T00:00:00+00:00
draft: false
author: "Jacob Tomlinson"
categories:
  - blog
tags:
  - GitHub Actions
  - Debugging
thumbnail: github
---

**Update**: This was a fun experiment and I recommend you check out the post for a fun read on setting up reverse shells. But I've since discovered [this awesome tmate action](https://github.com/mxschmitt/action-tmate) which lets you interactively debug in the browser or via SSH.

```yaml
- name: Debug with tmate on failure
  if: ${{ failure() }}
  uses: mxschmitt/action-tmate@v3
```

With this step if any previous step in your workflow fails a `tmate` session will be started and the connection info will be repeatedly printed in the workflow output.

```text
Created new session successfully
ssh xMMK8vwSQyCXdZfTCS9hN7fgx@nyc1.tmate.io

https://tmate.io/t/xMMK8vwSQyCXdZfTCS9hN7fgx
```

Much easier!

---

**Original post**

When a [GitHub Actions](https://github.com/features/actions) workflow fails it would be really nice to be able to interactively debug things with a shell. GitHub doesn't provide anything like a web console or SSH access to workflow runners so in this post we walk talk through throwing shells with [netcat](https://en.wikipedia.org/wiki/Netcat) and catching them with netcat and [ngrok](https://ngrok.com/).

## Throwing a reverse shell

The most common way to get a shell on a remote system is to log in via [SSH](https://en.wikipedia.org/wiki/SSH_(Secure_Shell)). This provides encryption and authentication and makes the whole process simple. However it requires you to run an SSH server on that system and have network and firewall rules configured to allow incoming traffic, and authentication credentials for that system.

Alternatively you can use a reverse shell, which is where a system will connect out to some other machine on the internet and then forward a shell over that connection. This technique is commonly used in the security community to open backdoors in compromised systems, but is also extremely useful for debugging on a restricted environment such as a CI worker.

## Catching a shell

In order to "throw" a shell to a remote system you first have to set up a machine to "catch" the connection.

For this example we are going to use netcat to catch our shell, `nc` is a standard linux utility that is available on most systems.

```console
$ nc -nlvp 4444
Listening on 0.0.0.0 4444
```

Now we are listening for incoming connections on port `4444`. Beware that this is an unauthenticated and unencrypted connection and we are going to expose it to the internet. For a bit of interactive debugging on open source projects on GitHub this is fine, but this shouldn't be used for sensitive information or long term solutions.

## Forwarding ports with ngrok

I'm also assuming here that the machine you are running this on (my developer laptop in my case) cannot receive traffic on port `4444` via the internet. So we can use ngrok to forward our ports.

Ngrok is a service which allows you to expose ports on your local machine to the internet, for the purposes of developing and testing software.

Once you've downloaded and authenticated ngrok you can set up the tunnel.

```console
$ ngrok tcp 4444
ngrok by @inconshreveable                                                                                                                                                                                                     (Ctrl+C to quit)

Session Status                online
Account                       Jacob Tomlinson (Plan: Free)
Version                       2.3.35
Region                        United States (us)
Web Interface                 http://127.0.0.1:4040
Forwarding                    tcp://2.tcp.ngrok.io:13604 -> localhost:4444

Connections                   ttl     opn     rt1     rt5     p50     p90
                              0       0       0.00    0.00    0.00    0.00
```

In this example we can see that port `4444` on `localhost` is now also available on port `13604` at `2.tcp.ngrok.io`. This will be different every time you create a connection.

## Configuring GitHub Actions

Now that we are listening for a shell connection we need to add a step to our GitHub Actions workflow to make the outbound connection.

We probably do not want to leak our connection information into our config. It's ephemeral so it's not a huge problem, but storing the connection info in a secret is still good practice.

In your repository head to `Settings > Secrets` and create `DEBUG_HOST` and `DEBUG_PORT` secrets with the hostname and port that `ngrok` gave us.

![Secrets](https://i.imgur.com/pc1Ldfz.png)

Then add a last step to your GitHub workflow.

```yaml
name: Interactive debugging example
on:
  push:

jobs:
  interactive:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@v2

      # Rest of my workflow steps

      - name: Thow interactive shell
        shell: bash -i {0}
        run: |
          rm /tmp/f>/dev/null 2>&1;mkfifo /tmp/f;cat /tmp/f|/bin/sh -i 2>&1|nc ${{ secrets.DEBUG_HOST }} ${{ secrets.DEBUG_PORT }} >/tmp/f
```

In this last step we use a combination of `mkfifo`, `cat`, `sh` and `nc` to forward a shell to our remote host.

When your workflow gets to this step it will appear to run indefinitely with no output.

![Workflow running](https://i.imgur.com/9LRAHpO.png)

But if we look at the `nc` session we have running on our local machine we should now see a shell prompt.

```console
Connection received on 127.0.0.1 57724
/bin/sh: 0: can't access tty; job control turned off
$
```

We can then run `bash` here to get a more useful shell.

```console
$ bash -i
bash: cannot set terminal process group (2507): Inappropriate ioctl for device
bash: no job control in this shell
runner@fv-az12-647:~/work/github-actions-shell/github-actions-shell$
```

## What can I do here?

Now that you have a shell on your remote system you can do whatever you like. Just be aware that `nc` will not forward control commands like `ctrl+c` and will instead close the `nc` connection. If this happens you will need to restart `nc` and restart your workflow.

This is also a simple shell so something like SSH which require a [pseudo-tty](https://unix.stackexchange.com/questions/21147/what-are-pseudo-terminals-pty-tty) may not work as expected.

But we can still do things like poke around the runner's system.

```console
runner@fv-az12-647:~/work/github-actions-shell/github-actions-shell$ whoami
runner

runner@fv-az12-647:~/work/github-actions-shell/github-actions-shell$ cat /etc/os-release
NAME="Ubuntu"
VERSION="18.04.5 LTS (Bionic Beaver)"
ID=ubuntu
ID_LIKE=debian
PRETTY_NAME="Ubuntu 18.04.5 LTS"
VERSION_ID="18.04"
HOME_URL="https://www.ubuntu.com/"
SUPPORT_URL="https://help.ubuntu.com/"
BUG_REPORT_URL="https://bugs.launchpad.net/ubuntu/"
PRIVACY_POLICY_URL="https://www.ubuntu.com/legal/terms-and-policies/privacy-policy"
VERSION_CODENAME=bionic
UBUNTU_CODENAME=bionic

runner@fv-az12-647:~/work/github-actions-shell/github-actions-shell$ hostname -f
fv-az12-647.gip0skj2w3au3jd4qdtkx3lorh.cx.internal.cloudapp.net
```

And most importantly we can now start debugging our CI steps to see what went wrong.

```console
runner@fv-az12-647:~/work/github-actions-shell/github-actions-shell$ pytest myapp
```