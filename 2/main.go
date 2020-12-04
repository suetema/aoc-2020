package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Password struct {
	min, max       int
	char, password string
}

func main() {
	var passwords []Password
	var correct int
	file, _ := os.Open("input")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	regex := regexp.MustCompile(`^(\d+)-(\d+) (\w): (\w+)$`)

	for scanner.Scan() {
		line := scanner.Text()
		min, _ := strconv.Atoi(regex.ReplaceAllString(line, "$1"))
		max, _ := strconv.Atoi(regex.ReplaceAllString(line, "$2"))
		password := Password{
			min:      min,
			max:      max,
			char:     regex.ReplaceAllString(line, "$3"),
			password: regex.ReplaceAllString(line, "$4"),
		}
		passwords = append(passwords, password)
	}

	for _, p := range passwords {
		count := strings.Count(p.password, p.char)
		if count >= p.min && count <= p.max {
			correct++
		}

	}
	fmt.Println(correct)
}
