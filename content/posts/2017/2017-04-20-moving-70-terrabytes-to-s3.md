---
title: Moving large volumes of data to S3
date: 2017-04-20T00:00:00+00:00
draft: false
categories:
- blog
tags:
- aws
author: Jacob Tomlinson
thumbnail: aws
canonical: http://archived.informaticslab.co.uk/aws/2017/04/20/moving-70-terrabytes-to-s3.html
canonical_title: the Informatics Lab Blog
---

_Originally published on the [Met Office Informatics Lab blog](http://archived.informaticslab.co.uk/aws/2017/04/20/moving-70-terrabytes-to-s3.html) on April 20th, 2017._

---

We just moved ~80TB of data to S3 (stay tuned to hear what we're doing with it).

### The problem

We are currently working on a project called [Jade][jade] which involves putting together a toolset for analysing very large amounts of data. We're doing lots of testing in [Amazon Web Services][aws] using scalable infrastructure to give these tools a thorough road test, but in order to test big data tools we need some big data.

One of our aims is to shift from analysing gigabytes of data in hours to terabytes in minutes or seconds. Therefore to do some reasonable tests we need multiple terabytes of data to practice on, but getting that much data into a cloud service like AWS is non-trivial.

We are fortunate to have some big pipes to the internet at the Met Office, however operational services get priority and the rest has to be shared between over a thousand employees. Therefore, an individual user is only able upload data at around 15-50mbps depending on what else is going on. Normally that feels like a decent amount of bandwidth for one person but assuming I can constantly use 50mbps I would only be able to upload just over half a terabyte of data per day. So uploading a 70TB dataset would take nearly six months.

### The solution

In a word, [sneakernet][xkcd-sneakernet].

> "Sneakernet is an informal term describing the transfer of electronic information by physically moving media such as magnetic tape, floppy disks, compact discs, USB flash drives or external hard drives from one computer to another; rather than transmitting the information over a computer network."
>
> _Source - [Wikipedia][sneakernet]_

It sounds like a joke at first, but when you do the maths it turns out that copying data to a hard drive and transporting it somewhere else is still a pretty efficient way to move large amounts of data around.

Luckily for us Amazon Web Services provide a robust and secure way of doing this called [AWS Snowball][aws-snowball]. Taking its name from their archive storage system [AWS Glacier][aws-glacier], Snowball is a service for importing big chunks of data into their cloud storage. It works by posting you a very large and durable hard drive appliance, you simply fill the drive up with data and then ship it back to them.

![AWS Snowball](https://i.imgur.com/J4jD23mh.jpg)

### Practicalities

#### Regions and availability

We decided we wanted to put the data into [S3][aws-s3] in the new [London region][aws-london-region] of AWS. As that region had only been open for a number of weeks we were unable to order the newer 100TB [Snowball Edge][aws-snowball-edge] and so went for the older 80TB Snowball.

There are two applications provided for interacting with the Snowball. The first is the [Snowball Client][aws-snowball-client] which is a command line utility for copying data to the device using syntax similar to the unix `cp`, `mv` and `rm` commands. The other is the [Snowball S3 Adapter][aws-snowball-s3-adapter] which connects to the device and then serves up an S3 API compatible endpoint, allowing you to use any S3 compatible tool such as the [aws-cli][aws-cli], or the [boto][boto] python client. The Snowball Edge runs the S3 adapter itself and therefore you do not require anything other than an S3 compatible tool, however the older models require you to install one or other of the supplied applications yourself on a separate machine.

#### Installing the device

To get the best performance from the device you need to ensure there are no bottlenecks to slow down the transfer. Every connection between the Snowball and the storage must be 10gbps or higher. Therefore it must connect to a 10gbps copper or optical SFP+ port on your network. You need a rather powerful workstation or client to run the software and that needs to be connected to the network with a 10gbps connection. The workstation/client also needs to have access to the data via a dedicated 10gbps link.

[Pricing for the device][aws-snowball-pricing] is an up front payment for the first 10 days you have it on site with additional charges for each day after that. Therefore we wanted to have engineers ready to install the device as soon as it arrived. As the data was on disk in our data centre we decided to install the Snowball in an adjacent rack where it could be connected to a high speed switch. This was not a problem technically, however it did involve some pre-planning. We calculated that transferring at full speed would take just under a day, but we wanted as larger margin for error as possible as we had never done this before.

#### Running the transfer

We decided to copy the data using the Snowball client, which in hindsight was probably as a mistake as it seems to be unofficially deprecated in favour of the s3 adapter. Sadly we came to this realisation too late and didn't want to start again by changing tool mid transfer.

Despite running the software on a machine with 24 CPU cores and 256GB or memory we repeatedly encountered "out of memory" errors. After troubleshooting this we discovered a hard limit of 8GB set in the shell wrapper for the tool. Once we raised this to something higher, the memory problems went away.

This also lead on to a few further issues with listing the contents of the Snowball and getting it to resume the transfer. We were copying ~2,000,000 files onto the device and ended up starting from scratch after a failure at 20%  of files transferred because we simply could not get the transfer to resume. In the Snowball documentation it [recommends][aws-snowball-recommendations] breaking your data down into smaller chunks separated by directory and doing multiple copy commands, this suggests they are aware of the resuming issues but have favoured a workaround. This is probably a pragmatic decision as doing a list operation on an object store can be very slow.

The final challenge we faced was around transfer speed. Despite the device being connected via a 10gbps link the data was only copying at around 1.5gbps despite having lots of cpu and memory available. We spent a substantial amount of time troubleshooting to check if there were any bottlenecks between our systems and the device but were unable to find any. This resulted in us just scraping past the 10 day deadline. Speaking with an AWS Solutions Architect it was suggested that the device may have been running at a reduced capacity due to some damage in transit and that it would be investigated on it's return to them. We were probably just unlucky.

### Summary

Overall we were very pleased with the relative ease and speed of using a Snowball rather than a broadband link. Next time we do this we will definitely check out the new Snowball Edge and the s3 adapter as it should simplify the process for us and should stop us from bumping into some of the issues we had.

[aws]: https://aws.amazon.com
[aws-cli]: https://aws.amazon.com/cli/
[aws-glacier]: https://aws.amazon.com/glacier/
[aws-london-region]: https://aws.amazon.com/blogs/aws/now-open-aws-london-region/
[aws-snowball]: https://aws.amazon.com/snowball/
[aws-snowball-client]: http://docs.aws.amazon.com/snowball/latest/ug/using-client.html
[aws-snowball-edge]: https://aws.amazon.com/snowball-edge/
[aws-snowball-pricing]: https://aws.amazon.com/snowball/pricing/
[aws-snowball-recommendations]: http://docs.aws.amazon.com/snowball/latest/ug/transfer-petabytes.html
[aws-snowball-s3-adapter]: http://docs.aws.amazon.com/snowball/latest/ug/snowball-transfer-adapter.html
[aws-s3]: https://aws.amazon.com/s3/
[boto]: https://github.com/boto/boto3
[jade]: http://www.informaticslab.co.uk/projects/jade.html
[sneakernet]: https://en.wikipedia.org/wiki/Sneakernet
[xkcd-sneakernet]: https://what-if.xkcd.com/31/
