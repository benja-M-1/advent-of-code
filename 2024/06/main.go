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
	if r1 != 5329 {
		fmt.Printf("puzzle 1: expected 5329 got %v\n", r1)
	} else {
		fmt.Printf("puzzle 1: %v\n", r1)
	}
	r2 := Two(string(input))
	fmt.Printf("puzzle 2: %v\n", r2)
}
func One(input string) int {
	input = strings.Trim(input, "\n")
	lines := strings.Split(input, "\n")

	m := Map{
		MaxX:           len(lines[0]) - 1,
		MaxY:           len(lines) - 1,
		Obstacles:      map[int][]int{},
		AddedObstacles: []Coord{},
	}

	p := Patrol{}

	parse(lines, m, &p)

	for m.IsInside(p.Position.Coord) {
		p.Next(&m)
	}

	return len(p.Path.Uniques)
}

func Two(input string) int {
	input = strings.Trim(input, "\n")
	lines := strings.Split(input, "\n")

	m := Map{
		MaxX:           len(lines[0]) - 1,
		MaxY:           len(lines) - 1,
		Obstacles:      map[int][]int{},
		AddedObstacles: []Coord{},
	}

	p := Patrol{}

	parse(lines, m, &p)

	for m.IsInside(p.Position.Coord) {
		p.Next(&m)
	}

	fmt.Println(m.AddedObstacles)
	return len(m.AddedObstacles)
}

func parse(lines []string, m Map, p *Patrol) {
	for y, line := range lines {
		for x, char := range line {
			switch string(char) {
			case ".":
				continue
			case "#":
				if _, ok := m.Obstacles[y]; !ok {
					m.Obstacles[y] = []int{}
				}
				m.Obstacles[y] = append(m.Obstacles[y], x)
			default:
				pos := Position{
					Coord: Coord{
						X: x,
						Y: y,
					},
					Direction: string(char),
				}
				p.Path.Add(pos)
				p.Position = pos
			}
		}
	}
}

type Map struct {
	MaxX           int
	MaxY           int
	Obstacles      map[int][]int
	AddedObstacles []Coord
}

func (m *Map) IsInside(c Coord) bool {
	if c.X < 0 || c.X >= m.MaxX {
		return false
	}
	if c.Y < 0 || c.Y >= m.MaxY {
		return false
	}
	return true
}

func (m *Map) IsObstacle(c Coord) bool {
	if _, ok := m.Obstacles[c.Y]; ok {
		if slices.Contains(m.Obstacles[c.Y], c.X) {
			return true
		}
	}
	return false
}

type Coord struct {
	X, Y int
}

type Path struct {
	Positions []Position
	Uniques   []Position
}

func (p *Path) Add(pos Position) {
	p.Positions = append(p.Positions, pos)
	if !slices.ContainsFunc(p.Uniques, func(i Position) bool {
		return pos.Y == i.Y && pos.X == i.X
	}) {
		p.Uniques = append(p.Uniques, pos)
	}
}

func (p *Path) Get(pos Position) *Position {
	for _, i := range p.Positions {
		if pos.Y == i.Y && pos.X == i.X {
			return &i
		}
	}

	return nil
}

type Position struct {
	Coord
	Direction string
}

type Patrol struct {
	Path     Path
	Position Position
}

func (p *Patrol) Next(m *Map) {
	switch p.Position.Direction {
	case "^":
		p.moveNorth(m)
	case ">":
		p.moveWest(m)
	case "v":
		p.moveSouth(m)
	case "<":
		p.moveEast(m)
	}
}

func (p *Patrol) moveNorth(m *Map) {
	for p.Position.Y > 0 {
		next := Position{
			Coord: Coord{
				X: p.Position.X,
				Y: p.Position.Y - 1,
			},
			Direction: p.Position.Direction,
		}

		if m.IsObstacle(next.Coord) {
			p.Position.Direction = ">"
			return
		} else {
			for _, pos := range p.Path.Positions {
				if pos.Direction == ">" && pos.Y == next.Y && pos.X > next.X {
					m.AddedObstacles = append(m.AddedObstacles, next.Coord)
				}
			}
		}

		// Move to next
		p.Position = next

		// Mark it as visited
		p.Path.Add(p.Position)
	}
}

func (p *Patrol) moveEast(m *Map) {
	for p.Position.X > 0 {
		next := Position{
			Coord: Coord{
				X: p.Position.X - 1,
				Y: p.Position.Y,
			},
			Direction: p.Position.Direction,
		}

		if m.IsObstacle(next.Coord) {
			p.Position.Direction = "^"
			return
		} else {
			for _, pos := range p.Path.Positions {
				if pos.Direction == "^" && pos.X == next.X && pos.Y < next.Y {
					m.AddedObstacles = append(m.AddedObstacles, next.Coord)
				}
			}
		}

		// Move to next
		p.Position = next

		// Mark it as visited
		p.Path.Add(p.Position)
	}
}

func (p *Patrol) moveWest(m *Map) {
	for p.Position.X < m.MaxX {
		next := Position{
			Coord: Coord{
				X: p.Position.X + 1,
				Y: p.Position.Y,
			},
			Direction: p.Position.Direction,
		}

		if m.IsObstacle(next.Coord) {
			p.Position.Direction = "v"
			return
		} else {
			for _, pos := range p.Path.Positions {
				if pos.Direction == "v" && pos.X == next.X && pos.Y > next.Y {
					m.AddedObstacles = append(m.AddedObstacles, next.Coord)
				}
			}
		}

		// Move to next
		p.Position = next

		// Mark it as visited
		p.Path.Add(p.Position)
	}
}

func (p *Patrol) moveSouth(m *Map) {
	for p.Position.Y < m.MaxY {
		next := Position{
			Coord: Coord{
				X: p.Position.X,
				Y: p.Position.Y + 1,
			},
			Direction: p.Position.Direction,
		}

		if m.IsObstacle(next.Coord) {
			p.Position.Direction = "<"
			return
		} else {
			for _, pos := range p.Path.Positions {
				if pos.Direction == "<" && pos.Y == next.Y && pos.X < next.X {
					m.AddedObstacles = append(m.AddedObstacles, next.Coord)
				}
			}
		}

		// Move to next
		p.Position = next

		// Mark it as visited
		p.Path.Add(p.Position)
	}
}
