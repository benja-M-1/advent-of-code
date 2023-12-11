package main

import (
	"embed"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
)

//go:embed *.txt
var f embed.FS

func main() {
	input, _ := f.ReadFile("input.txt")

	r1 := One(string(input))
	fmt.Printf("puzzle 1: %v\n", r1)
	r2 := Two(string(input), 1000000)
	fmt.Printf("puzzle 2: %v\n", r2)
}

type Position map[string]int

func (p Position) Recalculate(expansionFactor int, expandedColumns []int, expandedRows []int) Position {
	addedColumnsBefore := len(FindAll(expandedColumns, func(i, v int) bool { return v < p["x"] }))
	addedRowsBefore := len(FindAll(expandedRows, func(i, v int) bool { return v < p["y"] }))

	x := p["x"] + addedColumnsBefore*expansionFactor - addedColumnsBefore
	y := p["y"] + addedRowsBefore*expansionFactor - addedRowsBefore

	return Position{"x": x, "y": y}
}

func One(input string) int {
	return sum(strings.Trim(input, "\n"), 2)
}

func sum(input string, expansionFactor int) int {
	galaxies := map[string]Position{}
	expandedColumns := []int{}
	expandedRows := []int{}

	for y, row := range strings.Split(input, "\n") {
		expandedRows = append(expandedRows, y)
		for x, tile := range strings.Split(row, "") {
			if y == 0 {
				expandedColumns = append(expandedColumns, x)
			}

			if tile != "." {
				yi := slices.Index(expandedRows, y)
				if yi >= 0 {
					expandedRows = slices.Delete(expandedRows, yi, yi+1)
				}
				xi := slices.Index(expandedColumns, x)
				if xi >= 0 {
					expandedColumns = slices.Delete(expandedColumns, xi, xi+1)
				}
				tile = strconv.Itoa(len(maps.Keys(galaxies)) + 1)
				galaxies[tile] = Position{"x": x, "y": y}
			}
		}
	}

	var sum int

	rest := maps.Clone(galaxies)
	for i, g1 := range galaxies {
		g1 = g1.Recalculate(expansionFactor, expandedColumns, expandedRows)
		delete(rest, i)

		for _, g2 := range rest {
			g2 = g2.Recalculate(expansionFactor, expandedColumns, expandedRows)

			length := math.Abs(float64(g1["x"]-g2["x"])) + math.Abs(float64(g1["y"]-g2["y"]))

			sum += int(length)
		}
	}

	return sum
}

func Two(input string, expansionFactor int) int {
	return sum(strings.Trim(input, "\n"), expansionFactor)
}

func FindAll(cols []int, f func(i, v int) bool) []int {
	var n []int
	for i, v := range cols {
		if f(i, v) {
			n = append(n, i)
		}
	}
	return n
}
