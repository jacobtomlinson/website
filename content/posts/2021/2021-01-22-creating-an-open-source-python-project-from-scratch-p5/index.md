---
title: "Testing your Python package"
series: ["Creating an open source Python project from scratch"]
date: 2021-01-22T00:00:00+00:00
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

In this post we will cover testing our code.

# Testing

There are many many great resources out there for learning about testing software. In this post I'm going to try and focus on simple examples that you can use to get started quickly. Once you have a good foundation for your tests you can then dive into [mocking](https://www.toptal.com/python/an-introduction-to-mocking-in-python), [replaying HTTP requests](https://cassette.readthedocs.io/en/latest/) or even [hypothesis testing](https://hypothesis.readthedocs.io/en/latest/).

There are many different types of tests you can write: unit tests, integration tests, end-to-end tests, acceptance tests, system tests, functional tests, etc, etc. I'm not going to dig into each one of these and explain all the differences, instead I'm going to focus on the two categories of tests I commonly see in open source packages, **white box** and **black box**.

## White box tests

White box tests are where we can look inside the code and test how it works (I feel like clear box or transparent box would make more sense than white box). A common type of white box testing is unit testing, where you take each function in your code and write tests to ensure that function does exactly what it says it does.

To do this in our example `is-number` package we are going to use [pytest](https://docs.pytest.org/en/latest/). Pytest is a framework for writing and running tests in Python projects. We are only going to scratch the surface of what it can do, but it will be useful for getting up and running.

We also want to keep track of the dependencies required for developing and testing our project. These shouldn't go in `requirements.txt` because the end user doesn't need them. So instead we will create a new file called `requirements_test.txt` and put them in there.

```
pytest
```

Next we can install these testing tools.

```bash
pip install -r requirements_test.txt
```

Next we need to create some tests. By default `pytest` will explore the directory structure of a project looking for directories called `tests`. Then inside those directories it will look for Python files that begin with `test_`. Then inside those files it will look for functions whose name also start with `test_`. It then runs each of these functions and checks that it runs successfully.

Let's start with an example. Within our `is_number` directory will create a new `tests` directory and inside that create a file called `test_is_number.py`.

```bash
mkdir -p is_number/tests
touch is_number/tests/test_is_number.py
```

```python
from is_number import is_number


def test_is_number():
    assert is_number(1)

def test_is_not_number():
    assert not is_number("Hello world")
```

In our test file we are importing the `is_number` function from our `is_number` package and then defining two tests which make two assertions.

In Python the `assert` statement will raise an exception if the following expression is not `True`. This is very useful when testing, but can also be useful when writing Python code generally. Whenever you make an assumption in your code (this thing will always return a number) you can `assert` that assumption. Then if your assumption is ever wrong your users will get an assertion error instead of some other obscure issue. This can be very useful for tracking down bugs.

In our tests here we are asserting that `1` is a number and that `"Hello world"` is not a number.

Now we can run our tests with `pytest`.

```console
$ pytest is_number
========================== test session starts ==========================
platform darwin -- Python 3.7.3, pytest-5.0.1, py-1.8.0, pluggy-0.13.0
rootdir: /Users/jtomlinson/Projects/jacobtomlinson/is-number
collected 2 items

is_number/tests/test_is_number.py ..                                                                                                                                                                       [100%]

======================= 2 passed in 0.07 seconds ========================
```

We ran `pytest` on our `is_number` package and can see that it detected our two tests and that they passed.

These tests were unit tests because they ran on our one function, our one unit of code.

Because this is a white box test we can look at our function and try to come up with other tests.

```python
def is_number(in_value):
    """Checks if a value is a valid number.
    [truncated docstring]
    """
    try:
        float(in_value)
        return True
    except ValueError:
        return False
```

Our function tries to convert our `in_value` to a `float` and then returns `True` or `False` depending on whether this was successful. But what if we do something the `float()` method cannot deal with.

```python
def test_is_not_number():
    assert not is_number("Hello world")
    assert not is_number({"Hello": "world"})
```

The `float()` method expects a number or a string, but what if we pass it a dictionary? I've added another assertion to our `is_not_numnber` test and now our test fails.

```console
$ pytest is_number
========================== test session starts ==========================
platform darwin -- Python 3.7.3, pytest-5.0.1, py-1.8.0, pluggy-0.13.0
rootdir: /Users/jtomlinson/Projects/jacobtomlinson/is-number
collected 2 items

is_number/tests/test_is_number.py .F                                                                                                                                                                       [100%]

========================== FAILURES ==========================
_____________________ test_is_not_number _____________________
    def test_is_not_number():
>       assert not is_number({"Hello": "world"})

is_number/tests/test_is_number.py:10:
_ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _

in_value = {'Hello': 'world'}

    def is_number(in_value):
        try:
>           float(in_value)
E           TypeError: float() argument must be a string or a number, not 'dict'

is_number/__init__.py:9: TypeError
============= 1 failed, 1 passed in 0.09 seconds =============
```

Our test has failed here because the `float()` function raised an exception because it didn't get an object of the type it was expecting. We can fix our code so that if a `TypeError` is raised we also return `False`, this makes sense because a dictionary is not a number.

```python
def is_number(in_value):
    """Checks if a value is a valid number.
    [truncated docstring]
    """
    try:
        float(in_value)
        return True
    except (ValueError, TypeError):
        return False
```

Running our tests again they should pass.

```console
$ pytest is_number
============================== test session starts ==============================
platform darwin -- Python 3.7.3, pytest-5.0.1, py-1.8.0, pluggy-0.13.0
rootdir: /Users/jtomlinson/Projects/jacobtomlinson/is-number
collected 2 items

is_number/tests/test_is_number.py ..                                                                                                                                                                       [100%]

============================ 2 passed in 0.05 seconds ============================
```

## Black box tests

Now let's move onto black box tests. These are tests where we don't look inside the function and come up with tests based on the code, but instead look at the interface of the function and write tests based on that. I like to think of these as user tests, as we are testing what the user sees instead of what the developer sees.

We already know what the user sees because we wrote our docstring in the last post. Our user can see the documentation, but not the code itself.

### Doctest

We took the time to write some examples for our user, so let's test that these actually work. Luckily `pytest` already has a feature for that.

```console
$ pytest --doctest-modules is_number
============================== test session starts ==============================
platform darwin -- Python 3.7.3, pytest-5.0.1, py-1.8.0, pluggy-0.13.0
rootdir: /Users/jtomlinson/Projects/jacobtomlinson/is-number
plugins: env-0.6.2, asyncio-0.10.0, timeout-1.4.2
collected 3 items

is_number/is_number.py .                                                                                                                                                                                          [ 33%]
is_number/tests/test_is_number.py ..                                                                                                                                                                              [100%]

============================ 3 passed in 0.07 seconds ============================
```

By adding the `--doctest-modules` flag pytest has also collected the examples from the docstring and checked that each line of code gives the correct response.

For example in our docstring we show that `is_number(1)` is `True`, so pytest will run this code and verify that this is correct.

```python
>>> is_number(1)
True
```

We can add some config to our `setup.cfg` file to enable this flag by default.

```ini
[tool:pytest]
addopts = --doctest-modules
```

Now all of our docstrings will be tested when we run `pytest`.

### Writing more tests

We've written some tests after inspecting the code, and we've tested that our examples in our docstring work as expected. So lastly we could also write some more tests, but based purely on what the docstring tells us.

Our docstring tells us `You can also pass more complex objects, these will all be ``False``.`. So let's add some more lines to our tests to ensure this is valid. Specifically let's add a complex object like a `datetime` and also a callable.

```python
from datetime import datetime

from is_number import is_number


def test_is_number():
    assert is_number(1)


def test_is_not_number():
    assert not is_number("Hello world")
    assert not is_number({"Hello": "world"})
    assert not is_number(datetime.now())
    assert not is_number(lambda foo: foo)
```

```console
$ pytest --doctest-modules is_number
============================== test session starts ==============================
platform darwin -- Python 3.7.3, pytest-5.0.1, py-1.8.0, pluggy-0.13.0
rootdir: /Users/jtomlinson/Projects/jacobtomlinson/is-number
plugins: env-0.6.2, asyncio-0.10.0, timeout-1.4.2
collected 3 items

is_number/is_number.py .                                                                                                                                                                                          [ 33%]
is_number/tests/test_is_number.py ..                                                                                                                                                                              [100%]

============================ 3 passed in 0.13 seconds ============================
```

Running our tests again passes, which gives us confidence that what we wrote in the docstring is correct.

## Developer Documentation

The last thing we should do is ensure other developers can find and run our tests. The testing pattern we've done here is pretty standard, but there are still many ways to do this and it is common to include some information on how to run the tests.

Let's add a section to the README.

```rst
Testing
-------

This project uses ``pytest`` to run tests and also to test docstring examples.

Install the test dependencies.

.. code-block:: bash

   $ pip install -r requirements_test.txt

Run the tests.

.. code-block:: bash

    $ pytest
    === 3 passed in 0.13 seconds ===
```

## Summary

In this post we have covered:

- Common Python testing tools
- Testing our code
- Testing our docstrings

In future posts we will cover:

- Test driven development
- Automating our tests
- Automating future releases
- Generating documentation and hosting it
- Creating a community
- Handling future maintenance