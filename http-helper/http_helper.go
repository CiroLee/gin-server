package httphelper

import (
	"net/http"

	"github.com/CiroLee/go-server/constants/codes"
	"github.com/gin-gonic/gin"
)

func RequestError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"code": codes.ErrorCode["RequestError"].Code,
		"msg":  codes.ErrorCode["RequestError"].Msg,
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
