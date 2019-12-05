package util

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadLines(path string) (lines []string) {
	file, e := os.Open(path)
	if e != nil {
		log.Fatal(e)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func ReadIntLines(path string) []int {
	lines := ReadLines(path)
	numbers := make([]int, len(lines))
	for i, line := range lines {
		atoi, err := strconv.Atoi(line)
		Check(err)
		numbers[i] = atoi
	}
	return numbers
}

func ReadIntSlice(path string) []int {
	lines := ReadLines(path)
	split := StringSlice(strings.Split(lines[0], ","))
	return split.ToIntSlice()
}
