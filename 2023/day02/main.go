package main

import (
	"aoc/utils"
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	// "sort"
)

//go:embed input.txt
var input string
var maxes = []int{12, 13, 14}

func main() {
	fmt.Println("part 1:", part1(input))
	fmt.Println("part 2:", part2(input))
}

func part1(input string) int {
	games := []Game{}
	for _, line := range utils.SplitNewLines(input) {
		game, c_cube := parseGame(line)
		game, valid := cubeCheck(game, c_cube)
		if valid {
			games = append(games, game)
		}
	}
	total := 0
	for _, game := range games {
		total += game.id
	}
	return total
}

func part2(input string) int {
	games := []Game{}
	for _, line := range utils.SplitNewLines(input) {
		game, c_cube := parseGame(line)
		game, _ = cubeCheck(game, c_cube)
		games = append(games, game)
	}

	total := 0
	for _, game := range games {
		power := game.cubes[0] * game.cubes[1] * game.cubes[2]
		total += power
	}
	return total
}

func parseGame(input string) (Game, [3]int) {
	game := Game{}
	c_cube := [3]int{0, 0, 0}
	reGame := regexp.MustCompile(`Game (\d*)`)
	reRed := regexp.MustCompile(`(\d*)( red)`)
	reGreen := regexp.MustCompile(`(\d*)( green)`)
	reBlue := regexp.MustCompile(`(\d*)( blue)`)
	reggroup := []*regexp.Regexp{reRed, reGreen, reBlue}

	game.id, _ = strconv.Atoi(reGame.FindStringSubmatch(input)[1])
	// fmt.Println(game.id)
	for i, reg := range reggroup {
		// find all matches
		matches := reg.FindAllStringSubmatch(input, -1)
		// get the highest number
		for _, match := range matches {
			c, _ := strconv.Atoi(match[1])
			if c > c_cube[i] {
				c_cube[i] = c
			}
		}
	}
	return game, c_cube
}

// game struct
type Game struct {
	id    int
	cubes [3]int
}

func cubeCheck(game Game, cubes [3]int) (Game, bool) {
	var valid = true
	for i, cube := range cubes {
		if cube <= maxes[i] {
			if cube > game.cubes[i] {
				game.cubes[i] = cube
			}
		} else {
			game.cubes[i] = cube
			valid = false
		}
	}
	return game, valid
}
