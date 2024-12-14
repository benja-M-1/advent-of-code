package main

import (
	"slices"
	"strings"
)

func Two(input string) int {
	input = strings.Trim(input, "\n")
	m, antennas, bounds := parse(input)

	antinodes := traverseTwo(antennas, m, bounds)

	return len(antinodes)
}

func traverseTwo(antennas []map[string]int, m map[int]map[int]string, bounds [][]int) []map[string]int {
	antinodes := []map[string]int{}
	visited := []string{}

	for _, antenna := range antennas {
		antinodes, visited = findAntinodesTwo(antenna, antennas, m, bounds, antinodes, visited)
	}

	return antinodes
}

func findAntinodesTwo(antennaA map[string]int, antennas []map[string]int, m map[int]map[int]string, bounds [][]int, antinodes []map[string]int, visited []string) ([]map[string]int, []string) {
	for _, antennaB := range antennas {
		if antennaA["x"] == antennaB["x"] && antennaA["y"] == antennaB["y"] {
			continue
		}

		k := generateKey(antennaA, antennaB)
		if slices.Contains(visited, k) {
			continue
		}
		visited = append(visited, k)

		frequencyA := m[antennaA["y"]][antennaA["x"]]
		frequencyB := m[antennaB["y"]][antennaB["x"]]
		sameFrequency := frequencyA == frequencyB

		if !sameFrequency {
			continue
		}

		if !isAntinode(antennaA, antinodes) {
			antinodes = append(antinodes, antennaA)
		}

		if !isAntinode(antennaB, antinodes) {
			antinodes = append(antinodes, antennaB)
		}

		for _, antinode := range nextAntinodesDeep(antennaA, antennaB, bounds) {
			if isAntinode(antinode, antinodes) {
				continue
			}

			if sameFrequency {
				if m[antinode["y"]][antinode["x"]] == "." {
					m[antinode["y"]][antinode["x"]] = "#"
				}

				antinodes = append(antinodes, antinode)
			}
		}
	}

	return antinodes, visited
}

func nextAntinodesDeep(antennaA, antennaB map[string]int, bounds [][]int) []map[string]int {
	antinodes := []map[string]int{}

	padX := antennaA["x"] - antennaB["x"]
	padY := antennaA["y"] - antennaB["y"]

	antinodeA := map[string]int{"x": antennaA["x"] + padX, "y": antennaA["y"] + padY}
	for isWithinBounds(antinodeA["x"], antinodeA["y"], bounds) {
		antinodes = append(antinodes, map[string]int{"x": antinodeA["x"], "y": antinodeA["y"]})
		antinodeA["x"] += padX
		antinodeA["y"] += padY
	}

	antinodeB := map[string]int{"x": antennaB["x"] - padX, "y": antennaB["y"] - padY}
	for isWithinBounds(antinodeB["x"], antinodeB["y"], bounds) {
		antinodes = append(antinodes, map[string]int{"x": antinodeB["x"], "y": antinodeB["y"]})
		antinodeB["x"] -= padX
		antinodeB["y"] -= padY
	}

	return antinodes
}
