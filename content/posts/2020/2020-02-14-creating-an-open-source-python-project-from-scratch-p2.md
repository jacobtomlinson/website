---
title: "Versioning and formatting your Python code"
series: ["Creating an open source Python project from scratch"]
date: 2020-02-14T00:00:00+00:00
draft: false
author: "Jacob Tomlinson"
categories:
  - blog
tags:
  - Python
  - GitHub
  - Tutorial
  - Black
  - Versioneer
  - SemVer
thumbnail: python
---

In this post, we will cover a few project hygiene things that we may want to put into place to make our lives easier in the future.

## Versioning

At this point we should put some thought into versioning our library. When we make changes to our code we will want to release a new version, but what kind of naming scheme should we use for our version labels?

There are several commonly used versioning standards used in open-source software. The two that you will see most often are [Semantic Versioning](https://semver.org/) or SemVer for short, and [Calendar Versioning](https://calver.org/) or CalVer for short. The first increments the version numbers based on what changes are made to the software, the second increments the version number based on when those changes were made.

SemVer is commonly used when a project has an API which is being presented to the user, the developer can use the version information to communicate to the user whether or not new versions will impact their use of the library. CalVer is more commonly used for projects which do not necessarily provide an API to the user, an example may be a set of configuration files for a system which changes over time.

### SemVer

In this example, we are going to use SemVer. SemVer aims to communicate to your users what has changed between versions by incrementing one of three numbers separated by a dot.

Let's discuss each number but starting right to left, as that will take us from the smallest change to the biggest change. The righthand number is referred to as the `bugfix` version and should be incremented if we fix a bug in our library. For example, if there is a small problem that is breaking things for users in version `1.0.0` and we fix it the next release should be `1.0.1`.

The middle number is referred to as the `minor` version number. This should be incremented if we add a new feature to our library. Existing users should be able to continue using the library exactly as they always had done regardless of the new feature. Also when incrementing this number we should reset the `bugfix` number to `0`. For example, if we added a new function to our library, but left `is_number` totally alone, this is a `minor` modification and the next release should be `1.1.0`.

The left number is the `major` version number. This should be incremented if we have made a change which will break things for our users. After this version, they cannot just continue as usual and will probably have to read the documentation and change their behaviour accordingly. We also reset both other numbers to `0`. For example, if we modified `is_number` to return the strings `'yes'` and `'no'` instead of the booleans `True` and `False` this would be a breaking change and we should release `2.0.0`.

In our `setup.py` we have set the initial version to `0.0.1`. When the major version of a library is `0` this means it is in development mode and should be considered unstable. Generally in this mode, the `bugfix` version remains the same but the `minor` version is incremented for both `major` and `minor` changes, meaning anything could break at any time. Ideally once a library is "finished" and the functionality is stable you should move on to version `1.0.0` and adopt SemVer more strictly, however you will notice in the wild that many well-used projects have not yet had a version `1.0.0` release. This is because many open source projects are run by volunteers. Folks with limited time to work on projects do not want to commit to a stable API.

For more information on this schema see the [Semantic Versioning website](https://semver.org/).

## Automatic version detection

As we are using git for our version control we also have the ability to [tag specific commits with our version numbers](https://git-scm.com/book/en/v2/Git-Basics-Tagging). However it is easy to forget to update our `setup.py` file when creating a tag, so it would be great to enable our setup to automatically detect the current version from the git information.

To do this we can use [versioneer](https://github.com/warner/python-versioneer) which is a Python module that you package with your project. Instead of explicitly setting your version in `setup.py` versioneer gets it dynamically from the version control.

```bash
pip install versioneer
```

For versioneer to correctly identify our package version we need to provide some configuration. This is done by adding a `[versioneer]` section to a file called `setup.cfg`. This file is used for configuring many Python tools so let's create it.

```
[versioneer]
VCS = git
style = pep440
versionfile_source = is_number/_version.py
versionfile_build = is_number/_version.py
tag_prefix =
parentdir_prefix =
```

In our config, we have specified that we are using git for our version control and that our version numbers should follow [PEP 440](https://www.python.org/dev/peps/pep-0440/).

We also specify where our version file should live inside our package and that we are not using any prefixes.

Next, we need to install the versioneer module into our project.

```bash
versioneer install
```

Running this command will create a few new files for us in our project and also ask us to make some changes to our `setup.py` file. We need to import `versioneer` in the setup, replace the version keyword argument with `versioneer.get_version()` and add a new argument `cmdclass=versioneer.get_cmdclass()`.

```python
import setuptools
import versioneer

with open("README.rst", "r") as fh:
    long_description = fh.read()
with open("requirements.txt", "r") as fh:
    requirements = [line.strip() for line in fh]

setuptools.setup(
    name="is-number",
    version=versioneer.get_version(),
    cmdclass=versioneer.get_cmdclass(),
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

We should also go and have a look at the `__init__.py` file where our code lives as versioneer has added some extra lines in there. These lines use versioneer to set a `__version__` variable in the root scope of the file, this is common practice for Python libraries. However I prefer to have this section at the top and versioneer has added it to the bottom, so I'll readjust it to look like this.

```python
from ._version import get_versions
__version__ = get_versions()['version']
del get_versions


def is_number(in_value):
    try:
        float(in_value)
        return True
    except ValueError:
        return False
```

It has also created a `versioneer.py`, `MANIFEST.in` and `.gitattributes` file in the root of the project and a `_version.py` file in our library folder.

The `versioneer.py` and `_version.py` files contain the versioneer module code for calculating the version number from the git information.

The `.gitattributes` file tells git to run these files when an archive is being created during publishing, this will replace the contents of `_version.py` with a static version. This is important as the git history will not be included in the archive.

The `MANIFEST.in` file tells setuptools which files to include when publishing our package. We will come on to publishing in another post but as versioneer has created this file for us let's also put some other things into this file.

```
include README.rst
include LICENSE
include requirements.txt
graft is_number
recursive-exclude * *.py[co]
include versioneer.py
include is_number/_version.py
```

Here we are ensuring that other files like our `README.rst` and `requirements.txt` are going to be included in our published package, this is important because our setup reads them in. We set `graft is_number` to ensure all files in our library are included but then exclude all compiled `.pyc` and `.pyo` files as we do not want those to be packaged.

Now let's test that we can get the version of our package with versioneer.

```python
>>> import is_number
>>> is_number.__version__
'0+untagged.3.gf86897f.dirty'
```

We can see that we are getting a version number here, but it looks a little unusual. This is because we haven't created any git tags yet and we have uncommitted changes in our repo.

Let's commit all of our changes and tag our first version.

```bash
git add -A
git commit -m "Add versioneer"
git tag 0.0.1
```

Now if we run the same code again we should see the correct version.

```python
>>> import is_number
>>> is_number.__version__
'0.0.1'
```

## Formatting (Black)

A good practice for any open source project is to have a standard for formatting your code. Python has [PEP 8](https://www.python.org/dev/peps/pep-0008/) which is a style guide with recommendations on how to format your code. This guide is great and has lots of good advice and there are many linting tools, [flake8](https://flake8.pycqa.org/en/latest/) for example, to check that your code complies with this standard.

It has become commonplace for projects to use an automatic formatting tool such as [Black](https://black.readthedocs.io/en/stable/). Tools like Black will take a codebase and format it to comply with PEP 8, but it will also make some opinionated decisions when it is unclear what should be done. This results in code being formatted consistently across the codebase and reduces the amount of time developers will spend arguing about "correct" formatting.

You may not be 100% happy with the opinionated decisions that Black makes, but I encourage you to just embrace them as the reduction in mental overhead is worth the compromise. I even suggest you set your code editor to run Black automatically when you save a file.

To run black on your codebase you need to install it.

```bash
pip install black
```

Then you can format individual files or your entire project.

```bash
# Format everything in the current directory
black .
```

This will even format the versioneer files that were created, which I don't think is a terrible thing. Let's commit this formatting and move on.

```bash
git add -A
git commit -m "Apply Black"
```

## Summary

In this post we have covered:

- Versioning our package
- Automatically detecting the version from our git information
- PEP 8 and code style guides
- Formatting our code with Black

In future posts we will cover:

- Packaging our library and publishing it
- Adding tests
- Automating those tests
- Automating future releases
- Generating documentation and hosting it
- Creating a community
- Handling future maintenance
