package main

import (
	"fmt"
	"github.com/cznic/mathutil"
	"github.com/subchord/aoc-2019/util"
	"strconv"
)

func main() {
	ops := Opcodes(util.ReadIntSlice("inp"))

	seq := Phases{5, 6, 7, 8, 9}
	mathutil.PermutationFirst(seq)

	max := 0
	for {
		chans := []chan int{}
		for range seq {
			chans = append(chans, make(chan int, 1))
		}

		done := make(chan interface{})
		for i, v := range seq {
			go program(Opcodes(append([]int(nil), ops...)), chans[i], chans[(i+1)%len(chans)], done)
			chans[i] <- v
		}

		chans[0] <- 0
		for range seq {
			<-done
		}

		out := <-chans[0]
		if out > max {
			max = out
		}

		next := mathutil.PermutationNext(seq)
		if !next {
			break
		}
	}

	fmt.Println(max)

}

func program(o Opcodes, in <-chan int, out chan int, done chan interface{}) {
	ptr := 0

	for {
		instruction := o[ptr]
		inst, modes := parseInstruction(instruction)

		switch inst {
		case 1:
			o[o[ptr+3]] = o.getValue(ptr+1, modes[0]) + o.getValue(ptr+2, modes[1])
		case 2:
			o[o[ptr+3]] = o.getValue(ptr+1, modes[0]) * o.getValue(ptr+2, modes[1])
		case 3:
			o[o[ptr+1]] = <-in
		case 4:
			out <- o.getValue(ptr+1, modes[0])
		case 5:
			if o.getValue(ptr+1, modes[0]) > 0 {
				ptr = o.getValue(ptr+2, modes[1])
				continue
			}
		case 6:
			if o.getValue(ptr+1, modes[0]) == 0 {
				ptr = o.getValue(ptr+2, modes[1])
				continue
			}
		case 7:
			if o.getValue(ptr+1, modes[0]) < o.getValue(ptr+2, modes[1]) {
				o[o[ptr+3]] = 1
			} else {
				o[o[ptr+3]] = 0
			}
		case 8:
			if o.getValue(ptr+1, modes[0]) == o.getValue(ptr+2, modes[1]) {
				o[o[ptr+3]] = 1
			} else {
				o[o[ptr+3]] = 0
			}
		case 99:
			done <- struct{}{}
			return
		}
		ptr += []int{0, 4, 4, 2, 2, 3, 3, 4, 4}[inst]
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

func (o Opcodes) getValue(i int, m mode) int {
	switch m {
	case 0:
		return o[o[i]]
	case 1:
		return o[i]
	}
	panic("unknown mode")
}

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
