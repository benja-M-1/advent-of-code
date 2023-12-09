package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = `
0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45
`

func TestPuzzleOne(t *testing.T) {
	r := PuzzleOne(input)
	assert.Equal(t, 114, r)
}

func TestPuzzleTwo(t *testing.T) {
	r := PuzzleTwo(input)
	assert.Equal(t, 2, r)
}
