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
	if r1 != 308 {
		fmt.Printf("puzzle 1: expected 308 got %v\n", r1)
	} else {
		fmt.Printf("puzzle 1: %v\n", r1)
	}
	r2 := Two(string(input))
	if r2 != 1147 {
		fmt.Printf("puzzle 2: expected 1147 got %v\n", r2)
	} else {
		fmt.Printf("puzzle 2: %v\n", r2)
	}
}

func parse(input string) (map[int]map[int]string, []map[string]int, [][]int) {
	m := map[int]map[int]string{}
	antennas := []map[string]int{}

	lines := strings.Split(input, "\n")
	for y, line := range lines {
		for x, char := range strings.Split(line, "") {
			if _, ok := m[y]; !ok {
				m[y] = map[int]string{}
			}
			m[y][x] = char

			if char == "." {
				continue
			}

			antennas = append(antennas, map[string]int{"x": x, "y": y})
		}
	}

	bounds := [][]int{{0, 0}, {len(lines[0]) - 1, len(lines) - 1}}

	return m, antennas, bounds
}

func debug(m map[int]map[int]string) {
	yk := make([]int, 0, len(m))
	for k := range m {
		yk = append(yk, k)
	}
	slices.Sort(yk)

	for _, y := range yk {
		xk := make([]int, 0, len(m))
		for k := range m {
			xk = append(xk, k)
		}
		slices.Sort(xk)

		for _, x := range xk {
			fmt.Printf("%v", m[y][x])
		}
		fmt.Printf("\n")
	}
}

func isAntinode(antinode map[string]int, antinodes []map[string]int) bool {
	return slices.ContainsFunc(antinodes, func(m map[string]int) bool {
		return m["x"] == antinode["x"] && m["y"] == antinode["y"]
	})
}

func isWithinBounds(x, y int, bounds [][]int) bool {
	b0, b1 := bounds[0], bounds[1]
	if x < b0[0] || x > b1[0] || y < b0[1] || y > b1[1] {
		return false
	}

	return true
}

func generateKey(antennaA map[string]int, antennaB map[string]int) string {
	sortedAntennas := []map[string]int{antennaA, antennaB}
	slices.SortFunc(sortedAntennas, func(a, b map[string]int) int {
		if a["x"] < b["x"] && a["y"] < b["y"] {
			return -1
		} else if a["x"] > b["x"] && a["y"] > b["y"] {
			return 1
		} else {
			return 0
		}
	})
	k := fmt.Sprintf("%v", sortedAntennas)
	return k
}
