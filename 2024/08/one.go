package main

import (
	"slices"
	"strings"
)

func One(input string) int {
	input = strings.Trim(input, "\n")
	m, antennas, bounds := parse(input)

	antinodes := traverse(antennas, m, bounds)

	return len(antinodes)
}

func traverse(antennas []map[string]int, m map[int]map[int]string, bounds [][]int) []map[string]int {
	antinodes := []map[string]int{}
	visited := []string{}

	for _, antenna := range antennas {
		antinodes, visited = findAntinodes(antenna, antennas, m, bounds, antinodes, visited)
	}

	return antinodes
}

func findAntinodes(antennaA map[string]int, antennas []map[string]int, m map[int]map[int]string, bounds [][]int, antinodes []map[string]int, visited []string) ([]map[string]int, []string) {
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

		for _, antinode := range nextAntinodes(antennaA, antennaB) {
			if !isWithinBounds(antinode["x"], antinode["y"], bounds) {
				continue
			}

			if isAntinode(antinode, antinodes) {
				continue
			}
			antinodeFrequency := m[antinode["y"]][antinode["x"]]

			if sameFrequency {
				if antinodeFrequency == "." {
					m[antinode["y"]][antinode["x"]] = "#"
				}
				antinodes = append(antinodes, antinode)
			}
		}
	}

	return antinodes, visited
}

func nextAntinodes(antennaA, antennaB map[string]int) []map[string]int {
	padX := antennaA["x"] - antennaB["x"]
	padY := antennaA["y"] - antennaB["y"]

	return []map[string]int{
		{"x": antennaA["x"] + padX, "y": antennaA["y"] + padY},
		{"x": antennaB["x"] - padX, "y": antennaB["y"] - padY},
	}
}
