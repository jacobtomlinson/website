---
title: "How to highlight lines in a Hugo code block"
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

Sometimes when writing code in a blog post I want to emphasize a couple of lines in particular. Today I found out that Hugo has really nice syntax to do this in a regular markdown code-fence. 

```info
I prefer to use [code-fences](https://docs.github.com/en/get-started/writing-on-github/working-with-advanced-formatting/creating-and-highlighting-code-blocks#syntax-highlighting) over the [`highlight` shortcode](https://gohugo.io/content-management/shortcodes/#highlight) for code blocks because I get syntax highlighting of the code within the code-fence in my editor.
```

Here I am emphasizing line `5` that contains the the `print` statement.

```python {hl_lines=[5]}
import datetime

def main():
    current_time = datetime.datetime.now().strftime("%H:%M")
    print(f"Hello at {current_time}!")

if __name__ == "__main__":
    main()
```

To do this I included a little extra information as part of my triple-backtick code-fence.

````python {hl_lines=[1,10]}
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
