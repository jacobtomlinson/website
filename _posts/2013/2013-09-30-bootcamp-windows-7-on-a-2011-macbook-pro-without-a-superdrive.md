---
title: Bootcamp Windows 7 on a 2011 MacBook Pro without a SuperDrive
author: Jacob Tomlinson
layout: post
permalink: /2013/09/30/bootcamp-windows-7-on-a-2011-macbook-pro-without-a-superdrive/
category: Apple
thumbnail: bootcamp
tags:
  - mac
  - os x
  - paragon ntfs
  - rEFIt
  - ssd
  - USB SuperDrive
  - vmware
  - windows
---
So recently I swapped out the SuperDrive in my early 2011 MacBook pro for an additional HDD caddy. I then moved my 1TB HDD into the caddy and put a new SSD in the HDD slot. As I was messing around with hard drives I decided to go for a fresh install of OS X and Windows on the SSD.

However I found that reinstalling Windows without the internal SuperDrive was a major pain. For some reason Apple have put in some restrictions which means that if the machine originally had an internal SuperDrive then it is not possible to install from a USB stick or from a USB SuperDrive. However if the machine didn&#8217;t come with an internal drive like a MacBook air then the USB SuperDrive will work fine. This is most likely to be down to the firmware on each machine.

So in order to get Windows installed on my MacBook I had two options, put the SuperDrive back in (which wasn&#8217;t really an option as I had put the drive in an external caddy which doesn&#8217;t look like it will be easy top open again) or to find another method. So here are the steps I followed to successfully install Windows, I&#8217;ll document them exactly as I know they work but I&#8217;ll make a note on sections which are likely to be optional.

### What you&#8217;ll need

*   A windows 7 or 8 install media ISO file
*   <a title="VMWare Fusion" href="http://www.vmware.com/uk/products/fusion/" target="_blank">VMWare Fusion</a> v5 (free trial is sufficient) (v6 is known not to work with this method)
*   <a title="Paragon NTFS" href="http://www.paragon-software.com/home/ntfs-mac/" target="_blank">Paragon NTFS</a> (free trial is sufficient)
*   <a title="rEFIt" href="http://refit.sourceforge.net/" target="_blank">rEFIt</a> (an addition to the Mac EFI bootloader) (probably optional but removal not tested)
*   a USB DVD drive and physical Windows disc may be helpful

### Method

So initially I put a Windows 7 DVD in my USB SuperDrive and began the usual process of installing Windows using the bootcamp assistant. I launched the bootcamp assistant from Applications > Utilities and used it to partition the HDD. Now the assistant will complain if you don&#8217;t have the Windows media attached so it may be worth plugging in a USB DVD drive with a Windows disk in for this part, we won&#8217;t actually use this to install Windows. If you try to boot from the USB attached DVD you will most likely just get a flashing white cursor in the top left corner.

![VMware Fusion](http://i.imgur.com/Be3I5gL.png)

Now the posts I found online suggested using VMWare to use the partition as the HDD for a Virtual machine rather than a normal Virtual Disk file. This way you can partially install Windows onto the partition and then reboot natively to that partition and complete the installation. Now in order to do this you will need to install rEFIt which is an addition to the normal EFI bootloader, this is required as the partially installed Windows will not show up in the normal menu you get when holding alt when booting. So you will need to download rEFIt and install it.

Next you will need to install a trial of VMWare fusion 5 and create the VM linked to the partition. The guide I followed to do this can be found <a title="Installing Windows on a Macbook Pro without a Superdrive" href="http://www.andrewsavory.com/blog/2011/2156" target="_blank">here</a>. The important section from this article is as follows (see original article for screenshots)

You first need to configure VMWare to use the Boot Camp partition 'raw', as a real disk rather than a virtual disk. To do this, open up a terminal and type:

```bash
mkdir Virtual\ Machines
/Applications/VMware\ Fusion.app/Contents/Library/vmware-rawdiskCreator print /dev/disk0
```

The output will be something like:

```
Nr      Start       Size Type Id Sytem
-- ---------- ---------- ---- -- ------------------------
 1          1     409639 BIOS EE Unknown
 2     409640  583984384 BIOS AF HFS+
 3  584656896  392116224 BIOS  B Win95 FAT32
```

Identify the correct drive number on the left - as a guideline, it's probably listed as FAT32 or NTFS formatted under 'Sytem' (sic).

Next, create the raw disk description:

This will give you two raw disk description files

Next, start VMWare Fusion and select "New..." from the File menu. Select "Continue without disc"

On the installation media page, drag and drop the windows7.vmdk file previously created onto the "Use an existing virtual disk" drop-down (where it says "None"). You can't just select the disk by clicking the radio button and then the drop-down, because the disk files are in a different format and will show up greyed-out.You will be prompted to convert the virtual disk to a newer format. Select "Don't Convert". If you select "Convert", you will get an error about insufficient permissions. It works just fine without converting. If all went well, you should see your vmdk listed and be able to click Continue. Accept the defaults for the operating system and version, and click Continue. On the last page, click "Customize Settings", accept the default save filename and click "Save". You'll then be taken to the Customize Settings screen, where you want to configure the ISO to boot from. Click on the CD/DVD (IDE) option and then select the location of your Windows 7 ISO

Go through preliminary install up to the point that Windows does the first reboot. At this point, power down the VM, and reboot the computer.

Now once completing this you may find that this guide works perfectly for you and if you reboot and select Windows from the rEFIt menu the installation may continue, if that&#8217;s the case then fantastic, but if like me you encounter an error then read on.

When I followed that guide and rebooted I was confronted with the message &#8220;Could not find file winboot.exe&#8221;. Now this issue is caused by the Windows boot record pointing at the wrong partition. I guess this is caused by the fact that the number and possibly order of the partitions is different on the physical machine to when it is accessed form the VM. This problem can usually be fixed by loading up the recovery console from the installation DVD and repairing the MBR but the whole reason we&#8217;re doing this is because we can&#8217;t boot from the DVD so we&#8217;re going to have to blow away this failed installation and start again from a different angle, but don&#8217;t worry this wasn&#8217;t for nothing, what this has achieved is it has made this partition bootable so what we will do now is copy the installation media to this bootable partition and then it form there onto itself.

![Paragon NTFS](http://i.imgur.com/rQnMhnD.png)

To do this you will need to be able to modify the NTFS partition from OS X. So boot back into OS X and install Paragon NTFS, you can get a free trial of this which will be fine. Once you&#8217;ve installed that you need to open the Windows partition in finder and drag it&#8217;s entire contents to the trash and then empty the trash to delete everything. This will take a few minutes but you will be left with an empty bootable NTFS partition.

Next you need to open the Windows ISO image in disk utility (right click and open with disk utility) and then in disk utility select the ISO and click mount. Go back to finder and open a second window, one with the mounted ISO and one with the NTFS partition. Select everything in the ISO and drag it all to the NTFS partition. Once this has copied you should be able to reboot and select Windows in rEFIt again and now this time instead of getting the winload.exe issue it will boot up into the Windows installation media.

You can now go through the installation as normal selecting the NTFS partition you are currently booted form as the install location, it may whine about it not being empty but you are perfectly fine to install Windows along side the installation media. Once this completes it should reboot again and you will have to select Windows again in rEFIt. It will complete the installation and you should end up with a full Windows installation.

If you like now you can navigate to the C: drive in Windows and locate all the parts of the install media and delete them (just leave the boot directory well alone). You may also want to reduce the timeout of the boot menu as it originally has a 30 second count down allowing you to choose between Windows 7 and the install media.

Now that you have Windows installed you can boot back into OS X and uninstall VMWare fusion and Paragon NTFS as we don&#8217;t need these anymore.

Finally the default EFI bootloader should detect it when holding alt during boot so you could remove rEFIt if you like. I&#8217;ve chosen to keep it on mine though as it provides a nice boot interface to select your OS (especially if you customize it).

And that&#8217;s it. If you have any questions about this or need any help drop a comment below and I&#8217;ll do my best to give a helpful response.
