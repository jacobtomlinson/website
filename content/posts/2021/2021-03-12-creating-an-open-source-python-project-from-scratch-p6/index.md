---
title: "Test driven development in Python"
series: ["Creating an open source Python project from scratch"]
date: 2021-03-12T00:00:00+00:00
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

# What is test driven development (TDD)?

Test driven development is a style of development where you write your tests **before** you write your code.

There are some main benefits to doing this:

- It forces you to think more about the design of your code ahead of time.
- It ensures your code is tested, often tests come as an after thought.
- It makes testing easier because your code will be more modular and testable.

Many times in the past I've dived head first into writing code when I've had an idea. Then when it comes to writing tests it can be tricky to hook them in to what you already have. Doing a TDD approach has really helped me tackle these issues personally, but it did take some getting used to.

## Docstring driven development

Given that we already have our testing infrastructure set up to test the examples in our docstrings this gives us a really nice place to start.

This also gives us a fourth benefit:

- Ensure our functions have docstrings with working examples.

For this post we are going to add a new function to our package called `is_float`. This function is much like our `is_number` function but instead it will check if the passed object has a decimal value. So let's by importing this in our top level `__init__.py`.

```python
"""Utility functions to calculate if an object is a number."""
from .is_number import is_number
from .is_float import is_float

...
```

Then we will create a new file called `is_float.py` and define our function with a docstring. We will also add a couple of example usages, but we wont implement the function so these examples will not be valid yet.

```python
def is_float(in_value):
    """Checks if a value is a valid float.

    Parameters
    ----------
    in_value
        A variable of any type that we want to check is a float.

    Returns
    -------
    bool
        True/False depending on whether it was a float.

    Examples
    --------
    >>> is_float(1.5)
    True
    >>> is_float(1)
    False

    """
    pass
```

Next we can run `pytest` to see what happens.

```console
$ pytest
============================================= test session starts =============================================
platform darwin -- Python 3.7.4, pytest-5.0.1, py-1.8.0, pluggy-0.12.0
rootdir: /Users/jtomlinson/Projects/jacobtomlinson/is-number, inifile: setup.cfg
plugins: env-0.6.2, asyncio-0.10.0, timeout-1.4.2
collected 4 items

is_number/is_float.py F                                                                                 [ 25%]
is_number/is_number.py .                                                                                [ 50%]
is_number/tests/test_is_number.py ..                                                                    [100%]

================================================== FAILURES ===================================================
____________________________________ [doctest] is_number.is_float.is_float ____________________________________
007         A variable of any type that we want to check is a float.
008
009     Returns
010     -------
011     bool
012         True/False depending on whether it was a float.
013
014     Examples
015     --------
016     >>> is_float(1.5)
Expected:
    True
Got nothing

/Users/jtomlinson/Projects/jacobtomlinson/is-number/is_number/is_float.py:16: DocTestFailure
===================================== 1 failed, 3 passed in 0.09 seconds ======================================

```

We can see here that our test has failed. Our example showed that `is_float(1.5)` returns `True`, but instead it returned `None` because we haven't implemented it. A function which just passes returns `None`.

Now let's implement some code.

```python
def is_float(in_value):
    """Checks if a value is a valid float.

    Parameters
    ----------
    in_value
        A variable of any type that we want to check is a float.

    Returns
    -------
    bool
        True/False depending on whether it was a float.

    Examples
    --------
    >>> is_float(5)
    True
    >>> is_float(1)
    False

    """
    return isinstance(in_value, float)
```

This time if we run `pytest` things should pass.

```console
$ pytest
============================================= test session starts =============================================
platform darwin -- Python 3.7.4, pytest-5.0.1, py-1.8.0, pluggy-0.12.0
rootdir: /Users/jtomlinson/Projects/jacobtomlinson/is-number, inifile: setup.cfg
plugins: env-0.6.2, asyncio-0.10.0, timeout-1.4.2
collected 4 items

is_number/is_float.py .                                                                                 [ 25%]
is_number/is_number.py .                                                                                [ 50%]
is_number/tests/test_is_number.py ..                                                                    [100%]

========================================== 4 passed in 0.08 seconds ===========================================

```

Success! We completed our first bit of test driven development.

## Hunting bugs with TDD

Our function passes the tests but it isn't perfect yet. There are some differences from how the `is_number` function works that will cause our new `is_float` to not work as expected.

But that's the point here. We start our work by setting our expectations, then implement code that meets those expectations. The two examples we have started with check an integer and a float, but what about strings? We haven't set any expectations for strings, so we can't have any confidence that things will work.

This is the same process as hunting down a bug with TDD. If you are reviewing an issue from a user who says "your code didn't do what I expected" the first question we ask is "what did you expect?". If we agree that this is a valid expectation we can codify that expectation as a test.

For example I could raise an issue on the `is-number` project now which says:

> I expected `is_float("1.5")` to be `True`, but it is `False`.

This is a valid expectation, so in resolving this issue let's start by adding another example which demonstrates this expectation.

```python
def is_float(in_value):
    """Checks if a value is a valid float.

    Parameters
    ----------
    in_value
        A variable of any type that we want to check is a float.

    Returns
    -------
    bool
        True/False depending on whether it was a float.

    Examples
    --------
    >>> is_float(1.5)
    True
    >>> is_float(1)
    False
    >>> is_float("1.5")
    True

    """
    return isinstance(in_value, float)
```

Now if we run our tests again it will fail.

```console
$ pytest
============================================= test session starts =============================================
platform darwin -- Python 3.7.4, pytest-5.0.1, py-1.8.0, pluggy-0.12.0
rootdir: /Users/jtomlinson/Projects/jacobtomlinson/is-number, inifile: setup.cfg
plugins: env-0.6.2, asyncio-0.10.0, timeout-1.4.2
collected 4 items

is_number/is_float.py F                                                                                 [ 25%]
is_number/is_number.py .                                                                                [ 50%]
is_number/tests/test_is_number.py ..                                                                    [100%]

================================================== FAILURES ===================================================
____________________________________ [doctest] is_number.is_float.is_float ____________________________________
011     bool
012         True/False depending on whether it was a float.
013
014     Examples
015     --------
016     >>> is_float(1.5)
017     True
018     >>> is_float(1)
019     False
020     >>> is_float("1.5")
Expected:
    True
Got:
    False

/Users/jtomlinson/Projects/jacobtomlinson/is-number/is_number/is_float.py:20: DocTestFailure
===================================== 1 failed, 3 passed in 0.10 seconds ======================================
```

This is a really nice workflow. A user reported a problem, and we reproduced that problem in the form of a failing test. We updated our expectations and discovered they weren't quite right.

Next we need to update the code to get this test to pass, but without breaking the other tests.

```python
def is_float(in_value):
    """Checks if a value is a valid float.

    Parameters
    ----------
    in_value
        A variable of any type that we want to check is a float.

    Returns
    -------
    bool
        True/False depending on whether it was a float.

    Examples
    --------
    >>> is_float(1.5)
    True
    >>> is_float(1)
    False
    >>> is_float("1.5")
    True

    """
    try:
        return not float(in_value).is_integer()
    except (ValueError, TypeError):
        return False
```

Here we do the same as `is_number` and convert out input to a float. But then we use the `not ... is_integer()` pattern to check if the float has a decimal component.

Now our tests should pass.

```console
pytest
============================================= test session starts =============================================
platform darwin -- Python 3.7.4, pytest-5.0.1, py-1.8.0, pluggy-0.12.0
rootdir: /Users/jtomlinson/Projects/jacobtomlinson/is-number, inifile: setup.cfg
plugins: env-0.6.2, asyncio-0.10.0, timeout-1.4.2
collected 4 items

is_number/is_float.py .                                                                                 [ 25%]
is_number/is_number.py .                                                                                [ 50%]
is_number/tests/test_is_number.py ..                                                                    [100%]

========================================== 4 passed in 0.10 seconds ===========================================
```

## More tests

Now that we've written our user facing docstring and tested that it meets our expectations we may also want to add some more tests out of sight of the user. We probably want to throw a whole range of inputs at our function to ensure they come out with the result we expect, but we don't want to clutter our docstring up with this.

Let's create a new file called `is_number/tests/test_is_float.py` with some more tests in.

```python
from datetime import datetime

from is_number import is_float


def test_is_float():
    assert is_float(1.1)

    assert not is_float(1)
    assert not is_float(1.0)
    assert not is_float("Hello world")
    assert not is_float({"Hello": "world"})
    assert not is_float(datetime.now())
    assert not is_float(lambda foo: foo)
```

Then let's run `pytest` one last time to check things still pass.

```console
$ pytest
============================================= test session starts =============================================
platform darwin -- Python 3.7.4, pytest-5.0.1, py-1.8.0, pluggy-0.12.0
rootdir: /Users/jtomlinson/Projects/jacobtomlinson/is-number, inifile: setup.cfg
plugins: env-0.6.2, asyncio-0.10.0, timeout-1.4.2
collected 5 items

is_number/is_float.py .                                                                                 [ 20%]
is_number/is_number.py .                                                                                [ 40%]
is_number/tests/test_is_float.py .                                                                      [ 60%]
is_number/tests/test_is_number.py ..                                                                    [100%]

========================================== 5 passed in 0.16 seconds ===========================================
```

## Summary

In this post we have covered:

- Codifying our expectations as tests
- Writing our tests before our code
- Solving bugs by converting a user's issue into a failing test, then fixing it

In future posts we will cover:

- Automating our tests
- Automating future releases
- Generating documentation and hosting it
- Creating a community
- Handling future maintenance