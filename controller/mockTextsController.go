package controller

import (
	httphelper "github.com/CiroLee/go-server/http-helper"
	"github.com/CiroLee/go-server/service"
	"github.com/CiroLee/go-server/structs"
	"github.com/gin-gonic/gin"
)

func MockTextsController(ctx *gin.Context) {
	var q structs.TextsReq
	if err := ctx.ShouldBindQuery(&q); err != nil {
		httphelper.RequestError(ctx, err)
		return
	}

	mockTextsservice := service.MockTextsService{TextsReq: q}
	if q.Num > 1 {
		var r []string
		num := q.Num
		if q.Num > 100 {
			num = 100
		}
		for i := 0; i < num; i++ {
			r = append(r, mockTextsservice.TextGenerate())
		}
		httphelper.SuccessRes(ctx, r)
	} else {
		httphelper.SuccessRes(ctx, mockTextsservice.TextGenerate())
	}

}
