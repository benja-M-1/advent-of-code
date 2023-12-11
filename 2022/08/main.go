package main

import (
	"embed"
	"fmt"
	"slices"
	"strconv"
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
	input = strings.Trim(input, "\n")
	rows, cols := parse(input)

	visibles := 0
	for y, row := range rows {
		if y == 0 || y == len(rows)-1 {
			visibles += len(row)
			continue
		}

		for x, tree := range row {
			if x == 0 || x == len(row)-1 {
				visibles++
				continue
			}

			neighbours := [][]int{
				row[0:x],
				row[x+1:],
				cols[x][0:y],
				cols[x][y+1:],
			}

			visible := false
			for _, n := range neighbours {
				if !aocslices.ContainsGreaterOrEqual(n, tree) {
					visible = true
					break
				}
			}

			if visible {
				visibles++
			}
		}
	}

	return visibles
}

func Two(input string) int {
	input = strings.Trim(input, "\n")

	rows, cols := parse(input)
	scores := []int{}

	for y := 0; y < len(rows); y++ {
		if y == 0 || y == len(rows)-1 {
			continue
		}

		for x := 0; x < len(rows[y]); x++ {
			if x == 0 || x == len(rows[y])-1 {
				continue
			}

			a := rows[y][x]
			prevsX := append([]int{}, rows[y][0:x]...)
			slices.Reverse(prevsX)

			prevsY := append([]int{}, cols[x][0:y]...)
			slices.Reverse(prevsY)
			s := []int{
				score(a, prevsX),
				score(a, prevsY),
			}

			if x < len(rows[y])-1 {
				s = append(s, score(a, rows[y][x+1:]))
			}
			if y < len(rows)-1 {
				s = append(s, score(a, cols[x][y+1:]))
			}

			scores = append(scores, aocslices.Product(s))
		}
	}

	return slices.Max(scores)
}

func score(tree int, neighbours []int) int {
	score := 0
	for _, n := range neighbours {
		score++
		if tree <= n {
			break
		}
	}

	return score
}

func parse(input string) (map[int][]int, map[int][]int) {
	rows := map[int][]int{}
	cols := map[int][]int{}
	for y, row := range strings.Split(input, "\n") {
		trees := strings.Split(row, "")
		for x, tree := range trees {
			v, _ := strconv.Atoi(tree)
			rows[y] = append(rows[y], v)
			cols[x] = append(cols[x], v)
		}
	}
	return rows, cols
}
