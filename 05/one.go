package main

import (
	"strconv"
	"strings"
)

func PuzzleOne(input string) int {
	input = strings.Trim(input, "\n")

	var (
		seeds []int
	)
	currentMap := -1
	maps := make([][][]int, 7)
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "seeds: ") {
			parts := strings.Fields(line)
			for _, s := range parts[1:] {
				seed, _ := strconv.Atoi(s)
				seeds = append(seeds, seed)
			}
			continue
		}

		if strings.Contains(line, "map:") {
			currentMap++
			continue
		}

		parts := strings.Fields(line)
		destinationStart, _ := strconv.Atoi(parts[0])
		sourceStart, _ := strconv.Atoi(parts[1])
		rangeLength, _ := strconv.Atoi(parts[2])

		maps[currentMap] = append(maps[currentMap], []int{destinationStart, sourceStart, rangeLength})
	}

	location := 0
	for _, v := range seeds {
		for _, m := range maps {
			for _, line := range m {
				dest, src, rg := line[0], line[1], line[2]

				if src <= v && v <= src+rg {
					v = dest + v - src
					break
				}
			}
		}

		if location == 0 || v < location {
			location = v
		}
	}

	return location
}
