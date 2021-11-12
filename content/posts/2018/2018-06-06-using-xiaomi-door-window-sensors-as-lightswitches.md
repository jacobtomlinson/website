---
title: Using Xiaomi door/window sensors as light switches
date: 2018-06-06T00:00:00+00:00
draft: false
categories:
- blog
tags:
- Home Assistant
- Home Automation
- Xiaomi Aquara
author: Jacob Tomlinson
canonical: https://medium.com/jacob-tomlinson/using-xiaomi-door-window-sensors-as-light-switches-ce29d00aa981
---

## Introduction

For a while I've been searching for a decent light switch solution for my home automation setup. I've recently put in a pretty good solution using [Xiaomi door/window sensors](https://xiaomi-mi.com/mi-smart-home/xiaomi-mi-door-window-sensors/), I'm very happy with it and it ticks a lot of boxes.

My biggest challenge was due to living in a British house built during the 90s there are no neutral wires in the wall. This means that behind my lightswitch there is a live wire in, a live wire out and an earth wire for safety. When I press the switch it connects the two live wires and completes the circuit.

![](https://miro.medium.com/max/1000/1*VBKi6FUSR2RfUcqSObb8qA.png)

This rules out smart switches for me as they also need the neutral wire to allow the smart bit of the switch to be wired in parallel to the bulb. Without the neutral wire it would result in the smart switch being wired in series which would starve the switch and bulb of electricity and would also cause the switch to lose power when it turned off the light. There are some smart switches on the market which work without a neutral wire but I also wanted to try and keep my original switches if I could.

The way I see many people deal with this is to use smart bulbs. These are great and I am happy with my [Hue](https://www2.meethue.com/en-gb?origin=rHzPFKJg&pcrid=190888095499|mckv|srHzPFKJg_dc|plid||slid||&gclid=Cj0KCQjw6pLZBRCxARIsALaaY9ZVgDXUqx_BmRyGYvDrWACSV86BSBEUH3aV0qKSRNCGhklYG1I2lZQaAuLuEALw_wcB) and [Yeelight](https://www.yeelight.com/) bulbs, however there are a few problems. When the bulb is on you can dim it and turn it off using your phone, Alexa, etc. The bulb itself is still getting power from the wall switch but the smart controller has turned off the LEDs in the bulb. However if you turn it off at the switch you are cutting the power supply to the bulb and you then can't turn it back on using the smart methods until you press the switch again.

![](https://miro.medium.com/max/1000/1*bNNwq8BopCQKyYgh60iruw.jpeg "Image credit android central")

Some people seem to work around this by leaving the switches on and generally pretending they aren't there. They insist on only using the app or voice commands to control the light. Others even go as far as buying a remote and attaching it to the wall next to the now obsolete wall switch.

One of the main things in my smart home philosophy is that smart systems should follow the web design principle of [progressive enhancement](https://en.wikipedia.org/wiki/Progressive_enhancement). This means taking the most simple and minimal version of something and improving it in ways which will fall back to the simple version if there is a problem. So in the case of light switches I should have a fully functioning smart light which I can control from my phone or Echo, but I can still simply press the light switch without breaking the new functionality.

## Xiaomi

The solution I came up with to this problem was to bypass the light switch all together and then convert the switch into a remote control for the smart bulb. This means I don't have to worry about dealing with mains electricity and can also use any switch face I want.

The initial idea was to use a terminal block or wire nut to connect the two live wires, effectively turning the light on permanently. Then get some kind of smart battery powered wireless module which would connect to the back of the switch and inform [Home Assistant](https://www.home-assistant.io/) when the switch has been toggled.

To my surprise I was unable to find any nice small wireless binary sensors which I could simply connect to the back of the switch. There were many remotes and buttons but they all used momentary switches and were not suitable for adapting to a single throw switch.

![](https://miro.medium.com/max/1000/1*_vhx0VCyEdpN9BHGn4uoUA.jpeg "Xiaomi Door/Window Sensor")

Then I spotted the Xiaomi door/window sensors. They are very tiny little sensors which you stick to your doors or windows along with a small powerful magnet which opens and closes a magnetic reed switch when the door is opened and closed. It then communicates with a [hub](https://xiaomi-mi.com/mi-smart-home/xiaomi-mi-gateway-2/) via zigbee and is compatible with Home Assistant. This reed switch behaves in a similar way to what I want. When the switch closes it sends a message to the hub, and then when it opens again it send another message.

![](https://miro.medium.com/max/1400/1*PwS6qR-ws-zVlS3QyCQGDw.jpeg "Inside the sensor you can see the reed switch")

When opening up the sensor you can see the two ends of the reed switch are easily accessible. These are the two ends which I want to connect to my switch so that instead of the magnetic switch making and breaking the connection with the magnet's force it would be the manual light switch instead.

![](https://miro.medium.com/max/1400/1*Vsw0RZDcWbtYGPQ3oCIjFw.jpeg "Attaching wires to the switch")

I could've desoldered the switch and removed it completely however I opted to attach a couple of the wires directly to the legs, this gives me the option of desoldering them again in the future if I want to convert it back into a regular sensor.

Once I attached the wires I used a hobby knife to cut a couple of slots into the plastic housing and then clipped the sensor back into its case with the two wires sticking out of the bottom.

![](https://miro.medium.com/max/1400/1*PhRtSwXcSH6ApcLI0AX96Q.jpeg "Sensor back in the shell with wires exposed")

Doing this means I still have the nice looking sensor but it is now modified to be a light switch sensor.

![](https://miro.medium.com/max/1400/1*kUF35iEdZVDi-XC6n8BQEA.jpeg "Finished light switch sensor")

The next job was to turn off the power to the house, remove the existing light switch and hard wire the circuit on. It is highly recommended that you have a qualified electrician involved in this step as your wiring may vary depending on your property.

![](https://miro.medium.com/max/1400/1*WRF9FIvS66phaWeDH8wang.jpeg "Live and switched live connected with a terminal block, bypassing the light switch")

Now I can attach my custom Xiaomi light switch sensor to the light switch itself and put it back on the wall.

![](https://miro.medium.com/max/1400/1*tzMak2f20dcTRHKXL_zA2g.jpeg "Sensor attached to the switch")

The Xiaomi sensors come with a peelable sticky pad on the back so I stuck it straight to the switch and then screwed the wires into the terminal screws where the live wires had been. This results in a neat looking switch which takes up very little space in the wall cavity.

## Home Assistant

The last step was to configure Home Assistant to toggle my smart bulb when the state of the switch changed. I decided to toggle because I want the light to change state when the switch is pressed, whether that turned the switch on or off isn't really relevant.

```yaml
- alias: Light Switch
  trigger:
    platform: state
    entity_id: binary_sensor.door_window_sensor_158d000xxxxxc2
  action:
    service: light.toggle
    entity_id: light.somelight
```

## Conclusion

I'm very happy with this solution as I can now turn lights off at the wall switch and then turn them back on via Alexa or a Home Assistant automation. I've gone through most of the switches in my house and done this now.

The only downside with this approach is that if there is a problem with the Xiaomi Hub, Home Assistant or the smart bulb hub then the switches will fail to work, that's quite a few single points of failure. This also doesn't really fit with my progressive enhancement philosophy as the enhancement has made the original functionality of the switch dependent on the smart functionality. However I feel like this is an acceptable compromise until some better solution comes along.
