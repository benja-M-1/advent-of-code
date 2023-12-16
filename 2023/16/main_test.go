package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = `
.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....
`

func TestPuzzleOne(t *testing.T) {
	assert.Equal(t, 46, One(input))
}

func TestPuzzleTwo(t *testing.T) {
	assert.Equal(t, 51, Two(input))
}
