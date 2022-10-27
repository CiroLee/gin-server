package main

import (
	"github.com/CiroLee/go-server/controller"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/test", controller.TestController)
	r.GET("/mock/number", controller.MockNumberController)
	return r
}
