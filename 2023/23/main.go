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

type Coord struct {
	X, Y int
}

type Tile struct {
	Coord Coord
	Value string
}

func One(input string) int {
	grid, start := makeGrid(strings.Trim(input, "\n"))

	queue := []state{
		{*start, []Tile{}},
	}
	paths := [][]Tile{}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.Tile.Coord.Y == len(grid)-1 {
			p := append(current.Path, current.Tile)
			paths = append(paths, p)
			continue
		}

		directions := []Coord{}
		switch current.Tile.Value {
		case "<":
			directions = append(directions, []Coord{{-1, 0}}...)
		case ">":
			directions = append(directions, []Coord{{1, 0}}...)
		case "v":
			directions = append(directions, []Coord{{0, 1}}...)
		case "^":
			directions = append(directions, []Coord{{0, -1}}...)
		default:
			directions = append(directions, []Coord{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}...)
		}

		for _, dir := range directions {
			x := current.Tile.Coord.X + dir.X
			y := current.Tile.Coord.Y + dir.Y

			next, ok := grid[y][x]
			if !ok {
				continue
			} else if next.Value == "#" {
				continue
			} else if next.Value == "<" && dir.X == 1 {
				continue
			} else if next.Value == ">" && dir.X == -1 {
				continue
			} else if next.Value == "v" && dir.Y == -1 {
				continue
			} else if next.Value == "^" && dir.Y == 1 {
				continue
			} else if slices.Contains(current.Path, next) {
				continue
			}

			p := make([]Tile, len(current.Path))
			copy(p, current.Path)
			queue = append(queue, state{
				Tile: next,
				Path: append(p, current.Tile),
			})
		}
	}

	slices.SortFunc(paths, func(a, b []Tile) int {
		return len(b) - len(a)
	})

	return len(paths[0]) - 1
}

type state struct {
	Tile Tile
	Path []Tile
}

func Two(input string) int {
	grid, start := makeGrid(strings.Trim(input, "\n"))

	type item struct {
		tile Tile
		path []Tile
	}

	allSegments, firstSegment := segments(start, grid)
	queue := []item{{firstSegment[len(firstSegment)-1], firstSegment}}
	paths := [][]Tile{}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.path[len(current.path)-1].Coord.Y == len(grid)-1 {
			paths = append(paths, current.path)
			continue
		}

		for _, dir := range []Coord{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
			x := current.tile.Coord.X + dir.X
			y := current.tile.Coord.Y + dir.Y

			nextSegment, ok := allSegments[y][x]
			if !ok {
				continue
			}

			// The current path already goes through at least one tile on the segment which shouldn't be possible
			if len(current.path) > 1 && slices.ContainsFunc(current.path[:len(current.path)-1], func(tile Tile) bool {
				return slices.Contains(nextSegment, tile)
			}) {
				continue
			}

			// the end of the segment is the end of the current path, the segment must be reversed
			if nextSegment[len(nextSegment)-1] == current.path[len(current.path)-1] {
				s := append([]Tile{}, nextSegment...)
				slices.Reverse(s)
				nextSegment = s
			}

			if nextSegment[0] != current.path[len(current.path)-1] {
				continue
			}

			queue = append(queue, item{
				path: append(append([]Tile{}, current.path[:len(current.path)-1]...), nextSegment...),
				tile: nextSegment[len(nextSegment)-1],
			})
		}
	}

	slices.SortFunc(paths, func(a, b []Tile) int {
		return len(b) - len(a)
	})

	return len(paths[0]) - 1
}

func segments(start *Tile, grid map[int]map[int]Tile) (map[int]map[int][]Tile, []Tile) {
	paths := map[int]map[int][]Tile{}
	first := []Tile{}

	type item struct {
		tile Tile
		path []Tile
	}

	queue := []item{
		{
			tile: *start,
			path: []Tile{},
		},
	}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if _, ok := paths[current.tile.Coord.Y][current.tile.Coord.X]; ok {
			continue
		}
		//
		// if slices.ContainsFunc(paths, func(tiles []Tile) bool {
		// 	return (tiles[0] == current.path[0] && tiles[len(tiles)-1] == current.path[len(current.path)-1]) ||
		// 		(tiles[0] == current.path[len(current.path)-1] && tiles[len(tiles)-1] == current.path[0])
		// }) {
		// 	continue
		// }

		if current.tile.Coord.Y == len(grid)-1 {
			if _, ok := paths[current.tile.Coord.Y]; !ok {
				paths[current.tile.Coord.Y] = map[int][]Tile{}
			}

			paths[current.tile.Coord.Y][current.tile.Coord.X] = current.path
			continue
		}

		if strings.ContainsAny(current.tile.Value, "<>v^") {
			if _, ok := paths[current.path[0].Coord.Y]; !ok {
				paths[current.path[0].Coord.Y] = map[int][]Tile{}
			}

			if _, ok := paths[current.tile.Coord.Y]; !ok {
				paths[current.tile.Coord.Y] = map[int][]Tile{}
			}

			paths[current.tile.Coord.Y][current.tile.Coord.X] = current.path
			paths[current.path[0].Coord.Y][current.path[0].Coord.X] = current.path

			if len(first) == 0 {
				first = current.path
			}

			current.path = []Tile{}
		}

		for _, dir := range []Coord{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
			x := current.tile.Coord.X + dir.X
			y := current.tile.Coord.Y + dir.Y

			next, ok := grid[y][x]
			if !ok {
				continue
			} else if next.Value == "#" {
				continue
			} else if slices.Contains(current.path, next) {
				continue
			}

			p := append([]Tile{}, current.path...)
			queue = append(queue, item{next, append(p, next)})
		}
	}

	return paths, first
}

func makeGrid(input string) (map[int]map[int]Tile, *Tile) {
	var start *Tile
	grid := map[int]map[int]Tile{}
	for y, row := range strings.Split(input, "\n") {
		grid[y] = map[int]Tile{}
		for x, col := range strings.Split(row, "") {
			t := Tile{
				Coord{
					X: x,
					Y: y,
				},
				col,
			}
			if start == nil && t.Value == "." {
				start = &t
			}
			grid[y][x] = t
		}
	}
	return grid, start
}
