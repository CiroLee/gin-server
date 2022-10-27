package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// 生成随机整数
func RandomInterger(scope [2]int) (int, error) {
	rand.Seed(time.Now().UnixNano())
	min, max := scope[0], scope[1]
	if max <= min {
		return 0, fmt.Errorf("min must be less than max, min: %d, max:%d", scope[0], scope[1])
	}
	return rand.Intn(max-min) + min, nil
}

func RandomFloat32(scope [2]int) (float32, error) {
	rand.Seed(time.Now().UnixNano())
	min, max := scope[0], scope[1]
	if max <= min {
		return 0.0, fmt.Errorf("min must be less than max, min: %d, max:%d", scope[0], scope[1])
	}
	return rand.Float32() * (float32(max) - float32(min)), nil
}
