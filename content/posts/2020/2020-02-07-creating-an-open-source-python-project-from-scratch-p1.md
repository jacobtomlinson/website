---
title: "Creating an open source Python project from scratch"
series: ["Creating an open source Python project from scratch"]
date: 2020-02-07T00:00:00+00:00
draft: false
author: "Jacob Tomlinson"
categories:
  - blog
tags:
  - Python
  - GitHub
  - Tutorial
  - Git
  - OSS licensing
thumbnail: python
---

Have you had a great idea for an open-source Python library that you think people will find useful, but you don't know where to begin in creating and publishing it?

In this weekly series, we will walk through building, testing, documenting and publishing a Python package from start to finish and then explore building a community and maintaining the project into the future. We will discuss each step generally and then dive into a specific implementation so that in the end we will have a complete example package.

# Code

To start our project we will want to write some code. In these posts we will be writing a very simple bit of code and then mainly focus on structuring the project, creating supporting files and adding all the details that people expect in a Python package.

For our example project, we are going to create a Python version of the JavaScript [is-number](https://www.npmjs.com/package/is-number) package. We will provide a function which takes an input and returns `True` or `False` depending on whether the input is a valid number.

Let's start by making a directory for the project.

```bash
mkdir is-number
cd is-number
```

You should make a directory for your project and open it up in your favourite editor.

## README

The first file we are going to create is our `README` file. This file will often be the first thing people look at when they visit your project and should contain information on how they can get started.

In Python projects it is common to use [reStructuredText (RST)](https://en.wikipedia.org/wiki/ReStructuredText) for the documentation markup language, however [markdown](https://en.wikipedia.org/wiki/Markdown) is also very popular. For our example here we are going to stick with RST, so let's create a file called `README.rst`.

```rst
is-number
=========

A Python library to determine if something is a number.
```

We will just start with a basic file which describes what the project does in a sentence or two. We will come back to this and improve it as we go on.

## Functionality

Next, we need to write some code. You probably already have something in mind for your project. Perhaps you've had an idea for a new library, or may even have some code in another project that you want to extract into a standalone open-source library.

For our example library we are going to create a library called `is_number` with the following structure:

```
is_number/ Top-level package
          __init__.py Initialization file for the package
```

That file will contain our simple function for checking if something is a number.

```python
def is_number(in_value):
    try:
        float(in_value)
        return True
    except ValueError:
        return False
```

We can test our example package by running `python` in our project directory and importing and using the code.

```python
>>> from is_number import is_number
>>> is_number(10)
True
>>> is_number("10.0")
True
>>> is_number("Hello world")
False
```

Note that our overall project is called `is-number` with a hyphen and our Python library is called `is_number` with an underscore. This is a common pattern that you will see in open source Python projects.

We aren't going to go into any more detail here on how to write the Python module itself. The [Python documentation](https://docs.python.org/3/tutorial/modules.html#packages) already contains excellent information on how to structure your code so that others can import and use it.

Whatever library you want to create you will need to create a directory for your module and write your code in there. This is the fun bit so enjoy!

## Version control

Now that we have a couple of files and a little bit of code let's introduce some [version control](https://en.wikipedia.org/wiki/Version_control).

Version control helps us manage changes to our files over time. We will take snapshots (often called commits) of our code and our version control software will keep track of the changes between each commit. This way we can revisit old versions of code and keep track of things that change.

It also allows multiple people to work on the code at the same time. Each person takes a copy of the code at the latest revision and makes their changes. Then they can merge these changes back in. If neither person edited the same lines of the same files then they will be merged cleanly and both sets of changes will be applied. If they did edit the same lines then they will have a merge conflict and someone will have to decide which edit is correct (or maybe make a further edit to unite them).

For our example project, we are going to use [git](https://git-scm.com/) for our version control tool, but alternatives exist such as [Mercurial](https://www.mercurial-scm.org/) and [Apache Subversion](http://subversion.apache.org/).

```bash
# Make the current directory a git repository
git init .
```

Before we add any files to our repository it is good practice to add a `.gitignore` file which contains a list of files that shouldn't be included in the repository. For example, we wouldn't want temporary files, build artifacts or editor configurations or make their way into our version control.

You can find a nice Python example `.gitignore` file [here](https://github.com/github/gitignore/blob/master/Python.gitignore). I recommend you download it and use it as the base for your `.gitignore` file.

```bash
curl -sSL https://github.com/github/gitignore/raw/master/Python.gitignore > .gitignore
```

Now we can add all of our files and make our first commit.

```bash
# Stage all of our changes to be committed
git add -A

# Make our initial commit
git commit -m "Initial commit"
```

## Dependencies

When you write your library you may wish to use other libraries within it, these are called dependencies. It is common practice to list all of the libraries that your project depends on in a file called `requirements.txt` and to specify versions of those packages that are known to work.

In our example, we have such a simple piece of code that we don't have any dependencies, but if your code wanted to make HTTP requests for example and you wanted to use the `requests` library to do so you would add the following to your `requirements.txt` file.

```
requests==2.22.0
```

However, for our example project, we are just going to create an empty `requirements.txt` file.

```bash
touch requirements.txt
```

You then need to add it to your version control.

```bash
git add requirements.txt
git commit -m "Add requirements file"
```

## LICENSE

Now let's think about licensing our code. When you publish an open-source package it isn't enough to make the source code available to people. You also need to explicitly license the code to them.

When you create something like a piece of code you automatically own the copyright to that code (or your employer does, you should check your contract). Nobody else is allowed to use that code unless you legally permit them to do so, and you can do that with a license.

It is not recommended that you write a license from scratch, but instead that you choose an existing open source license. There are many to choose from and you can find out more about them at the [Open Source Initiative](https://opensource.org/licenses).

For our example project, we are going to choose the MIT license because that is what the JavaScript `is-number` package uses and also because it is a very common and minimal license.

To license our project we need to create a file called `LICENSE` that contains the content of the license. For the MIT license that would be:

```
Copyright <YEAR> <COPYRIGHT HOLDER>

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
```

You'll need to update the year and copyright holder name. Also if you are employed and your contract states that your work is the property of your employer then they are the copyright holder, not you, and you will need to get them to agree to license the project as open source.

Once you've added the license you'll need to commit it too.

```bash
git add LICENSE
git commit -m "Add license"
```

## setup.py

To enable people to install our library we need to create a `setup.py` file and use the `setuptools` package to tell Python how to install your code.

```python
import setuptools

with open("README.rst", "r") as fh:
    long_description = fh.read()
with open("requirements.txt", "r") as fh:
    requirements = [line.strip() for line in fh]

setuptools.setup(
    name="is-number",
    version="0.0.1",
    author="Jacob Tomlinson",
    author_email="jacob@tomlinson.email",
    description="A Python library to determine if something is a number.",
    long_description=long_description,
    long_description_content_type="text/x-rst",
    packages=setuptools.find_packages(),
    classifiers=[
        "Programming Language :: Python :: 3",
        "License :: OSI Approved :: MIT License",
        "Operating System :: OS Independent",
    ],
    python_requires='>=3.6',
    install_requires=requirements,
)
```

In our example setup, we start by reading in our `README.rst` and `requirements.txt` files so that we can reuse their contents during setup, this helps avoid duplication.

We then call the `setuptools.setup` method and provide information about the package, the author, descriptions, etc. We are using the `setuptools.find_packages()` method to automatically find our code and include it in the package and we have made sure to list our MIT license in the list of classifiers.

Let's commit this too.

```bash
git add setup.py
git commit -m "Add setup"
```

This is enough that we can install our package on our system.

```
pip install .
```

This will run the `setup.py` file in the current directory and install our Python module into our site-packages.

We could move to any other location in our filesystem and import and run our code.

```bash
cd ~  # Or anywhere else
python
```

```python
>>> from is_number import is_number
>>> is_number(10)
True
```

## Summary

We are now at the stage where we have a usable open source library that is installed on our system.

We have covered:

- Writing some simple code
- Creating a `README`
- Managing our code with version control
- Adding dependencies
- Adding an open-source license
- Setting up our module for installation

In future posts we will cover:

- Versioning our package
- Formatting our code
- Publishing our code (opening the source!)
- Packaging and publishing our library
- Adding tests
- Automating those tests
- Automating future releases
- Generating documentation and hosting it
- Creating a community
- Handling future maintenance
