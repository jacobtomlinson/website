---
title: 'Python script: Recursively remove empty folders/directories'
author: Jacob Tomlinson
layout: post
permalink: /2014/02/16/python-script-recursively-remove-empty-folders-directories/
category: Python
thumbnail: python
tags:
  - Module
  - Python
  - Script
---
So as part of a script I&#8217;m writing I needed the ability to recursively remove empty folders/directories from a filesystem. After a bit of googling I found <a title="Original python script" href="http://dev.enekoalonso.com/2011/08/06/python-script-remove-empty-folders" target="_blank">this very useful script</a> by <a title="Eneko Alonso's Blog" href="http://dev.enekoalonso.com" target="_blank">Eneko Alonso</a>. However the script isn&#8217;t really in a usable state for what I want so I decided to make a few changes to it and publish it on GitHub.

### Features

*   Recursively delete empty folders/directories
*   Run as a standalone script or import into existing script
*   Specify whether or not to delete the root folder (useful when you want to delete empty sub folders but not the folder itself even if empty)

### Usage


```bash
# Remove empty folders under /path/to/clean including /path/to/clean
$ /path/to/script/remove_empty_folders.py /path/to/clean

# Remove empty folders under /path/to/clean not including /path/to/clean
$ /path/to/script/remove_empty_folders.py /path/to/clean False
```


### Script

```python
#! /usr/bin/env python
'''
Module to remove empty folders recursively. Can be used as standalone script or be imported into existing script.
'''

import os, sys

def removeEmptyFolders(path, removeRoot=True):
  'Function to remove empty folders'
  if not os.path.isdir(path):
    return

  # remove empty subfolders
  files = os.listdir(path)
  if len(files):
    for f in files:
      fullpath = os.path.join(path, f)
      if os.path.isdir(fullpath):
        removeEmptyFolders(fullpath)

  # if folder empty, delete it
  files = os.listdir(path)
  if len(files) == 0 and removeRoot:
    print "Removing empty folder:", path
    os.rmdir(path)

def usageString():
  'Return usage string to be output in error cases'
  return 'Usage: %s directory [removeRoot]' % sys.argv[0]

if __name__ == "__main__":
  removeRoot = True

  if len(sys.argv) &lt; 1:
    print "Not enough arguments"
    sys.exit(usageString())

  if not os.path.isdir(sys.argv[1]):
    print "No such directory %s" % sys.argv[1]
    sys.exit(usageString())

  if len(sys.argv) == 2 and sys.argv[2] != "False":
    print "removeRoot must be 'False' or not set"
    sys.exit(usageString())
  else:
    removeRoot = False

  removeEmptyFolders(sys.argv[1], removeRoot)</code>
```
