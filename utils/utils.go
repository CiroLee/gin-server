package utils

import "strings"

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
