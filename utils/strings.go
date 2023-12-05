package utils

import (
	"bufio"
	"strings"
)

func SplitNewLines(input string) []string {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

/*
*
Split every two newlines, which would be a blank line that usually indicates another part of data
the entire region is returned as a string.
*/
func SplitRegions(input string) []string {
	input = strings.ReplaceAll(input, "\r\n\r\n", "\n\n")
	return strings.Split(input, "\n\n")
}

/*
*
this will return a slice of slices. Each slice is a region, and each region is a slice of strings
*/
func GetRegions(input string) [][]string {
	regions := SplitRegions(input)
	var result [][]string
	for _, region := range regions {
		result = append(result, SplitNewLines(region))
	}
	return result
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
