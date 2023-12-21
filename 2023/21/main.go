package main

import (
	"embed"
	"fmt"
	"slices"
	"strings"

	"golang.org/x/exp/maps"
)

// read https://github.com/villuna/aoc23/wiki/A-Geometric-solution-to-advent-of-code-2023,-day-21 to find a way to finish this puzzle

//go:embed *.txt
var f embed.FS

func main() {
	input, _ := f.ReadFile("input.txt")

	r1 := Puzzle(string(input), 64)
	fmt.Printf("puzzle 1: %v\n", r1)
	r2 := Puzzle(string(input), 26501365)
	fmt.Printf("puzzle 2: %v\n", r2)
}

func Puzzle(input string, steps int) int {
	g, start := parse(input)

	c := walk(g, start, steps)

	return c
}

func parse(input string) (*garden, *tile) {
	input = strings.Trim(input, "\n")
	g := &garden{layout: layout{}, value: layout{}}
	var start *tile

	for y, row := range strings.Split(input, "\n") {
		g.layout[y] = map[int]*tile{}
		g.value[y] = map[int]*tile{}
		for x, mark := range strings.Split(row, "") {
			t := tile{X: x, Y: y, Mark: mark}

			g.layout[y][x] = &t

			v := t
			if mark == "S" {
				start = &v
			}

			g.value[y][x] = &v
		}
	}

	return g, start
}

func walk(g *garden, start *tile, steps int) int {
	plots, notPlots := map[string]*tile{}, map[string]*tile{}
	queue := []queuedTile{
		{tile: start, steps: 0},
	}

	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]

		if current.steps%steps == 0 || steps%2 == current.steps%2 {
			current.tile.Mark = "O"
			plots[current.tile.Identifier()] = current.tile
		} else {
			notPlots[current.tile.Identifier()] = current.tile
		}

		if current.steps >= steps {
			continue
		}

		for _, next := range nexts(g, current.tile) {
			if next.Mark == "#" {
				continue
			}

			// Avoid to add the same tile twice in the queue
			if slices.ContainsFunc(queue, func(q queuedTile) bool { return q.tile.X == next.X && q.tile.Y == next.Y }) {
				continue
			}

			// Avoid to add a tile we know it can be a plot reached by the elf
			_, isPlot := plots[next.Identifier()]
			if isPlot {
				continue
			}

			// Avoid to add a tile we know it can't be a plot reached by the elf
			_, isNotPlot := notPlots[next.Identifier()]
			if isNotPlot {
				continue
			}

			queue = append(queue, queuedTile{tile: next, steps: current.steps + 1})
		}
	}

	return len(maps.Values(plots))
}

func nexts(g *garden, t *tile) []*tile {
	if t.Y == g.topEdge() ||
		t.Y == g.bottomEdge() ||
		t.X == g.leftEdge() ||
		t.X == g.rightEdge() {
		g.Expand()
	}

	ts := []*tile{}
	if up, ok := g.value[t.Y-1][t.X]; ok {
		ts = append(ts, up)
	}
	if down, ok := g.value[t.Y+1][t.X]; ok {
		ts = append(ts, down)
	}
	if left, ok := g.value[t.Y][t.X-1]; ok {
		ts = append(ts, left)
	}
	if right, ok := g.value[t.Y][t.X+1]; ok {
		ts = append(ts, right)
	}

	return ts
}

type layout map[int]map[int]*tile

type garden struct {
	layout layout
	value  map[int]map[int]*tile
}

func (g *garden) topEdge() int {
	ky := maps.Keys(g.value)
	slices.Sort(ky)

	return ky[0]
}

func (g *garden) bottomEdge() int {
	ky := maps.Keys(g.value)
	slices.Sort(ky)

	return ky[len(ky)-1]
}

func (g *garden) leftEdge() int {
	y := g.topEdge()
	kx := maps.Keys(g.value[y])
	slices.Sort(kx)

	return kx[0]
}

func (g *garden) rightEdge() int {
	kx := maps.Keys(g.value[0])
	slices.Sort(kx)

	return kx[len(kx)-1]
}

func (g *garden) Expand() {
	g.ExpandUp()
	g.ExpandDown()
	g.ExpandLeft()
	g.ExpandRight()
}

func (g *garden) ExpandUp() {
	initialRows := len(g.layout)
	topEdge := g.topEdge()

	for ny, y := topEdge-initialRows, topEdge; ny < topEdge; ny, y = ny+1, y+1 {
		g.value[ny] = map[int]*tile{}
		for x := range g.value[y] {
			g.value[ny][x] = g.copyTile(x, y, x, ny)
		}
	}
}

func (g *garden) ExpandDown() {
	initialRows := len(g.layout)
	bottomEdge := g.bottomEdge()

	for ny, y := bottomEdge+initialRows, bottomEdge; ny > bottomEdge; ny, y = ny-1, y-1 {
		g.value[ny] = map[int]*tile{}
		for x := range g.value[y] {
			g.value[ny][x] = g.copyTile(x, y, x, ny)
		}
	}
}

func (g *garden) ExpandLeft() {
	initialCols := len(g.layout[0])
	leftEdge := g.leftEdge()

	for y := range g.value {
		for nx, x := leftEdge-initialCols, leftEdge; nx < leftEdge; nx, x = nx+1, x+1 {
			g.value[y][nx] = g.copyTile(x, y, nx, y)
		}
	}
}

func (g *garden) ExpandRight() {
	initialCols := len(g.layout[0])
	rightEdge := g.rightEdge()

	for y := range g.value {
		for nx, x := rightEdge+initialCols, rightEdge; nx > rightEdge; nx, x = nx-1, x-1 {
			g.value[y][nx] = g.copyTile(x, y, nx, y)
		}
	}
}

func (g *garden) copyTile(x, y, toX, toY int) *tile {
	initial := g.value[y][x]
	t := tile{X: toX, Y: toY, Mark: initial.Mark}
	if t.Mark != "#" {
		t.Mark = "."
	}

	return &t
}

func (g *garden) String() string {
	ky := maps.Keys(g.value)
	slices.Sort(ky)

	var out string
	for _, y := range ky {
		kx := maps.Keys(g.value[y])
		slices.Sort(kx)

		for _, x := range kx {
			out += g.value[y][x].Mark
		}
		out += "\n"
	}

	return out
}

type tile struct {
	X, Y int
	Mark string
}

func (t tile) Identifier() string {
	return fmt.Sprintf("%v,%v", t.X, t.Y)
}

type queuedTile struct {
	tile  *tile
	steps int
}
