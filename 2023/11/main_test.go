package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = `
...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....
`

func TestPuzzleOne(t *testing.T) {
	assert.Equal(t, 374, One(input))
}

func TestPuzzleTwo(t *testing.T) {
	assert.Equal(t, 1030, Two(input, 10))
	assert.Equal(t, 8410, Two(input, 100))
}
