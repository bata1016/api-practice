package main

import (
	"github.com/bata1016/api-practice/db"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Hello, World")
	})
	db.Init()
	router.Run()

	db.Close()
}
