package main

import (
	"errors"
	"fmt"
	"github.com/subchord/aoc-2019/util"
	"strconv"
)

func main() {
	nums := util.ReadIntSlice("inp")
	part1(opcodes(nums))
	fmt.Println()
	nums = util.ReadIntSlice("inp")
	part2(opcodes(nums))
}

type parameterMode int

type opcodes []int

var err99 = errors.New("99")

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

func (ints opcodes) getValue(ptr int, mode parameterMode) int {
	switch mode {
	case 0:
		return ints[ints[ptr]]
	case 1:
		return ints[ptr]
	}
	panic("unknown mode")
}

func (ints opcodes) doInstructionPart1(ptr int, inp int) (int, error) {
	instruction := ints[ptr]
	if instruction == 99 {
		return 0, err99
	}

	inst, modes := parseInstruction(instruction)

	if inst == 1 {
		ints[ints[ptr+3]] = ints.getValue(ptr+1, modes[0]) + ints.getValue(ptr+2, modes[1])
		return ptr + 4, nil
	} else if inst == 2 {
		ints[ints[ptr+3]] = ints.getValue(ptr+1, modes[0]) * ints.getValue(ptr+2, modes[1])
		return ptr + 4, nil
	} else if inst == 3 {
		ints[ints[ptr+1]] = inp
		return ptr + 2, nil
	} else if inst == 4 {
		v := ints.getValue(ptr+1, modes[0])
		if v >0 {
			fmt.Print(v)
		}
		return ptr + 2, nil
	}

	panic(fmt.Sprintf("unknown op: %v", instruction))
}

func part1(nums opcodes) {
	ptr := 0
	var err error
	for {
		ptr, err = nums.doInstructionPart1(ptr, 1)
		if err != nil {
			break
		}
	}
}

func (ints opcodes) doInstructionPart2(ptr int, inp int) (int, error) {
	instruction := ints[ptr]
	if instruction == 99 {
		return 0, err99
	}

	inst, modes := parseInstruction(instruction)

	if inst == 1 {
		ints[ints[ptr+3]] = ints.getValue(ptr+1, modes[0]) + ints.getValue(ptr+2, modes[1])
		return ptr + 4, nil
	} else if inst == 2 {
		ints[ints[ptr+3]] = ints.getValue(ptr+1, modes[0]) * ints.getValue(ptr+2, modes[1])
		return ptr + 4, nil
	} else if inst == 3 {
		ints[ints[ptr+1]] = inp
		return ptr + 2, nil
	} else if inst == 4 {
		fmt.Print(ints.getValue(ptr+1, modes[0]))
		return ptr + 2, nil
	} else if inst == 5 {
		if ints.getValue(ptr+1, modes[0]) > 0 {
			return ints.getValue(ptr+2, modes[1]), nil
		}
		return ptr + 3, nil
	} else if inst == 6 {
		if ints.getValue(ptr+1, modes[0]) == 0 {
			return ints.getValue(ptr+2, modes[1]), nil
		}
		return ptr + 3, nil
	} else if inst == 7 {
		if ints.getValue(ptr+1, modes[0]) < ints.getValue(ptr+2, modes[1]) {
			ints[ints[ptr+3]] = 1
		} else {
			ints[ints[ptr+3]] = 0
		}
		return ptr + 4, nil
	}else if inst == 8 {
		if ints.getValue(ptr+1, modes[0]) == ints.getValue(ptr+2, modes[1]) {
			ints[ints[ptr+3]] = 1
		} else {
			ints[ints[ptr+3]] = 0
		}
		return ptr + 4, nil
	}

	panic(fmt.Sprintf("unknown op: %v", instruction))
}

func part2(nums opcodes) {
	ptr := 0
	var err error
	for {
		ptr, err = nums.doInstructionPart2(ptr, 5)
		if err != nil {
			break
		}
	}
}
