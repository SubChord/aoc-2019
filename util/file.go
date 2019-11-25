package util

import (
	"bufio"
	"log"
	"os"
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
