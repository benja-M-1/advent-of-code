package main

import (
	"strconv"
	"strings"
)

func PuzzleOne(input string) int {
	input = strings.Trim(input, "\n")

	total := 0
	for _, line := range strings.Split(input, "\n") {
		history := []int{}
		// Convert the slice of strings into a slice of ints
		for _, c := range strings.Fields(line) {
			n, _ := strconv.Atoi(c)
			history = append(history, n)
		}

		n := extrapolate(history, []int{})
		total += n + history[len(history)-1]
	}

	return total
}

func extrapolate(history []int, next []int) int {
	i := len(history) - 1
	diff := history[i] - history[i-1]
	next = append([]int{diff}, next...)
	remaining := history[:i]

	if len(remaining) > 1 {
		return extrapolate(remaining, next)
	}

	if allZeroes(next) {
		return 0
	}

	n := extrapolate(next, []int{})

	return next[len(next)-1] + n
}

func allZeroes(s []int) bool {
	zeroes := 0
	for _, n := range s {
		if n == 0 {
			zeroes++
		}
	}
	return len(s) == zeroes
}
