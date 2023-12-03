package utils

import (
	"testing"
)

func TestSplitNewLines(t *testing.T) {
	input := "one\ntwo\nthree"
	expected := []string{"one", "two", "three"}
	result := SplitNewLines(input)
	for i, line := range result {
		if line != expected[i] {
			t.Errorf("Expected %s, got %s", expected[i], line)
		}
	}
}
