---
author: Jacob Tomlinson
date: 2014-11-25T00:00:00+00:00
categories:
  - Python
tags:
  - python
  - easy_install
  - pip
  - proxy
thumbnail: python
title: How to easy_install and pip through a proxy
aliases:
  - /2014/11/25/easy-install-and-pip-through-a-proxy/
---


If you're trying to install a Python package using easy_install or pip and you
connect to the internet via a proxy you'll need to make a few changes to your
setup.

### easy_install

_easy_install_ requires you set the http_proxy and https_proxy environment
variables. You can either run the following commands in a terminal or add them
to your .bashrc.

```bash
export http_proxy=http://proxy_url:proxy_port
export https_proxy=http://proxy_url:proxy_port
```

### pip

_pip_ should follow the same rules as above, so you can specify the proxy config
in your environment variables.

However it also supports a `--proxy` flag if you want to specify it manually.

```bash
pip install --proxy=http://proxy_url:proxy_port package
```
