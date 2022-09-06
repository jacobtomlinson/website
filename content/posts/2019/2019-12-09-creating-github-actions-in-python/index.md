---
title: "Creating GitHub Actions in Python"
date: 2019-12-09T00:00:00+00:00
draft: false
author: "Jacob Tomlinson"
categories:
  - blog
tags:
  - Python
  - GitHub Actions
  - Tutorial
thumbnail: python
---

_Note: This post is also [available in Go flavour](/posts/2019/creating-github-actions-in-go/)._

[GitHub Actions](https://github.com/features/actions) provide a way to automate your software development workflows on GitHub. This includes traditional CI/CD tasks on all three major operating systems such as running test suites, building applications and publishing packages. But it also includes [automated greetings](https://github.com/actions/starter-workflows/blob/master/automation/greetings.yml) for new contributors, [labelling pull requests based on the files changed](https://github.com/actions/starter-workflows/blob/master/automation/label.yml), or even [creating cron jobs to perform scheduled tasks](https://help.github.com/en/actions/automating-your-workflow-with-github-actions/workflow-syntax-for-github-actions#onschedule).

Also GitHub Actions is **free for open source projects** 🎉!

In GitHub Actions you create [Workflows](https://help.github.com/en/actions/automating-your-workflow-with-github-actions/configuring-a-workflow), these workflows are made up of [Jobs](https://help.github.com/en/actions/automating-your-workflow-with-github-actions/workflow-syntax-for-github-actions#jobs), and these Jobs are made up of [Steps](https://help.github.com/en/actions/automating-your-workflow-with-github-actions/workflow-syntax-for-github-actions#jobsjob_idsteps).

> Steps can run commands, run setup tasks, or run an action in your repository, a public repository, or an action published in a Docker registry. Not all steps run actions, but all actions run as a step.

In this case an [Action](https://help.github.com/en/actions/automating-your-workflow-with-github-actions/about-actions) is a piece of code which will be run on the GitHub Actions infrastructure. It will have access to the workspace containing your project and can perform any task you wish. It can take inputs, provide outputs and access environment variables. This allows you to chain Actions together for endless possibilities. You can use prebuilt Actions from the [GitHub Marketplace](https://github.com/marketplace?type=actions) or write your own.

In this post we are going to walk through creating your own [Container Action](https://help.github.com/en/actions/automating-your-workflow-with-github-actions/creating-a-docker-container-action) in Python and publishing it to the GitHub Marketplace for others to use.

To give ourselves a head start we are going to use [this `python-container-action` template](https://github.com/jacobtomlinson/python-container-action) that I've published on GitHub. We will talk through each file we need to create so you don't strictly need to use it, but it will make our lives a little easier.

## Clone the template

To get started head to the [template on GitHub](https://github.com/jacobtomlinson/python-container-action) and press the big green "Use this template" button.

![Use this template](https://i.imgur.com/AduXSKU.png)

Give your new Action a repo name and optionally a description. In this example we are going to make an Action which lints YAML files, so I'm going to name it `gha-lint-yaml`.

![gha-lint-yaml](https://i.imgur.com/piYQfQP.png)

This will give us a new repository that is essentially a fork of the template.

![Our new Action repo](https://i.imgur.com/Bido10A.png)

Also our template has a couple of GitHub Actions configured to test the Action (woo GitHub Action recursion). We can visit the "Actions" tab and hopefully see both tests have passed.

![GitHub Actions running on our Action](https://i.imgur.com/rdeeANv.png)

We will come back to these tests later.

Lastly on the GitHub side for now let's clone our repository locally.

```console
$ git clone <your repo clone url>
```

![Repo clone dialog](https://i.imgur.com/trmfLEJ.png)

## Writing our Action code

If you look in the project you will find a few files. The two we are going to focus on now are the `Dockerfile` and `main.py`.

### Dockerfile

The `Dockerfile` here is a pretty minimal Python Docker build.

```docker
FROM python:3-slim AS builder
ADD . /app
WORKDIR /app

# We are installing a dependency here directly into our app source dir
RUN pip install --target=/app requests

# A distroless container image with Python and some basics like SSL certificates
# https://github.com/GoogleContainerTools/distroless
FROM gcr.io/distroless/python3-debian10
COPY --from=builder /app /app
WORKDIR /app
ENV PYTHONPATH /app
CMD ["/cmd/main.py"]
```

We start with a `builder` container based on `python:3-slim`. We copy in our code and run `pip install --target=/app <deps>` to install our dependancies into our app directory.

We then use a [multistage build](https://docs.docker.com/develop/develop-images/multistage-build/) step to move to a very [minimal distroless container](https://github.com/GoogleContainerTools/distroless). This container image has no operating system and just contains a couple of things like SSL certificates to ensure our application can function correctly when speaking to web services securely.

We copy our Python code and dependencies over from the `builder` container and set it as our runtime command.

The result of this is a fairly small container image. The example we have here should build to around 50MB.

```console
$ docker build -t jacobtomlinson/gha-lint-yaml:latest .
...
$ docker images | grep gha-lint-yaml
jacobtomlinson/gha-lint-yaml                               latest              4eb311726658        9 seconds ago       54.1MB
```

Of course this will grow a little as we add more code, but it'll still be small compared to a container with a full linux distro in it. It also is a very secure way to build containers as we have hugely reduced our attack surface if a bad actor were to exploit our application, they wouldn't even have a shell to get into in the container.

_You can make even smaller containers using a statically compiled language like Go. There's a [tutorial](/posts/2019/creating-github-actions-in-go/) and [template](https://github.com/jacobtomlinson/go-container-action) for that too!_

### main.py

Let's also take a minute to explore the example `main.py` file that we have here before we replace it with our YAML linting code.

```python
import os
import requests  # noqa We are just importing this to prove the dependency installed correctly


def main():
    my_input = os.environ["INPUT_MYINPUT"]

    my_output = f"Hello {my_input}"

    print(f"::set-output name=myOutput::{my_output}")


if __name__ == "__main__":
    main()
```

This is a little longer than a [standard Python hello world example](https://www.digitalocean.com/community/tutorials/how-to-write-your-first-python-3-program).

We have wrapped our logic in a `main()` function and are calling it from within an `if __name__ == "__main__"` block. We are also importing `requests` a little unnecessarily here just to prove that our dependency installation in our `Dockerfile` worked, we aren't actually going to use it.

The first line in our `main()` function is getting a GitHub Action input from an environment variable. GitHub actions can take arguments using a `with` statement and these get exposed via environment variables named `INPUT_` followed by the argument name in all caps.

```yaml
name: My Workflow
on: [push, pull_request]
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
    - uses: jacobtomlinson/python-container-action@master
      with:
        myInput: world
```

In the above configuration the variable `INPUT_MYINPUT` would be created with the value `world`. We can then grab this within our code.

The next line of Python code is formatting this variable into an output string, nothing fancy here apart from [f-strings](https://realpython.com/python-f-strings/) (woo f-strings 🎉). This means our output string will now be `Hello world` in our little example workflow.

Then lastly it is printing to the stdout using some specific GitHub Actions syntax. You can pass information back to the GitHub Actions workflow with [certain control phrases](https://help.github.com/en/actions/automating-your-workflow-with-github-actions/development-tools-for-github-actions). This is an example of setting an Action's output by printing `::set-output name=<output name>::<output value>`.

This means we can access our output value within other GitHub Actions later in the workflow.

### Writing our YAML linter

This article isn't intended to be a Python tutorial, so we won't go into the detail of the code here. But here is a general outline.

Our application is going to take some inputs:

- The `path` to the file we want to lint relative to the root of the project.
- Whether we should be `strict` in our linting (error on warning).

We will then use the [third-party library `yamllint`](https://yamllint.readthedocs.io/en/stable/index.html) to lint the code, expose errors and warnings using control phrases and exit with a `0` or `1` depending on whether the file linted successfully or not.

Our action will also print an output reporting the number of warnings if we pass linting.

```python
import os
import sys

from yamllint import linter
from yamllint.config import YamlLintConfig


def main():
    yaml_path = os.environ["INPUT_PATH"]
    strict = os.environ["INPUT_STRICT"] == "true"
    conf = YamlLintConfig("extends: default")
    warning_count = 0

    with open(yaml_path) as f:
        problems = linter.run(f, conf, yaml_path)

    for problem in problems:

        if problem.level == "warning" and strict:
            problem.level = "error"

        print(
            f"::{problem.level} file={yaml_path},line={problem.line},"
            f"col={problem.column}::{problem.desc} ({problem.rule})"
        )

        if problem.level == "warning":
            warning_count = warning_count + 1

        if problem.level == "error":
            sys.exit(1)

    print(f"::set-output name=warnings::{warning_count}")

    sys.exit(0)


if __name__ == "__main__":
    main()

```

Also as we are using the dependency `yamllint` in our code we need to update the `pip install` in our `Dockerfile`.

```docker
RUN pip install --target=/app yamllint
```

## Update action.yml

Now that we have some code we need to update our `action.yml` file to let GitHub Actions how to run our code and what inputs and outputs to expect.

```yaml
name: "YAML file linter"
description: "Lint a YAML file in your project"
author: "Jacob Tomlinson"
inputs:
  path:
    description: "Path to the YAML file to be linted"
    required: true
  strict:
    description: "Run the linter in strict mode (error on warnings)"
    required: false
    default: false
outputs:
  warnings:
    description: "Number of warnings raised if lint was successful"
runs:
  using: "docker"
  image: "Dockerfile"
```

We start here by setting some metadata to describe the Action, what it does and who wrote it.

We then list our inputs and outputs, give them a description, set any default values and whether they are optional or not. For our YAML linting Action we have specified a `path` input so we can set which file to lint, this is not optional. We also have a `strict` input where we can optionally set our linter into error on warning mode. Then we let GitHub Actions know to expect an output of `warnings` from our code.

Lastly we tell GitHub how to run our Action. In this case we are using `docker` and we can either specify an image by name, or set it to `Dockerfile` which will cause GitHub Actions to build our image itself when running the Action.

_A nice future enhancement could be to set up Continuous Deployment to build and push our image to Docker Hub. This would allow us to reference our image by name here and it will already be built, saving us time whenever our Action is used. Currently our image takes around 7 seconds to build which isn't terrible, but if our Action was more complex this would be longer._

## Update README

We should also update our `README.md` file to reflect this configuration to help folks understand how to use our Action.

I'm not going to include this file here as it is a bit long but you can [check it out on GitHub](https://github.com/jacobtomlinson/gha-lint-yaml/blob/master/README.md).

I recommend you include the following sections:

- Badges showing that it is a Marketplace Action
- An overview of what the Action does
- A table of inputs and outputs
- Examples of different ways to use your Action

## Write tests

Now that we have the code, configuration and documentation for our Action we should test it out.

### Local testing

While writing the code I did some local testing. I created a new file in a directory called `tests` named `valid.yaml`.

Then I ran my Action code with my environment variables set inline to test the change.

```console
$ INPUT_PATH=tests/valid.yaml INPUT_STRICT=false python main.py
::warning file=tests/valid.yaml,line=1,col=1::missing document start "---" (document-start)
::set-output name=warnings::1
```

Running this has reported one warning using the GitHub Actions control phrase format and reported `1` warnings as an output. Great!

I also created an invalid file called `tests/invalid.yaml` to check that the linting fails.

```console
$ INPUT_PATH=tests/invalid.yaml INPUT_STRICT=false python main.py
::warning file=tests/invalid.yaml,line=1,col=1::missing document start "---" (document-start)
::error file=tests/invalid.yaml,line=4,col=1::syntax error: could not find expected ':' (None)
exit 1
```

Here we can see that an error was raised and that the command exited with a status of `1` which would fail the GitHub Actions workflow. Super!

Now let's set up some Continuous Integration to test out code every time we push up changes or someone opens a Pull Request. We can set this up using GitHub Actions!

### Testing the build

The simplest thing we can do is to test that our Python code is valid. As we used the `python-container-action` template we already have a Workflow configured in `.github/workflows/python.yml` which should work just fine for our new code.

```yaml
name: Lint
on: [push, pull_request]
jobs:
  lint:
    name: Lint
    runs-on: ubuntu-latest
    steps:
      - name: Set up Python 3.7
        uses: actions/setup-python@v1
        with:
          python-version: "3.7"

      - uses: actions/checkout@v1

      - name: Lint
        run: |
          pip install flake8
          flake8 main.py
```

This Workflow runs three Steps. We get a working Python environment with [`actions/setup-python@v1`](https://github.com/actions/setup-python) and then we check out our code with [`actions/checkout@v1`](https://github.com/actions/checkout). These are both official Actions provided by GitHub and you can find the source for both of them on GitHub too.

We then have a `run` step where we have provided a short shell script to install `flake8` and then use it to lint our code.

We don't need to change anything here as this should be enough for most simple Python applications.

### Integration testing

Now that we have tested our code is valid we can also test that our Action runs correctly on the GitHub Actions infrastructure. We can do this by having our Action run itself on push and pull requests.

_This Action inception is one of my favorite features of GitHub Actions as we can do full end-to-end integration testing without any complex setup._

There is an example in `.github/workflows/integration.yml` which we will modify to look like this.

```yaml
name: Integration Test
on: [push]
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: Self test
        id: selftest
        uses: jacobtomlinson/gha-lint-yaml@master
        with:
          path: "tests/valid.yaml"
      - name: Check outputs and modified files
        run: |
          test "${{ steps.selftest.outputs.warnings }}" == "1"
```

In this workflow we start by checking out our Actions code. Then we run our own Action on itself and set our inputs to be the same as we were doing locally (note we have omitted the `strict` input as it is optional and the default value is fine).

This Action should lint our `valid.yaml` file and raise one warning.

Then we have a `run` Action which checks whether our Action behaved as we would expect by testing that the `warnings` output correctly shows `1`.

## Push to GitHub

Next we can push our Action to GitHub and let our checks run to ensure everything works as expected.

```console
$ git add -A

$ git commit -m "Initial commit"

$ git push origin master
```

We can head to the Actions tab on our repository and watch our `Lint` and `Integration Test` workflows run.

![Lint test](https://i.imgur.com/kMO2iYn.png)

Our lint step has passed, so we have definitely pushed valid Python code.

![Integration test](https://i.imgur.com/v43FfdW.png)

Hooray our integration test has also passed. We can have a look at the outputs from our Action and the checks to see that our Action raised a warning in orange as expected and then correctly reported `1` warnings. Note that is shows the default `strict` value was filled in for us too.



## Publishing to the Marketplace

Now that we have created our Action and verified that it works as expected we can publish it to the Marketplace so other folks can find it.

Technically anyone could use our Action right now, we have already used it ourselves in our integration test. But the `master` branch is a moving target and we want to make it a little easier for folks to find it.

To publish it we need to do a GitHub release. We can do this the usual way through the releases page, but GitHub may have already detected that we have written an Action and  included a prompt to on our repo. Click the "Draft a release" button on the "Publish this Action to Marketplace" dialog or head to the releases page and click it from there.

If you followed the publish dialog the "Publish this Action to the GitHub Marketplace" checkbox will already be checked. If you went in through the releases page you will need to check this box yourself.

![Drafting a release](https://i.imgur.com/8bcBNMH.png)

GitHub has already picked up some information from our `action.yml` file including the name and description. It is also warning us that we haven't set a logo or color for our Action. This is how you set the icon you will see next to your Action in the Marketplace.

![Actions with their icons](https://i.imgur.com/1BRca09.png)

We will leave ours out for now which will just use the defaults but you can go back to your `action.yml` and set yours however you like.

You also need to choose which categories your Action belongs in. For our YAML file linter Action I've gone for "Continuous Integration" and "Utilities".

You also need to set a version tag. My preference is to use SemVer and start at `0.1.0`.

Once you click "Publish Release" you will be taken to the release page. You should notice that on the left hand side it has a little blue "Marketplace" badge which means this release is available on the Marketplace.

![Our first release](https://i.imgur.com/CbbIbDU.png)

If you click that link you will be taken to the Marketplace page for your newly published action.

![Marketplace listing for YAML file linter](https://i.imgur.com/Ib5xRjO.png)

The Marketplace listing will use your `README.md` from your repository and provide some quick links for people to get started with your Action.

## Conclusion

That's it! We have successfully published our Action on the GitHub Marketplace for others to use. They can specify a tagged version so that they can ensure their workflows are deterministic and use it in any project they like.

So to quickly recap:

- We cloned the [`python-container-action` template](https://github.com/jacobtomlinson/python-container-action).
- We wrote a small utility application in Python to lint YAML files.
- We updated our `action.yml` file to describe the inputs and outputs of our application.
- We updated our `README.md` with documentation to help people get started with our Action.
- We wrote some tests to make sure our Action works as expected when run in a GitHub Actions Workflow.
- Finally we published it to the GitHub Marketplace by making a release.

You can have a look at the [YAML file linter Action here](https://github.com/jacobtomlinson/gha-lint-yaml) and even use it in your own workflows.

If you find this tutorial useful or even end up publishing your own GitHub Action on the Marketplace let me know in the comments below or on [social media](https://twitter.com/_JacobTomlinson)!
