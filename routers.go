package main

import (
	"github.com/CiroLee/go-server/controller"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/test", controller.TestController)
	mock := r.Group("/mock")
	{
		mock.GET("/number", controller.MockNumberController)
		mock.GET("/texts", controller.MockTextsController)
	}
	return r
}
