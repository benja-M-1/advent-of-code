package main

import (
	"embed"
	_ "embed"
	"fmt"
)

//go:embed *.txt
var f embed.FS

func main() {
	input, _ := f.ReadFile("input.txt")

	r1 := Puzzle(string(input))
	fmt.Printf("puzzle 1: %v\n", r1)
	if r1 != 532445 {
		fmt.Println("❌ puzzle 1")
	}

	r2 := PuzzleTwo(string(input))
	fmt.Printf("puzzle 2: %v\n", r2)
	if r2 != 79842967 {
		fmt.Println("❌ puzzle 2")
	}
}
