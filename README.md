# rollbar

[![Build Status](https://travis-ci.org/easonlin404/rollbar.svg)](https://travis-ci.org/easonlin404/rollbar)
[![codecov](https://codecov.io/gh/easonlin404/rollbar/branch/master/graph/badge.svg)](https://codecov.io/gh/easonlin404/rollbar)
[![Go Report Card](https://goreportcard.com/badge/github.com/easonlin404/rollbar)](https://goreportcard.com/report/github.com/easonlin404/rollbar)
[![GoDoc](https://godoc.org/github.com/easonlin404/rollbar?status.svg)](https://godoc.org/github.com/easonlin404/rollbar)

gin rollbar middleware

> Package pprof serves via its HTTP server runtime profiling data in the format expected by the pprof visualization tool.

## Usage

### Start using it

Download and install it:

```bash
$ go get github.com/easonlin404/rollbar
```

Import it in your code:

```go
import "github.com/easonlin404/rollbar"
```

### Example:

```go
package main

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()
  pprof.Register(router, nil)
  router.Run(":8080")
}
```

### change default path prefix:

```go
func main() {
	router := gin.Default()
	pprof.Register(router, &pprof.Options{
		// default is "debug/pprof"
		RoutePrefix: "debug/pprof",
	})
	router.Run(":8080")
}
```

### Use the pprof tool

Then use the pprof tool to look at the heap profile:

```bash
go tool pprof http://localhost:8080/debug/pprof/heap
```

Or to look at a 30-second CPU profile:

```bash
go tool pprof http://localhost:8080/debug/pprof/profile
```

Or to look at the goroutine blocking profile, after calling runtime.SetBlockProfileRate in your program:

```bash
go tool pprof http://localhost:8080/debug/pprof/block
```

Or to collect a 5-second execution trace:

```bash
wget http://localhost:8080/debug/pprof/trace?seconds=5
```
