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

func lineOfReflection(s []string, prev int) int {
	for i := 0; i < len(s)-1; i++ {
		l, r := s[0:i+1], s[i+1:]
		if len(l) < len(r) {
			r = r[:len(l)]
		} else if i+1-len(r) >= 0 {
			l = s[i+1-len(r) : i+1]
		}

		// Make some copies to ensure no side effects on the original slice
		right := make([]string, len(r))
		left := make([]string, len(l))
		copy(right, r)
		copy(left, l)
		slices.Reverse(left)

		if slices.Equal(left, right) && prev != i+1 {
			return i + 1
		}
	}

	return 0
}

func parse(input string) []map[string][]string {
	input = strings.Trim(input, "\n")
	patterns := []map[string][]string{}

	for _, pattern := range strings.Split(input, "\n\n") {
		cols, rows := []string{}, []string{}
		for _, row := range strings.Split(pattern, "\n") {
			rows = append(rows, row)

			for x, cell := range row {
				if len(cols)-1 < x {
					cols = append(cols, "")
				}
				cols[x] += string(cell)
			}
		}

		patterns = append(patterns, map[string][]string{"cols": cols, "rows": rows})
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

func fix(rows []string, l int) int {
	for r := range rows {
		for c := range rows[r] {
			old := string(rows[r][c])
			if old == "." {
				rows[r] = rows[r][:c] + "#" + rows[r][c+1:]
			} else {
				rows[r] = rows[r][:c] + "." + rows[r][c+1:]
			}

			l2 := lineOfReflection(rows, l)
			if l2 != 0 {
				return l2
			}

			// Put back the value
			rows[r] = rows[r][:c] + old + rows[r][c+1:]
		}
	}

	return 0
}
