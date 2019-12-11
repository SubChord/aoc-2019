package main

import (
	"fmt"
	"github.com/subchord/aoc-2019/util"
	"math"
	"sort"
	"strconv"
)

type Location struct {
	X int
	Y int
}

func (l Location) Move(direction int) Location {
	switch direction {
	case 0:
		l.Y++
	case 1:
		l.X++
	case 2:
		l.Y--
	case 3:
		l.X--
	}
	return l
}

func main() {
	opc := Opcodes(util.ReadIntSlice("inp"))
	part1(append([]int{}, opc...))
	part2(append([]int{}, opc...))
}

func part1(opc Opcodes) {
	direction := 0
	currentLocation := Location{}

	paint := make(map[Location]bool)

	done := make(chan interface{})
	out := make(chan int)

	go program(opc, func() int {
		b, ok := paint[currentLocation]
		if !ok {
			return 0
		}
		if !b {
			return 0
		} else {
			return 1
		}
	}, out, done)

	counter := 0
	for {
		select {
		case <-done:
			fmt.Println(len(paint))
			return
		case v := <-out:
			if counter%2 == 0 {
				paint[currentLocation] = v == 1
			} else {
				switch v {
				case 0:
					direction += 3
				case 1:
					direction += 1
				}
				direction %= 4
				currentLocation = currentLocation.Move(direction)
			}
			counter++
		}
	}
}

func part2(opc Opcodes) {
	direction := 0
	currentLocation := Location{}

	paint := make(map[Location]bool)

	done := make(chan interface{})
	out := make(chan int)

	paint[currentLocation] = true
	go program(opc, func() int {
		b, ok := paint[currentLocation]
		if !ok {
			return 0
		}
		if !b {
			return 0
		} else {
			return 1
		}
	}, out, done)

	counter := 0
out:
	for {
		select {
		case <-done:
			break out
		case v := <-out:
			if counter%2 == 0 {
				paint[currentLocation] = v == 1
			} else {
				switch v {
				case 0:
					direction += 3
				case 1:
					direction += 1
				}
				direction %= 4
				currentLocation = currentLocation.Move(direction)
			}
			counter++
		}
	}

	locations := []Location{}
	for l := range paint {
		locations = append(locations, l)
	}

	sort.Slice(locations, func(i, j int) bool {
		if locations[i].Y < locations[j].Y {
			return false
		}
		if locations[i].Y > locations[j].Y {
			return true
		}
		return locations[i].X < locations[j].X
	})

	minX := math.MaxInt64
	for _, location := range locations {
		if location.X < minX {
			minX = location.X
		}
	}

	lastY := locations[0].Y
	lastX := minX
	for _, location := range locations {
		if location.Y < lastY {
			lastY = location.Y
			lastX = minX
			fmt.Println()
		}

		for lastX < location.X {
			fmt.Print(" ")
			lastX ++
		}

		if paint[location] {
			fmt.Print("#")
		}else{
			fmt.Print(" ")
		}
		lastX++
	}
}

func program(opcode Opcodes, in func() int, out chan int, done chan interface{}) {
	o := make(Opcodes, len(opcode))
	copy(o, opcode)
	relativeBase := 0
	ptr := 0

	ns := make(Opcodes, len(o)+1000)
	o = append(o, ns...)

	for {
		instruction := o[ptr]
		inst, modes := parseInstruction(instruction)

		if ptr > len(o) {
			ns := make(Opcodes, len(o)+1000)
			o = append(o, ns...)
		}

		arg := func(i int) (addr int) {
			switch modes[i-1] {
			case 0:
				addr = o[ptr+i]
			case 1:
				addr = ptr + i
			case 2:
				addr = relativeBase + o[ptr+i]
			}
			return
		}

		switch inst {
		case 1:
			o[arg(3)] = o[arg(1)] + o[arg(2)]
		case 2:
			o[arg(3)] = o[arg(1)] * o[arg(2)]
		case 3:
			o[arg(1)] = in()
		case 4:
			out <- o[arg(1)]
		case 5:
			if o[arg(1)] != 0 {
				ptr = o[arg(2)]
				continue
			}
		case 6:
			if o[arg(1)] == 0 {
				ptr = o[arg(2)]
				continue
			}
		case 7:
			if o[arg(1)] < o[arg(2)] {
				o[arg(3)] = 1
			} else {
				o[arg(3)] = 0
			}
		case 8:
			if o[arg(1)] == o[arg(2)] {
				o[arg(3)] = 1
			} else {
				o[arg(3)] = 0
			}
		case 9:
			relativeBase += o[arg(1)]
		case 99:
			done <- true
			return
		}
		ptr += []int{0, 4, 4, 2, 2, 3, 3, 4, 4, 2}[inst]
	}
}

//Phases
type Phases []int

func (p Phases) Len() int {
	return len(p)
}

func (p Phases) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p Phases) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Opcodes
type Opcodes []int

type mode int

func parseInstruction(v int) (int, []mode) {
	s := strconv.Itoa(v)
	modes := make([]mode, 3)

	if len(s) <= 2 {
		return util.ToInt(s), modes
	}

	inst := util.ToInt(s[len(s)-2:])

	opModes := s[:len(s)-2]
	for i := range opModes {
		modes[len(opModes)-1-i] = mode(util.ToInt(string(opModes[i])))
	}

	return inst, modes
}
