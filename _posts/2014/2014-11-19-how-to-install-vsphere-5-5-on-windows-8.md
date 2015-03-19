---
title: How to install the vSphere 5.5 Client on Windows 8
author: Jacob Tomlinson
layout: post
permalink: /2014/11/19/how-to-install-vsphere-5-5-client-on-windows-8/
category: VMWare
thumbnail: vmware
tags:
  - vmware
  - vsphere
  - windows 8
  - guide
---

If you've tried installing the vSphere 5.5 client on Windows 8 you may have
received the following error message

> Error 28173.Setup failed to enable Microsoft
> .NET Framework 3.5. Enable this feature in Windows Server Manager
> before installing vCenter Server.

This is because Windows 8 doesn't enable the .NET Framework 3.5 by default and
vSphere 5.5 is dependant on it.

To correct this you simply need to enable .Net 3.5 in your Windows settings.
There are two ways of doing this depending on whether you have an internet connection
or not.

### Internet method

  * Click start on Windows 8.1 or press the Windows key on 8.0
  * Select "PC Settings"
  * Select "Control Panel"
  * Select "Programs and Features"
  * Select "Turn Windows features on or off"
  * Check ".NET Framework 3.5 (includes .NET 2.0 and 3.0)"
  * Hit OK
  * When prompted select install from Windows Update

### Offline method
  * Insert your Windows 8 installation disc
  * Click start on Windows 8.1 or press the Windows key on 8.0
  * Select the arrow for all programs
  * Locate "Command Prompt", right-click on it and select "Run as administrator"
  * Run this command `dism /online /enable-feature /all /featurename:NetFX3 /source:drive_letter:\sources\sxs`
    * Where `drive_letter` is the drive letter of your Windows 8 disc

Once you've done one of these methods you'll then be able to run the vSphere 5.5
install as normal.
