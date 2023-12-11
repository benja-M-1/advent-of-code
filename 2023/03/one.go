package main

import (
	_ "embed"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func Puzzle(input string) int {
	lines := strings.Split(strings.Trim(input, "\n"), "\n")
	numbers := map[int][]string{}
	for l, line := range lines {
		if line == "" {
			l++
			continue
		}
		lines[l] = line

		numbers[l] = strings.FieldsFunc(line, func(c rune) bool {
			return !unicode.IsNumber(c)
		})
		l++
	}

	keys := make([]int, 0, len(numbers))
	for k := range numbers {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	sum := 0
	for line, nbs := range numbers {
		for _, n := range nbs {
			index := strings.Index(lines[line], n)
			from, to := index-1, index+len(n)+1

			if from < 0 {
				from = 0
			}

			if to > len(lines[line]) {
				to = len(lines[line])
			}

			p := lines[line][from:to]

			// add previous line if possible
			if line > 0 {
				p = lines[line-1][from:to] + p
			}
			// add next line if possible
			if line < len(numbers)-1 {
				p += lines[line+1][from:to]
			}

			symbols := strings.FieldsFunc(p, func(c rune) bool {
				return unicode.IsLetter(c) || unicode.IsNumber(c) || string(c) == "."
			})
			if len(symbols) > 0 {
				i, _ := strconv.Atoi(n)
				sum += i
			}

			// Replace the number in the line with dots to avoid matching one
			// of the digits on the next iteration
			r := strings.Repeat(".", len(n))
			lines[line] = strings.Replace(lines[line], n, r, 1)
		}
	}

	return sum
}
