---
title: "Building a contributor community for your open source project"
series: ["Creating an open source Python project from scratch"]
date: 2021-04-30T00:00:00+00:00
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

With our open source project published on GitHub we probably want to allow folks to contribute changes. Some users of the project may find bugs, or desire extra features and will open issues to tell you. Users who have the skills required to make that change can open a Pull Request on GitHub to propose it. As the maintainer you can then review and merge those changes.

In this post we will discuss things you can do to make your project more approachable for contributors.

## contributing.rst

To get help with your project we want to make sure we provide guidance for contributors on how to do things. Even if you don't have strong feelings about how folks should go about contributing and are happy for things to be done the usual GitHub way it's best to write that down in a documentation file.

We already made a start on this in earlier posts when we documented how we lint, test and release our project. It's really important to document these things just for your own use, because you will forget! But these notes are now the foundations of our contributing guidelines.

GitHub recommends you create a CONTRIBUTING file either in the root directory, `docs/` directory or `.github/` directory of your project. Let's do this for our sample `is-number` project. The case and extension of the file doesn't really matter, so given our documentation is in rst let's create a file called `docs/contributing.rst`. To start the file let's move the `Developing` and `Testing` sections from our `README.rst` to this file.

```rst
Contributing
============

Developing
----------

This project uses ``black`` to format code and ``flake8`` for linting. We also support ``pre-commit`` to ensure
these have been run. To configure your local environment please install these development dependencies and set up
the commit hooks.

.. code-block:: bash

   $ pip install black flake8 pre-commit
   $ pre-commit install

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

We also need to include this in the table of contents in `docs/index.rst`. I'm going to create a second `toctree` section with the title `Developer` so we can separate our our user documentation from our developer documentation.

```rst
...

.. toctree::
   :maxdepth: 2
   :caption: Usage

   api
   faq


.. toctree::
   :maxdepth: 2
   :caption: Developer

   contributing
```

Now if we run `sphinx-autobuild docs docs/_build/html` we can preview our new documentation section.

![Screenshot of Sphinx preview showing the Contributing page](https://i.imgur.com/AOFNc9u.png)

This is a good foundation for our contributing docs, but we can add more.

Folks are going to read this documentation for a couple of reasons. First is folks who want to make a change and are looking for the "correct" way to go about things. The second group are folks who just generally want to contribute, but don't know where to start.

We want to make both of these groups feel welcome but also lay out the rules that they should follow to ensure they are submitting a quality contribution which can be reviewed easily. We then need to ensure they have all the information they need to get their local development environment set up and make changes which will pass our CI. Lastly we should talk them through how to make the pull request.

There is definitely a balance to be struck when writing this kind of documentation. We want to explain how to set up a local development environment for this project, but not how to set up Python in its entirety. We want to explain how to run the tests, but also accommodate folks who want to use IDEs and GUI tools for doing things. We want to explain how we want to receive Pull Requests, but not explain what Pull Requests are as a whole. Choosing the right place to draw a line in the sand when it comes to assumptions about user ability levels is tricky.

My preference is to assume that the contributor is competent enough with Python that they can make changes to code, but that they have a bare bones or minimal local setup. I also tend to assume we are using the CLI but provide enough context that IDE user's should be able to translate the steps to their tools. When writing documentation like this I often create a new conda environment containing just Python `conda create -n env-name python -y` and then walk through the steps to ensure I can get tests passing.

Here's an example of how we can update our `contributing.rst` file to cover all of this.

```rst
Contributing
============

We love contributions here in ``is-number``! If you're looking for something to work on then check out our
`issue tracker <https://github.com/jacobtomlinson/is-number/issues>`_ for open issues.

If you want to make a contribution to ``is-number`` then please raise a
`Pull Request <https://github.com/jacobtomlinson/is-number/pulls>`_ on GitHub.

To help speed up the review process please ensure the following:

- The PR addresses an open issue.
- All tests are passing locally with ``pytest``.
- The project passes linting with ``black`` and ``flake8``.
- If adding a new feature you also add documentation.

Developing
----------

To check out a local copy of the project you can `fork the project on GitHub <https://github.com/jacobtomlinson/is-number/fork>`_
and then clone it locally.

.. code-block:: bash

   $ git clone git@github.com:yourusername/is-number.git
   $ cd is-number

This project uses ``black`` to format code and ``flake8`` for linting. We also support ``pre-commit`` to ensure
these have been run. To configure your local environment please install these development dependencies and set up
the commit hooks.

.. code-block:: bash

   $ pip install black flake8 pre-commit
   $ pre-commit install

You can check that things are working correctly by calling pre-commit directly.

.. code-block:: bash

   $ pre-commit run --all-files
   black......................................Passed
   flake8.....................................Passed

These checks will be run automatically when you make a commit.

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

If you are working on a new feature please add tests to ensure the feature works as expected. If you are working on
a bug fix then please add a test to ensure there is no regression.

Tests are stored in ``is_number/tests`` and follow the pytest format.

.. code-block:: python

    from is_number import is_number

    def test_is_number():
        assert is_number(1)
        assert not is_number("hello world!")

Making a Pull Request
---------------------

Once you've made your changes and are ready to make a Pull Request please ensure tests and linting pass locally before pushing to GitHub.
When making your Pull Request please include a short description of the changes, but more importantly why they are important. Perhaps by
writing a before and after paragraph with user examples.

Also consider how your title look when it appears in a changelog. Does it full describe the change to an outside user? For example
``Add support for checking iterables contain all numbers`` is a much better title than ``Fixes #76``.

.. code-block:: markdown

    # Add support for checking iterables contain all numbers

    Closes #56

    **Changes**

    This PR allows the inspection of structures such as lists and sets to check if all elements are numbers.

    **Before**

    If a user passed a list of all numbers to `is_number` it would return `False`.

    ```python
    >>> from is_number import is_number
    >>> is_number([0,1,2])
    False
    ```

    **After**

    If a user passes a list of all numbers it will return true, unless they set the `strict` keyword argument to `True`.


    ```python
    >>> from is_number import is_number
    >>> is_number([0,1,2])
    True
    >>> is_number([0,1,2], strict=True)
    False
    ```
```

Let's recap the changes we've added to the file here.

- Added an introduction to make folks feel their contributions are welcome.
- Added a checklist of things you will be checked by CI or the reviewer.
- Added some more instructions on cloning the repo and running the linting checks manually.
- Added some advice on raising a good Pull Request including an example.

## maintaining.rst

In addition to adding contributing guidelines I also like to add maintaining guidelines too.

When you first start out these will likely just be notes for yourself, but as your project grows you may
want to give maintaining power to others and so it's great to have things written down.

Let's create a `docs/maintaining.rst` file with some basic review guidance. We can also move the `Releasing` section
from our `README.rst` here.

```rst
Maintaining
===========

Reviewing Pull Requests
-----------------------

This project generally accepts any Pull Request which improves the project.

Any small issues which fix typos or improve documentation can be merged straight in.

Any bug fix or enhancement PRs should reference an open issue which describes the problem. This gives
us and other users the opportunity to discuss the change before anyone invests time in implementing it.

Typically we have a "yes and" policy when it comes to reviewing where we generally accept whatever is contributed even
if it's not quite right. If you have a small amount of feedback during review, such as the user forgot to run ``black`` or
you want to reword something in a docstring it's preferable to just push extra commits to the PR, or just merge
and raise a follow up PR to tweak things.

For larger design or implementation feedback then feel free to push this back on to the contributor.

Releasing
---------

Releases are published automatically when a tag is pushed to GitHub.

.. code-block:: bash

   # Set next version number
   export RELEASE=x.x.x

   # Create tags
   git commit --allow-empty -m "Release $RELEASE"
   git tag -a $RELEASE -m "Version $RELEASE"

   # Push
   git push upstream --tags
```

Make sure you add this new file to the `toctree` in `docs/index.rst`.

```rst
...

.. toctree::
   :maxdepth: 2
   :caption: Developer

   contributing
   maintaining
```

Now let's branch, commit, push and merge our changes.

```console
$ git checkout -b contributing-maintaining-docs
$ git add -A
$ git commit -m "Add contributing and maintaining docs
$ git push --set-upstream origin contributing-maintaining-docs
```

![Screenshot of merged contributing-maintaining-docs PR](https://i.imgur.com/w4wZMsg.png)

With our changes merged we can see GitHub will add links to our contributing documentation in various places including the bottom right when users open a new issue.

![Screenshot of new issue with contributing guidelines link](https://i.imgur.com/g7A5s12.png)

### PR templates

Now that we've set out guidance for our contributors to follow let's try and make this as easy as possible.

We should assume that not everyone raising a PR will have read our `contributing.rst` documentation, so let's add
a Pull Request template to point them to that file and give them a framework to fill in.

This is pretty much the same as adding an issue template like we did in [part 10](https://jacobtomlinson.dev/posts/2021/building-a-user-community-for-your-open-source-project/) of this series.

We need to create a new file called `.github/PULL_REQUEST_TEMPLATE.md` with the following contents.

```markdown
<!-- If you've not read our contribution guidelines please take a look at https://is-number.readthedocs.io/en/latest/contributing.html -->

Closes #

**Checklist**

- [ ] The PR addresses an open issue.
- [ ] All tests are passing locally with ``pytest``.
- [ ] The project passes linting with ``black`` and ``flake8``.
- [ ] (optional) If adding a new feature also added documentation.

**Changes**

<!-- Describe the PR changes here -->

**Before**

<!-- What could a user not do or what bug did they experience before this PR -->

**After**

<!-- What has this PR enabled -->
```

Here we've pointed folks to the contributing guidelines in case they haven't seen them and then taken the example Pull Request from our contributing guidelines along with the contribution checklist and turned them into a real markdown checklist and a loose structure for folks to fill in.

Now when a user raises a pull request the body will be pre-populated to help them fill in all the info we need.

![Screenshot of new PR with template filled in](https://i.imgur.com/qQSXizZ.png)

## Summary

In this post we have covered:

- Adding contributing documentation
- Add maintainer documentation
- Adding a Pull Request template

In future posts we will cover:

- Branding your project
- Handling future maintenance