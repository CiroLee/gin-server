package controller

import (
	httphelper "github.com/CiroLee/go-server/http-helper"
	"github.com/CiroLee/go-server/service"
	"github.com/CiroLee/go-server/structs"
	"github.com/gin-gonic/gin"
)

func MockImageController(ctx *gin.Context) {
	var q structs.ImageReq
	if err := ctx.ShouldBindQuery(&q); err != nil {
		httphelper.RequestError(ctx, err)
		return
	}

	mockImageService := service.MockImageService{
		ImageReq: q,
		Ctx:      ctx,
	}

	if q.Num > 1 {
		var r = make([]string, 0)
		num := q.Num
		if q.Num > 100 {
			num = 100
		}
		for i := 0; i < num; i++ {
			img := mockImageService.ImageGenerate()
			r = append(r, img)
		}
		httphelper.SuccessRes(ctx, r)
	} else {
		httphelper.SuccessRes(ctx, mockImageService.ImageGenerate())
	}
}
