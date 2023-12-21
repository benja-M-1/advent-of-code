package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = `
...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........
`

func TestPuzzle(t *testing.T) {
	tests := []struct {
		plots int
		steps int
	}{
		{16, 6},
		// {50, 10},
		// {1594, 50},
		// {6536, 100},
		// {167004, 500},
		// {668697, 1000},
		// {16733044, 5000},
	}

	for _, test := range tests {
		assert.Equal(t, test.plots, Puzzle(input, test.steps))
	}
}

func Test_gridExpand(t *testing.T) {
	var expected = `
.................................
.....###.#......###.#......###.#.
.###.##..#..###.##..#..###.##..#.
..#.#...#....#.#...#....#.#...#..
....#.#........#.#........#.#....
.##...####..##...####..##...####.
.##..#...#..##..#...#..##..#...#.
.......##.........##.........##..
.##.#.####..##.#.####..##.#.####.
.##..##.##..##..##.##..##..##.##.
.................................
.................................
.....###.#......###.#......###.#.
.###.##..#..###.##..#..###.##..#.
..#.#...#....#.#...#....#.#...#..
....#.#........#.#........#.#....
.##...####..##..S####..##...####.
.##..#...#..##..#...#..##..#...#.
.......##.........##.........##..
.##.#.####..##.#.####..##.#.####.
.##..##.##..##..##.##..##..##.##.
.................................
.................................
.....###.#......###.#......###.#.
.###.##..#..###.##..#..###.##..#.
..#.#...#....#.#...#....#.#...#..
....#.#........#.#........#.#....
.##...####..##...####..##...####.
.##..#...#..##..#...#..##..#...#.
.......##.........##.........##..
.##.#.####..##.#.####..##.#.####.
.##..##.##..##..##.##..##..##.##.
.................................
`
	grid, _ := parse(input)
	grid.Expand()
	assert.Equal(t, strings.Trim(expected, "\n"), strings.Trim(grid.String(), "\n"))
}
