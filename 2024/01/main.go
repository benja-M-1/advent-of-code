package main

import (
	"embed"
	"fmt"
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

func One(input string) int {
	input = strings.Trim(input, "\n")

	a, b := []int{}, []int{}
	for _, line := range strings.Split(input, "\n") {
		n := strings.Split(line, "   ")
		l, _ := strconv.Atoi(n[0])
		r, _ := strconv.Atoi(n[1])

		a = append(a, l)
		b = append(b, r)
	}

	slices.Sort[[]int](a)
	slices.Sort[[]int](b)

	total := 0
	for i := 0; i < len(a); i++ {
		x := a[i] - b[i]
		if x < 0 {
			total += x * -1
		} else {
			total += x
		}
	}

	return total
}

func Two(input string) int {
	input = strings.Trim(input, "\n")

	a, b := []int{}, []int{}
	for _, line := range strings.Split(input, "\n") {
		n := strings.Split(line, "   ")
		l, _ := strconv.Atoi(n[0])
		r, _ := strconv.Atoi(n[1])

		a = append(a, l)
		b = append(b, r)
	}

	total := 0

	appearances := map[int]int{}
	for i := 0; i < len(a); i++ {
		if appearance, ok := appearances[a[i]]; ok {
			total += a[i] * appearance
			continue
		}

		appearances[a[i]] = 0
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				appearances[a[i]] += 1
			}
		}

		total += a[i] * appearances[a[i]]
	}

	return total
}
