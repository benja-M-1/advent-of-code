package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPuzzleOne(t *testing.T) {
	var input = `
RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)
`

	r := PuzzleOne(input)
	assert.Equal(t, 2, r)

	input = `
LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`

	r2 := PuzzleOne(input)
	assert.Equal(t, 6, r2)
}

func TestPuzzleTwo(t *testing.T) {
	var input = `
LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)
`

	r := PuzzleTwo(input)
	assert.Equal(t, 6, r)
}
