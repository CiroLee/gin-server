package controller

import (
	httphelper "github.com/CiroLee/go-server/http-helper"
	"github.com/CiroLee/go-server/service"
	"github.com/CiroLee/go-server/structs"
	"github.com/gin-gonic/gin"
)

func MockColorController(ctx *gin.Context) {
	var q structs.ColorReq
	if err := ctx.ShouldBindQuery(&q); err != nil {
		httphelper.RequestError(ctx, err)
		return
	}

	mockColorService := service.MockColorService{
		ColorReq: q,
		Ctx:      ctx,
	}
	if q.Num > 1 {
		var r []string
		num := q.Num
		if q.Num > 100 {
			num = 100
		}
		for i := 0; i < num; i++ {
			colors := mockColorService.ColorGenerate()
			r = append(r, colors)
		}
		httphelper.SuccessRes(ctx, r)
	} else {
		httphelper.SuccessRes(ctx, mockColorService.ColorGenerate())
	}
}
