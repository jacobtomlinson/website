---
title: 'Yet another Raspberry Pi Gameboy Advance'
author: Jacob Tomlinson
layout: post
category: Electronics
thumbnail: raspberry-pi
tags:
  - raspberry pi
  - gameboy
  - adafruit
---

### Display config

Full instructions [here][adafruit-composite].

```bash
sudo nano /boot/config.txt
```

Replace

```
#framebuffer_width=1280
#framebuffer_height=720
```

with

```
framebuffer_width=320
framebuffer_height=240
```

[adafruit-composite]: https://learn.adafruit.com/using-a-mini-pal-ntsc-display-with-a-raspberry-pi/configure-and-test
[lipopi]: https://github.com/NeonHorizon/lipopi
[example-build-1]: https://retropie.org.uk/forum/topic/960/neopigamer
