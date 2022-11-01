package service

import (
	"strconv"
	"time"

	"github.com/CiroLee/go-server/constants"
	httphelper "github.com/CiroLee/go-server/http-helper"
	"github.com/CiroLee/go-server/structs"
	"github.com/CiroLee/go-server/utils"
	"github.com/gin-gonic/gin"
)

type MockDateService struct {
	structs.DateReq
	Ctx *gin.Context
}

var (
	weekDict  = constants.WEEKDAY_DICT
	monthDict = constants.MONTH_DICT
)

const (
	layout        = "2006-01-02 15:04:05"
	timeStart int = 0
	timeEnd   int = 32502009599 // 2999-12-12 23:59:59
)

func (d *MockDateService) date() (string, error) {

	n, err := d.timestamp()
	if err != nil {
		return "", err
	}

	r := time.Unix(int64(n), 0).Format(layout)

	return r, nil
}

func (d *MockDateService) timestamp() (int, error) {
	var from, to = timeStart, timeEnd
	if d.From >= 0 {
		from = d.From
	}
	if d.To < timeEnd {
		to = d.To
	}

	n, err := utils.RandomInterger(from, to)
	if err != nil {
		return 0, err
	}

	return n, nil
}

func (d *MockDateService) weekday() string {
	item := utils.RandomELement(weekDict)
	if d.Language == "cn" {
		return item["cn"].(string)
	}
	if d.Abbr {
		return item["en"].([2]string)[1]
	}
	return item["en"].([2]string)[0]
}

func (d *MockDateService) month() string {
	item := utils.RandomELement(monthDict)
	if d.Language == "cn" {
		return item["cn"].(string)
	}
	if d.Abbr {
		return item["en"].([2]string)[1]
	}
	return item["en"].([2]string)[0]
}

func (d *MockDateService) DateGenerate() string {
	switch d.Type {
	case "date":
		r, err := d.date()
		if err != nil {
			httphelper.RequestError(d.Ctx, err)
		}
		return r
	case "timestamp":
		r, err := d.timestamp()
		if err != nil {
			httphelper.RequestError(d.Ctx, err)
		}
		return strconv.Itoa(r)
	case "weekday":

		return d.weekday()
	case "month":
		return d.month()
	default:
		return ""
	}
}
