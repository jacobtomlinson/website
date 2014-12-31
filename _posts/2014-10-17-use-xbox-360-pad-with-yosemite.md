---
title: How to use an Xbox 360 controller with OS X Yosemite
author: Jacob Tomlinson
layout: post
permalink: /2014/10/17/use-xbox-360-pad-with-yosemite/
category: Gaming
thumbnail: xbox-controller
tags:
  - xbox 360 controller
  - yosemite
  - OS X
  - gaming
---

### Introduction

Want to use an Xbox 360 controller to play games on Mac OS X 10.10 Yosemite? Here are
simple instructions to get you started in no time.

## Driver

To get your Xbox 360 pad working you'll need to install a driver. Drivers for OS X
are known as kexts and Apple provide a standard set to support all sorts of third
party hardware. But if you haven't guessed already they don't provide one for
the Xbox 360 controller so we'll have to look elsewhere.

Colin Munro at tattiebogle.net created a driver for the Xbox 360 controller,
however this hasn't been updated since Snow Leopard and it won't work with
OS X Yosemite. Luckily for us someone else has updated the code and [posted it on
GitHub](https://github.com/d235j/360Controller/releases).

To install it just download the .dmg file from the GitHub repository, open it
and run the installer.

You'll need to reboot once it's done for your Mac to recognise the controller.

## Configuration

This is enough to get your controller working in most situations but if you
want to change the configuration of your controller (e.g invert the axis) then
you just need to head to your "System Preferences" and you'll find a new option
called "Xbox 360 Controllers".

From this menu you'll be able to test your controller is working as well as
make changes to the config.

![Imgur](http://i.imgur.com/ikYGruw.png)

### Updates
 * Add disable kext sign check
 * Removed disable kext sign check, not needed as of v0.14
