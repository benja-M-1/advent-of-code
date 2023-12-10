package main

import (
	"embed"
	"fmt"
	"math"
	"slices"
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
	r22 := TwoSecond(string(input))
	fmt.Printf("puzzle 2: %v\n", r22)
}

const (
	Start  = "S"
	Ground = "."
	NS     = "|"
	EW     = "-"
	NE     = "L"
	NW     = "J"
	SW     = "7"
	SE     = "F"
)

type Position struct {
	X int
	Y int
}

func (p Position) Equal(position Position) bool {
	return p.X == position.X && p.Y == position.Y
}

func One(input string) int {
	input = strings.Trim(input, "\n")

	matrix, current := parse(input)

	m := Walk(matrix, current, Position{}, Matrix{})

	steps := 0
	for _, row := range m {
		steps += len(row)
	}

	return steps / 2
}

func parse(input string) (Matrix, Position) {
	matrix := Matrix{}
	start := Position{}

	for y, row := range strings.Split(input, "\n") {
		if _, ok := matrix[y]; !ok {
			matrix[y] = map[int]string{}
		}
		for x, tile := range strings.Split(row, "") {
			matrix[y][x] = tile
			if tile == Start {
				start.Y = y
				start.X = x
			}
		}
	}

	return matrix, start
}

func Walk(matrix Matrix, current Position, prev Position, loop Matrix) Matrix {
	var next Position

	tile := matrix.Get(current.X, current.Y)
	switch tile {
	case "|": // North - South
		next = Position{current.X, current.Y + 1}
		if prev.Y == next.Y {
			next.Y = current.Y - 1
		}
	case "-": // East - West
		next = Position{current.X + 1, current.Y}
		if prev.X == next.X {
			next.X = current.X - 1
		}
	case "L": // North - East
		next = Position{current.X, current.Y - 1}
		if prev.X == next.X && prev.Y == next.Y {
			next.X = current.X + 1
			next.Y = current.Y
		}
	case "J": // North - West
		next = Position{current.X - 1, current.Y}
		if prev.X == next.X && prev.Y == next.Y {
			next.X = current.X
			next.Y = current.Y - 1
		}
	case "7": // South - West
		next = Position{current.X - 1, current.Y}
		if prev.X == next.X && prev.Y == next.Y {
			next.X = current.X
			next.Y = current.Y + 1
		}
	case "F": // South - East
		next = Position{current.X + 1, current.Y}
		if prev.X == next.X && prev.Y == next.Y {
			next.X = current.X
			next.Y = current.Y + 1
		}
	case "S": // Start
		// Back to the start
		if loop.Has(current.X, current.Y) {
			return loop
		}

		loop.Set(current.X, current.Y, tile)

		nexts := []Position{}

		if current.Y < len(matrix[0]) {
			// up
			nexts = append(nexts, Position{current.X, current.Y + 1})
		}
		if current.Y > 0 {
			// down
			nexts = append(nexts, Position{current.X, current.Y - 1})
		}
		if current.X > 0 {
			// left
			nexts = append(nexts, Position{current.X - 1, current.Y})
		}
		if current.X < len(matrix) {
			// right
			nexts = append(nexts, Position{current.X + 1, current.Y})
		}

		for _, nxt := range nexts {
			path := Walk(matrix, nxt, current, loop)
			if len(path) == 0 {
				continue
			}
			return path
		}
		fallthrough
	default:
		return loop
	}

	loop.Set(current.X, current.Y, tile)

	return Walk(matrix, next, current, loop)
}

/**
 * Implemented with Stephen Turner's algorithm
 */
func TwoSecond(input string) int {
	input = strings.Trim(input, "\n")

	matrix, current := parse(input)

	loop := Walk(matrix, current, Position{}, Matrix{})

	var (
		m        float64
		n        int
		included []Position
	)

	ykeys := maps.Keys(matrix)
	slices.Sort(ykeys)

	for _, y := range ykeys {
		xkeys := maps.Keys(matrix[y])
		slices.Sort(xkeys)

		for _, x := range xkeys {
			tile := matrix.Get(x, y)
			if loop.Has(x, y) {
				if tile == "|" {
					m++
				} else if slices.Contains([]string{"J", "F", "S"}, tile) {
					m += 0.5
				} else if slices.Contains([]string{"L", "7"}, tile) {
					m -= 0.5
				}
			} else if math.Mod(m, 2) != 0 {
				n++
				included = append(included, Position{x, y})
			}
		}
	}

	return n
}

func Two(input string) int {
	input = strings.Trim(input, "\n")

	matrix, current := parse(input)

	loop := Walk(matrix, current, Position{}, Matrix{})

	var (
		n        int
		included []Position
	)

	for y := range matrix {
		for x := range matrix[y] {
			if !loop.Has(x, y) {
				matrix.Set(x, y, ".")
			}
		}
	}

	ykeys := maps.Keys(matrix)
	slices.Sort(ykeys)

	for _, y := range ykeys {
		xkeys := maps.Keys(matrix[y])
		slices.Sort(xkeys)

		for _, x := range xkeys {
			if matrix.Get(x, y) != Ground {
				continue
			}

			count := 0
			for _, xx := range xkeys[0:x] {
				v := matrix.Get(xx, y)
				if slices.Contains([]string{"|", "J", "L"}, v) {
					count++
				}
			}

			if count > 0 && count%2 != 0 {
				n++
				included = append(included, Position{x, y})
			}
		}
	}

	return n
}

type Matrix map[int]Row
type Row map[int]string

func (m Matrix) Has(x, y int) bool {
	_, ok := m[y][x]

	return ok
}

func (m Matrix) Get(x, y int) string {
	return m[y][x]
}

func (m Matrix) Set(x, y int, tile string) {
	if _, ok := m[y]; !ok {
		m[y] = Row{}
	}
	m[y][x] = tile
}

func (m Matrix) String() string {
	var s string

	ykeys := maps.Keys(m)
	slices.Sort(ykeys)

	for _, y := range ykeys {
		xkeys := maps.Keys(m[y])
		slices.Sort(xkeys)

		for _, x := range xkeys {
			s += m[y][x]
		}
		s += "\n"
	}

	return s
}
