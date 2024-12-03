package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	// "sort"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("part 1:", part1(input))
	fmt.Println("part 2:", part2(input))
}

func part1(input string) int {
	sum := 0
	r := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := r.FindAllStringSubmatch(input, -1)
	// for every 2 numbers, multiply them and add to sum
	for _, match := range matches {
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		sum += num1 * num2
	}
	return sum
}

func part2(input string) int {
	sum := 0
	r := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)
	matches := r.FindAllStringSubmatch(input, -1)
	run := true
	// for every 2 numbers, multiply them and add to sum
	for _, match := range matches {
		if match[0] == "do()" {
			run = true
		} else if match[0] == "don't()" {
			run = false
		} else if run {
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])
			sum += num1 * num2
		}
	}
	return sum
}
