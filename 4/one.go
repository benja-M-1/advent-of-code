package main

import (
	"slices"
	"strings"
)

func PuzzleOne(input string) int {
	points := 0
	for _, card := range strings.Split(strings.Trim(input, "\n"), "\n") {
		fields := strings.Fields(card)
		numbers := []string{}
		isWinningNumber := false
		cardPoints := 0
		// Skip the first 2 fields as they are "Card X:"
		for _, field := range fields[2:] {
			if field == "|" {
				isWinningNumber = true
				continue
			}

			if !isWinningNumber {
				numbers = append(numbers, field)
				continue
			}

			if slices.Contains(numbers, field) {
				if cardPoints == 0 {
					cardPoints = 1
				} else {
					cardPoints *= 2
				}
			}
		}

		points += cardPoints
	}

	return points
}
