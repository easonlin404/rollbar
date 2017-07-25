# rollbar

[![Build Status](https://travis-ci.org/easonlin404/rollbar.svg)](https://travis-ci.org/easonlin404/rollbar)
[![codecov](https://codecov.io/gh/easonlin404/rollbar/branch/master/graph/badge.svg)](https://codecov.io/gh/easonlin404/rollbar)
[![Go Report Card](https://goreportcard.com/badge/github.com/easonlin404/rollbar)](https://goreportcard.com/report/github.com/easonlin404/rollbar)
[![GoDoc](https://godoc.org/github.com/easonlin404/rollbar?status.svg)](https://godoc.org/github.com/easonlin404/rollbar)

Middleware to integrate with [rollbar](https://rollbar.com/) error monitoring. It uses [roll](https://github.com/stvp/roll) client for Go that reports errors and logs messages.

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

### Example

```go
package main

import (
	"github.com/easonlin404/rollbar"
	"github.com/gin-gonic/gin"
	"github.com/stvp/roll"
)

func main() {
	roll.Token = "POST_SERVER_ITEM_ACCESS_TOKEN"
	//roll.Environment = "production" // defaults to "development"

	r := gin.Default()
	r.Use(rollbar.Recovery(true))

	r.GET("/", func(*gin.Context) {
		panic("failed")
	})
	r.Run(":8080")
}
```