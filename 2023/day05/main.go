package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("part 1:", part1(input))
	fmt.Println("part 2:", part2(input))
}

func part1(input string) int {
	lines := strings.Split(input, "\n")
	seeds := parseSeedsPart1(lines[0])
	maps := parseMaps(lines[1:])

	lowestLocation := -1

	for _, seed := range seeds {
		location := mapSeedToLocation(seed, maps)
		if lowestLocation == -1 || location < lowestLocation {
			lowestLocation = location
		}
	}

	return lowestLocation
}

func part2(input string) int {
	lines := strings.Split(input, "\n")
	seeds := parseSeeds(lines[0])
	maps := parseMaps(lines[1:])

	numWorkers := 4
	seedsPerWorker := len(seeds) / numWorkers
	results := make(chan int, numWorkers)
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		start := i * seedsPerWorker
		end := start + seedsPerWorker
		if i == numWorkers-1 {
			end = len(seeds)
		}
		wg.Add(1)
		go func(seeds []int) {
			defer wg.Done()
			worker(seeds, maps, results)
		}(seeds[start:end])
	}

	wg.Wait()
	close(results)

	lowestLocation := -1
	for location := range results {
		if lowestLocation == -1 || location < lowestLocation {
			lowestLocation = location
		}
	}

	return lowestLocation
}

func worker(seeds []int, maps map[string][][]int, results chan<- int) {
	lowestLocation := -1
	for _, seed := range seeds {
		location := mapSeedToLocation(seed, maps)
		if lowestLocation == -1 || location < lowestLocation {
			lowestLocation = location
		}
	}
	results <- lowestLocation
}

func parseSeedsPart1(line string) []int {
	parts := strings.Split(line, " ")[1:]
	var seeds []int
	for _, part := range parts {
		seed, _ := strconv.Atoi(part)
		seeds = append(seeds, seed)
	}
	return seeds
}

func parseSeeds(line string) []int {
	parts := strings.Split(line, " ")[1:]
	var seeds []int
	for i := 0; i < len(parts); i += 2 {
		start, _ := strconv.Atoi(parts[i])
		length, _ := strconv.Atoi(parts[i+1])
		for j := 0; j < length; j++ {
			seeds = append(seeds, start+j)
		}
	}
	return seeds
}

func parseMaps(lines []string) map[string][][]int {
	maps := make(map[string][][]int)
	var currentMap string
	for _, line := range lines {
		if strings.Contains(line, "map:") {
			currentMap = strings.TrimSuffix(line, " map:")
			maps[currentMap] = [][]int{}
		} else if line != "" {
			parts := strings.Fields(line)
			startDest, _ := strconv.Atoi(parts[0])
			startSrc, _ := strconv.Atoi(parts[1])
			length, _ := strconv.Atoi(parts[2])
			maps[currentMap] = append(maps[currentMap], []int{startDest, startSrc, length})
		}
	}
	return maps
}

func mapSeedToLocation(seed int, maps map[string][][]int) int {
	categories := []string{"seed-to-soil", "soil-to-fertilizer", "fertilizer-to-water", "water-to-light", "light-to-temperature", "temperature-to-humidity", "humidity-to-location"}
	value := seed
	for _, category := range categories {
		value = mapValue(value, maps[category])
	}
	return value
}

func mapValue(value int, mappings [][]int) int {
	for _, mapping := range mappings {
		startDest, startSrc, length := mapping[0], mapping[1], mapping[2]
		if value >= startSrc && value < startSrc+length {
			return startDest + (value - startSrc)
		}
	}
	return value
}
