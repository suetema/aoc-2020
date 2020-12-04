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
	var correct2dot1, correct2dot2 int
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

	// 2.1
	for _, p := range passwords {
		count := strings.Count(p.password, p.char)
		if count >= p.min && count <= p.max {
			correct2dot1++
		}
	}

	// 2.2
	for _, p := range passwords {
		if (string(p.password[p.min-1]) == p.char) != (string(p.password[p.max-1]) == p.char) {
			correct2dot2++
		}
	}

	fmt.Println(correct2dot1)
	fmt.Println(correct2dot2)
}
