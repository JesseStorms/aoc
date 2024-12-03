package main

import (
	"aoc/utils"
	_ "embed"
	"fmt"
	"slices"

	// "sort"
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
	left, right := parseInput(input)
	//sort both from smallest to largest
	slices.Sort(left)
	slices.Sort(right)
	//for each element in both slices, calculate the distance and add it to the sum
	sum := 0
	for i := 0; i < len(left); i++ {
		sum += absDiffInt(left[i], right[i])
	}
	return sum
}

func part2(input string) int {
	left, right := parseInput(input)
	//sort both from smallest to largest
	slices.Sort(left)
	slices.Sort(right)

	// Map the left, where the key is the value of the left
	left_map := map[int]int{}
	for _, l := range left {
		for _, r := range right {
			if l == r {
				left_map[l]++
			}
		}
	}
	sum := 0
	for k, v := range left_map {
		// multiply the value of the key with the amount of times it appears in the right slice
		// and add it to the sum
		sum += k * v
	}

	return sum
}

func parseInput(input string) ([]int, []int) {
	left := []int{}
	right := []int{}

	for _, line := range utils.SplitNewLines(input) {
		// split line into two parts
		split := strings.Split(line, "   ")
		// parse left and right side to int
		left_int, _ := strconv.Atoi(split[0])
		right_int, _ := strconv.Atoi(split[1])
		// append to left and right
		left = append(left, left_int)
		right = append(right, right_int)
	}
	return left, right
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
