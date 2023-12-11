package main

import (
	"testing"
)

var i = `
2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8
`

func TestDayFourPuzzleOne(t *testing.T) {
	r := "2"

	got := One(string(i))
	if got != r {
		t.Errorf("One() = %v, want %v", got, r)
	}
}
func TestDayFourPuzzleTwo(t *testing.T) {
	r := "4"

	got := Two(string(i))
	if got != r {
		t.Errorf("Two() = %v, want %v", got, r)
	}
}
