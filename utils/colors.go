package utils

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func HexToRgb(hex string) ([]uint, error) {
	var hexReg = regexp.MustCompile("^([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$")
	hexLower := strings.ToLower(hex)
	if !hexReg.MatchString(hexLower) {
		return nil, fmt.Errorf("hexToRgb: invalid hex color %v", hex)
	}
	if len(hexLower) == 3 {
		hexLower += hexLower
	}

	var rgb = make([]uint, 0)
	for i := 0; i < 6; i += 2 {
		part := hexLower[i : i+2]
		c1, err := strconv.ParseUint(part, 16, 32)
		if err != nil {
			panic(err)
		}
		rgb = append(rgb, uint(c1))
	}

	return rgb, nil
}

func RgbToHsl(rgb []float64) []uint {
	r := float64(rgb[0] / 255)
	g := float64(rgb[1] / 255)
	b := float64(rgb[2] / 255)
	max := Max([]float64{r, g, b})
	min := Min([]float64{r, g, b})
	mid := (max + min) / 2

	var s float64
	h, l := mid, mid
	if min == max {
		h = 0.0
	} else {
		var delta = max - min
		if l > 0.5 {
			s = delta / (2.0 - max - min)
		} else {
			s = delta / (min + max)
		}
		switch max {
		case r:
			var temp float64
			if g < b {
				temp = 6.0
			} else {
				temp = 0.0
			}
			h = (g-b)/delta + temp
		case g:
			h = (b-r)/delta + 2
		case b:
			h = (r-g)/delta + 4
		}
	}

	h *= 60

	return []uint{
		uint(math.Round(h)),
		uint(math.Round(s * 100)),
		uint(math.Round(l * 100)),
	}
}
