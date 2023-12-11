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

func One(input string) int {
	input = strings.Trim(input, "\n")

	return markerPosition(input, 4)
}

func markerPosition(input string, length int) int {
	known := []string{}
	for i, c := range strings.Split(input, "") {
		k := slices.Index(known, c)
		if k >= 0 {
			if len(known)-1 == k {
				known = []string{}
			} else {
				known = known[k+1:]
			}
		}

		known = append(known, c)

		if len(known) == length {
			return i + 1
		}
	}

	return 0
}

func Two(input string) int {
	input = strings.Trim(input, "\n")

	return markerPosition(input, 14)
}
