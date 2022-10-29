package service

import (
	"strings"

	"github.com/CiroLee/go-server/constants"
	"github.com/CiroLee/go-server/utils"

	"github.com/CiroLee/go-server/structs"
)

const baseNum = 10

var zhChars = constants.ZH_CHARACTERS
var alphabet = constants.ALPHABET

type baseProps struct {
	Min, Max, Len int
	Upper         bool
}

// service类
type MockTextsService struct {
	structs.TextsReq
}

/** service methods **/

func (t *MockTextsService) letter() string {
	var source []string
	if t.Language == "en" {
		source = alphabet
	} else {
		source = zhChars
	}
	index, _ := utils.RandomInterger(0, len(source))

	return source[index]
}

func (t *MockTextsService) word(p baseProps) string {
	var str string
	len := 1
	if p.Len > 0 {
		len = p.Len
	} else {
		n, err := utils.RandomInterger(p.Min, p.Max)
		if err == nil {
			len = n
		}
	}
	for i := 0; i < len; i++ {
		str += t.letter()
	}

	if p.Upper {
		return utils.Capitalize(str)
	}
	return str
}

func (t *MockTextsService) sentence(p baseProps) string {
	var str string
	len := 1
	if p.Len > 0 {
		len = p.Len
	} else {
		n, err := utils.RandomInterger(p.Min, p.Max)
		if err == nil {
			len = n
		}
	}
	trailCode := "，"
	if t.Language == "en" {
		trailCode = " "
	}
	for i := 0; i < len; i++ {
		str += t.word(baseProps{
			Min: 1,
			Max: baseNum,
		}) + trailCode
	}

	if p.Upper {
		return utils.Capitalize(strings.TrimRight(str, trailCode))
	}
	return strings.TrimRight(str, trailCode)
}

// 对外方法

func (t *MockTextsService) TextGenerate() string {

	switch t.Type {
	case "word":
		return t.word(baseProps{
			Min:   t.Min,
			Max:   t.Max,
			Len:   t.Len,
			Upper: t.Upper,
		})
	case "sentence":
		return t.sentence(baseProps{
			Min:   t.Min,
			Max:   t.Max,
			Len:   t.Len,
			Upper: t.Upper,
		})
	default:
		return ""
	}

}
