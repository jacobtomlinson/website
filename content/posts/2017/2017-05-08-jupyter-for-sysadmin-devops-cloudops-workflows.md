---
title: Using Jupyter notebooks for SysAdmin, CloudOps and DevOps workflows.
date: 2017-05-08T00:00:00+00:00
draft: false
categories:
- blog
tags:
- process
author: Jacob Tomlinson
thumbnail: jupyter
canonical: https://archived.informaticslab.co.uk/process/2017/05/08/jupyter-for-sysadmin-devops-cloudops-workflows.html
canonical_title: the Informatics Lab Blog
---

_Originally published on the [Met Office Informatics Lab blog](https://archived.informaticslab.co.uk/process/2017/05/08/jupyter-for-sysadmin-devops-cloudops-workflows.html) on May 8th, 2017._

---

[Jupyter notebooks][jupyter] are awesome. If you speak to a data scientist or analyst who writes Python there's a very good chance that they use Jupyter notebooks. But I think there's another community that would benefit hugely from including them in their standard arsenal of tools, and that's folks in IT Infrastructure.

![Jupyter notebook](https://images.informaticslab.co.uk/articles/article-jupyter/29aec8999f72db598aa8a0b5b7433d9f.png)

I spend a vast amount of my time staring at a terminal window. I'm usually running Python scripts from my bash prompt (or zsh if I'm feeling like a hipster). I document lots of our processes in some kind of wiki and I occasionally create dashboards in HTML which are updated by a cronjob that generates a json data file. This is a pretty typical workflow for anyone who calls themselves a SysAdmin, CloudOps/DevOps Engineer or a similar infrastructure type role.

In the Informatics Lab we have spent a large amount of time building a data analysis platform for data scientists and analysts. The primary way of using the platform is through our hosted [Jupyter Notebooks][jade-notebooks] service. This has resulted in me spending quite a lot of time using notebooks and I can definitely say I've got the bug.

Notebooks are great for analysts as they can make some notes in markdown, perhaps to give an introduction to the work they are about to do. They can then write some Python (or bash, or Julia, or R, etc) which loads in some data and analyses it. They can mix in more markdown notes to give context and detail on what they are doing. The code they write may render a graph or a plot which Jupyter shows inline in the notebook. They may even go as far as including widgets like sliders or drop downs that change the value of a variable and rerun some chunk code. This results in a contained and coherent train of thought which can easily be shared with or demonstrated to others. Everything is in one place, the documentation, the code and the output. In recent versions of Jupyter there is even a [dashboard mode][jupyter-dashboards] that lets you hide all the code away and just present the notes, outputs and widgets for an end user to explore.

Hopefully this sounds similar to my infrastructure workflow above. Whether it's patching your servers, deploying new packages onto a system, building cloud infrastructure with a tool like cloudformation or terraform, it's the same mix of things. You need documentation, code and output.

Take patching for example. I may have a recurring task every 30 days to read through the new errata from RedHat, merge them into my local repository, ssh into a bunch of servers, run `yum update -y && reboot` and then keep an eye on my monitoring to check that all services come back as expected. It's also likely that I may have to run this a few times on patch day to do each environment individually or to do production environments in sections to maintain availability and SLAs.

I could very easily document the patching process in markdown in the notebook; which servers are done at which times, what to look out for when checking errata, who to notify before starting, etc. I could then have a drop down widget that lets me select a stack of servers and some code cells which run bash commands to merge the errata and then ssh into the machines and run the update and reboot command. I could have some Python cells which hit the API of my monitoring stack and display data about the service I am currently patching in a table or as a graph. The notebook could also be put into dashboard view showing the status of the patching and the dashboard could be shared with higher ups so they can have visibility.

I could then share this notebook with the rest of my team so that next month if the task falls to someone else they can just run through the notebook. No more copying and pasting commands from a wiki. No more alt+tab'ing between your terminal, monitoring page and documentation. Everything you need in one place.

Jupyter notebooks are definitely a tool which fit nicely in the infrastructure world. I highly recommend you try them out!

![Jupyter notebook doing AWS things](https://i.imgur.com/wr3ZABQh.png)

[jade-notebooks]: http://www.informaticslab.co.uk/technology/2016/09/12/try-jade.html
[jupyter]: http://jupyter.org/
[jupyter-dashboards]: https://github.com/jupyter/dashboards
