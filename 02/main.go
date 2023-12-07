package main

import (
	"bufio"
	"embed"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

func Puzzle(input string) int {
	rules := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	games := parse(input)
	sum := 0

	for id, sets := range games {
		possible := true
		for _, subset := range sets {
			for color, value := range subset {
				if value > rules[color] {
					possible = false
					break
				}
			}

			if !possible {
				break
			}
		}

		if possible {
			sum += id
		}
	}

	return sum
}

func PuzzleTwo(input string) int {
	games := parse(input)
	power := 0

	for _, sets := range games {
		maximums := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for _, subset := range sets {
			for color, value := range subset {
				n, ok := maximums[color]
				if !ok || value > n {
					maximums[color] = value
				}
			}
		}
		power += maximums["red"] * maximums["green"] * maximums["blue"]
	}

	return power
}

type Game map[string]int
type Games map[int][]Game

func parse(input string) Games {
	games := Games{}

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		parseLine(scanner.Text(), games)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error occurred: %v\n", err)
	}

	return games
}

func parseLine(line string, games Games) {
	line = strings.ReplaceAll(line, "Game ", "")
	parts := strings.Split(line, ": ")

	id, _ := strconv.Atoi(parts[0])
	sets := strings.Split(parts[1], "; ")

	games[id] = []Game{}
	for _, set := range sets {
		game := map[string]int{}
		parseSet(set, game)
		games[id] = append(games[id], game)
	}
}

func parseSet(set string, game Game) {
	for _, subset := range strings.Split(set, ", ") {
		p := strings.Split(subset, " ")
		v, _ := strconv.Atoi(p[0])
		game[p[1]] = v
	}
}

//go:embed *.txt
var f embed.FS

func main() {
	input, _ := f.ReadFile("input.txt")
	r1 := Puzzle(string(input))
	r2 := PuzzleTwo(string(input))

	fmt.Printf("puzzle 1: %v\npuzzle 2: %v\n", r1, r2)
}
