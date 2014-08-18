---
title: Additional HDD Philosphy
author: Jacob Tomlinson
layout: post
permalink: /2010/12/25/additional-hdd-philosphy/
category: Hardware
thumbnail: hard-drive
tags:
  - Hard Drive
  - Storage
---

A friend has asked me what I think the best way for him to set up his new HDD is. He currently has a 500GB Internal HDD and has purchased a second one. So here is what I would do.

### RAID  
This is the best way I would recommend setting up your PC if you&#8217;re using more than 1 HDD. But in order to do this you would need the 2 500GBs and an additional say 80GB-160GB HDD.  
In this configuration I would set up the 2 500GBs in RAID 0 so they would show up in your computer as a 1TB HDD but it would have double the read/write time of a normal 1TB drive, then I would install my operating system(s) on the additional drive. That way you can store all your files on a fast large drive and if you ever feel you want to reinstall your operating if it&#8217;s getting a bit cluttered then you won&#8217;t lose much. The only problem with this method is it takes a lot of messing about and getting additional drives.

### 3 Partitions  
The way I am recommending that my friend does it would be to have 2 500GB drives but with 3 partitions. He currently has 1 500GB drive which is full inside his pc with Windows 7, games, recorded TV, etc. And a new HDD waiting to be installed. Firstly I would backup any documents which you absolutely cannot lose, financial data etc to an online cloud storage system like <a title="Dropbox" href="https://www.dropbox.com/home" target="_blank">Dropbox</a> or <a title="Ubuntu One" href="https://one.ubuntu.com/" target="_blank">Ubuntu One</a>. Then I would install the new drive and use a Live Disc or OS installation disc to partition the drive into an 80GB and a 420GB. That way you can install your operating system on the smaller partition and wipe accordingly as mentioned above. Any games or large software suites you have installed and don&#8217;t mind leaving them on your old HDD we will discuss in a moment. I would now copy all files, documents, recorded TV etc, to your 420GB partiton on your new HDD. Then when you have moved all your files around have have everyting arranged in the way that you like I would delete the windows and users folders (documents and settings for pre-vista) from the old HDD. In theory this will give you an 80GB OS partition, a 420GB file storage partition and a 500GB file storage partition with your progam files folder still there. Lastly I would gather all your discs for games and software suites and begin to reinstall but when you are asked where you would like to install the software point it to the already existing installation on your old HDD. If the installer is clever like Steam it will see that the files are already there and skip copying them to complete the installation faster. This is only necessary because all the registry entries need to be recreated on your new OS (assuming your on Windows). Once you have complete this you should have all your files and OS in 2 partitions on your new HDD and all games and programs on your old HDD.

### Final Notes  
If you are adding drives to your computer and you are going to install a fresh copy of your OS to a new drive, while keeping the old drive, make sure that the MBR (Master Boot Record) is being put on the new HDD as well and that you have changed the BIOS to boot from this new drive. I only mention this because I have made the mistake of not doing this and having a small drive which I wanted to remove from my PC but couldn&#8217;t because it contained the MBR but was telling the pc to boot from my newer HDD anyway. The best solution I found to this happening is before installing your new OS disconnect all your other HDDs and install, then reconnect them afterwards.
