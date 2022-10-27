package controller

import (
	"fmt"
	"net/http"

	"github.com/CiroLee/go-server/constants/codes"
	"github.com/gin-gonic/gin"
)

type TestReq struct {
	Name   string `form:"name" binding:"required"`
	Gender string `form:"gender" binding:"oneof=male female"`
}

func TestController(ctx *gin.Context) {
	var q TestReq
	if err := ctx.ShouldBindQuery(&q); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": codes.ErrorCode["missParams"].Code,
			"msg":  codes.ErrorCode["missParams"].Msg,
			"data": err.Error(),
		})
		return
	}
	var helloGender string
	if ctx.Query("gender") == "male" {
		helloGender = "Mr."
	} else {
		helloGender = "Miss."
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": fmt.Sprintf("Hello, %s %s", helloGender, ctx.Query("name")),
	})

}
