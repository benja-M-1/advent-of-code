package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = `
O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....
`

func TestPuzzleOne(t *testing.T) {
	assert.Equal(t, 136, One(input))
}

func TestPuzzleTwo(t *testing.T) {
	assert.Equal(t, 64, Two(input))
}

func Test_tilt(t *testing.T) {
	var expected = parse(`
OOOO.#.O..
OO..#....#
OO..O##..O
O..#.OO...
........#.
..#....#.#
..O..#.O.O
..O.......
#....###..
#....#....
`)

	assert.Equal(t, expected, tilt(parse(input)))
}

func Test_Cycle(t *testing.T) {
	var expected1 = parse(`
.....#....
....#...O#
...OO##...
.OO#......
.....OOO#.
.O#...O#.#
....O#....
......OOOO
#...O###..
#..OO#....
`)
	var expected2 = parse(`
.....#....
....#...O#
.....##...
..O#......
.....OOO#.
.O#...O#.#
....O#...O
.......OOO
#..OO###..
#.OOO#...O
`)
	var expected3 = parse(`
.....#....
....#...O#
.....##...
..O#......
.....OOO#.
.O#...O#.#
....O#...O
.......OOO
#...O###.O
#.OOO#...O
`)

	cycle1 := cycle(parse(input))
	assert.Equal(t, expected1, cycle1)
	cycle2 := cycle(cycle1)
	assert.Equal(t, expected2, cycle2)
	cycle3 := cycle(cycle2)
	assert.Equal(t, expected3, cycle3)
}

func Test_rotate(t *testing.T) {
	var input = [][]string{
		{"00", "01", "02", "03"},
		{"10", "11", "12", "13"},
		{"20", "21", "22", "23"},
	}

	var rotated = [][]string{
		{"20", "10", "00"},
		{"21", "11", "01"},
		{"22", "12", "02"},
		{"23", "13", "03"},
	}

	assert.Equal(t, rotated, rotate(input))
}
