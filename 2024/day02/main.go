package main

import (
	"aoc/utils"
	_ "embed"
	"fmt"
	"strconv"

	// "sort"

	"strings"
)

//go:embed input.txt
var input string

func main() {
	// fmt.Println("part 1:", part1(input))
	fmt.Println("part 2:", part2(input))
}

func part1(input string) int {
	sum := 0
	for _, line := range utils.SplitNewLines(input) {
		//parse array of numbers
		numbers := []int{}
		for _, number := range strings.Fields(line) {
			num, _ := strconv.Atoi(number)
			numbers = append(numbers, num)
		}
		if isSafe(numbers) {
			sum++
			fmt.Println(numbers, "is safe")
		} else {
			fmt.Println(numbers, "is not safe")
		}
	}

	return sum
}

func part2(input string) int {
	sum := 0
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		// Parse array of numbers
		numbers := []int{}
		for _, number := range strings.Fields(line) {
			num, _ := strconv.Atoi(number)
			numbers = append(numbers, num)
		}
		if checkSafe(numbers) {
			sum++
			fmt.Println(numbers, "is safe")
		} else {
			fmt.Println(numbers, "is not safe")
		}
	}

	return sum
}

func isSafe(slice []int) bool {
	if len(slice) < 2 {
		return true
	}

	increasing := slice[1] > slice[0]
	for i := 1; i < len(slice); i++ {
		diff := slice[i] - slice[i-1]
		if (increasing && diff < 1) || (!increasing && diff > -1) || diff > 3 || diff < -3 {
			return false
		}
	}
	return true
}

// Helper function to check if a slice is safe
func checkSafe(slice []int) bool {
	if isSafe(slice) {
		return true
	}
	// Try removing each element one by one and check if the resulting slice is safe
	for i := range slice {
		// fmt.Printf("originally %v\n", slice)
		//copy the slice
		newSlice := make([]int, len(slice))
		copy(newSlice, slice)
		bruh := append(newSlice[:i], newSlice[i+1:]...)
		if isSafe(bruh) {
			return true
		}
	}

	return false
}
