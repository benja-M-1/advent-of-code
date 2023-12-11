package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = `
32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
`

func TestPuzzleOne(t *testing.T) {
	r := PuzzleOne(input)
	assert.Equal(t, 6440, r)
}

func TestPuzzleTwo(t *testing.T) {
	r := PuzzleTwo(input)
	assert.Equal(t, 5905, r)
}
