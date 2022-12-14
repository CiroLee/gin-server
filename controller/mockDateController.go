package controller

import (
	httphelper "github.com/CiroLee/go-server/http-helper"
	"github.com/CiroLee/go-server/service"
	"github.com/CiroLee/go-server/structs"
	"github.com/gin-gonic/gin"
)

func MockDateController(ctx *gin.Context) {
	var q structs.DateReq
	if err := ctx.ShouldBindQuery(&q); err != nil {
		httphelper.RequestError(ctx, err)
		return
	}

	mockDateService := service.MockDateService{
		DateReq: q,
		Ctx:     ctx,
	}

	if q.Num > 1 {
		var r []string
		num := q.Num
		if q.Num > 100 {
			num = 100
		}
		for i := 0; i < num; i++ {
			date := mockDateService.DateGenerate()
			r = append(r, date)
		}
		httphelper.SuccessRes(ctx, r)
	} else {
		httphelper.SuccessRes(ctx, mockDateService.DateGenerate())
	}

}
