package main

import (
	"strings"
)

func PuzzleTwo(input string) int {
	input = strings.Trim(input, "\n")

	lines := strings.Split(input, "\n")
	directions := lines[0]
	nodes := map[string]map[string]string{}
	current := []string{}
	for _, line := range lines[2:] {
		replacer := strings.NewReplacer("(", "", ")", "", "=", "", ",", "")
		parts := strings.Fields(replacer.Replace(line))
		nodes[parts[0]] = map[string]string{
			"L": parts[1],
			"R": parts[2],
		}
		if string(parts[0][2]) == "A" {
			current = append(current, parts[0])
		}
	}

	var (
		index int
		steps int
		ends  int
	)
	pathSteps := make([]int, len(current))
	for ends != len(current) {
		if index%len(directions) == 0 {
			index = 0
		}

		direction := string(directions[index])

		steps++
		index++

		for i := 0; i < len(current); i++ {
			node := nodes[current[i]]
			next := node[direction]

			if string(next[2]) == "Z" {
				ends++
				pathSteps[i] = steps
			}

			current[i] = next
		}

	}

	if len(pathSteps) == 2 {
		return lcm(pathSteps[0], pathSteps[1])
	}

	return lcm(pathSteps[0], pathSteps[1], pathSteps[2:]...)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func lcm(a, b int, n ...int) int {
	result := a * b / gcd(a, b)
	for i := 0; i < len(n); i++ {
		result = lcm(result, n[i])
	}

	return result
}
