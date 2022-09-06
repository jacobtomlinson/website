---
title: Cracking Enigma with Go
date: 2016-06-02T00:00:00+00:00
draft: false
categories:
- blog
tags:
- learning
author: Jacob Tomlinson
thumbnail: go
canonical: https://archived.informaticslab.co.uk/learning/2016/06/02/go-enigma.html
canonical_title: the Informatics Lab Blog
---

_Originally published on the [Met Office Informatics Lab blog](https://archived.informaticslab.co.uk/learning/2016/06/02/go-enigma.html) on June 2nd, 2016._

---
_Can I crack the Enigma code with Go on a MacBook?_ **Yes!**

I definitely can crack the [Enigma][enigma] code on a MacBook using an emulator I've written in [Golang][golang]. It might just take a while!

## Introduction
In the Informatics Lab we have been spending a lot of time exploring parallelisation technologies to find new ways of performing analysis on our ever expanding datasets. We have also been investigating scaling technologies for enterprise IT systems.

Whenever you are exploring technologies you need a few standard problems with known answers to try and solve. These problems need to range from a simple "Hello World", to something representative of what you want to run in production, and a few in between.

After watching the film [The Imitation Game][imitation-game] I have become very interested in how the Enigma code was cracked during WW2. The simplest solution to cracking any cipher is to try a brute force approach of testing every possible combination. Computers are very good at this as they can test many combinations per second and depending on the complexity of the cipher will return the answer in a short amount of time. However at the time there was no such thing as a computer, therefore [Alan Turing][alan-turing] and others designed and built a machine to brute force Enigma, this machine was a very early version of what we think of as a computer.

After learning about this I began to think that if the best minds of 1939 could crack Enigma by inventing the computer then in 2016 it would be possible for me to do the same with the aid of a modern computer.

## Step one: Learn how an Enigma machine works
Enigma Machines are an electronic device which looks similar to a type writer. It has a keyboard with the letters A-Z, some lights with the letters A-Z and some electronic mappings to jumble them up. When you press a key on the keyboard one of the lights turns on, giving you a ciphered letter.

![Enigma Machine](https://i.imgur.com/bh2zFIyh.jpg)

The groundbreaking thing about this machine was that the mapping between the letters is built into a series of rotating cogs called rotors. Whenever a key is pressed the rotors turn, changing the mapping between all the letters. Therefore if you encode the message `AAAAA` it will not give you `HHHHH`, you will in fact get something like `JKUFG`.

The other clever thing is that the machine includes a component called a reflector. This is a static rotor with a fixed mapping. Once the signal has passed through all rotors it goes through the reflector and then back through the rotors. This symmetry allows you to decode the message by simply typing it in, the same way you encoded it originally. You just reset the rotors to the positions they were in when you encoded and type the cipher text in and you will get the plain text message back.

The first version of the machine used during the war was called an Enigma M3. It came with 5 different rotors, of which you could use any three in any order and each one has 26 different starting positions. These rotors could also be taken apart and realigned internally in 26 different modes. There were two different reflectors and also a set of cables and holes called a plugboard which allowed you to manually swap letters. All of this together means there are over 158 million million million different starting positions for the machine. In order to send a message to someone you must both know the start position you are going to use. You encode the message starting at that position, transmit it to them, they decode it using the same setting and they can read your message. Anyone intercepting that message would have to guess the starting position from the possible 1.58x10<sup>20</sup> possibilities.

## Step two: Writing an emulator
The machine built during WW2 to crack the code was a large array of Enigma Machines wired together. Firstly I don't have any spare machines lying around and secondly I have a computer which is capable of simulating the movements of the machines much faster than they can physically move. Therefore I decided to write an application which given the start position and a plain text message would return the cipher text.

I decided to do this as a [JavaScript library][enigma-js] so that I could eventually use the library to build an interactive web page which would allow anyone to encode and decode Enigma messages. I also decided for simplicity that instead of figuring out the mathematical formulas the machine employs and encoding using them that I would emulate the movements of the machine itself, passing the letters from rotor to rotor and translating them.

As you can imagine when cracking this kind of problem performance is key. JavaScript is not a language which is known for performance. Before I even began trying to brute force codes with my library I realised that it would be too slow.

## Step three: Writing a better emulator
If I wanted to crack this using a standard laptop the code would need to be much faster. I realised that this could be a good opportunity to learn a new language, one based on performance. For the last year I've been working a lot with [Docker][docker] and [Kubernetes][kubernetes], these are both written in [Golang][golang] (also simply called Go). Go is a strongly-typed compiled language which was developed at Google to replace C++ in their applications.

I began porting code and writing a [Go Enigma library][engima-go] and once it was correctly implementing Enigma it became apparent that it would be much faster. My initial benchmarks showed that it could test around 21,000 settings per second. During the war the Axis forces would change their settings daily, therefore my cracker must be able to find the correct setting in less than 24 hours. My Go was fast but it still could only test around 1.8 billions settings per day which is about 11 orders of magnitude away from cracking the code.

I then spent some time optimising my code, this included removing all usage of string libraries, passing pointers to functions instead of values and general loop optimisations. This has increased performance from 21,000 to 262,000 attempts per second. This is an order of magnitude faster than before but still many orders away from cracking the code in a reasonable time.

## Step four: More understanding Enigma
My Go code can now test around a quarter of a million settings per second, but is still nowhere near fast enough. I think it is fair to assume that my Go code is far more efficient than the brute force machine created during WW2. Therefore it can't have been possible for them to crack the code using just that machine.

Reading more into the [history on the Enigma Machine][enigma-history] it seems that the mathematicians working on the problem found a way, given enough encrypted messages using the same setting, to remove the complexity added by the plugboard. I haven't read enough into this to fully understand how the maths works, or to figure out how to implement it programatically. However if we assume it is possible then that would reduce the number of settings to a more manageable 7,413,978,624.

Brute forcing a code which has already had the plugboard obfuscation removed mathematically would take a mere 7.8 hours using the Go emulator on a single core of my MacBook Pro. I could also easily pay a few dollars for a powerful [AWS instance][aws-ec2] with many more cores and crack the code in a few minutes.

## Conclusion
This has been an interesting side project so far. The problem is complex enough to stretch my brain and has therefore been enjoyable from a personal point of view.

It is also simple enough to implement in code without hitting too many gotchas. Next time I want to try a new language I will almost certainly implement an Enigma emulator in it.

When looking at the code for this project it can be split into two parts, the emulator and the brute forcer which implements the emulator. When teaching students or interns about programming they are both interesting examples to get stuck into. It would be very easy to say "Here is an Enigma library, write a brute forcer that utilises it" and then follow that exercise up with a "Now implement the library yourself".

But most importantly the binary produced from my Go code will be very useful for stress testing parallelisation tools and technologies. While we investigate things like [Docker Swarm][docker-swarm], [Kubernetes][kubernetes], [Hadoop/Spark][big-data] and more it will be a good benchmark to see how well they parallelise.

[alan-turing]: https://en.wikipedia.org/wiki/Alan_Turing
[aws-ec2]: https://aws.amazon.com/ec2/
[big-data]: http://www.informaticslab.co.uk/projects/hadoop.html
[docker]: http://www.informaticslab.co.uk/lab-school/2015/06/24/lab-school-docker.html
[docker-swarm]: http://www.informaticslab.co.uk/infrastructure/2015/12/09/raspberry-pi-docker-cluster.html
[enigma]: https://en.wikipedia.org/wiki/Enigma_machine
[engima-go]: https://github.com/jacobtomlinson/enigma-go
[enigma-history]: http://www.codesandciphers.org.uk/enigma/
[enigma-js]: https://github.com/jacobtomlinson/enigma-js
[golang]: https://golang.org/
[imitation-game]: http://www.imdb.com/title/tt2084970/
[kubernetes]: http://www.informaticslab.co.uk/infrastructure/2015/10/01/building-with-kubernetes.html
