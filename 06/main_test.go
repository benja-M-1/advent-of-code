package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = `
Time:      7  15   30
Distance:  9  40  200
`

func TestPuzzleOne(t *testing.T) {
	r := PuzzleOne(input)
	assert.Equal(t, 288, r)
}

func TestPuzzleTwo(t *testing.T) {
	r := PuzzleTwo(input)
	assert.Equal(t, 71503, r)
}
