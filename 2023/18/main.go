package main

import (
	"embed"
	"fmt"
	"math"
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

type point struct {
	X float64
	Y float64
}

type instruction struct {
	Dir    string
	Length float64
}

func One(input string) int {
	input = strings.Trim(input, "\n")

	instructions := []instruction{}
	for _, line := range strings.Split(input, "\n") {
		i := instruction{}
		parts := strings.Fields(line)
		i.Dir = parts[0]
		i.Length, _ = strconv.ParseFloat(parts[1], 64)

		instructions = append(instructions, i)
	}

	permtr := perimeter(instructions)
	points := dig(instructions)

	return int(area(points) + permtr/2 + 1)
}

func perimeter(instructions []instruction) float64 {
	permtr := float64(0)
	for i := range instructions {
		permtr += instructions[i].Length
	}

	return permtr
}

func dig(instructions []instruction) []point {
	points := []point{{X: 0, Y: 0}}
	for i := range instructions {
		current := points[len(points)-1]

		switch instructions[i].Dir {
		case "R", "0":
			current = point{current.X + instructions[i].Length, current.Y}
		case "L", "2":
			current = point{current.X - instructions[i].Length, current.Y}
		case "U", "3":
			current = point{current.X, current.Y - instructions[i].Length}
		case "D", "1":
			current = point{current.X, current.Y + instructions[i].Length}
		}

		points = append(points, current)
	}

	return points
}

// area calculates the area of a polygon using the shoelace formula.
// https://en.wikipedia.org/wiki/Shoelace_formula#Other_formulas
// A = 1/2 * sum(x[i] * (y[i+1] - y[i-1]))
func area(points []point) float64 {
	a := float64(0)
	for current := range points {
		next, prev := current+1, current-1
		if current == len(points)-1 {
			next = 1
		}
		if current == 0 {
			prev = len(points) - 1
		}

		a += points[current].X * (points[next].Y - points[prev].Y) / 2
	}

	return math.Abs(a)
}

func Two(input string) int {
	input = strings.Trim(input, "\n")

	instructions := []instruction{}

	// Each line looks like this:
	// U 2 (#7a21e3)
	for _, line := range strings.Split(input, "\n") {
		i := instruction{}
		// parts[0] is the direction
		// parts[1] is the length
		// parts[2] is the hex color
		parts := strings.Fields(line)

		// the 5th first chars of hex color contains the length
		// the remainder contains the direction
		// (#7a21e3) -> 7a21e 3
		i.Dir = parts[2][7:8] // (#7a21e3) -> 3
		hex := parts[2][2:7]  // (#7a21e3) -> 7a21e
		length, _ := strconv.ParseInt(hex, 16, 64)
		i.Length = float64(length)

		instructions = append(instructions, i)
	}

	permtr := perimeter(instructions)
	points := dig(instructions)

	return int(area(points) + permtr/2 + 1)
}
