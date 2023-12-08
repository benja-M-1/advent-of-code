package main

import (
	"sort"
	"strconv"
	"strings"
)

func PuzzleTwo(input string) int {
	input = strings.Trim(input, "\n")
	plays := ByWildcard{}

	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " ")
		hand := parts[0]
		bid, _ := strconv.Atoi(parts[1])

		plays = append(plays, NewWildcardPlay(hand, bid))
	}

	sort.Sort(plays)

	total := 0
	for rank, play := range plays {
		total += (rank + 1) * play.Bid
	}

	return total
}
