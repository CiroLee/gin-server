package service

import (
	"github.com/CiroLee/go-server/structs"
	"github.com/gin-gonic/gin"
)

const PLACEHOLDER_IMAGE_DOMAIN = "https://dummyimage.com"

type MokcImageService struct {
	structs.ImageReq
	Ctx *gin.Context
}
