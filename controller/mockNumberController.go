package controller

import (
	"net/http"
	"strconv"

	"github.com/CiroLee/go-server/constants/codes"
	"github.com/CiroLee/go-server/utils"
	"github.com/gin-gonic/gin"
)

type NumberReq struct {
	Type string `form:"type" binding:"oneof=int float"`
	Min  int    `form:"min" binding:"required,gte=-999999"`
	Max  int    `form:"max" binding:"required,lte=999999"`
}

// 生成随机数字 包括: 整数，浮点数
func MockNumberController(ctx *gin.Context) {
	var q NumberReq
	if err := ctx.ShouldBindQuery(&q); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": codes.ErrorCode["missParams"].Code,
			"msg":  codes.ErrorCode["missParams"].Msg,
			"data": err.Error(),
		})
		return
	}

	var (
		min, _ = strconv.Atoi(ctx.Query("min"))
		max, _ = strconv.Atoi(ctx.Query("max"))
	)

	r, err := utils.RandomInterger([2]int{min, max})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": codes.ErrorCode["invalidParams"].Code,
			"msg":  codes.ErrorCode["invalidParams"].Msg,
			"data": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": r,
	})
}
