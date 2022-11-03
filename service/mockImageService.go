package service

import (
	"fmt"
	"strings"

	httphelper "github.com/CiroLee/go-server/http-helper"
	"github.com/CiroLee/go-server/structs"
	"github.com/CiroLee/go-server/utils"
	"github.com/gin-gonic/gin"
)

const PLACEHOLDER_IMAGE_DOMAIN = "https://dummyimage.com"
const PICURM_IMAGE_DOMAIN = "https://picsum.photos"

type MockImageService struct {
	structs.ImageReq
	Ctx *gin.Context
}

func initSize(w, h int) (int, int) {
	var width, _ = utils.RandomInterger(320, 1025)
	var height, _ = utils.RandomInterger(320, 1025)
	if w != 0 && h == 0 {
		width = w
		height = w
	} else if w == 0 && h != 0 {
		width = h
		height = h
	} else if w != 0 && h != 0 {
		width = w
		height = h
	}

	return width, height
}

// note: hex不包含#号
func (m *MockImageService) isDark(hex string) bool {
	rgbColor, err := utils.HexToRgb(hex)
	if err != nil {
		httphelper.RequestError(m.Ctx, err)
		// TODO 优化错误统一处理
		panic(err.Error())
	}
	var float64RgbColor = make([]float64, 3)
	for i, v := range rgbColor {
		float64RgbColor[i] = float64(v)
	}
	hslColor := utils.RgbToHsl(float64RgbColor)
	return hslColor[2] < 50
}

func (m *MockImageService) randomHex() (hex string, dark bool) {
	mockColorService := MockColorService{
		ColorReq: structs.ColorReq{
			Type: "hex",
		},
	}
	// 生成的标准hex包含#号
	hexColor := mockColorService.ColorGenerate()
	hexColor = removeHexPreffix(hexColor)
	isDarkMode := m.isDark(hexColor)

	return hexColor, isDarkMode
}

func removeHexPreffix(hex string) string {
	if strings.HasPrefix(hex, "#") {
		return string(hex[1:])
	}
	return hex
}

func (m *MockImageService) placeholder() string {
	var isDarkMode bool
	width, height := initSize(m.Width, m.Height)
	bgColor := removeHexPreffix(m.BgColor)
	fontColor := removeHexPreffix(m.FontColor)

	if bgColor == "" {
		bgColor, isDarkMode = m.randomHex()
	} else {
		isDarkMode = m.isDark(bgColor)
	}

	if fontColor == "" {
		if isDarkMode {
			fontColor = "fff"
		} else {
			fontColor = "333"
		}
	}

	var imgUrl = fmt.Sprintf("%v/%vx%v/%v", PLACEHOLDER_IMAGE_DOMAIN, width, height, bgColor)

	if m.Text != "" {
		return fmt.Sprintf("%s/%s&text=%s", imgUrl, fontColor, m.Text)
	}
	return imgUrl

}

func (m *MockImageService) picsum() string {
	width, height := initSize(m.Width, m.Height)
	imgUrl := fmt.Sprintf("%v/%v/%v", PICURM_IMAGE_DOMAIN, width, height)
	if m.Grayscale {
		imgUrl += "?grayscale"
	}
	if m.Blur > 0 {
		if strings.HasSuffix(imgUrl, "?grayscale") {
			return fmt.Sprintf("%v&blur=%v", imgUrl, m.Blur)
		} else {
			return fmt.Sprintf("%v?blur=%v", imgUrl, m.Blur)
		}
	}

	return imgUrl
}

func (m *MockImageService) ImageGenerate() string {
	switch m.Type {
	case "placeholder":
		return m.placeholder()
	default:
		return m.picsum()
	}
}
