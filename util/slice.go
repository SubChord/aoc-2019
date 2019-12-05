package util

import "strconv"

type StringSlice []string

func (s StringSlice) ToIntSlice() []int {
	ints := make([]int, len(s))
	for i, st := range s {
		atoi, err := strconv.Atoi(st)
		Check(err)
		ints[i] = atoi
	}
	return ints
}
