package main

import (
	"fmt"
	"github.com/subchord/aoc-2019/util"
	"math"
	"strings"
)

type location struct {
	x, y int
}

func main() {
	lines := util.ReadLines("inp")
	l1 := strings.Split(lines[0], ",")
	l2 := strings.Split(lines[1], ",")

	//fmt.Println(part1(l1, l2))
	fmt.Println(part2(l1, l2))
}

func do(l location, s string) location {
	switch s[0] {
	case 'U':
		return location{
			x: l.x,
			y: l.y + 1,
		}
	case 'R':
		return location{
			x: l.x + 1,
			y: l.y,
		}
	case 'D':
		return location{
			x: l.x,
			y: l.y - 1,
		}
	case 'L':
		return location{
			x: l.x - 1,
			y: l.y,
		}
	}
	panic("unkown direction")
}

func part1(l1 []string, l2 []string) int {
	locs1 := []location{}
	locs2 := []location{}

	prevloc := location{}
	for _, v := range l1 {
		for i := 0; i < util.ToInt(v[1:]); i++ {
			l := do(prevloc, v)
			locs1 = append(locs1, l)
			prevloc = l
		}
	}

	prevloc = location{}
	for _, v := range l2 {
		for i := 0; i < util.ToInt(v[1:]); i++ {
			l := do(prevloc, v)
			locs2 = append(locs2, l)
			prevloc = l
		}
	}

	mhd := math.MaxInt64
	for _, v1 := range locs1 {
		for _, v2 := range locs2 {
			if v1 == v2 && !(v1.x == 0 && v1.y == 0) {
				mhd = int(math.Min(float64(mhd), math.Abs(float64(v1.x))+math.Abs(float64(v1.y))))
			}
		}
	}
	return mhd
}

func part2(l1 []string, l2 []string) int {
	locs1 := []location{}
	locs2 := []location{}

	locsmap1 := make(map[location]int)
	locsmap2 := make(map[location]int)

	prevloc := location{}
	dist := 0
	for _, v := range l1 {
		for i := 0; i < util.ToInt(v[1:]); i++ {
			dist++
			l := do(prevloc, v)
			locs1 = append(locs1, l)
			_, ok := locsmap1[l]
			if !ok {
				locsmap1[l] = dist
			}
			prevloc = l
		}
	}

	prevloc = location{}
	dist = 0
	for _, v := range l2 {
		for i := 0; i < util.ToInt(v[1:]); i++ {
			dist++
			l := do(prevloc, v)
			locs2 = append(locs2, l)
			_, ok := locsmap2[l]
			if !ok {
				locsmap2[l] = dist
			}
			prevloc = l
		}
	}

	v := math.MaxInt64
	for _, v1 := range locs1 {
		for _, v2 := range locs2 {
			if v1 == v2 && !(v1.x == 0 && v1.y == 0) {

				v = int(math.Min(float64(v), float64(locsmap1[v1]+locsmap2[v2])))

			}
		}
	}
	return v
}
