---
title: ChatOps - Automation via chat
date: 2017-12-19T00:00:00+00:00
draft: false
categories:
- blog
tags:
- worksops
- chatops
- opsdroid
- python
author: Jacob Tomlinson
thumbnail: chat
canonical: http://archived.informaticslab.co.uk/worksops/2017/12/19/chatops-workshop.html
---

_Originally published on the [Met Office Informatics Lab blog](http://archived.informaticslab.co.uk/worksops/2017/12/19/chatops-workshop.html) on December 19th, 2017._

# ChatOps - Automation via chat

This article is a companion to a workshop on using chat to automate ops workflows. This is a static version of a [Jupyter Notebook](http://jupyter.org/) which you can download [here](https://images.informaticslab.co.uk/misc/f7d8eb8d252480b0028a6246623b3d76.ipynb).

_This has been tested on the [Informatics Lab Jade platform](http://www.informaticslab.co.uk/projects/jade.html) but should work in any Jupyter Notebook server running Python 3.5._

## Why would you want to use a chatbot for ops workflows?

> ChatOps is a collaboration model that helps to connect people, process, tools, and automation into a transparent workflow (conversation-driven development). So it allows to automate tasks and collaborate, encourages teams to be transparent, working better, cheaper and faster.
>
> _https://github.com/exAspArk/awesome-chatops_

Often when working in an infrastructure engineering role you spend a lot of time working on your own, in your own terminal. Unless you are pairing with someone else or very good at documenting things (and let's face it none of us are) then that knowledge gained from solving that problem gets trapped in your head and isn't easy to share.

Working in a ChatOps paradigm moves this work into a communal chat room where others can see what is going on and learn from your work.

It also encourages you to write scripts which can be used in plain english, rather than expecting others to read a man page or help text which explains exactly what characters to type. There are downsides to this of course, often you want commands to be explicit and you don't want your bot to do unexpected things if it misunderstands you. This puts more pressure on the script writer to include logical checking, perhaps even getting the bot to ask the user to confirm the action before performing.

## What technologies are available for creating simple chatbots?

Popular options to create a ChatOps bot are [Hubot](https://hubot.github.com/), [Lita](https://www.lita.io/) and [Errbot](http://errbot.io/en/latest/) which allow you to write scripts in JavaScript, Ruby and Python respectively.

For a while in the Informatics Lab we used [Hubot](https://hubot.github.com/) with good success, however there were limitations and missing features that I wanted to use. So in my spare time I decided to create [opsdroid](http://opsdroid.github.io/) which is a Python bot which tries to address the [limitations](https://github.com/opsdroid/opsdroid/issues/1) I found in the other technologies. Therefore I am totally biased and will be focusing on opsdroid for the rest of this workshop, however the other technologies are also good and I would recommend them.

## Creating a simple opsdroid chatbot in Python

_This guide assumes you are running through this in a notebook on a system which has never run opsdroid. If you're following the blog post please amend the scripts accordingly to your environment._

First we need to install opsdroid. You will need to have Python 3.5 and pip already installed on your system.


```bash
pip install opsdroid
```

### Create a workspace
Next we should create an opsdroid directory for us to work in. We'll clone a "Hello World" skill here to get us started and also use it as a place to put logs.


```bash
# Make the directory
mkdir -p $HOME/opsdroid

# Clone the "Hello world skill"
git clone https://github.com/opsdroid/skill-hello.git $HOME/opsdroid/skill-myskill

# Create the log file
touch $HOME/opsdroid/opsdroid.log
```

### Add some configuration
When you run opsdroid for the first time a [yaml](http://yaml.org/) config file will be created for you called `~/.opsdroid/configuration.yaml` with some sensible defaults. However for this workshop let's create our own config file which loads the example skill we cloned a second ago.


```bash
# Create the opsdroid config directory
mkdir -p $HOME/.opsdroid

# Cat a config file into it using a heredoc
cat <<EOF > $HOME/.opsdroid/configuration.yaml

################################################
##
## Our opsdroid configuration file
################################################

## Set the logging level and location
logging:
  level: debug
  path: $HOME/opsdroid/opsdroid.log
  console: false

## Connector modules
connectors:
  - name: shell

## Skill modules
skills:

  ## Our custom skill
  # This is the example skill we cloned from GitHub before that we will customise later. We are setting it
  # not to cache so that when we edit the code later opsdroid will reload it correctly.
  - name: myskill
    path: $HOME/opsdroid/skill-myskill
    no-cache: true

  ## Developer Tools - a skill which reloads all skills when you say `reload` to the bot.
  # NOTE: You don't need to specify a path for official skills, opsdroid will assume they are on GitHub
  # and try to download them automatically.
  - name: devtools

EOF
```

### Open a terminal
To test our bot we are going to use the `shell` connector. There are a whole range of connectors you can configure in opsdroid, for example `Twitter`, `Facebook Messenger` and `Slack`. The `shell` connector is the most basic and provides a simple command line interface to chat with your bot, this is useful for testing and development.

To run our bot and get the opsdroid shell we'll need to open a terminal. Switch back to the directory view tab and click `New > Terminal`.

The terminal will open in a new tab, you may want to place it side-by-side with this one so you can continue following the instructions.

### Run opsdroid and say hello

In the terminal simply run `opsdroid`. This should immediately drop you into the opsdroid shell. You can test it out by typing `hello`, the bot should say hello back.

_If you are using the official Jupyter docker image you will have a default username of `jovyan` which is the [name given to a user of Jupyter Notebooks](https://github.com/jupyter/docker-stacks/issues/358), so that's why it will call you `jovyan`._

![Hello opsdroid](https://i.imgur.com/4967YRs.png)

### Customise our skill

Now that we have tested our skill let's change it to do some more things. If you switch back to the directory view tab and refresh the page you should see our `opsdroid` directory. If you click into it and then into the `skill-myskill` directory you should see the basic files which make up an opsdroid skill.

_If you have experience with Python development you may notice that a skill is just a Python module._

![Example skill layout](https://i.imgur.com/Ea0SHca.png
)

Click the `__init__.py` file to open up the skill code.

![Example skill code](https://i.imgur.com/CMtK1TJ.png)

### The layout of a skill

An opsdroid skill is a Python function which takes three arguments:

- `opsdroid` a reference to the core opsdroid object.
- `config` the section from the config file which relates to this skill.
- `message` an object containing the message from the user long with the users name, which connector the message came from and a `respond` method for sending messages back.

_Technically it is a [Python coroutine](https://docs.python.org/3/library/asyncio-task.html#example-hello-world-coroutine) rather than a function, hence the `async` at the begining._

These functions must be decorated with an opsdroid **matcher**, this is how opsdroid decides which function to call when a message is received. In this workshop we are going to focus on the regex matcher which simple matches the message against a regular expression, however there are more complex matchers which use third party Natural Language Understanding services like [LUIS](https://www.luis.ai/home) or [Lex](https://aws.amazon.com/lex/).

When a message comes into opsdroid it is tested against each matcher to see if the function applies to that message, if multiple functions match they are ranked based on match quality and complexity and then the one with the highest score is called.

_There is also a `setup` function which is called when your skill is loaded just in case you need to do some prep._

### Hello world

Let's focus on the hello world function.

```python
from opsdroid.matchers import match_regex
import random

...

@match_regex(r'hi|hello|hey|hallo')
async def hello(opsdroid, config, message):
    text = random.choice(["Hi {}", "Hello {}", "Hey {}"]).format(message.user)
    await message.respond(text)
```

This function makes a random choice between three strings `Hi`, `Hello`, and `Hey`. Each is followed by a placeholder for the user's name which is inderted by the `format(message.user)` call.

Then when it has chosen a string and formatted it responds to the message with this string.

This function matches against any message which contains the words `hi`, `hello`, `hey` or `hallo`. This is a very basic regular expression which doesn't do much checking. For example those words could exist within other words and would make no sense to respond to.

![Bad match](https://i.imgur.com/Fc1Ded2.png
)

In the chat above I said `do not say hello to this message` which not only matches `hello` in that sentence but also the word `hi` within the word `this`. Causing the bot to respond.

Let's make this regex more strict by limiting the match to only be when the message is **only** one of those words, not containing one of them.

Switch back to the tab where you opened `__init__.py` and update the regex to look like this:

```python
# Old
@match_regex(r'hi|hello|hey|hallo')

# New
@match_regex(r'^(hi|hello|hey|hallo)$')
```

The `^` and `$` operators ensure that there is no other text between the beginning and end of the line and the brackets make the group of words explicit.

Now if you tell the bot to `reload` and then try that message again it will not match because there are other characters either side of the words `hello` and `hi`.

![HiHello](https://i.imgur.com/AWLdW3O.png)

## Automating interesting things with your bot

To do anything truly useful with your bot yo're going to need to be able to access and control other things. The best way to do this is via an API. Below is a simple demo of requesting some information from and API and returning it to the user.

_Fun fact this API is totally not safe for work, and I didn't know that before doing a live version of this workshop._

```python
import aiohttp

@match_regex(r'tell me about Chuck')
async def chuck_facts(opsdroid, config, message):

    # Get a new http client session
    session = aiohttp.ClientSession()

    # Request a new Chuck Norris fact
    response = await session.get('https://api.chucknorris.io/jokes/random')

    # Parse the json response
    body = await response.json()

    # Respond with the fact that was returned
    await message.respond(body["value"])
```

![Chuck fact](https://i.imgur.com/GYQFGHY.png)

