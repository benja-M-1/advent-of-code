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

	r1 := One(string(input))
	fmt.Printf("puzzle 1: %v\n", r1)
	r2 := Two(string(input))
	fmt.Printf("puzzle 2: %v\n", r2)
}

func One(input string) int {
	patterns := parse(input)
	sum := 0
	for _, pattern := range patterns {
		l := lineOfReflection(pattern["rows"], 0) * 100
		if l == 0 {
			l = lineOfReflection(pattern["cols"], 0)
		}

		sum += l
	}

	return sum
}

func lineOfReflection(s [][]string, prev int) int {
	for i := 0; i < len(s); i++ {
		if i+1 >= len(s) || !slices.Equal(s[i], s[i+1]) {
			continue
		}

		similar := true
		for a, b := i-1, i+2; a >= 0 && b < len(s); a, b = a-1, b+1 {
			if !slices.Equal(s[a], s[b]) {
				similar = false
				break
			}
		}

		if similar && prev != i+1 {
			return i + 1
		}
	}

	return 0
}

func parse(input string) []map[string][][]string {
	input = strings.Trim(input, "\n")
	patterns := []map[string][][]string{}

	for _, pattern := range strings.Split(input, "\n\n") {
		cols, rows := [][]string{}, [][]string{}
		for _, row := range strings.Split(pattern, "\n") {

			cells := strings.Split(row, "")
			rows = append(rows, cells)

			for x, cell := range cells {
				if x > len(cols)-1 {
					cols = append(cols, []string{})
				}
				cols[x] = append(cols[x], cell)
			}
		}

		patterns = append(patterns, map[string][][]string{"cols": cols, "rows": rows})
	}

	return patterns
}

func Two(input string) int {
	patterns := parse(input)
	sum := 0
	for _, pattern := range patterns {
		l := lineOfReflection(pattern["rows"], 0)
		l2 := fix(pattern["rows"], l) * 100
		if l2 == 0 {
			l = lineOfReflection(pattern["cols"], 0)
			l2 = fix(pattern["cols"], l)
		}
		sum += l2
	}

	return sum
}

func fix(rows [][]string, l int) int {
	for r := range rows {
		for c := range rows[r] {
			old := rows[r][c]
			if rows[r][c] == "." {
				rows[r][c] = "#"
			} else {
				rows[r][c] = "."
			}

			l2 := lineOfReflection(rows, l)
			if l2 != 0 {
				return l2
			}

			// Put back the value
			rows[r][c] = old
		}
	}

	return 0
}
