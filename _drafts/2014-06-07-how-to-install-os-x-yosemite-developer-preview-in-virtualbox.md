---
title: How to install OS X Yosemite Developer Preview Beta in Virtualbox
author: Jacob Tomlinson
layout: post
permalink: /2014/06/07/how-to-install-os-x-yosemite-developer-preview-in-virtualbox/
categories:
  - Apple
tags:
  - apple
  - Developer Preview
  - Guide
  - os x
  - terminal
  - VirtualBox
  - Yosemite
---
Like me you may be excited about the Developer Preview Beta of OS X 10.10 Yosemite and want to try it out, but you don&#8217;t want to deal with a buggy system between now and the general release. If that&#8217;s the case you&#8217;ll want to install Yosemite as a virtual machine on your Mac. Here&#8217;s how I&#8217;ve done it on mine using VirtualBox.

**[Update]** Performance is pretty poor at the moment.

&nbsp;

*   First you&#8217;ll want to log into your Apple developer account and download the &#8220;OS X Yosemite Developer Preview&#8221;, it will open up the Mac App Store and automatically redeem a code for you to download the Developer Preview.

<a href="http://www.jacobtomlinson.co.uk/2014/06/07/how-to-install-os-x-yosemite-developer-preview-in-virtualbox/screen-shot-2014-06-07-at-18-11-08-2/" rel="attachment wp-att-849"><img class="alignnone size-full wp-image-849" src="http://www.jacobtomlinson.co.uk/wp-content/uploads/2014/06/Screen-Shot-2014-06-07-at-18.11.081.png" alt="Screen Shot 2014-06-07 at 18.11.08" width="1011" height="398" /></a>

*   Once you have the installation app downloaded make sure you don&#8217;t let it run or it will upgrade your Mac to the Developer preview.
*   Quit the application.
*   Before we can make Yosemite work with VirtualBox we need to make a couple of changes to the installation media.
*   Open up a terminal and run the following commands.

```
sudo gem install iesd
iesd -i /Applications/Install\ OS\ X\ 10.10\ Developer\ Preview.app -o YosemitePreview.dmg -t BaseSystem
hdiutil convert YosemitePreview.dmg -format UDSP -o YosemitePreview.sparseimage
hdiutil mount /Applications/Install\ OS\ X\ 10.10\ Developer\ Preview.app/Contents/SharedSupport/InstallESD.dmg
hdiutil mount YosemitePreview.sparseimage
cp /Volumes/OS\ X\ Install\ ESD/BaseSystem.* /Volumes/OS\ X\ Base\ System/
hdiutil detach /Volumes/OS\ X\ Install\ ESD/
hdiutil detach /Volumes/OS\ X\ Base\ System/
hdiutil convert YosemitePreview.sparseimage -format UDZO -o YosemitePreviewVirtualBox.dmg
```


*   You&#8217;ll probably then want to clean up the extra \`YosemitePreview.dmg\` and \`YosemitePreview.sparseimage\` with the following commands.


```
rm YosemitePreview.sparseimage
rm YosemitePreview.dmg
```


*   Open VirtualBox and click &#8220;New&#8221;.
    *   Give your new VM a name.
    *   Select &#8220;Mac OS X&#8221; as the type.
    *   Select &#8220;Mac OS X (64 bit)&#8221; for the version.

<a href="http://www.jacobtomlinson.co.uk/2014/06/07/how-to-install-os-x-yosemite-developer-preview-in-virtualbox/screen-shot-2014-06-07-at-18-35-53/" rel="attachment wp-att-850"><img class="alignnone size-full wp-image-850" src="http://www.jacobtomlinson.co.uk/wp-content/uploads/2014/06/Screen-Shot-2014-06-07-at-18.35.53.png" alt="Screen Shot 2014-06-07 at 18.35.53" width="653" height="399" /></a>

*   Create a disk in your desired location.
    *   Make sure you give it at least the 20GB default size.
*   <span style="line-height: 1.5em;">Give it a couple of processor cores.</span>
*   <span style="line-height: 1.5em;">Allocate at least 2GB of RAM (I gave mine 4GB).</span>

<a href="http://www.jacobtomlinson.co.uk/2014/06/07/how-to-install-os-x-yosemite-developer-preview-in-virtualbox/screen-shot-2014-06-07-at-18-05-45/" rel="attachment wp-att-842"><img class="alignnone size-full wp-image-842" src="http://www.jacobtomlinson.co.uk/wp-content/uploads/2014/06/Screen-Shot-2014-06-07-at-18.05.45.png" alt="Screen Shot 2014-06-07 at 18.05.45" width="884" height="751" /></a>

*   Now you&#8217;ll need to go into the settings and under system change the Chipset to PIIX3.
*   If you have a newer MacBook Pro open Terminal and do the following


```
cd /path/to/your/vm
VBoxManage modifyvm –cpuidset 00000001 000306a9 00020800 80000201 178bfbff
```


<a href="http://www.jacobtomlinson.co.uk/2014/06/07/how-to-install-os-x-yosemite-developer-preview-in-virtualbox/screen-shot-2014-06-07-at-18-05-53/" rel="attachment wp-att-843"><img class="alignnone size-full wp-image-843" src="http://www.jacobtomlinson.co.uk/wp-content/uploads/2014/06/Screen-Shot-2014-06-07-at-18.05.53.png" alt="Screen Shot 2014-06-07 at 18.05.53" width="692" height="605" /></a>

*   Then go into storage, select the empty disk image
*   Click the disk icon
*   Select &#8220;Choose a virtual CD/DVD disk file&#8230;&#8221;
    *   You can see my yosemite disk and a ubuntu disk from previous use.
*   Select your disk and click ok.

<a href="http://www.jacobtomlinson.co.uk/2014/06/07/how-to-install-os-x-yosemite-developer-preview-in-virtualbox/screen-shot-2014-06-07-at-18-40-39/" rel="attachment wp-att-851"><img class="alignnone size-full wp-image-851" src="http://www.jacobtomlinson.co.uk/wp-content/uploads/2014/06/Screen-Shot-2014-06-07-at-18.40.39.png" alt="Screen Shot 2014-06-07 at 18.40.39" width="955" height="667" /></a>

*   You&#8217;re now ready to boot up your VM, so click the &#8220;Start&#8221; button.
*   You&#8217;ll notice that you get a messy looking boot screen with lots of unix output rather than the normal Apple logo and loading icon, this is because it&#8217;s a Developer Preview. Eventually you&#8217;ll end up at the installation screen.

<a href="http://www.jacobtomlinson.co.uk/2014/06/07/how-to-install-os-x-yosemite-developer-preview-in-virtualbox/screen-shot-2014-06-07-at-18-06-39/" rel="attachment wp-att-844"><img class="alignnone size-full wp-image-844" src="http://www.jacobtomlinson.co.uk/wp-content/uploads/2014/06/Screen-Shot-2014-06-07-at-18.06.39.png" alt="Screen Shot 2014-06-07 at 18.06.39" width="1052" height="841" /><br /> </a>

*   From here you need to select Utilities -> Disk Utility from the top menu bar.
*   In disk utility you need to select your VBOX HARDDRIVE
*   Erase it as Mac OS Extended (Journaled) and give it a name.

<a href="http://www.jacobtomlinson.co.uk/2014/06/07/how-to-install-os-x-yosemite-developer-preview-in-virtualbox/screen-shot-2014-06-07-at-18-07-21/" rel="attachment wp-att-845"><img class="alignnone size-full wp-image-845" src="http://www.jacobtomlinson.co.uk/wp-content/uploads/2014/06/Screen-Shot-2014-06-07-at-18.07.21.png" alt="Screen Shot 2014-06-07 at 18.07.21" width="1052" height="841" /></a> <a href="http://www.jacobtomlinson.co.uk/2014/06/07/how-to-install-os-x-yosemite-developer-preview-in-virtualbox/screen-shot-2014-06-07-at-18-07-26/" rel="attachment wp-att-846"><img class="alignnone size-full wp-image-846" src="http://www.jacobtomlinson.co.uk/wp-content/uploads/2014/06/Screen-Shot-2014-06-07-at-18.07.26.png" alt="Screen Shot 2014-06-07 at 18.07.26" width="1052" height="841" /></a>

*   Once you&#8217;ve done this quit Disk Utility and start the installation.
*   Select your newly formatted disk and wait for it to install.

<a href="http://www.jacobtomlinson.co.uk/2014/06/07/how-to-install-os-x-yosemite-developer-preview-in-virtualbox/screen-shot-2014-06-07-at-18-08-08/" rel="attachment wp-att-847"><img class="alignnone size-full wp-image-847" src="http://www.jacobtomlinson.co.uk/wp-content/uploads/2014/06/Screen-Shot-2014-06-07-at-18.08.08.png" alt="Screen Shot 2014-06-07 at 18.08.08" width="1052" height="841" /></a>

Once this completed you should be greeted with something that looks like this.

<a href="http://www.jacobtomlinson.co.uk/2014/06/07/how-to-install-os-x-yosemite-developer-preview-in-virtualbox/screen-shot-2014-06-07-at-18-44-44/" rel="attachment wp-att-852"><img class="alignnone size-full wp-image-852" src="http://www.jacobtomlinson.co.uk/wp-content/uploads/2014/06/Screen-Shot-2014-06-07-at-18.44.44.png" alt="Screen Shot 2014-06-07 at 18.44.44" width="1052" height="841" /></a>

&nbsp;

Now sadly the performance of Yosemite in VirtualBox is pretty poor. But it&#8217;s good if you want to take a quick look at the features.
