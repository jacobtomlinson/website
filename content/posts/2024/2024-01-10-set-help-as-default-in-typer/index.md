---
title: "How to get typer to show help by default"
date: 2024-01-10T00:00:00+00:00
draft: false
author: "Jacob Tomlinson"
categories:
  - blog
tags:
  - typer
  - python
  - cli
  - quick tips
---

I love using [typer](https://typer.tiangolo.com/) for creating CLI tools in Python. It makes creating complex trees of subcommands really straightforward.

However, I often don't want the command on its own to have functionality. Instead I just want it to print out the help text, like if I run `git` or `docker` on their own I expect them to print out their help text.

```console
$ git
These are common Git commands used in various situations:

start a working area (see also: git help tutorial)
   clone     Clone a repository into a new directory
   init      Create an empty Git repository or reinitialize an existing one
...
```

This is not the default behaviour in typer so to set this we need to use the `no_args_is_help` flag.

```python
app = typer.Typer(no_args_is_help=True)
```

## Example

Here's an example from the docs with two subcommands.

```python
# main.py
import typer

app = typer.Typer()


@app.command()
def hello(name: str):
    "Say hello"
    print(f"Hello {name}")


@app.command()
def goodbye(name: str, formal: bool = False):
    "Say goodbye"
    if formal:
        print(f"Goodbye Ms. {name}. Have a good day.")
    else:
        print(f"Bye {name}!")


if __name__ == "__main__":
    app()

```

However, if I run `main.py` without specifying a subcommand I get an error message that isn't very useful.

```console
$ python main.py
Usage: main.py [OPTIONS] COMMAND [ARGS]...
Try 'main.py --help' for help.

Error: Missing command.
```

But if I set the `no_args_is_help` option when creating my `app` I get a more useful help.

```python {hl_lines=[4]}
# main.py
import typer

app = typer.Typer(no_args_is_help=True)


@app.command()
def hello(name: str):
    "Say hello"
    print(f"Hello {name}")


@app.command()
def goodbye(name: str, formal: bool = False):
    "Say goodbye"
    if formal:
        print(f"Goodbye Ms. {name}. Have a good day.")
    else:
        print(f"Bye {name}!")


if __name__ == "__main__":
    app()
```

```console
$ python main.py
Usage: main.py [OPTIONS] COMMAND [ARGS]...

Options:
  --install-completion [bash|zsh|fish|powershell|pwsh]
                                  Install completion for the specified shell.
  --show-completion [bash|zsh|fish|powershell|pwsh]
                                  Show completion for the specified shell, to
                                  copy it or customize the installation.
  --help                          Show this message and exit.

Commands:
  goodbye  Say goodbye
  hello    Say hello
```

Now I can easily see that I can use the `hello` or `goodbye` subcommand.
