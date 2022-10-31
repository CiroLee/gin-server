package controller

import (
	httphelper "github.com/CiroLee/go-server/http-helper"
	"github.com/CiroLee/go-server/service"
	"github.com/CiroLee/go-server/structs"
	"github.com/gin-gonic/gin"
)

// 生成随机数字 包括: 整数，浮点数
func MockNumberController(ctx *gin.Context) {
	var q structs.NumberReq
	if err := ctx.ShouldBindQuery(&q); err != nil {
		httphelper.InvalidParamsRes(ctx, err)
		return
	}

	mockNumberService := service.MockNumberService{NumberReq: q, Ctx: ctx}

	if q.Num > 1 {
		var r []any
		num := q.Num
		if q.Num > 100 {
			num = 100
		}

		for i := 0; i < num; i++ {
			r = append(r, mockNumberService.NumGenerate())
		}

		httphelper.SuccessRes(ctx, r)
	} else {
		httphelper.SuccessRes(ctx, mockNumberService.NumGenerate())
	}

}
