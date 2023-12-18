package main

import (
	"embed"
	"fmt"
	"math"
	"slices"
	"strconv"
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

type block struct {
	X                 int
	Y                 int
	HeatLoss          float64
	CumulatedHeatLoss float64
}

type orientedBlock struct {
	*block
	Direction direction
}

type direction struct {
	Value  string
	Length int
}

func One(input string) int {
	input = strings.Trim(input, "\n")

	graph := parse(input)
	return traverse(graph)
}

func traverse(graph map[int]map[int]*block) int {
	visited := []*orientedBlock{}
	queue := []*orientedBlock{{
		block: graph[0][0],
	}}
	for len(queue) > 0 {
		slices.SortFunc(queue, func(i, j *orientedBlock) int { return int(i.HeatLoss - j.HeatLoss) })
		current := queue[0]
		queue = queue[1:]

		if slices.ContainsFunc(visited, func(v *orientedBlock) bool {
			return v.block == current.block
		}) {
			continue
		}

		var prevNode *block
		switch current.Direction.Value {
		case "left":
			if pr, ok := graph[current.Y][current.X+1]; ok {
				prevNode = pr
			}
		case "right":
			if pl, ok := graph[current.Y][current.X-1]; ok {
				prevNode = pl
			}
		case "up":
			if pd, ok := graph[current.Y+1][current.X]; ok {
				prevNode = pd
			}
		case "down":
			if pu, ok := graph[current.Y-1][current.X]; ok {
				prevNode = pu
			}
		}

		if prevNode != nil {
			current.CumulatedHeatLoss = current.HeatLoss + prevNode.CumulatedHeatLoss
		} else {
			current.CumulatedHeatLoss = current.HeatLoss
		}

		fmt.Println(current.X, current.Y, current.Direction.Value, current.CumulatedHeatLoss)
		if current.X == len(graph[0])-1 && current.Y == len(graph)-1 {
			return int(current.CumulatedHeatLoss - graph[0][0].HeatLoss)
		}

		nexts := map[string]*block{}
		if ln, ok := graph[current.Y][current.X-1]; ok {
			nexts["left"] = ln
		}
		if un, ok := graph[current.Y-1][current.X]; ok {
			nexts["up"] = un
		}
		if rn, ok := graph[current.Y][current.X+1]; ok {
			nexts["right"] = rn
		}
		if dn, ok := graph[current.Y+1][current.X]; ok {
			nexts["down"] = dn
		}

		for d, n := range nexts {
			dir := current.Direction
			if dir.Value == d && dir.Length == 3 {
				continue
			}

			if dir.Value == "left" && d == "right" ||
				dir.Value == "right" && d == "left" ||
				dir.Value == "up" && d == "down" ||
				dir.Value == "down" && d == "up" {
				continue
			}

			if dir.Value != d {
				dir.Value = d
				dir.Length = 0
			}

			dir.Length++
			queue = append(queue, &orientedBlock{
				block:     n,
				Direction: dir,
			})

		}

		visited = append(visited, current)
	}

	return 0
}

func parse(input string) map[int]map[int]*block {
	graph := map[int]map[int]*block{}
	for y, line := range strings.Split(input, "\n") {
		graph[y] = map[int]*block{}
		for x, char := range line {
			v, _ := strconv.ParseFloat(string(char), 64)
			n := &block{
				X:                 x,
				Y:                 y,
				HeatLoss:          v,
				CumulatedHeatLoss: math.Inf(1),
			}

			if x == 0 && y == 0 {
				n.CumulatedHeatLoss = 0
			}

			graph[y][x] = n
		}
	}

	return graph
}

func Two(input string) int {
	input = strings.Trim(input, "\n")

	return len(input)
}
