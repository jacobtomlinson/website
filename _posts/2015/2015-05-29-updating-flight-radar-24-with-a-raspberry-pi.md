---
title: Updating flightradar24 with a Raspberry Pi
author: Jacob Tomlinson
layout: post
category: Projects
thumbnail: plane
tags:
- flightradar24
- fr24
- raspberry pi
excerpt: 'Using a Raspberry Pi and a DVB-T tuner to feed into flightradar24'
---

You can feed data into [flightradar24][fr24] (from now on referred to as fr24) simply using a Raspberry Pi and a cheap USB DVB-T tuner.

## How fr24 tracks planes

All the data in fr24 is collected by volunteers around the world using some very simple radio equipment. Most modern planes constantly broadcast information about themselves, including their callsign, altitude, longitude and latitude. The radio broadcast signal is called [ADS-B][ads-b], which is broadcast on a radio frequency right next to the frequency range used by terrestrial digital video broadcasting ([DVB-T][dvb-t]). DVB-T is commonly used in europe and other parts of the world and therefore it is very easy to get hold of a USB tuner. These can be modified (in software) to pick up data from ADS-B to send on to fr24.

## Equipment

All you need to collect data and send it to fr24 is a computer capable of utilising a USB DVB-T tuner and the tuner itself. One of the cheapest computers on the market at the moment, both in terms of upfront cost and electricity usage, is the [Raspberry Pi][raspberry-pi]. The Pi is perfectly capable of running a DVB-T tuner so it will be ideal for this project but there is nothing stopping you using a Beaglebone, linux server, Windows PC or whatever you have lying around and are happy to keep on 24/7.

Sadly not all USB DVB-T tuners are capable of picking up the 1090Mhz broadcast so you need to make sure your tuner has either the RTL2832U/E4000 or RTL2832U/R820T chip/tuner combinations. These are pretty cheap, I got mine for &pound;6.98 on [eBay][tuner-ebay-listing]. The antenna which comes with the tuner is fine but you might want to upgrade to a [bigger one][external-antenna].

## Preparing up your Pi

This guide is based on a fresh install of [raspbian][raspbian] so make sure you've got that set up already.

## Installing the drivers

To receive ADS-B signals through your DVB-T tuner you need to install some custom drivers.

```
apt-get update && apt-get install cmake gcc pkg-config libusb-1.0 make git-core libc-dev
git clone git://git.osmocom.org/rtl-sdr.git
cd rtl-sdr
mkdir build
cd build
cmake ../ -DINSTALL_UDEV_RULES=ON
make && make install
ldconfig
```

You'll also have to blacklist the default driver otherwise you'll get conflicts.

```
echo "blacklist dvb_usb_rtl28xxu" > /etc/modprobe.d/dvb-t.conf
```

## Installing fr24feed

Now your tuner is set up you'll need the monitoring software created by fr24 called fr24feed. In the past you will have needed to install dump1090 (the software which speaks to the tuner) and then install fr24feed which takes the data from dump1090 and sends it on to fr24's servers. However fr24 have now packaged dump1090 into fr24feed so the only thing you'll need to do is download the `.deb` and install it. Check the [fr24 website][fr24feed-repo] for the latest version.

```
wget http://feed.flightradar24.com/raspberry-pi/fr24feed_x.x.x-x_armv6l.deb
sudo dpkg -i fr24feed_*_armv6l.deb
```

## Configure fr24feed

Finally you need to register your receiver with fr24 and configure the software.

```
sudo fr24feed --signup
```

You'll be asked a few questions, just fill in your details, select yes for the data feeds and choose `Malcolm Robb's fork` for the dump1090 variant.

Once this is done you can test whether your radar is working by running `service fr24feed status`. You should see output similar to this

```
pi@raspberrypi ~ $ service fr24feed status
[ ok ] FR24 Feeder/Decoder Process: running.
[ ok ] FR24 Stats Timestamp: 2015-06-03 20:44:05.
[ ok ] FR24 Link: connected [UDP].
[ ok ] FR24 Radar: T-EGTE25.
[ ok ] FR24 Tracked AC: 4.
[ ok ] Receiver: connected (3117 MSGS/0 SYNC).
```

Congratulations, if you can see output similar to the above then you have a working feed into flightradar24 running on your Raspberry Pi.


[ads-b]: http://en.wikipedia.org/wiki/Automatic_dependent_surveillance-broadcast
[dvb-t]: http://en.wikipedia.org/wiki/DVB-T
[external-antenna]: http://shop.jetvision.de/epages/64807909.sf/en_GB/?ObjectPath=/Shops/64807909/Products/67100
[fr24]: http://www.flightradar24.com/
[fr24feed-repo]: http://feed.flightradar24.com/raspberry-pi/\
[raspberry-pi]: https://www.raspberrypi.org/
[raspbian]: https://www.raspbian.org/
[tuner-ebay-listing]: http://www.ebay.co.uk/itm/201338088784
