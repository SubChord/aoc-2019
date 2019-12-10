package main

import (
	"fmt"
	"github.com/subchord/aoc-2019/util"
	"math"
	"sort"
)

type Location struct {
	X int
	Y int
}

func (l Location) manhattan(l2 Location) int {
	return int(math.Abs(float64(l.X-l2.X)) + math.Abs(float64(l.Y-l2.Y)))
}

func main() {
	lines := util.ReadLines("inp")
	comets := parseLines(lines)
	location, max := part1(comets)
	fmt.Println(max)
	part2(location, comets)
}

func part2(location Location, comets []Location) {
	distances := make(map[float64][]Location)

	// Group comets by atan2 (same line)
	for i, comet := range comets {
		if comet == location {
			continue
		}

		atan2 := math.Atan2(float64(location.Y-comet.Y), float64(location.X-comet.X))

		// correct atan2 so that 90deg == 0
		atan2 += 1.5 * math.Pi
		atan2 = math.Mod(atan2, 2*math.Pi)

		_, ok := distances[atan2]
		if !ok {
			distances[atan2] = []Location{}
		}
		distances[atan2] = append(distances[atan2], comets[i])
	}

	// Sort slices based on manhattan distance
	for _, comets := range distances {
		sort.Slice(comets, func(i, j int) bool {
			return comets[i].manhattan(location) < comets[j].manhattan(location)
		})
	}

	atans := []float64{}
	for f := range distances {
		atans = append(atans, f)
	}
	sort.Float64s(atans)

	count := 0
	i := 0
	var l Location
	for count < 200 {
		for {
			key := atans[i % len(atans)]
			locations := distances[key]
			if len(locations) == 0 {
				i++
				continue
			}
			distances[key] = locations[1:]
			l = locations[0]
			i++
			break
		}

		count++
	}

	fmt.Println(l.X*100 + l.Y)
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
