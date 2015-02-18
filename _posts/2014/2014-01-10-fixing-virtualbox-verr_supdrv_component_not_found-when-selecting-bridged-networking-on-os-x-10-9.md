---
title: Fixing VirtualBox verr_supdrv_component_not_found when selecting bridged networking on OS X 10.9
author: Jacob Tomlinson
layout: post
thumbnail: command-line
permalink: /2014/01/10/fixing-virtualbox-verr_supdrv_component_not_found-when-selecting-bridged-networking-on-os-x-10-9/
category: Apple
tags:
  - kexts
  - mavericks
  - os x
  - VirtualBox
---
While installing CentOS in VirtualBox (version 4.2.4) on OS X (version 10.9.1) I came across the following error message when selecting bridged networking

```
virtualbox verr_supdrv_component_not_found
```

The solution I found to this is to reload the VirtualBox kexts on the host operating system. To do this I used a script which I found [here][1], I've put it on [GitHub as a gist][2] for future reference.


```bash
#!/bin/bash
# Reload Virtual Box Kexts

unload() {
        kextstat | grep "org.virtualbox.kext.VBoxUSB" > /dev/null 2>&1 && sudo kextunload -b org.virtualbox.kext.VBoxUSB
        kextstat | grep "org.virtualbox.kext.VBoxNetFlt" > /dev/null 2>&1 && sudo kextunload -b org.virtualbox.kext.VBoxNetFlt
        kextstat | grep "org.virtualbox.kext.VBoxNetAdp" > /dev/null 2>&1 && sudo kextunload -b org.virtualbox.kext.VBoxNetAdp
        kextstat | grep "org.virtualbox.kext.VBoxDrv" > /dev/null 2>&1 && sudo kextunload -b org.virtualbox.kext.VBoxDrv
}

load() {
        sudo kextload /Library/Extensions/VBoxDrv.kext -r /Library/Extensions/
        sudo kextload /Library/Extensions/VBoxNetFlt.kext -r /Library/Extensions/
        sudo kextload /Library/Extensions/VBoxNetAdp.kext -r /Library/Extensions/
        sudo kextload /Library/Extensions/VBoxUSB.kext -r /Library/Extensions/
}

case "$1" in
        unload|remove)
                unload
                ;;
        load)
                load
                ;;
        *|reload)
                unload
                load
                ;;
esac
```

[1]: https://forums.virtualbox.org/viewtopic.php?f=8&t=56013&start=15#p272403
[2]: https://gist.github.com/jacobtomlinson/8361722
