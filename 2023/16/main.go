package main

import (
	"embed"
	"errors"
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

var (
	Right = []int{1, 0}
	Up    = []int{0, -1}
	Left  = []int{-1, 0}
	Down  = []int{0, 1}
)

func One(input string) int {
	input = strings.Trim(input, "\n")

	grid := [][]string{}
	for _, values := range strings.Split(input, "\n") {
		grid = append(grid, strings.Split(values, ""))
	}

	energized := walk(grid, [][]int{}, [][]int{{0, 0}}, [][]int{Right})

	return len(energized)
}

func walk(grid [][]string, energized [][]int, currents [][]int, currentDirections [][]int) [][]int {
	nexts := [][]int{}
	directions := [][]int{}

	for c := range currents {
		cur := currents[c]
		dir := currentDirections[c]

		s := grid[cur[1]][cur[0]]

		// If the splitting cell is already energized, it is a cycle
		i := slices.IndexFunc(energized, func(cel []int) bool {
			return slices.Equal(cel, cur)
		})
		if i != -1 && (s == "|" || s == "-") {
			continue
		}

		switch s {
		case ".":
			if nxt, err := next(cur, dir, grid); err == nil {
				directions = append(directions, dir)
				nexts = append(nexts, nxt)
			}
		case "|":
			if dir[1] == Up[1] || dir[1] == Down[1] {
				if nxt, err := next(cur, dir, grid); err == nil {
					directions = append(directions, dir)
					nexts = append(nexts, nxt)
				}
			} else {
				// Split the beam to up and down
				for _, d := range [][]int{Up, Down} {
					if nxt, err := next(cur, d, grid); err == nil {
						directions = append(directions, d)
						nexts = append(nexts, nxt)
					}
				}
			}
		case "-":
			if dir[0] == Left[0] || dir[0] == Right[0] {
				if nxt, err := next(cur, dir, grid); err == nil {
					directions = append(directions, dir)
					nexts = append(nexts, nxt)
				}
			} else {
				// Split the beam to right and left
				for _, d := range [][]int{Left, Right} {
					if nxt, err := next(cur, d, grid); err == nil {
						directions = append(directions, d)
						nexts = append(nexts, nxt)
					}
				}
			}
		case "\\":
			if slices.Equal(dir, Up) {
				dir = Left
			} else if slices.Equal(dir, Down) {
				dir = Right
			} else if slices.Equal(dir, Left) {
				dir = Up
			} else if slices.Equal(dir, Right) {
				dir = Down
			}

			if nxt, err := next(cur, dir, grid); err == nil {
				directions = append(directions, dir)
				nexts = append(nexts, nxt)
			}
		case "/":
			if slices.Equal(dir, Up) {
				dir = Right
			} else if slices.Equal(dir, Down) {
				dir = Left
			} else if slices.Equal(dir, Left) {
				dir = Down
			} else if slices.Equal(dir, Right) {
				dir = Up
			}

			if nxt, err := next(cur, dir, grid); err == nil {
				directions = append(directions, dir)
				nexts = append(nexts, nxt)
			}
		}

		if i := slices.IndexFunc(energized, func(cel []int) bool {
			return slices.Equal(cel, cur)
		}); i == -1 {
			energized = append(energized, cur)
		}
	}

	if len(nexts) == 0 {
		return energized
	}

	return walk(grid, energized, append([][]int{}, nexts...), append([][]int{}, directions...))
}

func next(current []int, direction []int, grid [][]string) ([]int, error) {
	nxt := append([]int{}, current...)
	if direction[0] == Right[0] || direction[0] == Left[0] {
		nxt[0] += direction[0]

		// The beam reached the edge of the grid
		if nxt[0] < 0 || nxt[0] >= len(grid[0]) {
			return nil, errors.New("beam reached the edge of the grid")
		}
	} else {
		nxt[1] += direction[1]

		// The beam reached the edge of the grid
		if nxt[1] < 0 || nxt[1] >= len(grid) {
			return nil, errors.New("beam reached the edge of the grid")
		}
	}

	return nxt, nil
}

func Two(input string) int {
	input = strings.Trim(input, "\n")

	grid := [][]string{}
	for _, values := range strings.Split(input, "\n") {
		grid = append(grid, strings.Split(values, ""))
	}

	energized := 0
	for y := range grid {
		for x := range grid[y] {
			if y > 0 && y < len(grid)-1 && x > 0 && x < len(grid[y])-1 {
				continue
			}

			current := []int{x, y}
			var directions [][]int
			if y == 0 {
				directions = append(directions, Down)
			} else if y == len(grid)-1 {
				directions = append(directions, Up)
			}
			if x == 0 {
				directions = append(directions, Right)
			} else if x == len(grid[y])-1 {
				directions = append(directions, Left)
			}

			for d := range directions {
				e := walk(grid, [][]int{}, [][]int{current}, [][]int{directions[d]})
				if len(e) > energized {
					energized = len(e)
				}
			}
		}
	}

	return energized
}
