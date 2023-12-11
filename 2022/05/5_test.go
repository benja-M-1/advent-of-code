package main

import (
	"testing"
)

var i = `
    [D]
[N] [C]
[Z] [M] [P]
 1   2   3

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`

func TestDayFivePuzzleOne(t *testing.T) {
	r := "CMZ"

	got := One(string(i))
	if got != r {
		t.Errorf("One() = %v, want %v", got, r)
	}
}

func TestDayFivePuzzleTwo(t *testing.T) {
	r := "MCD"

	got := Two(string(i))
	if got != r {
		t.Errorf("Two() = %v, want %v", got, r)
	}
}
