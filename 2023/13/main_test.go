package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = `
#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#
`

func TestPuzzleOne(t *testing.T) {
	assert.Equal(t, 405, One(input))
}

func TestPuzzleDebug(t *testing.T) {
	var input = `
#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.
`
	assert.Equal(t, 5, One(input))
}

func TestPuzzleTwo(t *testing.T) {
	assert.Equal(t, 400, Two(input))
}
