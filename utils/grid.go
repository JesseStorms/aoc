package utils

// this handles walking through a grid of chars, will probably be useful for a lot of problems

type Grid struct {
	Runes [][]rune
}

func BuildGrid(input string) Grid {
	lines := SplitNewLines(input)
	grid := [][]rune{}
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}
	return Grid{Runes: grid}
}

func (g Grid) At(x, y int) rune {
	return g.Runes[y][x]
}

func (g Grid) Set(x, y int, char rune) {
	g.Runes[y][x] = char
}

/**
 * Returns a slice of the grid, starting at x,y and going width,height
 * Elements that are out of bounds are ignored
 */
func (g Grid) Slice(x, y, width, height int) [][]rune {
	slice := [][]rune{}
	for i := y; i < y+height; i++ {
		if i < 0 || i >= len(g.Runes) {
			continue
		}
		row := []rune{}
		for j := x; j < x+width; j++ {
			if j < 0 || j >= len(g.Runes[i]) {
				continue
			}
			row = append(row, g.Runes[i][j])
		}
		slice = append(slice, row)
	}
	return slice
}

/**
Returns a square around x,y with radius r
Elements that are out of bounds are ignored
Take this grid for example
a b c
d e f
g h i
Square(1, 1, 1) would return
a b c
d e f
g h i
Square (0, 0, 1) would return
a b
d e
Out of bounds elements are dropped
 **/
func (g Grid) Square(x, y, r int) [][]rune {
	return g.Slice(x-r, y-r, 2*r+1, 2*r+1)
}
