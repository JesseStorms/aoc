package main

import (
	"aoc/utils"
	_ "embed"
	"fmt"
	"regexp"
	"unicode"

	// "sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

var words = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func main() {
	fmt.Println("part 1:", part1(input))
	fmt.Println("part 2:", part2(input))
}

func part1(input string) int {
	values := parseInput(input)
	total := 0
	for _, line := range values {
		total += line
	}
	return total
}

func part2(input string) int {
	values := parseInput2(input)
	total := 0
	for _, line := range values {
		total += line
	}
	return total
}

func parseInput(input string) []int {
	res := []int{}
	re := regexp.MustCompile(`[A-z]`)

	for _, line := range utils.SplitNewLines(input) {
		match := re.ReplaceAllString(line, "")

		i, e := strconv.Atoi(match[0:1] + match[len(match)-1:])
		if e != nil {
			fmt.Println("Error converting to int???????????")
		}
		if i < 10 {
			i = i + i*10
		}
		res = append(res, i)
	}
	return res
}

func parseInput2(input string) []int {
	res := []int{}
	re := regexp.MustCompile(`[A-z]`)

	for _, line := range utils.SplitNewLines(input) {
		parsing := ""
		newl := ""
		for _, char := range line {
			if unicode.IsLetter(char) {
				parsing += string(char)
			} else {
				parsing = ""
				newl += string(char)
			}
			for i, word := range words {
				if strings.Contains(parsing, word) {
					newl += strconv.Itoa(i + 1)
					// retain last char of parsing
					parsing = string(parsing[len(parsing)-1])
				}
			}
		}
		line = newl

		match := re.ReplaceAllString(line, "")

		i, e := strconv.Atoi(match[0:1] + match[len(match)-1:])
		if e != nil {
			fmt.Println("Error converting to int???????????")
		}
		if i < 10 {
			i = i + i*10
		}
		// fmt.Println(i)
		res = append(res, i)
	}
	return res
}
