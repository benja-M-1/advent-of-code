package main

import (
	"testing"
)

var i = `
1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
`

func TestDayOnePuzzleOne(t *testing.T) {
	r := "24000"

	got := One(string(i))
	if got != r {
		t.Errorf("One() = %v, want %v", got, r)
	}
}

func TestDayOnePuzzleTwo(t *testing.T) {
	r := "45000"

	got := Two(string(i))
	if got != r {
		t.Errorf("Two() = %v, want %v", got, r)
	}
}
