package main

import (
	"fmt"
	"github.com/subchord/aoc-2019/util"
	"math"
	"strings"
)

type Layer [][]int

func main() {
	lines := util.ReadLines("inp")
	digits := util.StringSlice(strings.Split(lines[0], "")).ToIntSlice()

	part1(digits)
	part2(digits)
}

func part1(digits []int) {
	layers := digitsToLayers(digits, 25, 6)

	digs := func(l Layer, dig int) int {
		count := 0
		for _, ints := range l {
			for _, v := range ints {
				if v == dig {
					count++
				}
			}
		}
		return count
	}

	least := math.MaxInt64
	var idx int
	for i, layer := range layers {
		zeroCount := digs(layer, 0)
		if zeroCount < least {
			least = zeroCount
			idx = i
		}
	}

	fmt.Println(digs(layers[idx], 1) * digs(layers[idx], 2))
}

func part2(digits []int) {
	layers := digitsToLayers(digits, 25, 6)
	stackedLayer := [6][25][]int{}

	for _, layer := range layers {
		for i, row := range layer {
			for j, pixel := range row {
				stackedLayer[i][j] = append(stackedLayer[i][j], pixel)
			}
		}		
	}

	for _, row := range stackedLayer {
		for _, cell := range row {
			for _, v := range cell {
				if v == 1 {
					fmt.Print("x")
					break
				}
				if v == 0 {
					fmt.Print(" ")
					break
				}
			}
		}
		fmt.Println()
	}
}

func digitsToLayers(digits []int, width int, height int) []Layer {
	layers := []Layer{}
	ptr := 0
	for ptr < len(digits) {
		l := Layer{}
		for h := 0; h < height; h++ {
			l = append(l, digits[ptr:ptr+width])
			ptr += width
		}
		layers = append(layers, l)
	}
	return layers
}
