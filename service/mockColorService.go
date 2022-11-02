package service

import (
	"fmt"
	"strconv"

	"github.com/CiroLee/go-server/structs"
	"github.com/CiroLee/go-server/utils"
	"github.com/gin-gonic/gin"
)

type MockColorService struct {
	structs.ColorReq
	Ctx *gin.Context
}

func generateColorSlice(range1, range2, range3 [2]int) [3]int {
	var s [3]int
	s[0], _ = utils.RandomInterger(range1[0], range1[1])
	s[1], _ = utils.RandomInterger(range2[0], range2[1])
	s[2], _ = utils.RandomInterger(range3[0], range3[1])

	return s
}

func (c *MockColorService) hex() string {
	var _hex = "#"
	for i := 0; i < 6; i++ {
		n, _ := utils.RandomFloat(0, 1, 6)
		_hex += strconv.FormatInt(int64(n*16), 16)
	}
	return _hex
}

func (c *MockColorService) rgb() string {
	var s = generateColorSlice([2]int{0, 256}, [2]int{0, 256}, [2]int{0, 256})
	return fmt.Sprintf("rgb(%v, %v, %v)", s[0], s[1], s[2])
}

func (c *MockColorService) rgba() string {
	alpha, _ := utils.RandomFloat(0, 1, 2)
	s := generateColorSlice([2]int{0, 256}, [2]int{0, 256}, [2]int{0, 256})
	return fmt.Sprintf("rgba(%v, %v, %v, %v)", s[0], s[1], s[2], alpha)
}

func (c *MockColorService) hsl() string {
	s := generateColorSlice([2]int{0, 101}, [2]int{0, 361}, [2]int{0, 101})
	fmt.Printf("s: %v\n", s)
	return fmt.Sprintf("hsl(%v, %v%%, %v%%)", s[0], s[1], s[2])
}

func (c *MockColorService) ColorGenerate() string {
	switch c.Type {
	case "hex":
		return c.hex()
	case "rgb":
		return c.rgb()
	case "rgba":
		return c.rgba()
	case "hsl":
		return c.hsl()
	default:
		return ""
	}
}
