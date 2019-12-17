package main

import (
	"fmt"
	"github.com/subchord/aoc-2019/util"
	"strconv"
	"strings"
)

func main() {
	lines := util.ReadLines("inp")
	inp := lines[0]

	part1(inp)
	part2(inp)
}

func part1(inp string) {
	pattern := []int{0, 1, 0, -1}
	numbers := util.StringSlice(strings.Split(inp, "")).ToIntSlice()

	for i := 0; i < 100; i++ {
		out := make([]int, len(numbers))
		for idx := range out {
			ptrn := getPattern(pattern, len(numbers), idx+1)
			sum := 0
			for j, v := range ptrn {
				sum += numbers[j] * v
			}
			sumStr := strconv.Itoa(sum)
			out[idx] = util.ToInt(string(sumStr[len(sumStr)-1]))
		}
		numbers = out
	}
	fmt.Println(numbers[:8])
}

func part2(inp string) {
	inpAsIntSlice := util.StringSlice(strings.Split(inp, "")).ToIntSlice()

	pattern := []int{0, 1, 0, -1}
	numbers := make([]int, 0, len(inp)*10000)
	for i := 0; i < 10000; i++ {
		numbers = append(numbers, inpAsIntSlice...)
	}

	for i := 0; i < 100; i++ {
		fmt.Printf("%v ", i)
		out := make([]int, len(numbers))
		for idx := range out {
			ptrn := getPattern(pattern, len(numbers), idx+1)
			sum := 0
			for j, v := range ptrn {
				sum += numbers[j] * v
			}
			sumStr := strconv.Itoa(sum)
			out[idx] = util.ToInt(string(sumStr[len(sumStr)-1]))
		}
		numbers = out
	}
	fmt.Println()

	first7str := ""
	for i := 0; i < 7; i++ {
		first7str += strconv.Itoa(numbers[i])
	}

	first7 := util.ToInt(first7str)

	fmt.Println(numbers[first7 : first7+8])
}

func getPattern(pattern []int, length, position int) []int {
	out := []int{}
	for len(out) < length+1 {
		for i := range pattern {
			for j := 0; j < position; j++ {
				out = append(out, pattern[i])
			}
		}
	}
	return out[1 : length+1]
}
