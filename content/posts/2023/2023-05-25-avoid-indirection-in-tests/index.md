---
title: "Avoid indirection in tests at all costs"
date: 2023-05-25T00:00:00+00:00
draft: false
author: "Jacob Tomlinson"
categories:
  - blog
tags:
  - work
  - coding
  - testing
---

When writing tests the balance between avoiding indirection and DRY-ness should be much more weighted towards avoiding indirection than in the code it is testing.

I regularly find myself pointing folks to [this blog post by Matt Rocklin about avoiding indirection in your code](https://matthewrocklin.com/blog/work/2019/06/23/avoid-indirection) which extends [this post on writing dumb code](https://matthewrocklin.com/blog/work/2018/01/27/write-dumb-code). Both are excellent reads.

The general thesis of those posts is that your code should be readable by novice programmers. Being clever makes code hard to debug and maintain in the future.

I regularly see folks try and [DRY](https://en.wikipedia.org/wiki/Don%27t_repeat_yourself) out their code as much as possible. Matt's post suggests that there is a balance to be found between avoiding repetition and reducing how much future developers need to jump around a code base when debugging some future problem.

I'd like to extend this concept and say that **tests should avoid indirection at all costs**.

When writing many similar tests it is tempting to use lots of fixtures and utility functions to reduce repetition. However, when debugging a test and trying to decide whether your changes have broken the test or the test itself is broken you need simplicity.

The last thing you want to see is something that looks like this:

```python
import pytest
from library import Thing

@pytest.fixture
def some_thing():
    thing = Thing.create(*args, **kwargs)
    yield thing
    thing.delete()

# Insert 10s-100s of lines of fixtures here

def assert_some_thing_works(thing):
    assert expected_stuff in thing
    return True

# Insert 10s-100s of lines of helper functions here

# Insert 10s-100s of lines of other tests here

def test_some_thing_works(some_thing):
    assert assert_some_thing_works(some_thing)
```

This is an extreme example but now imagine it with multiple fixtures, multiple helpers and maybe a couple of extra lines in the test that makes the test unique. I see this pattern often and it means you constantly have to jump back and forth between the test, the fixtures and the helpers to figure out what on earth is going on.

The actual test itself is all but hollowed out, it is just a shell that glues together fixtures and helpers.

Here are a few suggestions for writing tests that are pleasant to debug:

- Forget about them being DRY
- Try and keep all executed lines within the test function
- Only use fixtures when they are very necessary
- Keep tests short but free from indirection
- The whole test should fit on your screen

I would much rather see 100 tests that are all slight variations with duplication. Because ultimately when debugging I only want to read and run one of them.

{{< highlight python "hl_lines=8-11" >}}
import pytest

def test_some_thing_works():
    thing = Thing.create(*args, **kwargs)
    assert expected_stuff in thing
    thing.delete()

def test_some_thing_doesnt_work():  # ❌ Failing test
    thing = Thing.create(*args, **bad_kwargs)
    assert expected_stuff not in thing
    thing.delete()

def test_some_method_of_some_thing():
    thing = Thing.create(*args, **bad_kwargs)
    assert thing.some_method() == something
    thing.delete()
{{< /highlight >}}
