package main

import (
	"testing"
)

var i = `
vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
`

func TestDayThreePuzzleOne(t *testing.T) {
	r := "157"

	got := One(string(i))
	if got != r {
		t.Errorf("DayTwoPuzzleTwo() = %v, want %v", got, r)
	}
}

func TestDayThreePuzzleTwo(t *testing.T) {
	r := "70"

	got := Two(string(i))
	if got != r {
		t.Errorf("Two() = %v, want %v", got, r)
	}
}
