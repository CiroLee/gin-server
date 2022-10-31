package service

import (
	"fmt"
	"time"

	httphelper "github.com/CiroLee/go-server/http-helper"
	"github.com/CiroLee/go-server/structs"
	"github.com/CiroLee/go-server/utils"
	"github.com/gin-gonic/gin"
)

type MockDateService struct {
	structs.DateReq
	Ctx *gin.Context
}

const layout = "2006-01-02 15:04:05"
const timeStart int = 0
const timeEnd int = 32502009599 // 2999-12-12 23:59:59

func (d *MockDateService) date() (string, error) {
	var from, to = timeStart, timeEnd
	if d.From >= 0 {
		from = d.From
	}
	if d.To < timeEnd {
		to = d.To
	}

	n, err := utils.RandomInterger(from, to)
	fmt.Printf("n: %v\n", n)
	fmt.Printf("from: %v\n", from)
	fmt.Printf("to: %v\n", d.To)
	if err != nil {
		return "", err
	}

	r := time.Unix(int64(n), 0).Format(layout)

	return r, nil
}

func (d *MockDateService) DateGenerate() string {
	switch d.Type {
	case "date":
		r, err := d.date()
		if err != nil {
			httphelper.InvalidParamsRes(d.Ctx, err)
		}
		return r
	default:
		return ""
	}
}
