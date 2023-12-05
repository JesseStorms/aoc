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
	fmt.Println("part 1:", part1(input))
	// fmt.Println("part 2:", part2(input))
}

func part1(input string) int {
	regions := utils.GetRegions(input)
	seeds := parseSeeds(regions[0][0])
	maps := []Routemap{}
	for _, region := range regions[1:] {
		maps = append(maps, parseMap(region))
	}
	result := []Step{}
	for _, seed := range seeds {
		// solve(seed, maps, Location)
		result = append(result, solve(seed, maps, Location))
	}
	min := result[0].Value
	// get the lowest value in result
	for _, step := range result {
		// get the lowest value in result
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
func solve(seed Step, maps []Routemap, stage Stage) Step {
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
