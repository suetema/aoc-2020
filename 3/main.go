package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	tree rune = rune('#')
)

func countTrees(xi, yi int, matrix [][]rune) int {

	var x, y, trees int = 0, 0, 0
	for y < len(matrix) {
		if y < len(matrix) {
			if matrix[y][x%len(matrix[0])] == tree {
				trees++
			}
		}
		x += xi
		y += yi
	}
	return trees

}

func inputToMatrix(input string) [][]rune {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var matrix [][]rune
	index := 0
	for scanner.Scan() {
		matrix = append(matrix, []rune(scanner.Text()))
		index++
	}
	return matrix
}

func main() {
	var trees int = 0
	var slope [][]rune

	slope = inputToMatrix("input")
	trees = countTrees(3, 1, slope)
	fmt.Println(trees)
	trees *= countTrees(1, 1, slope)
	trees *= countTrees(5, 1, slope)
	trees *= countTrees(7, 1, slope)
	trees *= countTrees(1, 2, slope)
	fmt.Println(trees)

}
