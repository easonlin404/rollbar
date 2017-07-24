package main

import (
	"github.com/easonlin404/rollbar"
	"github.com/gin-gonic/gin"
)

func main() {
	token := "POST_SERVER_ITEM_ACCESS_TOKEN"
	environment := "development"

	r := gin.Default()
	r.Use(rollbar.Recovery(token, environment))

	r.Run(":8080")
}
