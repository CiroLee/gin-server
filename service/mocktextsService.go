package service

import (
	"github.com/CiroLee/go-server/constants"
	"github.com/CiroLee/go-server/utils"

	"github.com/CiroLee/go-server/structs"
)

var zhChars = constants.ZH_CHARACTERS
var alphabet = constants.ALPHABET

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

func (t *MockTextsService) Word() string {
	var str string
	length := 1
	if t.Len > 0 {
		length = t.Len
	} else {
		n, err := utils.RandomInterger(t.Min, t.Max)
		if err != nil {
			length = n
		}
	}

	for i := 0; i < length; i++ {
		str += t.letter()
	}

	return str
}

// 对外方法

func (t *MockTextsService) TextGenerate() string {

	if t.Type == "word" {
		return t.Word()
	}
	return ""
}
