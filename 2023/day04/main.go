package main

import (
	"aoc/utils"
	_ "embed"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
	// "sort"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("part 1:", part1(input))
	fmt.Println("part 2:", part2(input))
}

func part1(input string) int {
	cards := parseInput(input)
	// fmt.Println(cards)
	total := 0
	for _, card := range cards {
		total += card.value()
	}
	return total
}

func part2(input string) int {
	cards := parseInput(input)
	// Part 2
	for i, c := range cards {
		for k := 0; k < c.Instances; k++ {
			hits := 0
			for _, n := range c.playing_numbers {
				if slices.Contains(c.winning_numbers, n) {
					hits += 1
				}
			}

			for j := i + 1; j <= i+hits; j++ {
				cards[j].Instances++
			}
		}
	}

	totalCards := 0
	for _, c := range cards {
		totalCards += c.Instances
	}
	return totalCards
}

func parseInput(input string) []card {
	lines := utils.SplitNewLines(input)
	cards := []card{}
	for _, line := range lines {
		cards = append(cards, parseCard(line))
	}
	return cards
}

func parseCard(line string) card {
	reCard := regexp.MustCompile(`(Card +\d+)`)
	reNumbers := regexp.MustCompile(`(\d*)`)

	card := card{}
	card.id, _ = strconv.Atoi(reCard.FindStringSubmatch(line)[1])

	// split on |
	nums := strings.Split(line, "|")
	nums[0] = reCard.ReplaceAllString(nums[0], "")

	for _, n := range reNumbers.FindAllStringSubmatch(nums[0], -1) {
		num, _ := strconv.Atoi(n[1])
		if num != 0 {
			card.winning_numbers = append(card.winning_numbers, num)

		}
	}

	for _, n := range reNumbers.FindAllStringSubmatch(nums[1], -1) {
		num, _ := strconv.Atoi(n[1])
		if num != 0 {
			card.playing_numbers = append(card.playing_numbers, num)
		}
	}
	card.Instances = 1
	return card
}

type card struct {
	id              int
	winning_numbers []int
	playing_numbers []int
	Instances       int
}

func (c card) winning_matches() int {
	total := 0
	for _, n := range c.winning_numbers {
		for _, m := range c.playing_numbers {
			if n == m {
				// fmt.Println(n, m, total)
				total++
			}
		}
	}

	return total
}

func (c card) value() int {
	return double(1, card.winning_matches(c))

}

// double the input n times (1->2->4->8...)
func double(input int, n int) int {
	if n == 0 {
		return 0
	}
	for i := 1; i < n; i++ {
		// fmt.Println(input)
		input = input * 2
	}
	return input
}
