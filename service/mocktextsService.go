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
var zhNames = constants.COMMON_ZH_NAMES
var enNames = constants.COMMON_EN_NAMES

type baseProps struct {
	Min, Max, Len int
	Upper         bool
}

// service类
type MockTextsService struct {
	structs.TextsReq
}

/** service methods **/

func initLength(l, min, max int) int {
	var len = 1
	if l > 0 {
		len = l
	} else {
		n, err := utils.RandomInterger(min, max)
		if err == nil {
			len = n
		}
	}
	return len
}

// PRIVATE return a random letter
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

func (t *MockTextsService) name() string {
	var source []string
	if t.Language == "en" {
		source = enNames
	} else {
		source = zhNames
	}
	index, _ := utils.RandomInterger(0, len(source))

	if t.Upper {
		return utils.Capitalize(source[index])
	}
	return source[index]

}

func (t *MockTextsService) str() string {
	var str string
	s := []rune(constants.STRINGS)
	length := initLength(t.Len, t.Min, t.Max)

	for i := 0; i < length; i++ {
		index, _ := utils.RandomInterger(0, len(s))
		str += string(s[index])
	}
	return str
}

// return a random word, support Chinese and English
func (t *MockTextsService) word(p baseProps) string {
	var str string
	length := initLength(t.Len, t.Min, t.Max)
	for i := 0; i < length; i++ {
		str += t.letter()
	}

	if p.Upper {
		return utils.Capitalize(str)
	}
	return str
}

// return a random a sentence
func (t *MockTextsService) sentence(p baseProps) string {
	var str string
	length := initLength(t.Len, t.Min, t.Max)
	trailCode := "，"
	if t.Language == "en" {
		trailCode = " "
	}
	for i := 0; i < length; i++ {
		str += t.word(baseProps{
			Min: 1,
			Max: baseNum,
		}) + trailCode
	}

	return utils.Capitalize(strings.TrimRight(str, trailCode))

}

// return a random paragraph
func (t *MockTextsService) paragraph(p baseProps) string {
	var str string
	length := initLength(t.Len, t.Min, t.Max)

	trailCode := "。"
	if t.Language == "en" {
		trailCode = "."
	}

	for i := 0; i < length; i++ {
		str += t.sentence(baseProps{
			Min: 1,
			Max: baseNum,
		}) + trailCode
	}
	return utils.Capitalize(strings.TrimRight(str, trailCode))
}

// PUBLIC methods

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
			Min: t.Min,
			Max: t.Max,
			Len: t.Len,
		})
	case "paragraph":
		return t.paragraph(baseProps{
			Min: t.Min,
			Max: t.Max,
			Len: t.Len,
		})
	case "name":
		return t.name()
	case "string":
		return t.str()
	default:
		return ""
	}

}
