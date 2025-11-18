---
author: Jacob Tomlinson
title: "Python package managers: uv vs pixi?"
date: 2025-11-18T00:00:00+00:00
draft: false
categories:
  - blog
tags:
  - Python
  - Packaging
  - Open Source
---

When I talk to people about Python package management in 2025 I see the following tools in active use; [`uv`](https://docs.astral.sh/uv/), [`pixi`](https://pixi.sh/latest/), [`pip`](https://pypi.org/project/pip/), [`conda`](https://docs.conda.io/en/latest/), [`mamba`](https://mamba.readthedocs.io/en/latest/), [`micromamba`](https://mamba.readthedocs.io/en/latest/user_guide/micromamba.html) and [`poetry`](https://python-poetry.org/). There may be others, but I don't hear much about them.

Two of these tools in particular, `uv` and `pixi`, are part of the new and shiny VC-backed Rust-based "modern Python tooling". It's not obvious at first glance what the differences are and when to choose one over the other. While there seems to be lots of getting started content out there about these tools there isn't much in the way of a history or deep comparison so I thought I'd write something down. 

To really understand the differences it helps to have some historical context, so we are going to start at the beginning. 

```info
The history part of this post got pretty long. If you would prefer to just read my opinions on which tool to choose then skip ahead to [So what do I use?](#so-what-do-i-use).
```

## Pre-history: `easy_install` vs `pip`

For as long as I can remember there have been multiple Python package managers to choose from. Back in the day it was `easy_install` and `pip`. As you probably can guess `pip` emerged as the most popular package manager from that time.

These tools work by downloading Python packages from [PyPI](https://pypi.org/), an online repository of packages, and putting them in the right place on your filesystem for them to be importable in Python. In the early days the packages themselves were just gzip files with Python code inside. Your package manager downloaded the file, unzipped it to the right place and you're done!

So what was wrong with `easy_install`? The short answer is it didn't have a dependency solver. When you download a Python package it may depend on other packages, and these dependencies may have constraints. For example the latest release of the package `foo` may need `bar>=3`, so when `pip` downloads `foo` it needs to be sure to also download a version of `bar` that is `v3` or higher. The `easy_install` tool just didn't support this. You may have another dependency in your project that also depends on `bar` but has different constraints. As your project gets more complicated you may have more dependencies, and in turn those dependencies may have their own. 

!["A graph of the dependency tree for the kr8s project"](./kr8s-pydeps.svg "The dependency tree for [kr8s](https://kr8s.org), created by [pydeps](https://github.com/thebjorn/pydeps) <br />`pydeps kr8s --max-module-depth=1 --rankdir BT --reverse`")

As you can see in this diagram[^1] of the dependency tree for [kr8s](https://kr8s.org) there is a complex interdependency between `kr8s`, `httpx`, `httpx_ws`, `anyio` and `sniffio` each with specific requirements around which versions are needed. And this is a relatively simple dependency graph. Solving the optimal version of every package you need is an [NP-hard](https://research.swtch.com/version-sat) problem, and it's a core goal of every package manager from `pip` onwards.

## Non-Python dependencies

`pip` was an amazing tool for many people, but Python is a special language where it's extremely popular to write libraries of code in compiled languages like C, C++, Fortran or Rust and bind them into Python. While Python is a relatively slow language it can call into these fast compiled dependencies and use them in the same way it can use Python dependencies. Many languages can do this, but this practice has taken off hugely in the Python community because it allows users to trade off performant code with a friendly and flexible programming language and often get the best of both worlds.

One big problem with `pip` in the early days was that it only handled source distributions. This means it could download a gzip file of source code and put it in the right place, call some hooks that was it. If you wanted to package some C code that could be used from Python you would need to zip up the C code along with it's Python bindings and publish it to PyPI, then when `pip` installed the package it would download the code and then run the compiler locally on your machine to turn the C code into something that could actually be executed. If you didn't have all the C compilers and related tools on your machine you were in for a bad time. 

Python is extremely popular as a beginner/intermediate language and you can become very productive quickly without needing loads of computer science knowledge. So the proportion of Python developers with no understanding of C, compilers, linkers and that whole toolchain and ecosystem was high, and those users would have a bad experience installing compiled code via `pip`.

```bash
# Source - https://stackoverflow.com/q/79413114
# Posted by Motorhead, modified by community. See post 'Timeline' for change history
# Retrieved 2025-11-18, License - CC BY-SA 4.0

  [194/535] Generating numpy/_core/_multiarray_tests.cpython-39-x86_64-cygwin.dll.p/_multiarray_tests.c

  [195/535] Compiling C object numpy/_core/libnpymath.a.p/meson-generated_ieee754.c.o
  FAILED: numpy/_core/libnpymath.a.p/meson-generated_ieee754.c.o
  /usr/bin/x86_64-w64-mingw32-gcc -Inumpy/_core/libnpymath.a.p -Inumpy/_core -I../numpy/_core -Inumpy/_core/include -I../numpy/_core/include -I../numpy/_core/src/npymath -I../numpy/_core/src/common -I/usr/include/python3.9 -I/tmp/pip-install-2nk4smdu/numpy_d0c3008da8224ebc9f1bede0e4cba273/.mesonpy-y0j8jkqq/meson_cpu -fdiagnostics-color=always -DNDEBUG -D_FILE_OFFSET_BITS=64 -Wall -Winvalid-pch -std=c11 -O3 -fno-strict-aliasing -msse -msse2 -msse3 -DNPY_HAVE_SSE2 -DNPY_HAVE_SSE -DNPY_HAVE_SSE3 -MD -MQ numpy/_core/libnpymath.a.p/meson-generated_ieee754.c.o -MF numpy/_core/libnpymath.a.p/meson-generated_ieee754.c.o.d -o numpy/_core/libnpymath.a.p/meson-generated_ieee754.c.o -c numpy/_core/libnpymath.a.p/ieee754.c
  In file included from /usr/include/python3.9/Python.h:50,
                   from ../numpy/_core/src/npymath/npy_math_common.h:4,
                   from ../numpy/_core/src/npymath/ieee754.c.src:7:
  /usr/include/python3.9/pyport.h:230:10: fatal error: sys/select.h: No such file or directory

    230 | #include <sys/select.h>
        |          ^~~~~~~~~~~~~~
  compilation terminated.
```

I'm not exactly sure on the history here but around this time the folks over at [Anaconda](https://anaconda.org) (formerly known as Continuum Analytics) were trying to solve this problem. They worked a lot with Python users in the sciences where there was heavy use of libraries with compiled code like `numpy` and `scipy`. At this time the creators of PyPI and `pip` felt that it should remain as a pure source code repository, this would change later on, but there was enough friction here that the Anaconda folks decided to create the [`conda`](https://docs.conda.io/en/latest/) package manager.

## Binary packages

The `conda` package manager handles a different kind of package. While you can still put pure Python code into a conda package you can also include pre-compiled binaries. When you build a conda package you run the compiler for all the common operating systems you expect it to run on, Windows, Linux, macOS and the common CPU architectures like x86 and ARM. This is a lot more work for the developers to build all these packages, but it hugely simplifies things for the end user as `conda` can just download the right binaries for their system without needing to compile anything.

Conda packages are downloaded from a repository separate to PyPI hosted by Anaconda at [anaconda.org](https://anaconda.org/). The packages uploaded to PyPI need to also be uploaded to [anaconda.org](https://anaconda.org/). Thankfully the community over at [conda-forge](https://conda-forge.org/) have built incredible tooling around automating this process.

Another thing `conda` does differently is it can look at your computer and find things that have been installed by other means through [virtual packages](https://docs.conda.io/projects/conda/en/stable/user-guide/tasks/manage-virtual.html). Nearly all compiled code depends on core libraries like `glibc` or `musl` which are included with the operating system, `conda` can figure out what versions of these packages you have and then include that in it's package dependency solve. This has been especially useful in the [CUDA Python](https://anaconda.org/conda-forge/cuda-python) ecosystem where all Python CUDA packages depend on specific NVIDIA GPU driver and CUDA versions.

One interesting consequence of `conda` supporting compiled binaries as well as source code is that it can package code written in any language, and compiled for any target hardware. This has resulted in `conda` becoming a popular package manager in the [R](https://anaconda.org/conda-forge/r-base) and [Julia](https://anaconda.org/conda-forge/julia) communities. You can also use it to install entire other ecosystems like nodejs or golang. You can even [package Python itself](https://anaconda.org/conda-forge/python) as a conda package and install it with `conda`.

For many years `conda` grew in popularity with communities that heavily rely on compiled dependencies for their Python code, like Data Science, AI, Machine Learning, Geosciences, Biosciences, Robotics, etc.

## Environments and solving

The `conda` tool also wrapped up similar functionality to [virtual environments `venv`](https://docs.python.org/3/library/venv.html) where you can create many isolated development environments. But it took a different road and instead of always creating your environment in the current directory it keeps them all centrally where you could activate them by name and reuse them across multiple projects.

```bash
conda create -n my-awesome-environment python=3.14
conda activate my-awesome-environment
```

Another difference is that whenever you install a package in an environment with `conda` it does the entire solve again. It looks at all the packages you've installed in your environment in the past and then tries to find compatible versions of all shared dependencies. One consequence of this is that you can end up in a state where your environment is unsolvable. If `foo` depends on `fizz>=3` and `baz` depends on `fizz<3` you cannot install both `foo` and `baz` in the same conda environment.

With `pip` if you try and do it in one command the solve will fail in the same way. 

```console
$ pip install foo baz
failed to solve...
```

However, if you do it in two separate commands this will succeed and `fizz` will end up being the version that `baz` needed because it was installed second. 

```console
$ pip install foo && pip install baz
success! (but foo is broken)
```

This means `pip` can happily create broken environments where `foo` is totally broken in a way that `conda` will refuse to do.

There are also a few other things I like about the conda package spec. For example in `pip` you can mark a dependency as optional, a user can choose to install it or not, but if they choose not then `pip` pays no attention to it at all. In `conda` you say a package is optional, but if it is installed directly or by some other package then its version will be taken into account. An example is where your code checks for an import and if it exists then it uses it, but it uses a feature from a specific version.

```python
try:
    import foo
except ImportError:
    foo = None

if foo:
    foo.some_feature() # this method was added in foo 1.5
```

Then in your `pyproject.toml` you specify the extra.

```toml
[project.optional-dependencies]
bar = ["bar>=1.5"]
```

The `pip` package may optionally depend on `bar>=1.5` via the `bar` extra. If you run install `foo` with the `bar` extra you will get `bar` at the expected version. 

```console
$ pip install foo[bar]
success!
```

If you run something conflicting it will fail.

```console
$ pip install foo[bar] bar==1.4
failed to solve...
```
However, you can easily leave out the extra and just install the broken package and `pip` will happily install it. 

```console
$ pip install foo bar==1.4
success! (but now foo is broken because bar==1.4 doesn't have foo.some_feature())
```

In conda packages you can use the [run_constrained](https://docs.conda.io/projects/conda-build/en/stable/resources/define-metadata.html#run-constrained) option in your `recipe.yml` to specify that `bar` isn't directly depended on, but if it is installed it must be at least `1.5`.

```yaml
requirements:
  run_constrained:
    - bar >=1.5
```

If you try and do the same conflicting command in `conda` it will fail.

```console
$ conda install foo bar==1.4
failed to solve...
```

The goal of describing these differences isn't to say `conda` is better. Just that `conda` set out to solve a bunch of problems that `pip` had. Some of which have been fixed by both projects in different and opinionated ways, and others which haven't be fixed in `pip` at all. These differences keep the `conda` community alive and relevant.

## Reinventing the wheel

In the years since conda was created the `pip` community have added binary packages. We now have the [Python `wheel`](https://packaging.python.org/en/latest/specifications/binary-distribution-format/) which is a new package type that allows you to include compiled code and publish it on PyPI. This means `pip` no longer needs to compile these dependencies locally, closing this feature gap with `conda`. However, the `conda` community is now big enough and mature enough that it has momentum and will likely not go away despite this early unique selling point becoming irrelevant.

There are some differences between wheels used by `pip` and conda packages used by `conda`. One being that wheels must include all compiled dependencies in a statically linked binary. Because PyPI is for Python code and its dependencies it doesn't want to think about the dependencies of those dependencies. So if your C code depends on something else then you need to bundle all of that together. This means C libraries with shared dependencies can't dynamically link to them and share the dependency causing environments created with `pip` to be much larger than those created with `conda` due to this duplication.

## Conda got slow

As the conda ecosystem grew and environments got more complex the way it solved the environment got slower and slower. I remember around 2019 this felt especially painful when working in large environments. To solve this a small group of `conda` users created [`mamba`](https://mamba.readthedocs.io/en/latest/), a reimplementation of `conda` in C++. The `mamba` tool behaved exactly the same as `conda` but was much faster.

The way you install `conda` is via a script which creates a base conda environment which contains `conda` itself. The `mamba` install was exactly the same, but with a fast compiled tool this felt unecessary. As a result the [`micromamba`](https://mamba.readthedocs.io/en/latest/user_guide/micromamba.html) tool was also published which was a static version of `mamba` that could be installed as a single binary and didn't depend on having a base environment. Other than that it worked exactly the same.

## Conda got fast

With many but not all `conda` users switching over to `mamba` there was now a fast and slow conda experience. To solve this the folks at Anaconda helped refactor the `mamba` solver out into a separate library and then [utilized that library in `conda`](https://www.anaconda.com/blog/a-faster-conda-for-a-growing-community). This meant that `conda` got all the performance from `mamba` without users needing to change tools. This effectively killed the demand for `mamba`, although `micromamba` is still popular in environments where having a single binary and a simple install is useful, like in CI.

## Locking and declarative environments

Another topic commonly discussed in package management communities is locking. Once you've solved your environment you can store that solve somewhere in your project so that you don't need to do it again unless you specifically need to upgrade something. This is helpful because if you do the same solve again next week there may be new packages released which affects the results of the solve. For reproducible environments having a lock file is critical, whether you want to reproduce a science experiment 5 years from now, or just be confident that your production web server isn't going to pick up a new version of something and break unexpectedly.

In the Python community a popular solution to this was [`poetry`](https://python-poetry.org/). This package manager effectively wrapped `pip` but also stored a lock file of the solve and allowed you to reproduce your environment easily. While I'm sure `poetry` has some other great features I tend to think of it as `pip` + locking.

For me `poetry` was my first introduction to declarative Python environments. In `conda` you could create an `environment.yaml` and in `pip` a `requirements.txt`. While the best practice was to update those files and recreate your environment from scratch each time it rarely happened that way. Most people would run multiple install commands like `pip install foo`, then `pip install bar`, etc and have an environment grow organically and potentially introduce conflicts over time. With `poetry` you would run `poetry add foo`, and `poetry add bar`. The list of dependencies was stored in the project and was solved from scratch each time. This way you guarantee you don't end up with a broken environment. You can see a parallel here with `conda` where solves are done from scratch, but updating and storing the environment explicitly felt refreshing and new.

## Managing Python itself

A challenge in the `pip` style ecosystem is that `pip` and derived tools like `poetry` depends on Python, so before you get your package manager you need to figure out where to get Python from in the first place. This is solved in `conda` because `python` is just another package for you to include in your dependencies, each environment can have it's own version of Python. 

A few tools popped up to be used alongside `pip` to handle this including [`pyenv`](https://github.com/pyenv/pyenv) and [`rye`](https://rye.astral.sh/). While `pyenv` is a simple tool for installing multiple Python versions on your system `rye` had a greater goal. It wanted to provide a way to manage Python projects entirely from installing Python, to running linters and tests, to managing dependencies and creating lock files. Inspired by [`cargo`](https://doc.rust-lang.org/stable/cargo/) from the Rust community it had a bunch of interesting features but did not displace `pip`.

## Next-gen conda with `pixi`

With the improvements from `mamba` being rolled back into `conda`, and `micromamba` just solving a niche use case, the `mamba` developers started thinking about what the next generation of `conda` tooling would look like. Also inspired by the `cargo` package manager in the Rust community they formed a new company called [Prefix](https://prefix.dev/) and built [`pixi`](https://pixi.sh/latest/). 

The `pixi` package manager is designed to take the best features from all the tools and combine them in a cohesive way. It allows you to install both conda-forge and PyPI packages in a single environment. It has a filesystem-based approach similar to `venv` where environments live in your projects instead of the central `conda` environment location. You install `pixi` as a single binary, similar to `micromamba`. The new `pixi.toml` spec was also introduced to enable storing declarative environments in config files and they added `pixi.lock` to handle locking. 

The target audience of `pixi` is folks who use a lot of compiled or non-Python dependencies in their projects, value locking and reproducibility and want fast modern tooling.

## Next-gen pip with `uv`

The environment solves in `pip` are less compute intensive than `conda` because it doesn't try and resolve the entire environment every time. But there was still a lot of performance left on the table. Installing Python packages via `pip` felt very slow compared to using `pixi`, `mamba` or modern versions of `conda`.

The folks working on `rye` saw this opportunity and did for `pip` as `mamba`/`pixi` had done for `conda`. They built [`uv`](https://docs.astral.sh/uv/), which started as a reimplementation of `pip` in Rust. They built a really fast solver for PyPI packages and also bundled environment management similar to `venv` and `poetry` directly into `uv`. This now meant that the pip ecosystem via `uv` has even more of a feature overlap with `conda`, it's fast, declarative, has locking, can handle compiled dependencies and can manage virtual environments. However, by the nature of reimplementing `pip` directly you will find that `uv pip` still has many of the problems that `pip` has.

Over time almost all the project management features from `rye` were ported over to `uv`, and `rye` has since been deprecated.

You can also specify a Python version when creating a `uv` environment. The Astral folks have set up a special build pipeline so you can quickly and easily download and manage multiple Python versions. It's not the same as the way the `conda`/`pixi` ecosystem distribute `python` as just another package, but functionally the feature feels the same.

## Converging tools

Hopefully this dive through the history of Python package tooling has given you an idea of how things have evolved over time. New requirements and features have been added, but due to the nature of the Python community new features often arrived as a new alternative tool. Then the next tools would include those previous features and more.

To recap, these are the things that the Python community has wanted over time that has driven this evolution of tools:

- A performant Python package dependency solver
- Virtual environments
- Binary packages for compiled code and bindings
- Python version management
- Declarative environments
- Locking

Both `pixi` and `uv` can do all of these things, but the way they achieve them are slightly different. I would argue that `pixi` has a superset of `uv`'s functionality because it also supports conda-forge packages and all the additional behaviour differences from `conda`. It also has a [task runner](https://pixi.sh/latest/workspace/advanced_tasks/) which [`uv` is yet to implement](https://github.com/astral-sh/uv/issues/5903). 

However, it seems that because the majority of Python users don't need the same nuanced differences that the `conda` community needs the slightly simpler design of `uv` has won them the biggest market share at the moment.

## So what do I use?

Python tooling is complex and it's not immediately clear what to use, as you can see there are lots of tools with overlapping functionality. I wish I could say that there was one clear tool for the future, but unfortunately the answer in 2025 is "it depends".

It might help you make your mind up if I describe how I use these tools in various situations.

### Scripting: `uv`

If I am working on a standalone script, I will use `uv`'s script functionality to contain my environment in the script itself.

This [example from the `uv` docs](https://docs.astral.sh/uv/guides/scripts/#creating-a-python-script) uses `requests` to read some data and `rich` to format it nicely in the terminal.

```python
# example.py
import requests
from rich.pretty import pprint

resp = requests.get("https://peps.python.org/api/peps.json")
data = resp.json()
pprint([(k, v["title"]) for k, v in data.items()][:10])
```

With `uv` you can run this directly with an ephemeral virtual environment. Creating the environment with `uv` is so fast it runs almost instantly.

```bash
uv run --with requests,rich example.py
```

You can also declaratively add these dependencies to the script do you can omit the `--with` flag in the future.

```bash
uv add --script example.py 'requests<3' 'rich'
```

This adds a comment to the top of the script with the environment stored.

```python
# /// script
# dependencies = [
#   "requests<3",
#   "rich",
# ]
# ///

import requests
from rich.pretty import pprint

resp = requests.get("https://peps.python.org/api/peps.json")
data = resp.json()
pprint([(k, v["title"]) for k, v in data.items()][:10])
```

```bash
uv run example.py
```

### Application Development: `uv`

If I am working on an application like an HTTP API server, a personal CLI tool or a desktop application I use `uv`. In these projects I am most likely consuming packages from PyPI, I want to develop in a virtual environment, I care about lock files and I want the best performance.

We can start by making a new project.

```bash
mkdir myproject
cd myproject
uv init
```

This creates a `pyproject.toml` to store the project name and dependencies, a `uv.lock` to keep the locked environment and `hello.py` to get me started.

I can then add new dependencies, both for the project but also for testing and development

```bash
uv add requests
uv add --dev pytest
```

I can run the project code, or any other tools I want to use during development.

```bash
uv run hello.py
uv run pytest
```

### AI or CUDA Application Development: `pixi`

The exception is when working on projects that require local GPU acceleration. Some AI related projects are just calling an external API so if all you are doing is making HTTP requests then see Application Development above. But if you're training and/or inferencing models locally you will probably be using a GPU, and that means you are using compiled CUDA dependencies. For anything that uses CUDA I would much prefer to use conda packages because they are smaller than their wheel counterparts, my environments are smaller because compiled code can use dynamic libraries and the tooling is more robust when it comes to solving these complex environments. This is also a fast moving space and reproducibility is key, so the locking that `pixi` provides over standard `conda` is very important for application development.

We can start by making a new project.

```bash
mkdir myproject
cd myproject
pixi init
```

This creates a `pixi.toml` to store the project name and dependencies. I can then add new conda-forge dependencies.

```bash
pixi add cupy-xarray torchgeo # CUDA accelerated libraries for geospatial analysis
pixi add pytest
```

Then you can do the same things with `pixi run`.

```bash
pixi run mycode.py
pixi run pytest
```

### Science, Machine Learning, Analysis, Research: `pixi`

If I am working on a Data Science project, something which uses Jupyter Notebooks or anything where I will benefit from conda packages then I use `pixi`. The kind of code written by Data Scientists, Traditional Scientists, Machine Learning Researchers, etc will usually not be depended on by another project. It might be run once to get a single answer, or run many times by peers reproducing or updating results. These kinds of projects generally have compiled dependencies so again the conda package ecosystem is going to be more mature. Reproducibility means locking, so `pixi` really shines here.

```bash
mkdir myresearch
cd myresearch
pixi init
```

Then you can install `jupyter` and related dependencies for your work.

```bash
pixi add jupyterlab
pixi add pandas
```

Then launch Jupyter

```bash
pixi run jupyter lab
```

### Library Development: `uv`

I personally spend a lot of my time working on software libraries that are depended on by other projects. When working on these I specifically _do not_ want a lock file, I want loose and flexible dependencies with a broad CI suite testing many different versions to ensure the highest compatibility in people's projects. I want fast, lightweight and often ephemeral virtual environments for developing in, running tests and comparing against different dependency versions. Both `uv` and `pixi` are great for this, but most of the time I care more about PyPI dependencies and `pixi` feels awkward being a conda first project.

The way that `pixi` either stores it's dependencies in a `pixi.toml` or a `[tool.pixi]` namespace in `pyproject.toml` also means there is a bunch of duplication of developer dependencies if you want to support both `uv` and `pixi` for development. I understand that `pixi` stores dependencies separately because conda packages are not the same as PyPI packages, but for most situations when developing a library you need to support PyPI anyway.

I also like to encourage contributions to my projects and the reality is that `uv` is more popular than `pixi` right now, so it makes sense for me to provide `uv` instructions in my `CONTRIBUTING.md` guide.

For example I can clone my `kr8s` project.

```bash
git clone https://github.com/kr8s-org/kr8s.git
cd kr8s
```

Then create a new virtual environment with a specific Python version to develop in.

```bash
uv venv --python 3.13
source .venv/bin/activate
uv sync --dev
```

Then run the project's tests.

```bash
uv run pytest kr8s
```

I cannot understate how fast all of the above commands run when using `uv` compared to the same thing with `pyenv`/`pip`/`venv`.

### Interdependent Library Development: `conda` + `uv pip`

An exception to the above is when I'm working on multiple libraries with interconnected dependencies. For example the other day I was debugging a failing test in `dask.distributed`. To investigate this I had both `dask` and `distributed` installed from source. I then discovered the bug was in a dependency of `distributed` called `tblib`. So I wanted to have all three installed from source, then I needed to [make changes to `tblib`](https://github.com/ionelmc/python-tblib/pull/85) while running the test suite from `distributed`, all in the same environment.

For these situations I find the `conda` environment model of having named environments stored in a central location more helpful. I didn't want environments stored within the `dask`, `distributed` or `tblib` directories, which is how `pixi` and `uv` feel most natural. I just want to create an environment, specify a Python version, and use `uv pip install -e .` to install each project from their locally cloned git repo. I use `uv pip` because the functionality is the same as `pip` just way faster.

For example let's start by cloning a bunch of interdependent projects.

```bash
git clone https://github.com/dask/dask.git
git clone https://github.com/dask/distributed.git
git clone https://github.com/ionelmc/python-tblib.git
```

Then I would create a new named environment with `conda`.

```bash
conda create -n dask-dev python=3.13 -y
conda activate dask-dev
```

Then I would go through each project and install it from source along with any test dependencies.

```bash
cd dask
git checkout 2025.11.0  # Due to some pinning shenanigans in dask we need to install a tagged version
uv pip install -e .[complete] .[test]
git checkout main

cd ../distributed
uv pip install -e .

cd ../tblib
uv pip install -e .
```

Then finally I could run the `distributed` tests against my local versions of `dask` and `tblib`.

```bash
pytest
```

````info
I have played around with using `pixi` to do this, mainly to reduce the number of tools I use to just `pixi` and `uv`. You can create environments anywhere on your system, and you can activate an environment from any path with but it just feels clunky. 

```bash
pixi init /path/to/env
pixi shell --manifest-path /path/to/env
```

I went as far as making some [shell aliases](https://github.com/jacobtomlinson/dotfiles/blob/9a35055e975d8622282279f32ea505b89957de5c/.zshrc.d/python_pixi.zsh) to try and reproduce `conda activate <name>`, `conda create` and `conda env list` with `pixi` but it just feels a bit hacky. I could probably also do something equally hacky with `uv`, but as a long time `conda` user it just feels more natural to me to keep using it for these central environments.
````

## Final thoughts on `pip`

This post started out explaining how `pip` came and solved an important problem solving dependencies for Python more than 15 years ago. It is the tool recommended by the [Python Software Foundation](https://www.python.org/psf-landing/) to install Python packages. However, in 2025 I do not use it at all. The needs of the community have moved on again and now I'm jumping between `uv`, `pixi` and `conda` until the next thing comes along.

[^1]: Funnily enough drawing this diagram was a perfect use case for `pixi`/`conda`, which you'll learn more about in the article. The diagram is drawn with `graphviz` which is a non-Python binary package which isn't available on PyPI and therefore can't be installed with `pip`, `uv`, `poetry` or any other tools that consumes from PyPI. It must either be installed with your system package manager like `apt` or `brew`, or be installed with `conda`.
