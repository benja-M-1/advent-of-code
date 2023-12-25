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
	if r1 >= 1155 || r1 == 37 || r1 == 54 || r1 == 119 || r1 == 80 || r1 == 112 {
		fmt.Printf("puzzle 1 ❌❌❌❌❌❌❌❌: %v\n", r1)
	} else {
		fmt.Printf("puzzle 1: %v\n", r1)
	}
	r2 := Two(string(input))
	fmt.Printf("puzzle 2: %v\n", r2)
}

type Coord struct {
	X, Y, Z int
}

type Brick struct {
	ID    string
	Left  Coord
	Right Coord
}

func One(input string) int {
	input = strings.Trim(input, "\n")
	identifiers := map[int]string{
		0:  "A",
		1:  "B",
		2:  "C",
		3:  "D",
		4:  "E",
		5:  "F",
		6:  "G",
		7:  "H",
		8:  "I",
		9:  "J",
		10: "K",
		11: "L",
		12: "M",
		13: "N",
		14: "O",
		15: "P",
		16: "Q",
		17: "R",
		18: "S",
		19: "T",
		20: "U",
		21: "V",
		22: "W",
		23: "X",
		24: "Y",
		25: "Z",
		26: "a",
		27: "b",
		28: "c",
		29: "d",
		30: "e",
		31: "f",
		32: "g",
		33: "h",
		34: "i",
		35: "j",
		36: "k",
		37: "l",
		38: "m",
		39: "n",
		40: "o",
		41: "p",
		42: "q",
		43: "r",
		44: "s",
		45: "t",
		46: "u",
		47: "v",
		48: "w",
		49: "x",
		50: "y",
		51: "z",
		52: "0",
		53: "1",
		54: "2",
		55: "3",
		56: "4",
		57: "5",
		58: "6",
		59: "7",
		60: "8",
		61: "9",
		62: "+",
		63: "/",
		64: "=",
		65: "!",
		66: "@",
		67: "#",
		68: "$",
		69: "%",
		70: "^",
		71: "&",
		72: "*",
		73: "(",
		74: ")",
		75: "_",
		76: "-",
		77: "+",
		78: "=",
		79: "{",
		80: "}",
		81: "[",
		82: "]",
		83: "|",
		84: "\\",
		85: ":",
		86: ";",
		87: "\"",
		88: "'",
		89: "<",
		90: ">",
		91: ",",
		92: "?",
	}

	graph := map[int][]*Brick{}

	identifier := 0

	for _, line := range strings.Split(input, "\n") {
		id := identifiers[identifier]
		identifier++
		if identifier > 92 {
			identifier = 0
		}
		b := &Brick{ID: id}
		parts := strings.Split(line, "~")
		coords := []Coord{}
		for _, c := range parts {
			v := strings.Split(c, ",")
			x, _ := strconv.Atoi(v[0])
			y, _ := strconv.Atoi(v[1])
			z, _ := strconv.Atoi(v[2])
			coords = append(coords, Coord{x, y, z})
		}
		b.Left = coords[0]
		b.Right = coords[1]

		for z := b.Left.Z; z <= b.Right.Z; z++ {
			if _, ok := graph[z]; !ok {
				graph[z] = []*Brick{}
			}
			graph[z] = append(graph[z], b)
		}
	}

	fall(graph)

	return len(disintegrateables(graph))
}

func disintegrateables(graph map[int][]*Brick) []*Brick {
	d := []*Brick{}
	levels := maps.Keys(graph)
	slices.Sort(levels)

	debug(graph)

	for l := 0; l < len(levels); l++ {
		currentLevel := levels[l]
		// If it is the top level, every block can be disintegrated
		if l == len(levels)-1 {
			d = append(d, graph[currentLevel]...)
			continue
		}

		// If there is only one block on the line it can't be disintegrated
		if len(graph[currentLevel]) < 2 {
			continue
		}

		supporteds := map[*Brick][]*Brick{}
		for _, b := range graph[currentLevel] {
			if currentLevel != b.Right.Z {
				continue
			}

			upperLevel := levels[l+1]
			queue := graph[upperLevel]
			supporteds[b] = []*Brick{}
			for len(queue) > 0 {
				b2 := queue[0]
				queue = queue[1:]

				if (b.Left.X > b2.Right.X || b.Right.X < b2.Left.X) ||
					(b.Left.Y > b2.Right.Y || b.Right.Y < b2.Left.Y) {
					continue
				}
				supporteds[b] = append(supporteds[b], b2)
			}
		}

		for b, s := range supporteds {
			if len(s) == 0 {
				d = append(d, b)
				continue
			}

			mutual := 0
			for b2, s2 := range supporteds {
				if b == b2 || len(s2) == 0 {
					continue
				}

				for _, supported := range s {
					if slices.Contains(s2, supported) {
						mutual++
					}
				}
			}
			if mutual == len(s) {
				d = append(d, b)
			}
		}
	}

	return d
}

func fall(graph map[int][]*Brick) {
	levels := maps.Keys(graph)
	slices.Sort(levels)

	// Make the bricks fall downward as far as they can go
	for l := 0; l < len(levels); l++ {
		if l == 0 {
			continue
		}

		currentLevel := levels[l]
		currentLevelBricks := make([]*Brick, len(graph[currentLevel]))
		copy(currentLevelBricks, graph[currentLevel])
		for _, b := range currentLevelBricks {
			// There may be bricks that span over the current level
			if b.Left.Z < currentLevel {
				continue
			}
			newLevel := currentLevel
		LowerLevel:
			for ll := 1; l-ll >= 0; ll++ {
				lowerLevel := levels[l-ll]
				queue := graph[lowerLevel]
				for len(queue) > 0 {
					b2 := queue[0]
					queue = queue[1:]
					if (b.Left.X > b2.Right.X || b.Right.X < b2.Left.X) ||
						(b.Left.Y > b2.Right.Y || b.Right.Y < b2.Left.Y) {
						continue
					}
					newLevel = b2.Right.Z + 1
					break LowerLevel
				}

				newLevel = lowerLevel
			}

			if newLevel != currentLevel {
				previousToppestLevel := b.Right.Z
				b.Left.Z, b.Right.Z = b.Left.Z-(currentLevel-newLevel), b.Right.Z-(currentLevel-newLevel)
				for ld := previousToppestLevel; ld > b.Right.Z; ld-- {
					i := slices.Index(graph[ld], b)
					if i != -1 {
						graph[ld] = append(graph[ld][:i], graph[ld][i+1:]...)
					}
				}
				graph[newLevel] = append(graph[newLevel], b)
			}
		}

		if len(graph[currentLevel]) == 0 {
			delete(graph, currentLevel)
			continue
		}
	}
}

func Two(input string) int {
	input = strings.Trim(input, "\n")

	return len(input)
}

func debug(graph map[int][]*Brick) {
	levels := maps.Keys(graph)
	slices.Sort(levels)
	graphX := map[int][]string{}
	graphY := map[int][]string{}
	for l := len(levels) - 1; l >= 0; l-- {
		cubesX := strings.Split(strings.Repeat(".", 10), "")
		cubesY := strings.Split(strings.Repeat(".", 10), "")

		level := levels[l]
		graphX[level] = append(graphX[level], cubesX...)
		graphY[level] = append(graphY[level], cubesY...)
	}

	for _, bricks := range maps.Values(graph) {
		for _, b := range bricks {
			for z := b.Left.Z; z <= b.Right.Z; z++ {
				if _, ok := graphX[z]; !ok {
					graphX[z] = strings.Split(strings.Repeat(".", 10), "")
					graphY[z] = strings.Split(strings.Repeat(".", 10), "")
				}
				for x := b.Left.X; x <= b.Right.X; x++ {
					if graphX[z][x] != "." {
						graphX[z][x] = "?"
					} else {
						graphX[z][x] = b.ID
					}
				}
				for y := b.Left.Y; y <= b.Right.Y; y++ {
					if graphY[z][y] != "." {
						graphY[z][y] = "?"
					} else {
						graphY[z][y] = b.ID
					}
				}
			}
		}
	}

	for l := len(levels) - 1; l >= 0; l-- {
		x := graphX[levels[l]]
		fmt.Println(strings.Join(x, ""))
	}
	fmt.Println("-----------------------------")
	for l := len(levels) - 1; l >= 0; l-- {
		x := graphY[levels[l]]
		fmt.Println(strings.Join(x, ""))
	}
	for l := len(levels) - 1; l >= 0; l-- {
		bricks := graph[levels[l]]
		slices.SortFunc(bricks, func(a, b *Brick) int {
			return a.Left.X - b.Left.X
		})
		for _, b := range bricks {
			fmt.Println(b.Left.X, ",", b.Left.Y, ",", b.Left.Z, "~", b.Right.X, ",", b.Right.Y, ",", b.Right.Z, b.ID)
		}
	}

}
