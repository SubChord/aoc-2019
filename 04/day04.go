package main

import (
	"fmt"
	"github.com/subchord/aoc-2019/util"
	"strconv"
	"strings"
)

var inp = "138241-674034"

func main() {
	nums, _ := util.StringSlice(strings.Split(inp, "-")).ToIntSlice()
	fmt.Println(part1(nums[0], nums[1]))
	fmt.Println(part2(nums[0], nums[1]))
}

func part1(a int, b int) int {
	isValid := func(s string) bool {
		f := true
		f = f && len(s) == 6

		incr := true
		double := false
		var prev *int
		for _, c := range s {
			cint := util.ToInt(string(c))
			if prev != nil && cint < *prev {
				incr = false
			}

			if prev != nil && *prev == cint {
				double = true
			}
			prev = &cint
		}

		f = f && incr
		f = f && double
		return f
	}

	count := 0
	for i := 0; i < b-a; i++ {
		if isValid(strconv.Itoa(a + i)) {
			count++
		}
	}
	return count
}

func part2(a int, b int) int {
	isValid := func(s string) bool {
		f := true
		f = f && len(s) == 6

		incr := true
		double := false
		var prev *int
		for i := 0; i < len(s); i++ {
			cint := util.ToInt(string(s[i]))
			if prev != nil && cint < *prev {
				incr = false
			}

			if i+1 < len(s) && util.ToInt(string(s[i+1])) == cint {
				count := 0
				for i+1 < len(s) && util.ToInt(string(s[i+1])) == cint {
					i++
					count++
				}
				if count == 1 {
					double = true
				}
			}

			prev = &cint
		}

		f = f && incr
		f = f && double
		return f
	}

	count := 0
	for i := 0; i < b-a; i++ {
		if isValid(strconv.Itoa(a + i)) {
			count++
		}
	}
	return count
}
