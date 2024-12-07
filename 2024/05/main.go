package main

import (
	"embed"
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"

	aocslices "adventofcode/pkg/slices"
	aocstrconv "adventofcode/pkg/strconv"
)

//go:embed *.txt
var f embed.FS

func main() {
	input, _ := f.ReadFile("input.txt")

	r1 := One(string(input))
	if r1 != 5374 {
		fmt.Printf("puzzle 1: %v expected, %v given\n", 5374, r1)
	} else {
		fmt.Printf("puzzle 1: %v\n", r1)
	}
	r2 := Two(string(input))
	fmt.Printf("puzzle 2: %v\n", r2)
}

func One(input string) int {
	input = strings.Trim(input, "\n")

	s := strings.Split(input, "\n\n")
	rules, updates := s[0], s[1]

	orders := map[int][]int{}
	for _, order := range strings.Split(rules, "\n") {
		v := strings.Split(order, "|")
		x, _ := strconv.Atoi(v[0])
		y, _ := strconv.Atoi(v[1])

		if _, ok := orders[x]; !ok {
			orders[x] = []int{}
		}
		orders[x] = append(orders[x], y)
	}

	middles := []int{}
	for _, update := range strings.Split(updates, "\n") {
		v := aocstrconv.StringstoI(strings.Split(update, ","))
		isOrdered := true
		for i := 0; i < len(v)-1 && isOrdered; i++ {
			isOrdered = slices.Contains(orders[v[i]], v[i+1])
		}

		if isOrdered {
			i := (len(v) - 1) / 2
			middles = append(middles, v[i])
		}
	}

	return aocslices.Sum(middles)
}

func Two(input string) int {
	input = strings.Trim(input, "\n")

	s := strings.Split(input, "\n\n")
	rules, updates := s[0], s[1]

	orders := map[int][]int{}
	for _, order := range strings.Split(rules, "\n") {
		v := strings.Split(order, "|")
		x, _ := strconv.Atoi(v[0])
		y, _ := strconv.Atoi(v[1])

		if _, ok := orders[x]; !ok {
			orders[x] = []int{}
		}
		orders[x] = append(orders[x], y)
	}

	middles := []int{}
	for _, update := range strings.Split(updates, "\n") {
		v := aocstrconv.StringstoI(strings.Split(update, ","))
		unordered := false
		for i := 0; i < len(v)-1 && !unordered; i++ {
			unordered = !slices.Contains(orders[v[i]], v[i+1])
		}

		if unordered {
			sort.Slice(v, func(i, j int) bool {
				return slices.Contains(orders[v[i]], v[j])
			})
			i := (len(v) - 1) / 2
			middles = append(middles, v[i])
		}
	}

	return aocslices.Sum(middles)
}
