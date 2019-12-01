package main

import (
	"fmt"
	"github.com/subchord/aoc-2019/util"
)

func main() {
	numbers := util.ReadIntLines("inp")
	fmt.Println(part1(numbers))
	fmt.Println(part2(numbers))
}

func part1(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n/3 - 2
	}
	return sum
}

func part2(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		f := n
		for f > 6 {
			f = f/3 - 2
			sum += f
		}
	}
	return sum
}
