package utils

import (
	"testing"
)

func TestSplitNewLinesLinux(t *testing.T) {
	input := "one\ntwo\nthree"
	expected := []string{"one", "two", "three"}
	result := SplitNewLines(input)
	for i, line := range result {
		if line != expected[i] {
			t.Errorf("Expected %s, got %s", expected[i], line)
		}
	}
}

func TestSplitNewLinesWindows(t *testing.T) {
	input := "one\r\ntwo\r\nthree"
	expected := []string{"one", "two", "three"}
	result := SplitNewLines(input)
	for i, line := range result {
		if line != expected[i] {
			t.Errorf("Expected %s, got %s", expected[i], line)
		}
	}
}

func TestReverse(t *testing.T) {
	input := "hello"
	expected := "olleh"
	result := Reverse(input)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestReverseEmptyString(t *testing.T) {
	input := ""
	expected := ""
	result := Reverse(input)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestReverseSingleCharacter(t *testing.T) {
	input := "a"
	expected := "a"
	result := Reverse(input)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestReversePalindrome(t *testing.T) {
	input := "racecar"
	expected := "racecar"
	result := Reverse(input)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
func TestSplitRegions(t *testing.T) {
	input := "one\ntwo\n\nthree\n\nfour"
	expected := []string{"one\ntwo", "three", "four"}
	result := SplitRegions(input)
	if len(result) != len(expected) {
		t.Errorf("Expected %d regions, got %d", len(expected), len(result))
	}
	for i, region := range result {
		if region != expected[i] {
			t.Errorf("Expected region %d to be %s, got %s", i, expected[i], region)
		}
	}
}

func TestGetRegions(t *testing.T) {
	input := "one\ntwo\n\nthree\n\nfour"
	expected := [][]string{{"one", "two"}, {"three"}, {"four"}}
	result := GetRegions(input)
	if len(result) != len(expected) {
		t.Errorf("Expected %d regions, got %d", len(expected), len(result))
	}
	for i, region := range result {
		if len(region) != len(expected[i]) {
			t.Errorf("Expected region %d to have %d lines, got %d", i, len(expected[i]), len(region))
		}
		for j, line := range region {
			if line != expected[i][j] {
				t.Errorf("Expected line %d in region %d to be %s, got %s", j, i, expected[i][j], line)
			}
		}
	}
}
