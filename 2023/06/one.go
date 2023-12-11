package main

import (
	"strconv"
	"strings"
)

func PuzzleOne(input string) int {
	input = strings.Trim(input, "\n")
	lines := strings.Split(input, "\n")

	times := strings.Fields(strings.TrimPrefix(lines[0], "Time: "))
	distances := strings.Fields(strings.TrimPrefix(lines[1], "Distance: "))
	total := 0

	for race := 0; race < len(times); race++ {
		pos := 0
		time, _ := strconv.Atoi(times[race])
		dist, _ := strconv.Atoi(distances[race])

		for s, hold := time, 0; s > 0; s, hold = s-1, hold+1 {
			d := hold * (time - hold)
			if d > dist {
				pos++
			}
		}

		if pos > 0 {
			total *= pos
			if total == 0 {
				total = pos
			}
		}
	}

	return total
}
