package utils

import (
	"reflect"
	"testing"
)

func TestBuildGrid(t *testing.T) {
	input := "abc\ndef\nghi"
	expected := Grid{
		Runes: [][]rune{
			{'a', 'b', 'c'},
			{'d', 'e', 'f'},
			{'g', 'h', 'i'},
		},
	}
	result := BuildGrid(input)
	if len(result.Runes) != len(expected.Runes) {
		t.Errorf("Expected grid with %d rows, got %d rows", len(expected.Runes), len(result.Runes))
	}
	for i := 0; i < len(result.Runes); i++ {
		if len(result.Runes[i]) != len(expected.Runes[i]) {
			t.Errorf("Expected row with %d columns, got %d columns", len(expected.Runes[i]), len(result.Runes[i]))
		}
		for j := 0; j < len(result.Runes[i]); j++ {
			if result.Runes[i][j] != expected.Runes[i][j] {
				t.Errorf("Expected %c, got %c", expected.Runes[i][j], result.Runes[i][j])
			}
		}
	}
}

func TestGridAt(t *testing.T) {
	grid := Grid{
		Runes: [][]rune{
			{'a', 'b', 'c'},
			{'d', 'e', 'f'},
			{'g', 'h', 'i'},
		},
	}
	tests := []struct {
		x, y     int
		expected rune
	}{
		{0, 0, 'a'},
		{1, 1, 'e'},
		{2, 2, 'i'},
	}
	for _, test := range tests {
		result := grid.At(test.x, test.y)
		if result != test.expected {
			t.Errorf("Expected %c, got %c", test.expected, result)
		}
	}
}

func TestGridSlice(t *testing.T) {
	grid := Grid{
		Runes: [][]rune{
			{'a', 'b', 'c'},
			{'d', 'e', 'f'},
			{'g', 'h', 'i'},
		},
	}
	tests := []struct {
		x, y, width, height int
		expected            [][]rune
	}{
		{0, 0, 2, 2, [][]rune{
			{'a', 'b'},
			{'d', 'e'},
		}},
		{1, 1, 3, 2, [][]rune{
			{'e', 'f'},
			{'h', 'i'},
		}},
		{2, 2, 1, 1, [][]rune{
			{'i'},
		}},
	}
	for _, test := range tests {
		result := grid.Slice(test.x, test.y, test.width, test.height)
		if len(result) != len(test.expected) {
			t.Errorf("Expected slice with %d rows, got %d rows", len(test.expected), len(result))
		}
		for i := 0; i < len(result); i++ {
			if len(result[i]) != len(test.expected[i]) {
				t.Errorf("Expected row with %d columns, got %d columns", len(test.expected[i]), len(result[i]))
			}
			for j := 0; j < len(result[i]); j++ {
				if result[i][j] != test.expected[i][j] {
					t.Errorf("Expected %c, got %c", test.expected[i][j], result[i][j])
				}
			}
		}
	}
}

func TestGridSquare(t *testing.T) {
	grid := Grid{
		Runes: [][]rune{
			{'a', 'b', 'c'},
			{'d', 'e', 'f'},
			{'g', 'h', 'i'},
		},
	}
	tests := []struct {
		x, y, r  int
		expected [][]rune
	}{
		{1, 1, 1, [][]rune{
			{'a', 'b', 'c'},
			{'d', 'e', 'f'},
			{'g', 'h', 'i'},
		}},
		{0, 0, 1, [][]rune{
			{'a', 'b'},
			{'d', 'e'},
		}},
	}
	for _, test := range tests {
		result := grid.Square(test.x, test.y, test.r)
		if len(result) != len(test.expected) {
			t.Errorf("Expected square with %d rows, got %d rows", len(test.expected), len(result))
		}
		for i := 0; i < len(result); i++ {
			if len(result[i]) != len(test.expected[i]) {
				t.Errorf("Expected row with %d columns, got %d columns", len(test.expected[i]), len(result[i]))
			}
			for j := 0; j < len(result[i]); j++ {
				if result[i][j] != test.expected[i][j] {
					t.Errorf("Expected %c, got %c", test.expected[i][j], result[i][j])
				}
			}
		}
	}
}
func TestGridSet(t *testing.T) {
	tests := []struct {
		name     string
		x, y     int
		char     rune
		expected Grid
	}{
		{
			name: "Set top left corner",
			x:    0,
			y:    0,
			char: 'x',
			expected: Grid{
				Runes: [][]rune{
					{'x', 'b', 'c'},
					{'d', 'e', 'f'},
					{'g', 'h', 'i'},
				},
			},
		},
		{
			name: "Set middle",
			x:    1,
			y:    1,
			char: 'x',
			expected: Grid{
				Runes: [][]rune{
					{'a', 'b', 'c'},
					{'d', 'x', 'f'},
					{'g', 'h', 'i'},
				},
			},
		},
		{
			name: "Set bottom right corner",
			x:    2,
			y:    2,
			char: 'x',
			expected: Grid{
				Runes: [][]rune{
					{'a', 'b', 'c'},
					{'d', 'e', 'f'},
					{'g', 'h', 'x'},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			grid := Grid{
				Runes: [][]rune{
					{'a', 'b', 'c'},
					{'d', 'e', 'f'},
					{'g', 'h', 'i'},
				},
			}

			grid.Set(test.x, test.y, test.char)

			if !reflect.DeepEqual(grid, test.expected) {
				t.Errorf("Expected grid %v, got %v", test.expected, grid)
			}
		})
	}
}
