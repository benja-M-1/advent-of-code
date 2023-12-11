package main

import (
	"testing"
)

var input = `
A Y
B X
C Z
`

func TestDayTwoPuzzleOne(t *testing.T) {
	r := "15"

	got := One(input)
	if got != r {
		t.Errorf("DayTwoPuzzleOne() = %v, want %v", got, r)
	}
}

func TestDayTwoPuzzleTwo(t *testing.T) {
	r := "12"

	got := Two(input)
	if got != r {
		t.Errorf("DayTwoPuzzleTwo() = %v, want %v", got, r)
	}
}
