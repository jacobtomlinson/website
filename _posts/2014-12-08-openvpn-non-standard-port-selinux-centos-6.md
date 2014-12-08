---
title: Run OpenVPN on non-standard port with SELinux and Centos 6
author: Jacob Tomlinson
layout: post
permalink: /2014/12/08/openvpn-non-standard-port-selinux-centos-6/
category: Centos
thumbnail: centos
tags:
- centos 6
- SELinux
- OpenVPN
---

I recently installed OpenVPN on a Centos 6 server but found that I couldn't get
the service to start. Running `service openvpn start` failed despite being
able to run `openvpn --config /path/to/config` without errors.

When looking in `/var/log/messages` after a failed start I found the following
error message

```
TCP/UDP: Socket bind failed on local address [undef]: Permission denied
Exiting due to fatal error
```

OpenVPN was failing to bind to the port and this was because I had configured it
to run on a non-standard port. By default in Centos 6 [SELinux][1] is set to
`enforcing` and it will block any services which try to start on an unusual port.

To get around this you must tell SELinux that you're happy for it to run on a
non-standard port with the `semanage` command. This is not installed by default
so you'll need to run this install `policycoreutils-python`.


```bash
yum install policycoreutils-python
```

You can then use the following command to tell SELinux you're happy for OpenVPN
to run on your specified port.

```bash
semanage port -a -t openvpn_port_t -p udp port
```

Once you've done this you `service openvpn start` should now succeed.

[1]: http://selinuxproject.org/page/Main_Page
