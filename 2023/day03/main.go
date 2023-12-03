package main

import (
	"aoc/utils"
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"unicode"
	// "sort"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("part 1:", part1(input))
	fmt.Println("part 2:", part2(input))
}

func part1(input string) int {
	partNumbers := parseInput(input)
	total := 0
	for _, part := range partNumbers {
		if part.symbol != 'N' {
			total += part.part
		}
	}
	return total
}

func part2(input string) int {
	partNumbers := parseInput(input)
	gears, _ := []partNumber{}, []partNumber{}
	for _, part := range partNumbers {
		if part.symbol != 'N' {
			if part.symbol == '*' {
				gears = append(gears, part)
			}
		}
	}
	total := 0
	sharedGears := map[[2]int][]partNumber{}
	for _, gear := range gears {
		sharedGears[gear.symbolPos] = append(sharedGears[gear.symbolPos], gear)
	}
	for _, gear := range sharedGears {
		if len(gear) == 2 {
			total += gear[0].part * gear[1].part
		}
	}

	return total
}

func parseInput(input string) []partNumber {
	currentCharSequence := ""
	partNumbers := []partNumber{}
	currentPartNumber := partNumber{}
	partMap := buildMap(input)

	// Function to handle end of number sequence
	handleEndOfSequence := func(x, y int) {
		if currentCharSequence != "" {
			currentPartNumber.posEnd = [2]int{x, y}
			currentPartNumber.posStart = [2]int{x - len(currentCharSequence), y}
			currentPartNumber.part, _ = strconv.Atoi(currentCharSequence)
			currentPartNumber.symbolPos, currentPartNumber.symbol = findSymbol(partMap, currentPartNumber)
			partNumbers = append(partNumbers, currentPartNumber)
			currentCharSequence = ""
		}
	}

	for y, line := range partMap {
		for x, char := range line {
			if _, err := strconv.Atoi(string(char)); err == nil {
				currentCharSequence += string(char)
			} else {
				handleEndOfSequence(x, y)
			}
		}
		handleEndOfSequence(len(line), y)
	}
	return partNumbers
}

func buildMap(input string) []string {
	// build map
	return utils.SplitNewLines(input)
}

func findSymbol(partMap []string, part partNumber) ([2]int, rune) {
	// Convert part number to string
	// string_num := strconv.Itoa(part.part)
	string_num := int(math.Log10(float64(part.part))) + 1

	// Check each row in the vicinity of the part's starting position
	for dy := -1; dy <= 1; dy++ {
		y := part.posStart[1] + dy
		// Skip if y is out of bounds
		if y < 0 || y >= len(partMap) {
			// fmt.Println("y out of bounds")
			continue
		}

		// Check each character in the part number
		// for i := range string_num {
		for i := 0; i < string_num; i++ {
			// Check each column in the vicinity of the part's starting position
			for dx := -1; dx <= 1; dx++ {
				x := part.posStart[0] + dx + i
				// Skip if x is out of bounds
				if x < 0 || x >= len(partMap[y]) {
					// fmt.Println("x out of bounds")
					continue
				}

				// fmt.Println("checking", string_num, " specifically", string(char), "at", x, y, "which is", string(partMap[y][x]))
				// If the current position is not a dot and not a digit, return it
				if partMap[y][x] != '.' && !unicode.IsDigit(rune(partMap[y][x])) {
					return [2]int{x, y}, rune(partMap[y][x])
				}
			}
		}
	}
	// Return -1, -1 and 'N' if no symbol is found
	return [2]int{-1, -1}, 'N'
}

type partNumber struct {
	posStart  [2]int // x, y
	posEnd    [2]int // x, y
	part      int
	symbol    rune
	symbolPos [2]int // x, y
}
