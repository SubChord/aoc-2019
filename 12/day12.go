package main

import (
	"fmt"
	"github.com/subchord/aoc-2019/util"
	"math"
	"regexp"
	"strings"
)

var inp = `<x=-9, y=-1, z=-1>
<x=2, y=9, z=5>
<x=10, y=18, z=-12>
<x=-6, y=15, z=-7>`

type Vector struct {
	X int
	Y int
	Z int
}

type Moon struct {
	Location Vector
	Velocity Vector
}

func (m *Moon) Energy() int {
	pot := 0.0
	pot += math.Abs(float64(m.Location.X))
	pot += math.Abs(float64(m.Location.Y))
	pot += math.Abs(float64(m.Location.Z))

	kin := 0.0
	kin += math.Abs(float64(m.Velocity.X))
	kin += math.Abs(float64(m.Velocity.Y))
	kin += math.Abs(float64(m.Velocity.Z))

	return int(pot * kin)
}

type Moons []*Moon

func (m Moons) Energy() (s int) {
	for _, moon := range m {
		s += moon.Energy()
	}
	return
}

func (m Moons) Tick() {
	for _, m1 := range m {
		for _, m2 := range m {
			if m1 == m2 {
				continue
			}

			if m2.Location.X > m1.Location.X {
				m1.Velocity.X++
			} else if m2.Location.X < m1.Location.X {
				m1.Velocity.X--
			}

			if m2.Location.Y > m1.Location.Y {
				m1.Velocity.Y++
			} else if m2.Location.Y < m1.Location.Y {
				m1.Velocity.Y--
			}

			if m2.Location.Z > m1.Location.Z {
				m1.Velocity.Z++
			} else if m2.Location.Z < m1.Location.Z {
				m1.Velocity.Z--
			}
		}
	}

	for _, m := range m {
		m.Location.X += m.Velocity.X
		m.Location.Y += m.Velocity.Y
		m.Location.Z += m.Velocity.Z
	}
}

func main() {
	moons := parseMoons()
	part1(append([]*Moon{}, moons...))
	moons = parseMoons()
	part2(append([]*Moon{}, moons...))
}

func part1(moons Moons) {
	for i := 0; i < 1000; i++ {
		moons.Tick()
	}
	fmt.Println(moons.Energy())
}

func part2(moons Moons) {
	initial := append([]*Moon{}, moons...)
	axisCount := make([]int, 3)
	done := make([]bool, 3)

	for !done[0] || !done[1] || !done[2] {
		moons.Tick()

		for i := 0; i < len(axisCount); i++ {
			if !done[i] {
				d := true
				for m, moon := range moons {
					switch i {
					case 0:
						d = d && moon.Location.X == initial[m].Location.X && moon.Velocity.X == initial[m].Velocity.X
					case 1:
						d = d && moon.Location.Y == initial[m].Location.Y && moon.Velocity.Y == initial[m].Velocity.Y
					case 2:
						d = d && moon.Location.Z == initial[m].Location.Z && moon.Velocity.Z == initial[m].Velocity.Z
					}
				}

				if d {
					done[i] = true
				} else {
					axisCount[i]++
				}
			}
		}
	}
	fmt.Println(LCM(axisCount[0], axisCount[1], axisCount[2]))
}

func parseMoons() []*Moon {
	moons := []*Moon{}
	for _, v := range strings.Split(inp, "\n") {
		r := regexp.MustCompile(`<x=(-?\d*), y=(-?\d*), z=(-?\d*)>`)
		submatch := r.FindAllStringSubmatch(v, -1)
		x := util.ToInt(submatch[0][1])
		y := util.ToInt(submatch[0][2])
		z := util.ToInt(submatch[0][3])
		moons = append(moons, &Moon{
			Location: Vector{
				X: x,
				Y: y,
				Z: z,
			},
			Velocity: Vector{},
		})
	}
	return moons
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
