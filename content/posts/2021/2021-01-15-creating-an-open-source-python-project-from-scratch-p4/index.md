---
title: "Documenting your Python code"
series: ["Creating an open source Python project from scratch"]
date: 2021-01-15T00:00:00+00:00
draft: false
author: "Jacob Tomlinson"
categories:
  - blog
tags:
  - Python
  - GitHub
  - Tutorial
thumbnail: python
---

This post will cover documenting our code. Specifically adding documentation within the code itself.

# Docstrings

Right now our code is undocumented, so if the user inspects our function they will only see the interface (the way you call it) but with no other context. We can use [IPython](https://ipython.org/) to quickly inspect this.

```
$ ipython
Python 3.7.4 (default, Aug 13 2019, 15:17:50)
Type 'copyright', 'credits' or 'license' for more information
IPython 7.8.0 -- An enhanced Interactive Python. Type '?' for help.

In [1]: from is_number import is_number

In [2]: is_number?

Signature: is_number(in_value)
Docstring: <no docstring>
File:      ~/Projects/jacobtomlinson/is-number/is_number/is_number.py
Type:      function
```

IPython has a very useful feature where you can call a function name with a `?` at the end and it will show you information about the function.

We can see here that we call the function with an `in_value`, but it has no docstring.

The docstring is a string that you place at the start of your function to describe how it works. Typically you can put whatever you want in this string, but a couple of standards have emerged; [Numpy style docstrings](https://numpydoc.readthedocs.io/en/latest/format.html) and [Google style docstrings](https://sphinxcontrib-napoleon.readthedocs.io/en/latest/example_google.html). For this example we are going to use Numpy style.

## Adding a docstring

To start let's head back to our `is_number.py` file and add a simple docstring.

```python
def is_number(in_value):
    """Checks if a value is a valid number."""
    try:
        float(in_value)
        return True
    except ValueError:
        return False
```

Here we've added a string to our function that describes what the function does.

Now if we run `ipython` again and inspect our function we can see our docstring.

```python
In [2]: is_number?

Signature: is_number(in_value)
Docstring: Checks if a value is a valid number.
File:      ~/Projects/jacobtomlinson/is-number/is_number/is_number.py
Type:      function
```

We can also find this string in the `__doc__` attribute of our function.

```python
>>> from is_number import is_number
>>> is_number.__doc__
'Checks if a value is a valid number.'
```

## Inputs and outputs

This is a good start, we are telling the user what the function does, but we can tell them more about what goes in and what comes out of our function.

```python
def is_number(in_value):
    """Checks if a value is a valid number.

    Parameters
    ----------
    in_value
        A variable of any type that we want to check is a number.

    Returns
    -------
    bool
        True/False depending on whether it was a number.

    """
    try:
        float(in_value)
        return True
    except ValueError:
        return False
```

We've added a couple of numpydoc headings; `Parameters` and `Returns`. This tells the user what they can put into the function and what will be returned.

These headings are also special, they are part of the numpydoc standard which means that other tools know how to interpret them. This will be really useful in a future post when we come to building a documentation website.

## Seeing is believing

For most users seeing our function in action will really solidify what it does in their mind.

To solve this we can add an `Examples` section, which I would argue is the most important section in any bit of documentation.

```python
def is_number(in_value):
    """Checks if a value is a valid number.

    Parameters
    ----------
    in_value
        A variable of any type that we want to check is a number.

    Returns
    -------
    bool
        True/False depending on whether it was a number.

    Examples
    --------
    >>> is_number(1)
    True
    >>> is_number(1.0)
    True
    >>> is_number("1")
    True
    >>> is_number("1.0")
    True
    >>> is_number("Hello")
    False

    You can also pass more complex objects, these will all be ``False``.

    >>> is_number({"hello": "world"})
    False
    >>> from datetime import datetime
    >>> is_number(datetime.now())
    False

    Even something which contains all numbers will be ``False``, because it is not itself a number.

    >>> is_number([1, 2, 3, 4])
    False

    """
    try:
        float(in_value)
        return True
    except ValueError:
        return False
```

Here we've added a bunch of examples and also some slightly more complex ones with comments on why they give the value that they do.

Docstrings are written in [reStructuredText (RST)](https://en.wikipedia.org/wiki/ReStructuredText) just like our README.

Also note how our docstring is now much longer than the code itself. The code is very trivial in this example, but it is still common for good docstrings to get longer than the code it describes as we communicate to our users how to use our code.

## Module level docstrings

As well as documenting our functions it is also important to document our module too. This makes it easy for users to figure out what our module does over all.

You can do this by adding a string to the very top of each of your Python files. For example our `__init__.py` could look like this

```python
"""Utility functions to calculate if an object is a number."""
from .is_number import is_number
from ._version import get_versions

__version__ = get_versions()["version"]
del get_versions
```

## Summary

In this post we have covered:

- Documenting our code with docstrings
- Different docstring styles
- Writing examples

In future posts we will cover:

- Adding tests
- Automating those tests
- Automating future releases
- Generating documentation and hosting it
- Creating a community
- Handling future maintenance