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
    - [`cp`](#cp) - Copies a file or directory
    - `dd` - Copies and converts a file
    - `df` - Shows disk free space on file systems
    - [`dir`](#dir) - Is exactly like `ls -C -b`. (Files are by default listed in columns and sorted vertically.)
    - `dircolors` - Set up color for ls
    - `install` - Copies files and set attributes
    - `ln` - Creates a link to a file
    - [`ls`](#ls) - Lists the files in a directory
    - `mkdir` - Creates a directory
    - `mkfifo` - Makes named pipes (FIFOs)
    - `mknod` - Makes block or character special files
    - `mktemp` - Creates a temporary file or directory
    - [`mv`](#mv) - Moves files or rename files
    - `realpath` - Returns the resolved absolute or relative path for a file
    - [`rm`](#rm) - Removes (deletes) files, directories, device nodes and symbolic links
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
    - [`true`](#true) - Does nothing, but exits successfully
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

[`cat`](https://linux.die.net/man/1/cat) or "concatenate" takes one or more text files, concatenates them together and prints them out to the terminal.

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

### Day 2: `cp` {#cp}

The copy command [`cp`](https://linux.die.net/man/1/cp) allows you to copy a file from one place to another.

```console
$ cp /tmp/foo /tmp/bar
```

I commonly use the `-R` flag to copy a directory recursively.

```console
$ cp -R /tmp/dir1/ /tmp/dir2/
```

An interesting more advanced use case I've used before is to also specify the `-L` flag which tells `cp` to follow symbolic links instead of copying them.

```console
$ cp -LR /tmp/dir1/ /tmp/dir2/  # Any symbolic links found in dir1 will create copies of the files, not the links in dir2
```

### Day 3: `ls` {#ls}

The "list" command or [`ls`](https://linux.die.net/man/1/ls) lists the contents of a directory.

The default behaviour on most systems is to print out a tab separated list of names.

```console
$ ls ~/
Applications     Documents        Library          Music            Pictures
Desktop          Downloads        Movies           Notes            Projects
Public           Scratch
```

I commonly use the `-l` flag to list detailed information.

```console
$ ls -l ~/
total 0
drwx------@   3 jtomlinson  staff    96 11 Jan  2024 Applications
drwx------@  12 jtomlinson  staff   384 12 Feb 14:08 Desktop
drwx------@   4 jtomlinson  staff   128 30 Jul  2024 Documents
drwx------@ 227 jtomlinson  staff  7264  5 Mar 14:05 Downloads
drwx------@  99 jtomlinson  staff  3168 12 Feb 14:07 Library
drwx------    5 jtomlinson  staff   160  4 Mar  2024 Movies
drwx------+   4 jtomlinson  staff   128 10 Jan  2024 Music
drwx------   10 jtomlinson  staff   320  6 Mar 14:46 Notes
drwx------+  13 jtomlinson  staff   416 21 Jan 13:59 Pictures
drwx------@  16 jtomlinson  staff   512 15 Nov 13:04 Projects
drwx------@   4 jtomlinson  staff   128 10 Jan  2024 Public
drwx------@  11 jtomlinson  staff   352 27 Feb 12:35 Scratc
```

I also often use the `-a` flag to list hidden files.

```console
$ ls -la ~/
total 824
drwxr-x---+  73 jtomlinson  staff    2336  7 Mar 10:59 .
drwxr-xr-x    7 root        admin     224 25 Feb 10:11 ..
-rw-r--r--@   1 jtomlinson  staff   10244  7 Mar 10:32 .DS_Store
drwx------+  60 jtomlinson  staff    1920  6 Mar 17:01 .Trash
lrwxr-xr-x@   1 jtomlinson  staff      64 10 Jan  2024 .bash_profile
lrwxr-xr-x@   1 jtomlinson  staff      58 10 Jan  2024 .bashrc
drwxr-xr-x@  12 jtomlinson  staff     384  6 Mar 11:34 .cache
drwxr-xr-x@  11 jtomlinson  staff     352  9 Dec 14:03 .config
drwxr-xr-x@   9 jtomlinson  staff     288 27 Feb 12:36 .kube
drwxr-xr-x@   6 jtomlinson  staff     192 22 Feb  2024 .local
drwx------@  22 jtomlinson  staff     704 30 Sep 14:42 .ssh
lrwxr-xr-x@   1 jtomlinson  staff      61 10 Jan  2024 .tmux.conf
lrwxr-xr-x@   1 jtomlinson  staff      57 10 Jan  2024 .vimrc
drwxr-xr-x@   5 jtomlinson  staff     160 10 Jan  2024 .vscode
-rw-------@   1 jtomlinson  staff   93970  7 Mar 10:59 .zsh_history
lrwxr-xr-x@   1 jtomlinson  staff      57 10 Jan  2024 .zshrc
drwx------@   3 jtomlinson  staff      96 11 Jan  2024 Applications
drwx------@  12 jtomlinson  staff     384 12 Feb 14:08 Desktop
drwx------@   4 jtomlinson  staff     128 30 Jul  2024 Documents
drwx------@ 227 jtomlinson  staff    7264  5 Mar 14:05 Downloads
drwx------@  99 jtomlinson  staff    3168 12 Feb 14:07 Library
drwx------    5 jtomlinson  staff     160  4 Mar  2024 Movies
drwx------+   4 jtomlinson  staff     128 10 Jan  2024 Music
drwx------   10 jtomlinson  staff     320  6 Mar 14:46 Notes
drwx------+  13 jtomlinson  staff     416 21 Jan 13:59 Pictures
drwx------@  16 jtomlinson  staff     512 15 Nov 13:04 Projects
drwx------@   4 jtomlinson  staff     128 10 Jan  2024 Public
drwx------@  11 jtomlinson  staff     352 27 Feb 12:35 Scratch
```

I also like the `-h` flag which displays the file sizes in a human readable way.

```console
$ ls -lah ~/.zsh_history
-rw-------@ 1 jtomlinson  staff    92K  7 Mar 11:02 .zsh_history
```

In my [dotfiles](https://github.com/jacobtomlinson/dotfiles) I set some aliases around these to make things easier. I originally cribbed these from the Red Hat Enterprise Linux 6 (RHEL6) default `.bashrc` file and have used them ever since.

```bash
# Set ls colours
export LSCOLORS=ExFxBxDxCxegedabagacad

# Aliases to make ls easier to use in different modes, taken from RHEL 6
alias ls='ls -GFh'
alias ll="ls -l"
alias lh="ls -lh"
alias la="ls -la"
alias lah="ls -lah"
alias sl="ls"
```

### Day 4: `mv` {#mv}

Moving a file with `mv` is another common operation. If you are moving a file or folder within a filesystem (e.g on the same hard drive partition) then you can think of it more like renaming because it doesn't actually move any of the bytes on the disk, it just updates the filesystem with a new name.

```bash
mv /tmp/foo /tmp/bar  # Renames the file foo to bar in /tmp
```

You might want to move a bunch of files into a different folder, but if a file exists with the same name you don't want to overwite it.

You can use the `-n` flag to not overwite files, or the `-i` flag to prompt you before overwriting. Generally `-i` is the default and you can set `-f` which forcibly overwites files without asking.

```bash
mv -f /tmp/baz/* /tmp/buzz/.  # Move all the files in baz into buzz and overwite any that already exist without prompting
```

### Day 5: `dir` {#dir}

The `dir` command is effectively an alias to `ls -C -b`. The `-C` flag forces multiple columns and `-b` will escape special characters. You wont find this command on all systems, for example it's not on macOS or in busybox but you will find it on Ubuntu.

This command exists mostly to be compatible with other operating systems like DOS and Windows.

```console
$ dir /
bin  boot  dev  etc  home  lib  media  mnt  opt  proc  root  run  sbin  srv  sys  tmp  usr  var
```

### Day 6: `rm` {#rm}

The `rm` command removes a file. You can also recursively do this with `rm -r`. This will prompt you for each file and check if you want to delete it. You can override this with `rm -f`.

Combining these two flags gives you `rm -rf` which is often considered one of the most dangerous linux commands because you can delete your whole filesystem this way. 

```bash
rm /tmp/foo
```

When I was a linux sysadmin there was a user who had a script which cleaned up files in a directory. The directory was set as an environment variable `$SOME_DIR`. So their script had a line like this:

```bash
rm -rf $SOME_DIR/*
```

If an environment variable is unset it just resolves to an empty string, and this script had a bug which caused that environment variable to be unset. So when their script ran it ran `rm -rf /*` instead. The script was run via a cron job in the middle of the night, so when they cam in the next day they found everything in their home directory had been deleted, along with every file on the shared network storage with `777` permissions. We spent a lot of time restoring files from backups that day.

### Day 7: `true` {#true}

The `true` command does nothing, but always returns successfully with exit code `0`. 

You most commonly see this used within a script with `set -e` which will exit the script if any commands fail. You might want to continue executing the script if a command fails, so you can use the or opetator with `true` to make it always appear to have succeeded.

```bash
#!/bin/bash
# somescript.sh

set -e

# The rm fails but the or true makes the line succeed and the script will continue
rm /tmp/filethatdoesntexist || true
```
