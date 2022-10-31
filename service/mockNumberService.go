package service

import (
	httphelper "github.com/CiroLee/go-server/http-helper"
	"github.com/CiroLee/go-server/structs"
	"github.com/CiroLee/go-server/utils"
	"github.com/gin-gonic/gin"
)

type MockNumberService struct {
	structs.NumberReq
	Ctx *gin.Context
}

func (s *MockNumberService) NumGenerate() any {
	switch s.Type {
	case "int":
		r, err := utils.RandomInterger(s.Min, s.Max)
		if err != nil {
			httphelper.InvalidParamsRes(s.Ctx, err)
		}
		return r

	default:
		r, err := utils.RandomFloat(s.Min, s.Max, s.Demical)
		if err != nil {
			httphelper.InvalidParamsRes(s.Ctx, err)
		}
		return r
	}
}
