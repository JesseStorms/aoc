package utils

import "strings"

func SplitNewLines(input string) []string {
	return strings.Split(input, "\n")
}
