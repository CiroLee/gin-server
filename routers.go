package main

import (
	"github.com/CiroLee/go-web-gin/controller"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/test", controller.TestController)

	return r
}
