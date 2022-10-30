package controller

import (
	"net/http"

	"github.com/CiroLee/go-server/constants/codes"
	"github.com/CiroLee/go-server/service"
	"github.com/CiroLee/go-server/structs"
	"github.com/gin-gonic/gin"
)

func MockTextsController(ctx *gin.Context) {
	var q structs.TextsReq
	if err := ctx.ShouldBindQuery(&q); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": codes.ErrorCode["invalidParams"].Code,
			"msg":  codes.ErrorCode["invalidParams"].Msg,
			"data": err.Error(),
		})
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
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "success",
			"data": r,
		})
	} else {
		r := mockTextsservice.TextGenerate()
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "success",
			"data": r,
		})
	}

}
