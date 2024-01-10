---
title: "How highlight lines in a Hugo code block"
date: 2024-01-11T00:00:00+00:00
draft: false
author: "Jacob Tomlinson"
categories:
  - blog
tags:
  - hugo
  - syntax
  - static site generators
  - blogging
  - quick tips
---

Sometimes when writing some code in a blog post I want to emphasize a couple of lines in particular. Today I found out that Hugo has some really nice syntax to do this in a regular codeblock.

Here I am emphasizing line `5` that contains the the `print` statement.

```python {hl_lines=[5]}
import datetime

def main():
    current_time = datetime.datetime.now().strftime("%H:%M")
    print(f"Hello at {current_time}!")

if __name__ == "__main__":
    main()
```

To do this I just include a little extra information as part of my triple-backtick code-fence.

````markdown
```python {hl_lines=[5]}
import datetime

def main():
    current_time = datetime.datetime.now().strftime("%H:%M")
    print(f"Hello at {current_time}!")

if __name__ == "__main__":
    main()
```
````

To learn more check out the [Hugo syntax highlighting documentation](https://gohugo.io/content-management/syntax-highlighting/#highlighting-in-code-fences).
