---
title: "Python version epochs are broken"
date: 2024-05-01T00:00:00+00:00
draft: false
author: "Jacob Tomlinson"
categories:
  - blog
tags:
  - python
  - versioning
---

In [PEP440](https://peps.python.org/pep-0440) Python introduced [Version Epochs](https://packaging.python.org/en/latest/specifications/version-specifiers/#version-epochs) as a mechanism to allow projects to change versioning scheme. Unfortunately there's no way I could see a project actually making use of this without confusing their users.

I [very much regret using CalVer](https://jacobtomlinson.dev/posts/2023/sometimes-i-regret-using-calver/) in some of the projects I work on. When building software libraries it can be very valuable to be able to communicate information via a version number, and CalVer doesn't allow you to do that. In [Dask we've been discussing moving away from CalVer](https://github.com/dask/community/issues/372) to a different scheme like [EffVer](https://jacobtomlinson.dev/effver/) and Python version epochs appeared to be the machanism by which we could do this.

> "If a project is using date based versions like `2014.04` and would like to switch to semantic versions like `1.0`, then the new releases would be identified as older than the date based releases when using the normal sorting scheme [...] **by specifying an explicit epoch, the sort order can be changed appropriately**, as all versions from a later epoch are sorted after versions from an earlier epoch"
>
> _[PEP440 and the Python Packaging Guide](https://packaging.python.org/en/latest/specifications/version-specifiers/#version-epochs)_

## Version Epochs

In Python an epoch is a numeric prefix for version numbers separated by an `!`.

```text
E!X.Y  # Version identifier with epoch
```

I've worked in the Python ecosystem for over a decade and I've never seen a project use an epoch, so here's a quick overview as you probably haven't seen them being used either. 

If you don't specify one in your version number then it is assumed to be epoch `0`. For example Dask version `2.30.1`, which was the final SemVer release in late 2020, is technically version `0!2.30.1`. Nobody really needs to worry about this though as you can omit the epoch and most projects never increment their epoch.

The next release after that was when Dask switched to CalVer and was version `2020.12.0`. This was technically version `0!2020.12.0` and was part of the same epoch, but it sorted correctly because `2020` is greater than `2`. 

When moving on from CalVer to a different scheme the Dask maintainers could argue that "all CalVer releases were the _third major version_ of Dask". So in our new scheme we would probably jump to `4.0.0`.

The Python Packaging Guide states that projects using date based versions can switch to versions like `1.0` using an epoch, so moving to `4.0.0` should be fine right?

## Epochs greater than zero can't be omitted

To understand how this would work in practice [I created a dummy package called `epochexperiments`](https://github.com/jacobtomlinson/epochexperiments) and pushed it to the [PyPI testing index](https://test.pypi.org/) so that I could play around.

I made a few releases to mimic Dask's historic tags with the major versions `1`, `2` and `2024`. Then I explored incrementing the epoch and releasing `1!4` to begin the new scheme.

Installing the `epochexperiments` package with no constraints correctly installed the most recent `1!4.0.1` tag of the project, so sorting definitely works as expected. However my assumption going into this experiment was that we could switch to a scheme like `1.0`, so I had hoped that users would just be able to install `epochexperiments==4.0.1`.

Unfortunately once the epoch has incremented past zero it is no longer optional, so that install fails. Instead you have to install `epochexperiments==1!4.0.1` and explicitly state which epoch the version is in.

I can understand why it has been implemented this way. If you have a hostoric version `0!4.0.0` and new version `1!4.0.0`, then if a user pinned to `==4.0.0` they should be able to be confident that this is a hard pin and will resolve deterministically forever in the future. But this behaviour means you can never switch to a `1.0` version scheme, only a `1!1.0` scheme.

## Wildcards don't behave as expected

Another oddity that I noticed when experimenting is that wildcards changed the behaviour of constraints in an unexpected way.

If I install `epochexperiments>=2` I get the very latest release, which in this example is `1!4.0.1`. However if I install `epochexperiments>=2.*` I get `2024.4.1`, which is the latest release from the `0!` epoch. Adding a wildcard seems to have added an implicit `<1!` constraint that isn't there otherwise.

I'm not sure if this is by design or just a bug in the implementation.

## Should we use epochs in Dask?

If moving to a fourth major version in Dask means that all version numbers have to be prefixed with a `1!` then there's no way we would do this, it would just be too confusing for users. Every time a user had to think about the Dask version number they wanted they would have to jump the hurdle of learning about epochs, and that's a lot for us to ask. In my opinion I think it's also just rather ugly.

To me this means that epochs in Python are not fit for purpose at all. They are intended to help projects move to different version schemes. But by making them an explicit prefix to the version number this breaks versioning conventions in many languages and would create confusion for most users.

## Could this be fixed?

My assumption when starting this exploration was that we would need to specify the epoch when publishing the package, but users would be able to omit it when they consumed the package. 

```bash
pip install 'epochexperiments>=4'  # In my opinion this should install 1!4.0.1
```

This would require some changes to how version ordering is calculated when resolving packages, but as a user it feels more intuitive. It would also allow projects to truly switch back to a `1.0` scheme from CalVer.

To make this work there would also have to be some restrictions around what version numbers could be published. For example having `0!4.0.0` and `1!4.0.0` shouldn't be possible. When pushing the `1!4.0.0` version PyPI should reject it with an error stating this version already exists. Perhaps a restriction could be added so that if a major version has been used in one epoch it is not possible to use it in another. So with Dask we would be able to publish `1!4.0.0` but `1!1.x`, `1!2.x` and `1!2020-2024.x` would be off limits.

Without changes along these lines I don't see projects making use of Python version epochs.
