package main

import (
	"fmt"
	"net/http"

	"github.com/CiroLee/go-web-gin/constants/codes"
	"github.com/gin-gonic/gin"
)

func testHandler(c *gin.Context) {
	query := c.Query("name")
	if query == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": codes.ErrorCode["missParams"].Code,
			"msg":  codes.ErrorCode["missParams"].Msg,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "success",
			"data": fmt.Sprintf("this is a %s", query),
		})
	}
}

func SetUpRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/test", testHandler)

	return r
}
