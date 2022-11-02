package utils

import (
	"strings"

	"golang.org/x/exp/constraints"
)

func Capitalize(s string) string {
	var r []rune
	var rs = []rune(s)
	for i, v := range rs {
		if i == 0 {
			f := strings.ToUpper(string(v))
			r = append(r, []rune(f)[0])
		} else {
			r = append(r, v)
		}
	}

	return string(r)
}

func RandomELement[T any](slice []T) T {
	index, _ := RandomInterger(0, len(slice))
	return slice[index]
}

func Max[T constraints.Ordered](s []T) T {
	var max T
	if len(s) == 0 {
		return max
	}
	max = s[0]
	for i := 1; i < len(s); i++ {
		item := s[i]
		if item > max {
			max = item
		}
	}

	return max
}

func Min[T constraints.Ordered](s []T) T {
	var min T
	if len(s) == 0 {
		return min
	}
	min = s[0]
	for i := 1; i < len(s); i++ {
		item := s[i]
		if item < min {
			min = item
		}
	}

	return min
}
