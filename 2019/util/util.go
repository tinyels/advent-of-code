package util

import "strconv"

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func GetIntFromString(s string) int {
	i, err := strconv.Atoi(s)
	Check(err)
	return i
}

func MinInt(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
