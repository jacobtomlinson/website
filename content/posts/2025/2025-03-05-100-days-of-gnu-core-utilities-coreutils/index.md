---
title: "100 Days of Coreutils"
date: 2025-03-05T00:00:00+00:00
draft: false
author: "Jacob Tomlinson"
url: /blog/100-days-of-coreutils
categories:
  - blog
tags:
  - GNU Core Utilities
  - coreutils
  - 100 days challenge
---

I consider myself an advanced Linux and macOS user. I'm currently a software engineer developing primarily for Linux systems, and I've previously worked as a Linux and Mac System Administrator. Over the years I've spent tons of time on the command line, however I bet there are a bunch of [GNU Core Utilities (coreutils)](https://en.wikipedia.org/wiki/GNU_Core_Utilities) commands I've never used before.

I'm setting myself a challenge to work my way through the [list of coreutils commands](https://en.wikipedia.org/wiki/List_of_GNU_Core_Utilities_commands) and try one out every day. There are around 100 commands give or take, so this should make a nice 100 days challenge. I'll update this post with a little background on the command, what you would use it for and some example usage. I'll also post these on my [Bluesky account](https://bsky.app/profile/jacobtomlinson.dev).

My goal here is to try out every command and hopefully discover a few new ones to add to my every day CLI usage.

## Background

The [GNU Core Utilities (coreutils)](https://en.wikipedia.org/wiki/GNU_Core_Utilities) are an open-source set of tools that you will find on nearly any unix-based system. If you've ever used the command line or terminal you will have used some of these commands before, even if you didn't know they were part of a package called GNU coreutils. These include common commands like `cat`, `ls` and `rm` and less common ones like `expand` and `readlink`.

## List of commands

Here's a table of contents taken from the [wikipedia list of coreutils commands](https://en.wikipedia.org/wiki/GNU_Core_Utilities). As I work through the list (not necessarily in order) I'll update this with links to the page sections.

- File Utilities
    - `chcon` - Changes file security context (SELinux)
    - `chgrp` - Changes file group ownership
    - `chown` - Changes file ownership
    - `chmod` - Changes the permissions of a file or directory
    - `cp` - Copies a file or directory
    - `dd` - Copies and converts a file
    - `df` - Shows disk free space on file systems
    - `dir` - Is exactly like `ls -C -b`. (Files are by default listed in columns and sorted vertically.)
    - `dircolors` - Set up color for ls
    - `install` - Copies files and set attributes
    - `ln` - Creates a link to a file
    - `ls` - Lists the files in a directory
    - `mkdir` - Creates a directory
    - `mkfifo` - Makes named pipes (FIFOs)
    - `mknod` - Makes block or character special files
    - `mktemp` - Creates a temporary file or directory
    - `mv` - Moves files or rename files
    - `realpath` - Returns the resolved absolute or relative path for a file
    - `rm` - Removes (deletes) files, directories, device nodes and symbolic links
    - `rmdir` - Removes empty directories
    - `shred` - Overwrites a file to hide its contents, and optionally deletes it
    - `sync` - Flushes file system buffers
    - `touch` - Changes file timestamps; creates file
    - `truncate` - Shrink or extend the size of a file to the specified size
    - `vdir` - Is exactly like `ls -l -b`. (Files are by default listed in long format.)
- Text Utilities
    - `b2sum` - Computes and checks BLAKE2b message digest
    - `base32` - Encodes or decodes Base32, and prints result to standard output
    - `base64` - Encodes or decodes Base64, and prints result to standard output
    - `basenc` - Encodes or decodes various encodings like Hexadecimal, Base32, Base64, Z85 etc., and prints result to standard output
    - [`cat`](#cat) - Concatenates and prints files on the standard output
    - `cksum` - Checksums (IEEE Ethernet CRC-32) and count the bytes in a file.
    - `comm` - Compares two sorted files line by line
    - `csplit` - Splits a file into sections determined by context lines
    - `cut` - Removes sections from each line of files
    - `expand` - Converts tabs to spaces
    - `fmt` - Simple optimal text formatter
    - `fold` - Wraps each input line to fit in specified width
    - `head` - Outputs the first part of files
    - `join` - Joins lines of two files on a common field
    - `md5sum` - Computes and checks MD5 message digest
    - `nl` - Numbers lines of files
    - `numfmt` - Reformat numbers
    - `od` - Dumps files in octal and other formats
    - `paste` - Merges lines of files
    - `ptx` - Produces a permuted index of file contents
    - `pr` - Converts text files for printing
    - `sha1sum`/`sha224sum`/`sha256sum`/`sha384sum`/`sha512sum` - Computes and checks SHA-1/SHA-2 message digests
    - `shuf` - generate random permutations
    - `sort` - sort lines of text files
    - `split` - Splits a file into pieces
    - `sum` - Checksums and counts the blocks in a file
    - `tac` - Concatenates and prints files in reverse order line by line
    - `tail` - Outputs the last part of files
    - `tr` - Translates or deletes characters
    - `tsort` - Performs a topological sort
    - `unexpand` - Converts spaces to tabs
    - `uniq` - Removes duplicate lines from a sorted file
    - `wc` - Prints the number of bytes, words, and lines in files
- Shell Utilities
    - `arch` - Prints machine hardware name (same as `uname -m`)
    - `basename` - Removes the path prefix from a given pathname
    - `chroot` - Changes the root directory
    - `date` - Prints or sets the system date and time
    - `dirname` - Strips non-directory suffix from file name
    - `du` - Shows disk usage on file systems
    - `echo` - Displays a specified line of text
    - `env` - Displays and modifies environment variables
    - `expr` - Evaluates expressions
    - `factor` - Factors numbers
    - `false` - Does nothing, but exits unsuccessfully
    - `groups` - Prints the groups of which the user is a member
    - `hostid` - Prints the numeric identifier for the current host
    - `id` - Prints real or effective UID and GID
    - `link` - Creates a link to a file
    - `logname` - Print the user's login name
    - `nice` - Modifies scheduling priority
    - `nohup` - Allows a command to continue running after logging out
    - `nproc` - Queries the number of (active) processors
    - `pathchk` - Checks whether file names are valid or portable
    - `pinky` - A lightweight version of finger
    - `printenv` - Prints environment variables
    - `printf` - Formats and prints data
    - `pwd` - Prints the current working directory
    - `readlink` - Displays value of a symbolic link
    - `runcon` - Run command with specified security context
    - `seq` - Prints a sequence of numbers
    - `sleep` - Delays for a specified amount of time
    - `stat` - Returns data about an inode
    - `stdbuf` - Controls buffering for commands that use stdio
    - `stty` - Changes and prints terminal line settings
    - `tee` - Sends output to multiple files
    - `test` - Evaluates an expression
    - `timeout` - Run a command with a time limit
    - `true` - Does nothing, but exits successfully
    - `tty` - Prints terminal name
    - `uname` - Prints system information
    - `unlink` - Removes the specified file using the unlink function
    - `uptime` - Tells how long the system has been running
    - `users` - Prints the user names of users currently logged into the current host
    - `who` - Prints a list of all users currently logged in
    - `whoami` - Prints the effective userid
    - `yes` - Prints a string repeatedly
    - `[` - 	A synonym for test; this program permits expressions like `[ expression ]`. 

## Commands

### Day 1: `cat` {#cat}

`cat` or "concatenate" takes one or more text files, concatenates them together and prints them out to the terminal.

I rarely use the concatenation functionality, I usually just use it to print a single file out to the screen.

```console
$ cat /etc/os-release
PRETTY_NAME="Ubuntu 22.04.5 LTS"
NAME="Ubuntu"
VERSION_ID="22.04"
VERSION="22.04.5 LTS (Jammy Jellyfish)"
VERSION_CODENAME=jammy
ID=ubuntu
ID_LIKE=debian
HOME_URL="https://www.ubuntu.com/"
SUPPORT_URL="https://help.ubuntu.com/"
BUG_REPORT_URL="https://bugs.launchpad.net/ubuntu/"
PRIVACY_POLICY_URL="https://www.ubuntu.com/legal/terms-and-policies/privacy-policy"
UBUNTU_CODENAME=jammy
```

But if you have multiple files you can `cat` them out in a single command.

```console
$ cat /tmp/foo
foo
$ cat /tmp/bar
bar
$ cat /tmp/foo /tmp/bar
foo
bar
```
