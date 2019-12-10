package main

import (
	"fmt"
	"github.com/subchord/aoc-2019/util"
	"math"
)

type Location struct {
	X int
	Y int
}

func main() {
	lines := util.ReadLines("inp")
	comets := parseLines(lines)
	location, max := part1(comets)
	fmt.Println(max)
	part2(location)
}

func part2(location Location) {

}

func part1(comets []Location) (Location, int) {
	canSee := make(map[Location]int)
	for i, comet := range comets {
		seen := make(map[float64]bool)
		for j, otherComet := range comets {
			if i == j {
				continue
			}
			atan2 := math.Atan2(float64(comet.Y-otherComet.Y), float64(comet.X-otherComet.X))
			seen[atan2] = true
		}
		canSee[comets[i]] = len(seen)
	}

	max := 0
	var bestLoc Location
	for loc, v := range canSee {
		if v > max {
			max = v
			bestLoc = loc
		}
	}

	return bestLoc, max
}

func parseLines(lines []string) []Location {
	comets := []Location{}
	for i, line := range lines {
		for j, v := range line {
			if v == '#' {
				comets = append(comets, Location{
					X: j,
					Y: i,
				})
			}
		}
	}
	return comets
}
