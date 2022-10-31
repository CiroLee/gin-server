package httphelper

import (
	"net/http"

	"github.com/CiroLee/go-server/constants/codes"
	"github.com/gin-gonic/gin"
)

func InvalidParamsRes(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"code": codes.ErrorCode["invalidParams"].Code,
		"msg":  codes.ErrorCode["invalidParams"].Msg,
		"data": err.Error(),
	})
}

func SuccessRes(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": data,
	})
}
