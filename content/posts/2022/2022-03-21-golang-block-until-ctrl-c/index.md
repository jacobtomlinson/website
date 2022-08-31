---
title: "Golang block until interrupt with ctrl+c"
date: 2022-03-21T00:00:00+00:00
draft: false
author: "Jacob Tomlinson"
categories:
  - blog
tags:
  - Golang
  - Snippet
---

Today I found myself needing a Go application's main thread to stop and wait until the user wants it to exit with a `ctrl+c` keyboard interrupt.

To do this we can use the `os/signal` package to put signals into a channel, and then block until that channel receives a signal.

```go
import (
  "fmt"
  "os"
  "os/signal"
  "syscall"
)

func init() {
  done := make(chan os.Signal, 1)
  signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
  fmt.Println("Blocking, press ctrl+c to continue...")
  <-done  // Will block here until user hits ctrl+c
}
```
