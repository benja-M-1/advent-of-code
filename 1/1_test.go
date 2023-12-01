package main

import (
	"testing"
)

func TestDayOnePuzzleOne(t *testing.T) {
	input := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
`
	expected := 142

	got := DayOne(input)
	if got != expected {
		t.Errorf("DayOne() = %v, want %v", got, expected)
	}
}

func TestDayOnePuzzleTwo(t *testing.T) {
	input := `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
`
	expected := 281

	got := DayOne(input)
	if got != expected {
		t.Errorf("DayOnePuzzleTwo() = %v, want %v", got, expected)
	}
}
