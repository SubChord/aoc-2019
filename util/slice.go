package util

import "strconv"

type StringSlice []string

func (s StringSlice) ToIntSlice() ([]int, error) {
	ints := make([]int, len(s))
	for i, st := range s {
		atoi, err := strconv.Atoi(st)
		if err != nil {
			return nil, err
		}
		ints[i] = atoi
	}
	return ints, nil
}
