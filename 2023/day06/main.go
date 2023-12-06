package main

import (
	"aoc/utils"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("part 1:", part1(input))
	// fmt.Println("part 2:", part2(input))
}

func part1(input string) int {
	races := parseRaces(input)
	fmt.Println(races)
	return 0
}

func parseRaces(input string) []Race {
	// replace all multiple spaces with a single space using regex

	lines := utils.SplitNewLines(input)
	var result []Race
	for _, thing := range strings.Split(lines[0], " ")[1:] {
		time, _ := strconv.Atoi(thing)
		result = append(result, Race{Time: time})
	}
	for i, thing := range strings.Split(lines[1], " ")[1:] {
		fmt.Println(thing)
		distance, _ := strconv.Atoi(thing)
		result[i].Distance = distance
	}
	return result
}
