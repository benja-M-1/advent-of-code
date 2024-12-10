package main

import (
	"embed"
	"fmt"
	"strconv"
	"strings"

	aocstrconv "adventofcode/pkg/strconv"
)

//go:embed *.txt
var f embed.FS

func main() {
	input, _ := f.ReadFile("input.txt")

	r1 := One(string(input))
	if r1 != 3312271365652 {
		fmt.Printf("puzzle 1: expecyed 3312271365652, given %v\n", r1)
	} else {
		fmt.Printf("puzzle 1: %v\n", r1)
	}

	r2 := Two(string(input))
	if r2 != 509463489296712 {
		fmt.Printf("puzzle 2: expecyed 509463489296712, given %v\n", r2)
	} else {
		fmt.Printf("puzzle 2: %v\n", r2)
	}
}

func One(input string) int {
	input = strings.Trim(input, "\n")

	sum := 0
	for _, line := range strings.Split(input, "\n") {
		expected, _ := strconv.Atoi(line[:strings.Index(line, ":")])
		numbers := aocstrconv.StringstoI(strings.Split(strings.Trim(line[strings.Index(line, ":")+2:], " "), " "))

		solved := solvable(numbers[0], numbers[1], expected, numbers[2:], mul, add)

		if solved {
			sum += expected
		}
	}

	return sum
}

func mul(a, b int) int {
	return a * b
}

func add(a, b int) int {
	return a + b
}

func concat(a, b int) int {
	conc, _ := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
	return conc
}

type operation func(a, b int) int

func solvable(a, b, expected int, rest []int, operations ...operation) bool {
	for _, op := range operations {
		r := op(a, b)

		if len(rest) == 0 {
			if r == expected {
				return true
			}
			continue
		}

		if solvable(r, rest[0], expected, rest[1:], operations...) {
			return true
		}
	}

	return false
}

func Two(input string) int {
	input = strings.Trim(input, "\n")

	sum := 0
	for _, line := range strings.Split(input, "\n") {
		expected, _ := strconv.Atoi(line[:strings.Index(line, ":")])
		numbers := aocstrconv.StringstoI(strings.Split(strings.Trim(line[strings.Index(line, ":")+2:], " "), " "))

		solved := solvable(numbers[0], numbers[1], expected, numbers[2:], mul, add, concat)

		if solved {
			sum += expected
		}
	}

	return sum
}
