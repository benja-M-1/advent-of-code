package main

import (
	"embed"
	"fmt"
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
	r2 := Two(string(input))
	fmt.Printf("puzzle 2: %v\n", r2)
}

type Node struct {
	X      int
	Y      int
	Edge   bool
	Inside bool
	Value  string
}

func One(input string) int {
	input = strings.Trim(input, "\n")

	topLeft, bottomRight := []int{0, 0}, []int{0, 0}
	grid := map[int]map[int]*Node{0: {0: {X: 0, Y: 0, Edge: true, Value: "#"}}}
	current := []int{0, 0}
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Fields(line)
		dir := parts[0]
		length, _ := strconv.Atoi(parts[1])

		for length != 0 {
			switch dir {
			case "R":
				current = []int{current[0] + 1, current[1]}
			case "L":
				current = []int{current[0] - 1, current[1]}
			case "U":
				current = []int{current[0], current[1] - 1}
			case "D":
				current = []int{current[0], current[1] + 1}
			}

			n := &Node{
				X:     current[0],
				Y:     current[1],
				Edge:  true,
				Value: "#",
			}

			if _, ok := grid[current[1]]; !ok {
				grid[current[1]] = map[int]*Node{
					current[0]: n,
				}
			} else {
				grid[current[1]][current[0]] = n
			}

			if topLeft[0] > current[0] {
				topLeft[0] = current[0]
			}

			if topLeft[1] > current[1] {
				topLeft[1] = current[1]
			}

			if bottomRight[0] < current[0] {
				bottomRight[0] = current[0]
			}

			if bottomRight[1] < current[1] {
				bottomRight[1] = current[1]
			}

			length--
		}
	}

	surface := 0
	for y := topLeft[1]; y <= bottomRight[1]; y++ {
		k := maps.Keys(grid[y])
		slices.SortFunc(k, func(i, j int) int {
			if i < j {
				return -1
			} else if i > j {
				return 1
			} else {
				return 0
			}
		})

		ledge, redge := k[0], k[len(k)-1]

		for x := topLeft[0]; x < ledge; x++ {
			grid[y][x] = &Node{X: x, Y: y, Inside: false, Value: "."}
		}

		inside := false

		for x := ledge; x <= redge; x++ {
			cur, cok := grid[y][x]
			prev, pok := grid[y][x-1]
			up, uok := grid[y-1][x]

			if !cok {
				cur = &Node{X: x, Y: x, Value: "."}
				grid[y][x] = cur
			}

			if cok && cur.Edge {
				if uok && up.Inside && pok && prev.Inside {
					inside = true
				} else if uok && up.Edge && pok && prev.Inside {
					inside = true
				} else if pok && uok && !up.Inside && !up.Edge {
					inside = prev.Inside
				}

				cur.Inside = inside

				surface++
				continue
			}

			if pok && prev.Edge && !prev.Inside {
				if uok && up.Edge && !up.Inside {
					inside = true
				} else if uok && !up.Edge && up.Inside {
					inside = true
				}
			} else if pok && prev.Edge && prev.Inside {
				if uok && up.Edge && up.Inside {
					inside = false
				} else if uok && !up.Edge {
					inside = up.Inside
				}
			}

			if inside {
				cur.Inside = true
				if !cur.Edge {
					cur.Value = "-"
				}

				surface++
			}
		}

		for x := redge + 1; x <= bottomRight[0]; x++ {
			grid[y][x] = &Node{X: x, Y: y, Inside: false, Value: "."}
		}

	}

	ky := maps.Keys(grid)
	slices.Sort(ky)
	for _, y := range ky {
		fmt.Print(y, "\t")
		kx := maps.Keys(grid[y])
		slices.Sort(kx)
		for _, x := range kx {
			fmt.Print(grid[y][x].Value)
		}
		fmt.Println()
	}

	return surface
}

type Corner struct {
	X int64
	Y int64
}

type Rectangle struct {
	TopLeft     *Corner
	BottomRight *Corner
}

func Two(input string) int {
	return 0
}
