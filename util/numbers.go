package util

import "strconv"

func IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func ToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
