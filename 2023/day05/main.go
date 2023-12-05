package main

import (
	"aoc/utils"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	// "sort"
)

//go:embed input.txt
var input string

func main() {
	// fmt.Println("part 1:", part1(input))
	fmt.Println("part 2:", part2(input))
}

func part1(input string) int {
	regions := utils.GetRegions(input)
	seeds := parseSeeds(regions[0][0])
	maps := []Routemap{}
	for _, region := range regions[1:] {
		maps = append(maps, parseMap(region))
	}

	return solve(seeds, maps)
}

func part2(input string) int {
	regions := utils.GetRegions(input)
	seeds := parseSeeds2(regions[0][0])
	maps := []Routemap{}
	for _, region := range regions[1:] {
		maps = append(maps, parseMap(region))
	}

	return solve(seeds, maps)
}

func solve(seeds []Step, maps []Routemap) int {
	result := []Step{}
	results := make(chan Step, len(seeds))
	for _, seed := range seeds {
		go func(seed Step) {
			results <- solveSeed(seed, maps, Location)
		}(seed)
	}

	for i := 0; i < len(seeds); i++ {
		result = append(result, <-results)
	}
	close(results)

	min := result[0].Value

	for _, step := range result {

		if step.Value < min {
			min = step.Value
		}
	}
	return min
}

/*
*
Run through the seeds, and for each seed, run through the maps until the step's type matches the stage's type.
*/
func solveSeed(seed Step, maps []Routemap, stage Stage) Step {
	for _, routemap := range maps {
		seed = routemap.Solve(seed)
	}
	return seed
}

func parseSeeds(input string) []Step {
	var result []Step
	for _, part := range strings.Split(input, " ")[1:] {
		value, _ := strconv.Atoi(part)
		result = append(result, Step{Value: value, Type: Seed})
	}
	return result
}

/*
the first number is the seed, the second is the length of the seed.
example: 3 2 means the seed is 3, and the length is 2.
this means that the seed is 3, 4, 5.
*/
func parseSeeds2(input string) []Step {
	var result []Step
	var values []int
	for _, part := range strings.Split(input, " ")[1:] {
		value, _ := strconv.Atoi(part)
		values = append(values, value)
	}
	for i := 0; i < len(values); i += 2 {
		result = append(result, Step{Value: values[i], Type: Seed})
		for j := 1; j < values[i+1]; j++ {
			result = append(result, Step{Value: values[i] + j, Type: Seed})
		}
	}
	return result
}

func parseMap(input []string) Routemap {
	var result Routemap
	// header parsed here
	mapType := StageMap[strings.Split(strings.Split(input[0], " ")[0], "-")[2]]
	result.Type = mapType
	for _, line := range input[1:] {
		// 50 98 2
		var route Route
		parts := strings.Split(line, " ")
		route.DestStart, _ = strconv.Atoi(parts[0])
		route.SourceStart, _ = strconv.Atoi(parts[1])
		route.Length, _ = strconv.Atoi(parts[2])
		route.Type = mapType
		result.Entries = append(result.Entries, route)
	}
	return result
}
