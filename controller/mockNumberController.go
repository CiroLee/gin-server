package controller

import (
	"net/http"

	"github.com/CiroLee/go-server/constants/codes"
	"github.com/CiroLee/go-server/service"
	"github.com/CiroLee/go-server/structs"
	"github.com/gin-gonic/gin"
)

// 生成随机数字 包括: 整数，浮点数
func MockNumberController(ctx *gin.Context) {
	var q structs.NumberReq
	if err := ctx.ShouldBindQuery(&q); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": codes.ErrorCode["invalidParams"].Code,
			"msg":  codes.ErrorCode["invalidParams"].Msg,
			"data": err.Error(),
		})
		return
	}

	mockNumberService := service.MockNumberService{NumberReq: q}
	r, err := mockNumberService.Generate()

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
