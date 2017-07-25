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