package main

import (
	"fmt"
	"github.com/subchord/aoc-2019/util"
	"strconv"
)

func main() {
	ops := Opcodes(util.ReadIntSlice("inp"))

	done := make(chan interface{})
	out := make(chan int)

	go program(ops, func() int {
		return 0
	}, out, done)

	c := 0
	for {
		select {
		case o := <-out:
			switch o {
			case 35:
				c++
				fmt.Print("#")
			case 46:
				c++
				fmt.Print(".")
			case 10:
				fmt.Printf("\t%v\n", c)
				c = 0
			}
		case <-done:
			return
		}
	}

}

func program(opcode Opcodes, in func() int, out chan int, done chan interface{}) {
	o := make(Opcodes, len(opcode))
	copy(o, opcode)
	relativeBase := 0
	ptr := 0

	ns := make(Opcodes, len(o)+10000)
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
