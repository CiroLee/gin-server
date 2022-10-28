package service

import (
	"github.com/CiroLee/go-server/structs"
	"github.com/CiroLee/go-server/utils"
)

type MockNumberService struct {
	structs.NumberReq
}

func (s *MockNumberService) NumGenerate() (any, error) {
	switch s.Type {
	case "int":
		return utils.RandomInterger(s.Min, s.Max)
	default:
		return utils.RandomFloat(s.Min, s.Max, s.Demical)
	}
}
