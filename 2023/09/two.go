package main

import (
	"strconv"
	"strings"
)

func PuzzleTwo(input string) int {
	input = strings.Trim(input, "\n")

	total := 0
	for _, line := range strings.Split(input, "\n") {
		history := []int{}
		// Convert the slice of strings into a slice of ints
		for _, c := range strings.Fields(line) {
			n, _ := strconv.Atoi(c)
			history = append(history, n)
		}

		n := extrapolateBackward(history, []int{})
		total += history[0] - n
	}

	return total
}

func extrapolateBackward(history []int, next []int) int {
	i := len(history) - 1
	diff := history[i] - history[i-1]
	next = append([]int{diff}, next...)
	remaining := history[:i]

	if len(remaining) > 1 {
		return extrapolateBackward(remaining, next)
	}

	if allZeroes(next) {
		return 0
	}

	n := extrapolateBackward(next, []int{})

	return next[0] - n
}
