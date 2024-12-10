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

	topo := map[int]map[int]int{}
	trailheads := [][]int{}

	for y, line := range strings.Split(input, "\n") {
		for x, char := range strings.Split(line, "") {
			if char == "." {
				continue
			}

			height, _ := strconv.Atoi(char)

			if _, ok := topo[y]; !ok {
				topo[y] = map[int]int{}
			}
			topo[y][x] = height

			if height == 0 {
				trailheads = append(trailheads, []int{x, y})
			}
		}
	}

	sum := 0
	for _, trailhead := range trailheads {
		x, y := trailhead[0], trailhead[1]
		thsum := score(x, y, map[int]map[int]int{}, topo, [][]int{trailhead})
		sum += thsum
	}

	return sum
}

func score(x, y int, ends map[int]map[int]int, topo map[int]map[int]int, path [][]int) int {
	sum := 0
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for _, dir := range directions {
		nx := x + dir[0]
		ny := y + dir[1]

		if _, ok := topo[ny][nx]; !ok {
			continue
		}

		if slices.ContainsFunc(path, func(e []int) bool {
			return e[0] == nx && e[1] == ny
		}) {
			continue
		}

		if topo[ny][nx] != topo[y][x]+1 {
			continue
		}

		if topo[ny][nx] == 9 {
			if _, ok := ends[ny][nx]; !ok {
				if _, ok := ends[ny]; !ok {
					ends[ny] = map[int]int{}
				}
				ends[ny][nx] = 9
				sum += 1
			}

			continue
		}

		sum += score(nx, ny, ends, topo, append(path, []int{x, y}))
	}

	return sum
}

func Two(input string) int {
	input = strings.Trim(input, "\n")

	topo := map[int]map[int]int{}
	trailheads := [][]int{}

	for y, line := range strings.Split(input, "\n") {
		for x, char := range strings.Split(line, "") {
			if char == "." {
				continue
			}

			height, _ := strconv.Atoi(char)

			if _, ok := topo[y]; !ok {
				topo[y] = map[int]int{}
			}
			topo[y][x] = height

			if height == 0 {
				trailheads = append(trailheads, []int{x, y})
			}
		}
	}

	sum := 0
	for _, trailhead := range trailheads {
		x, y := trailhead[0], trailhead[1]
		thsum := rating(x, y, topo, [][]int{trailhead})
		sum += thsum
	}

	return sum
}

func rating(x, y int, topo map[int]map[int]int, path [][]int) int {
	count := 0
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for _, dir := range directions {
		nx := x + dir[0]
		ny := y + dir[1]

		if _, ok := topo[ny][nx]; !ok {
			continue
		}

		if slices.ContainsFunc(path, func(e []int) bool {
			return e[0] == nx && e[1] == ny
		}) {
			continue
		}

		if topo[ny][nx] != topo[y][x]+1 {
			continue
		}

		if topo[ny][nx] == 9 {
			count += 1

			continue
		}

		count += rating(nx, ny, topo, append(path, []int{x, y}))
	}

	return count
}
