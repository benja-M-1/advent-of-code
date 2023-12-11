package main

import (
	"embed"
	"fmt"
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

func One(input string) string {
	input = strings.Trim(input, "\n")

	sum := 0
	rucksacks := strings.Split(input, "\n")

	for _, items := range rucksacks {
		shared := map[string]int{}
		firstCompartment := items[0 : len(items)/2]
		secondCompartment := items[len(items)/2:]

		// Get the shared items
		for _, item := range firstCompartment {
			if strings.Contains(secondCompartment, string(item)) {
				shared[string(item)]++
			}
		}

		for item := range shared {
			priority := strings.Index("abcdefghijklmnopqrstuvwxyz", string(item)) + 1
			if priority == 0 {
				priority = strings.Index("ABCDEFGHIJKLMNOPQRSTUVWXYZ", string(item)) + 27
			}

			sum += priority
		}
	}

	return strconv.Itoa(sum)
}

func Two(input string) string {
	input = strings.Trim(input, "\n")

	sum := 0
	rucksacks := strings.Split(input, "\n")
	groups := map[int][]string{}
	group := 0

	for i, rucksack := range rucksacks {
		if i > 0 && i%3 == 0 {
			group++
		}
		groups[group] = append(groups[group], rucksack)
	}

	for _, r := range groups {
		shared := map[string]int{}
		for _, item := range r[0] {
			if strings.Contains(r[1], string(item)) && strings.Contains(r[2], string(item)) {
				shared[string(item)]++
			}
		}

		for item := range shared {
			priority := strings.Index("abcdefghijklmnopqrstuvwxyz", string(item)) + 1
			if priority == 0 {
				priority = strings.Index("ABCDEFGHIJKLMNOPQRSTUVWXYZ", string(item)) + 27
			}

			sum += priority
		}
	}

	return strconv.Itoa(sum)
}
