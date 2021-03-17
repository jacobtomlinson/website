---
title: "Awaitable Objects and Async Context Managers in Python"
date: 2021-03-17T00:00:00+00:00
draft: false
author: "Jacob Tomlinson"
categories:
  - blog
tags:
  - Python
  - AsyncIO
  - Tutorial
thumbnail: python
---

Python objects are synchronous by default. When working with `asyncio` if we create an object the `__init__` is a regular function and we cannot do any async work in here.

```python
import asyncio


class Hello:
    def __init__(self):
        print("init")
        # We cannot await anything in here

    async def a_method(self):
        print("a_method")
        # We can await in here


async def main():
    h = Hello()
    await h.a_method()


if __name__ == "__main__":
    loop = asyncio.get_event_loop()
    loop.run_until_complete(main())
```

Running this code will print:

```
init
a_method
```

## Awaitable objects

We can make our object directly awaitable by giving it a `__await__` method. This method [must return an iterator](https://docs.python.org/3/reference/datamodel.html?highlight=__await__#object.__await__).

When defining an async function the `__await__` method is created for us, so we can use an async closure and use the `__await__` method from that.

```python
import asyncio


class Hello:
    def __init__(self):
        print("init")
        # We cannot await anything in here

    def __await__(self):
        async def closure():
            print("await")
            # We can await in here
            return self

        return closure().__await__()

    async def a_method(self):
        print("a_method")
        # We can await in here


async def main():
    h = await Hello()
    await h.a_method()


if __name__ == "__main__":
    loop = asyncio.get_event_loop()
    loop.run_until_complete(main())
```

Here we've added the `__await__` method and updated our object creation to be `h = await Hello()`.

Running this code will print:

```
init
await
a_method
```

## Async context managers

We can also turn out object into an async context manager with `__aenter__` and `__aexit__` coroutines.

```python
import asyncio


class Hello:
    def __init__(self):
        print("init")

    def __await__(self):
        async def closure():
            print("await")
            return self

        return closure().__await__()

    async def __aenter__(self):
        print("enter")
        return self

    async def __aexit__(self, *args):
        print("exit")

    async def a_method(self):
        print("a_method")


async def main():
    async with Hello() as h:
        print("context")
        await h.a_method()


if __name__ == "__main__":
    loop = asyncio.get_event_loop()
    loop.run_until_complete(main())
```

Here we've added the context manager methods and updated our object creation to be an `async with` statement.

Running this code will print:

```
init
enter
context
a_method
exit
```

Notice that while our enter and exit coroutines were called as expected our object is never awaited.

We can fix this by awaiting it ourselves within the `__aenter__` method.

```python
import asyncio


class Hello:
    def __init__(self):
        print("init")

    def __await__(self):
        async def closure():
            print("await")
            return self

        return closure().__await__()

    async def __aenter__(self):
        print("enter")
        await self
        return self

    async def __aexit__(self, *args):
        print("exit")

    async def a_method(self):
        print("a_method")


async def main():
    async with Hello() as h:
        print("context")
        await h.a_method()


if __name__ == "__main__":
    loop = asyncio.get_event_loop()
    loop.run_until_complete(main())
```

Running this code will print:

```
init
enter
await
context
a_method
exit
```

## Wrap up

With our new context manager class the only place we cannot use async code is within the `__init__` method. Which is completely reasonable as we should only ever be setting up our object's attributes in there anyway.

Hopefully this article has given a quick overview on creating awaitable objects and async context managers and also shown the order of operations when using one.