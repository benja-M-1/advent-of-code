package main

import (
	"strconv"
	"strings"
)

func PuzzleTwo(input string) int {
	input = strings.Trim(input, "\n")
	lines := strings.Split(input, "\n")

	replacer := strings.NewReplacer(" ", "", "Time: ", "", "Distance: ", "")
	time, _ := strconv.Atoi(replacer.Replace(lines[0]))
	dist, _ := strconv.Atoi(replacer.Replace(lines[1]))
	total := 0

	for s, hold := time, 0; s > 0; s, hold = s-1, hold+1 {
		d := hold * (time - hold)
		if d > dist {
			total++
		}
	}

	return total
}
