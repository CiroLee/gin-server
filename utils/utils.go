package utils

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

// 生成随机整数
func RandomInterger(min, max int) (int, error) {
	rand.Seed(time.Now().UnixNano())
	if max <= min {
		return 0, fmt.Errorf("min must be less than max, min: %d, max:%d", min, max)
	}
	return rand.Intn(max-min) + min, nil
}

func RandomFloat(min, max int, demical int) (float64, error) {
	if max <= min {
		return 0.0, fmt.Errorf("min must be less than max, min: %d, max:%d", min, max)
	}
	d := math.Pow10(4)
	fmt.Printf("demical: %v\n", demical)
	if demical > 0 {
		d = math.Pow10(demical)
	}
	rand.Seed(time.Now().UnixNano())
	fmt.Printf("d: %v\n", d)
	num := rand.Float64() * (float64(max) - float64(min))
	temp := strconv.FormatFloat(math.Trunc(num*d)/d, 'f', -1, 64)
	num, err := strconv.ParseFloat(temp, 64)
	if err != nil {
		return 0.0, err
	}
	return num, nil
}
