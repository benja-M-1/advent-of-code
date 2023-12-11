package main

import (
	"bufio"
	"embed"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

//go:embed *.txt
var f embed.FS

func main() {
	input, _ := f.ReadFile("input.txt")

	r1 := One(string(input))
	fmt.Printf("puzzle 1: %v\n", r1)
	r2 := Two(string(input))
	fmt.Printf("puzzle 2: %v\n", r2)
}

func One(input string) string {
	s := bufio.NewScanner(strings.NewReader(input))

	stacks := map[int][]string{}

	isProcedure := false

	for s.Scan() {
		line := s.Text()

		// There is an empty line just before the instructions
		if line == "" {
			isProcedure = true

			continue
		}

		if !isProcedure {
			stack := 0
			for i, s := range line {
				if i >= 0 && i%4 == 0 {
					stack++
					continue
				}

				if !unicode.IsLetter(s) {
					continue
				}

				stacks[stack] = append([]string{string(s)}, stacks[stack]...)
			}

			continue
		}

		words := strings.Split(line, " ")
		crates, _ := strconv.Atoi(words[1])
		from, _ := strconv.Atoi(words[3])
		to, _ := strconv.Atoi(words[5])

		for i := 0; i < crates; i++ {
			stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-1])
			stacks[from] = stacks[from][:len(stacks[from])-1]
		}
	}

	letters := make([]string, len(stacks))
	for i, s := range stacks {
		letters[i-1] = s[len(s)-1]
	}

	return strings.Join(letters, "")
}

func Two(input string) string {
	s := bufio.NewScanner(strings.NewReader(input))

	stacks := map[int][]string{}

	isProcedure := false

	for s.Scan() {
		line := s.Text()

		// There is an empty line just before the instructions
		if line == "" {
			isProcedure = true

			continue
		}

		if !isProcedure {
			stack := 0
			for i, s := range line {
				if i >= 0 && i%4 == 0 {
					stack++
					continue
				}

				if !unicode.IsLetter(s) {
					continue
				}

				stacks[stack] = append([]string{string(s)}, stacks[stack]...)
			}

			continue
		}

		words := strings.Split(line, " ")
		crates, _ := strconv.Atoi(words[1])
		from, _ := strconv.Atoi(words[3])
		to, _ := strconv.Atoi(words[5])

		start := len(stacks[from]) - crates
		end := len(stacks[from])

		stacks[to] = append(stacks[to], stacks[from][start:end]...)
		stacks[from] = stacks[from][:start]
	}

	letters := make([]string, len(stacks))
	for i, s := range stacks {
		letters[i-1] = s[len(s)-1]
	}

	return strings.Join(letters, "")
}
