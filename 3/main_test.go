package main

import (
	"testing"
)

var input = `
467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

func TestPuzzleOne(t *testing.T) {
	expected := 4361
	got := Puzzle(input)

	if got != expected {
		t.Errorf("Puzzle() = %v, want %v", got, expected)
	}
}

func TestPuzzleTow(t *testing.T) {
	expected := 467835
	got := PuzzleTwo(input)

	if got != expected {
		t.Errorf("PuzzleTwo() = %v, want %v", got, expected)
	}
}

func TestPuzzleTow2(t *testing.T) {
	var input = `
...253
..*...
...876
`
	expected := 253 * 876
	got := PuzzleTwo(input)

	if got != expected {
		t.Errorf("PuzzleTwo() = %v, want %v", got, expected)
	}
}

func TestPuzzleTow3(t *testing.T) {
	var input = `
.714.840.
....*....
......595
`
	expected := 714 * 840
	got := PuzzleTwo(input)

	if got != expected {
		t.Errorf("PuzzleTwo() = %v, want %v", got, expected)
	}
}
