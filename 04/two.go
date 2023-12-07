package main

import (
	"slices"
	"strconv"
	"strings"
)

func PuzzleTwo(input string) int {
	instances := map[int]int{}
	for _, card := range strings.Split(strings.Trim(input, "\n"), "\n") {
		fields := strings.Fields(card)
		numbers := []string{}
		isWinningNumber := false
		cardPoints := 0
		cardId, _ := strconv.Atoi(strings.Replace(fields[1], ":", "", 1))
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
				cardPoints++
				instances[cardId+cardPoints] = instances[cardId+cardPoints] + 1

				cardCopies := instances[cardId]
				for ; 0 < cardCopies; cardCopies-- {
					instances[cardId+cardPoints] = instances[cardId+cardPoints] + 1
				}
			}

		}

		instances[cardId] = instances[cardId] + 1
	}

	points := 0
	for _, v := range instances {
		points += v
	}

	return points
}
