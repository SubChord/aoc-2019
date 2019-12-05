package main

import (
	"errors"
	"fmt"
	"github.com/subchord/aoc-2019/util"
	"strings"
)

func main() {
	lines := util.ReadLines("inp")
	parts := strings.Split(lines[0], ",")
	numbers := util.StringSlice(parts).ToIntSlice()

	fmt.Println(part1(append([]int(nil), numbers...)))
	fmt.Println(part2(append([]int(nil), numbers...)))
}

func do(ints []int) error {
	pt := 0
	for ints[pt] != 99 {
		if ints[pt] == 1 {
			ints[ints[pt+3]] = ints[ints[pt+1]] + ints[ints[pt+2]]
		} else if ints[pt] == 2 {
			ints[ints[pt+3]] = ints[ints[pt+1]] * ints[ints[pt+2]]
		} else {
			return errors.New("unkown opt")
		}
		pt += 4
	}
	return nil
}

func part1(ints []int) int {
	ints[1] = 12
	ints[2] = 2
	err := do(ints)
	if err != nil {
		return -1
	}
	return ints[0]
}

func part2(ints []int) int {
	a, b := 0, 0
out:
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			a, b = i, j
			nums := append([]int(nil), ints...)
			nums[1] = a
			nums[2] = b
			_ = do(nums)
			if nums[0] == 19690720 {
				break out
			}
		}
	}

	return 100*a + b
}
