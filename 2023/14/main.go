package main

import (
	"embed"
	"fmt"
	"strings"

	aocslices "adventofcode/pkg/slices"
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
	return weight(tilt(parse(input)))
}

func weight(platform [][]string) int {
	sum := 0

	load := len(platform[0])
	for _, row := range platform {
		sum += load * aocslices.CountIf(row, func(v string) bool { return v == "O" })
		load--
	}

	return sum
}

func parse(input string) [][]string {
	input = strings.Trim(input, "\n")

	rows := [][]string{}
	for _, values := range strings.Split(input, "\n") {
		rows = append(rows, strings.Split(values, ""))
	}

	return rows
}

func Two(input string) int {
	seen := map[int][][]string{}
	platform := parse(input)
	n := 1000000000
	for i := 0; i < n; i++ {
		platform = cycle(platform)
		hash := fmt.Sprintf("%v", platform)
		found := []int{}
		for k, v := range seen {
			h := fmt.Sprintf("%v", v)
			if h != hash {
				continue
			}
			found = append(found, k)

			if len(found) == 2 {
				cycleLen := k - found[0]
				final := k + (n-k)%cycleLen
				return weight(seen[final])
			}
		}

		c := make([][]string, len(platform))
		for i, row := range platform {
			c[i] = make([]string, len(row))
			copy(c[i], row)
		}
		seen[i+1] = c
	}

	return 0
}

func cycle(platform [][]string) [][]string {
	for i := 0; i < 4; i++ {
		platform = rotate(tilt(platform))
	}

	return platform
}

func tilt(platform [][]string) [][]string {
	p := make([][]string, len(platform))

	emptySpots := make([]int, len(platform[0]))
	for row, values := range platform {
		pr := make([]string, len(values))
		copy(pr, values)
		p[row] = pr

		for col := 0; col < len(values); col++ {
			if row == 0 {
				emptySpots[col] = 0
			}

			v := p[row][col]
			if v == "." {
				continue
			}

			if v == "#" {
				emptySpots[col] = row + 1
				continue
			}

			p[row][col] = "."
			p[emptySpots[col]][col] = v

			emptySpots[col] = emptySpots[col] + 1
		}
	}

	return p
}

func rotate(platform [][]string) [][]string {
	p := [][]string{}

	for row, values := range platform {
		for col, value := range values {
			if row == 0 {
				p = append(p, []string{})
			}

			p[col] = append([]string{value}, p[col]...)
		}
	}

	return p
}
