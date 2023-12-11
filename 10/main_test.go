package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPuzzleOne(t *testing.T) {
	tests := []struct {
		input string
		steps int
	}{
		{

			input: `
.....
.S-7.
.|.|.
.L-J.
.....
		`,
			steps: 4,
		},
		{
			input: `
..F7.
.FJ|.
SJ.L7
|F--J
LJ...
`,
			steps: 8,
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			r := One(test.input)
			assert.Equal(t, test.steps, r)
		})
	}
}

func TestPuzzleTwo(t *testing.T) {
	tests := []struct {
		input string
		tiles int
	}{
		{
			input: `
...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........
`,
			tiles: 4,
		},
		{
			input: `
..........
.S------7.
.|F----7|.
.||....||.
.||....||.
.|L-7F-J|.
.|..||..|.
.L--JL--J.
..........
		`,
			tiles: 4,
		},
		{
			input: `
.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...
		`,
			tiles: 8,
		},
		// 		{
		// 			input: `
		// FF7FSF7F7F7F7F7F---7
		// L|LJ||||||||||||F--J
		// FL-7LJLJ||||||LJL-77
		// F--JF--7||LJLJ7F7FJ-
		// L---JF-JLJ.||-FJLJJ7
		// |F|F-JF---7F7-L7L|7|
		// |FFJF7L7F-JF7|JL---7
		// 7-L-JL7||F7|L7F-7F7|
		// L.L7LFJ|||||FJL7||LJ
		// L7JLJL-JLJLJL--JLJ.L
		// `,
		// 			tiles: 10,
		// 		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			r := Two(test.input)
			assert.Equal(t, test.tiles, r)
		})
	}
}

func TestPuzzleTwoSecond(t *testing.T) {
	tests := []struct {
		input string
		tiles int
	}{
		{
			input: `
...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........
`,
			tiles: 4,
		},
		{
			input: `
..........
.S------7.
.|F----7|.
.||....||.
.||....||.
.|L-7F-J|.
.|..||..|.
.L--JL--J.
..........
		`,
			tiles: 4,
		},
		{
			input: `
.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...
		`,
			tiles: 8,
		},
		// 		{
		// 			input: `
		// FF7FSF7F7F7F7F7F---7
		// L|LJ||||||||||||F--J
		// FL-7LJLJ||||||LJL-77
		// F--JF--7||LJLJ7F7FJ-
		// L---JF-JLJ.||-FJLJJ7
		// |F|F-JF---7F7-L7L|7|
		// |FFJF7L7F-JF7|JL---7
		// 7-L-JL7||F7|L7F-7F7|
		// L.L7LFJ|||||FJL7||LJ
		// L7JLJL-JLJLJL--JLJ.L
		// 		`,
		// 			tiles: 10,
		// 		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			r := TwoSecond(test.input)
			assert.Equal(t, test.tiles, r)
		})
	}
}
