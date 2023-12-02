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

	var sum int
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := strings.ReplaceAll(scanner.Text(), "Game ", "")
		parts := strings.Split(line, ": ")
		id, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Printf("error occurred: %v\n", err)
		}

		possible := true
		sets := strings.Split(parts[1], "; ")
		for _, set := range sets {
			for _, s := range strings.Split(set, ", ") {
				p := strings.Split(s, " ")
				v, err := strconv.Atoi(p[0])
				if err != nil {
					fmt.Printf("error occurred: %v\n", err)
				}
				if v > rules[p[1]] {
					possible = false
				}
			}
		}

		if possible {
			sum += id
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error occurred: %v\n", err)
	}

	return sum
}

func PuzzleTwo(input string) int {
	var power int
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := strings.ReplaceAll(scanner.Text(), "Game ", "")
		parts := strings.Split(line, ": ")

		max := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		sets := strings.Split(parts[1], "; ")
		for _, set := range sets {
			for _, s := range strings.Split(set, ", ") {
				p := strings.Split(s, " ")
				v, err := strconv.Atoi(p[0])
				if err != nil {
					fmt.Printf("error occurred: %v\n", err)
				}
				n, ok := max[p[1]]
				if !ok || v > n {
					max[p[1]] = v
				}
			}
		}

		power += max["red"] * max["green"] * max["blue"]
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error occurred: %v\n", err)
	}

	return power
}

//go:embed *.txt
var f embed.FS

func main() {
	input, _ := f.ReadFile("input.txt")
	r1 := Puzzle(string(input))
	r2 := PuzzleTwo(string(input))

	fmt.Printf("puzzle 1: %v\npuzzle 2: %v\n", r1, r2)
}
