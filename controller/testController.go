package controller

import (
	"fmt"
	"net/http"

	"github.com/CiroLee/go-web-gin/constants/codes"
	"github.com/gin-gonic/gin"
)

type TestReq struct {
	Name   string `form:"name" binding:"required"`
	Gender string `form:"gender" binding:"oneof=male female"`
}

func TestController(c *gin.Context) {
	var q TestReq
	if err := c.ShouldBindQuery(&q); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": codes.ErrorCode["missParams"].Code,
			"msg":  codes.ErrorCode["missParams"].Msg,
			"data": err.Error(),
		})
		return
	}
	var helloGender string
	if c.Query("gender") == "male" {
		helloGender = "Mr."
	} else {
		helloGender = "Miss."
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": fmt.Sprintf("Hello, %s %s", helloGender, c.Query("name")),
	})

}
