package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("part 1:", part1(input))
	fmt.Println("part 2:", part2(input))
}

func part1(input string) int {
	elves := parseInput(input)

	sort.Slice(elves, func(i, j int) bool {
		return elves[i] > elves[j]
	})
	return elves[0]
}

func part2(input string) int {
	elves := parseInput(input)

	sort.Slice(elves, func(i, j int) bool {
		return elves[i] > elves[j]
	})
	return elves[0] + elves[1] + elves[2]
}

func parseInput(input string) []int {
	res := []int{}
	for _, cals := range strings.Split(input, "\r\n\r\n") {
		total := 0
		for _, cal := range strings.Split(cals, "\r\n") {
			i, e := strconv.Atoi(cal)
			if e != nil {
				fmt.Println("Error converting to int???????????")
			}
			total += i
		}
		res = append(res, total)
	}
	return res
}
