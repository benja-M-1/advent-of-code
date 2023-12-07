package main

import (
	"embed"
	"fmt"
)

//go:embed *.txt
var f embed.FS

func main() {
	input, _ := f.ReadFile("input.txt")

	r1 := PuzzleOne(string(input))
	fmt.Printf("puzzle 1: %v\n", r1)
	r2 := PuzzleTwo(string(input))
	fmt.Printf("puzzle 2: %v\n", r2)
}
