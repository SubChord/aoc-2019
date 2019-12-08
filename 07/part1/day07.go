package main

import (
	"errors"
	"fmt"
	"github.com/cznic/mathutil"
	"github.com/subchord/aoc-2019/util"
	"strconv"
)

func main() {
	ops := opcodes(util.ReadIntSlice("inp"))
	part1(ops)
}

func part1(ops opcodes) {

	results := []int{}

	phases := opcodes([]int{0, 1, 2, 3, 4})
	mathutil.PermutationFirst(phases)

	for {
		prevResult := 0
		for _, phase := range phases {
			opcopy := make(opcodes, len(ops))
			copy(opcopy, ops)
			prevResult, _, _ = opcopy.doAmp(0, phase, prevResult, false)
		}
		results = append(results, prevResult)

		next := mathutil.PermutationNext(phases)
		if !next {
			break
		}
	}

	fmt.Println(mathutil.MaxVal(0, results...))

}

func (ops opcodes) doAmp(pointer int, phase int, inpt int, useInpOnce bool) (int, int, error) {
	ptr := pointer
	inputted := useInpOnce
	for {
		inp := phase
		if inputted {
			inp = inpt
		}
		p, err, result, input := ops.doInstruction(ptr, inp)

		ptr = p
		if !inputted && input {
			inputted = true
		}

		if result > 0 {
			return result, ptr, nil
		}

		if err != nil {
			return 0, 0, err
		}
	}
}

type opcodes []int

func (o opcodes) Len() int {
	return len(o)
}

func (o opcodes) Less(i, j int) bool {
	return o[i] < o[j]
}

func (o opcodes) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}

func (ints opcodes) doInstruction(ptr int, inp int) (int, error, int, bool) {
	instruction := ints[ptr]
	if instruction == 99 {
		return 0, err99, -1, false
	}

	inst, modes := parseInstruction(instruction)

	if inst == 1 {
		ints[ints[ptr+3]] = ints.getValue(ptr+1, modes[0]) + ints.getValue(ptr+2, modes[1])
		return ptr + 4, nil, -1, false
	} else if inst == 2 {
		ints[ints[ptr+3]] = ints.getValue(ptr+1, modes[0]) * ints.getValue(ptr+2, modes[1])
		return ptr + 4, nil, -1, false
	} else if inst == 3 {
		ints[ints[ptr+1]] = inp
		return ptr + 2, nil, -1, true
	} else if inst == 4 {
		return ptr + 2, nil, ints.getValue(ptr+1, modes[0]), false
	} else if inst == 5 {
		if ints.getValue(ptr+1, modes[0]) > 0 {
			return ints.getValue(ptr+2, modes[1]), nil, 0, false
		}
		return ptr + 3, nil, -1, false
	} else if inst == 6 {
		if ints.getValue(ptr+1, modes[0]) == 0 {
			return ints.getValue(ptr+2, modes[1]), nil, 0, false
		}
		return ptr + 3, nil, -1, false
	} else if inst == 7 {
		if ints.getValue(ptr+1, modes[0]) < ints.getValue(ptr+2, modes[1]) {
			ints[ints[ptr+3]] = 1
		} else {
			ints[ints[ptr+3]] = 0
		}
		return ptr + 4, nil, -1, false
	} else if inst == 8 {
		if ints.getValue(ptr+1, modes[0]) == ints.getValue(ptr+2, modes[1]) {
			ints[ints[ptr+3]] = 1
		} else {
			ints[ints[ptr+3]] = 0
		}
		return ptr + 4, nil, -1, false
	}

	panic(fmt.Sprintf("unknown op: %v", instruction))
}

type parameterMode int

var err99 = errors.New("99")

func (ints opcodes) getValue(ptr int, mode parameterMode) int {
	switch mode {
	case 0:
		return ints[ints[ptr]]
	case 1:
		return ints[ptr]
	}
	panic("unknown mode")
}

func parseInstruction(instValue int) (int, []parameterMode) {
	s := strconv.Itoa(instValue)
	modes := make([]parameterMode, 3)

	if len(s) <= 2 {
		return util.ToInt(s), modes
	}

	inst := util.ToInt(s[len(s)-2:])

	opModes := s[:len(s)-2]
	for i := range opModes {
		modes[len(opModes)-1-i] = parameterMode(util.ToInt(string(opModes[i])))
	}

	return inst, modes
}
