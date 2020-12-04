package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, num)
	}
	return lines, scanner.Err()
}

func compute(lines []int, target int) int {
	for i := 0; i < len(lines); i++ {
		for j := i + 1; j < len(lines); j++ {
			for k := j + 1; k < len(lines); k++ {
				if lines[j]+lines[i]+lines[k] == target {
					return lines[j] * lines[i] * lines[k]
				}
			}
		}
	}
	return 0
}

func main() {
	target := 2020
	lines, err := readLines("input")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	fmt.Println(compute(lines, target))
}
