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
