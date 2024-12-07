package main

import (
	"embed"
	"fmt"
	"slices"
	"strings"
)

//go:embed *.txt
var f embed.FS

func main() {
	input, _ := f.ReadFile("input.txt")

	if r1 := One(string(input)); r1 != 2662 {
		fmt.Printf("puzzle 1 should be %v, %v given\n", 2662, r1)
	} else {
		fmt.Printf("puzzle 1: %v\n", r1)
	}

	if r2 := Two(string(input)); r2 != 2034 {
		fmt.Printf("puzzle 2 should be %v, %v given\n", 2034, r2)
	} else {
		fmt.Printf("puzzle 2: %v\n", r2)
	}
}

const xmas = "XMAS"
const samx = "SAMX"

func One(input string) int {
	input = strings.Trim(input, "\n")

	matches := 0
	m := map[int]map[int]string{}
	for y, line := range strings.Split(input, "\n") {
		for x, char := range line {
			if _, ok := m[y]; !ok {
				m[y] = map[int]string{}
			}
			m[y][x] = string(char)

			if string(char) == "X" {
				if _, ok := topLeft(x, y, "X", m, xmas); ok {
					matches++
				}
				if _, ok := top(x, y, "X", m, xmas); ok {
					matches++
				}
				if _, ok := topRight(x, y, "X", m, xmas); ok {
					matches++
				}
				if _, ok := left(x, y, "X", m, xmas); ok {
					matches++
				}
			}

			if string(char) == "S" {
				if _, ok := topLeft(x, y, "S", m, samx); ok {
					matches++
				}
				if _, ok := top(x, y, "S", m, samx); ok {
					matches++
				}
				if _, ok := topRight(x, y, "S", m, samx); ok {
					matches++
				}
				if _, ok := left(x, y, "S", m, samx); ok {
					matches++
				}
			}
		}
	}

	return matches
}

func topLeft(x, y int, w string, m map[int]map[int]string, expected string) (string, bool) {
	if x < 0 || y < 0 {
		return "", false
	}

	w += m[y-1][x-1]
	if w == expected {
		return w, true
	}

	if strings.Contains(expected, w) {
		return topLeft(x-1, y-1, w, m, expected)
	}

	return "", false
}

func topRight(x, y int, w string, m map[int]map[int]string, expected string) (string, bool) {
	if x > len(m[y])-1 || y < 0 {
		return "", false
	}

	w += m[y-1][x+1]
	if w == expected {
		return w, true
	}

	if strings.Contains(expected, w) {
		return topRight(x+1, y-1, w, m, expected)
	}

	return "", false
}

func left(x, y int, w string, m map[int]map[int]string, expected string) (string, bool) {
	if x < 0 || y < 0 || y > len(m)-1 {
		return "", false
	}

	w += m[y][x-1]
	if w == expected {
		return w, true
	}

	if strings.Contains(expected, w) {
		return left(x-1, y, w, m, expected)
	}

	return "", false
}

func top(x, y int, w string, m map[int]map[int]string, expected string) (string, bool) {
	if x < 0 || x > len(m[y])-1 || y < 0 {
		return "", false
	}

	w += m[y-1][x]
	if w == expected {
		return w, true
	}

	if strings.Contains(expected, w) {
		return top(x, y-1, w, m, expected)
	}

	return "", false
}

const mas = "MAS"
const sam = "SAM"

func Two(input string) int {
	input = strings.Trim(input, "\n")
	input = strings.Trim(input, "\n")

	matches := 0
	m := map[int]map[int]string{}
	for y, line := range strings.Split(input, "\n") {
		for x, char := range line {
			if _, ok := m[y]; !ok {
				m[y] = map[int]string{}
			}
			m[y][x] = string(char)

			if x < 2 || y < 2 {
				continue
			}

			starters := []string{"S", "M"}
			br := m[y][x]
			bl := m[y][x-2]
			if !slices.Contains(starters, br) && !slices.Contains(starters, bl) {
				continue
			}

			w := mas
			if br == "S" {
				w = sam
			}

			if _, ok := topLeft(x, y, br, m, w); !ok {
				continue
			}

			w = mas
			if bl == "S" {
				w = sam
			}
			if _, ok := topRight(x-2, y, bl, m, w); ok {
				matches++
			}
		}
	}

	return matches
}
