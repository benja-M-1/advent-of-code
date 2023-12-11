package main

import (
	"strings"
)

func PuzzleOne(input string) int {
	input = strings.Trim(input, "\n")

	lines := strings.Split(input, "\n")
	directions := lines[0]
	nodes := map[string]map[string]string{}
	for _, line := range lines[2:] {
		replacer := strings.NewReplacer("(", "", ")", "", "=", "", ",", "")
		parts := strings.Fields(replacer.Replace(line))
		nodes[parts[0]] = map[string]string{
			"L": parts[1],
			"R": parts[2],
		}
	}

	var (
		index int
		steps int
	)
	next := "AAA"
	for next != "ZZZ" {
		if steps%len(directions) == 0 {
			index = 0
		}

		direction := string(directions[index])

		node := nodes[next]
		next = node[direction]

		steps++
		index++
	}

	return steps
}
