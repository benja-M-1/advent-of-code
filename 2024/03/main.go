package main

import (
	"embed"
	"fmt"
	"strconv"
	"strings"

	"adventofcode/pkg/slices"
)

//go:embed *.txt
var f embed.FS

func main() {
	input, _ := f.ReadFile("input.txt")

	r1 := One(string(input))
	if r1 != 155955228 {
		fmt.Printf("puzzle 1 should be %v, %v given\n", 155955228, r1)
	} else {
		fmt.Printf("puzzle 1: %v\n", r1)
	}

	r2 := Two(string(input))
	if r2 != 100189366 {
		fmt.Printf("puzzle 2 should be %v, %v given\n", 100189366, r2)
	} else {
		fmt.Printf("puzzle 2: %v\n", r2)
	}
}

func One(input string) int {
	input = strings.Trim(input, "\n")

	results := []int{}

	for _, line := range strings.Split(input, "\n") {
		chars := strings.Split(line, "")
		var cur, function, variable string
		var isMultiply bool
		var x, y int

		reset := func() {
			function, variable = "", ""
			x, y = 0, 0
			isMultiply = false
		}

		for i := 0; i < len(chars); i++ {
			switch chars[i] {
			case "(":
				function, cur = cur, ""
				if len(function) >= 3 && function[len(function)-3:] == "mul" {
					isMultiply = true
				}
			case ")":
				if !isMultiply {
					continue
				}
				y, _ = strconv.Atoi(variable)
				results = append(results, x*y)
				reset()
			case ",":
				if !isMultiply {
					continue
				}
				x, _ = strconv.Atoi(variable)
				variable = ""
			case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
				if !isMultiply {
					continue
				}
				variable += chars[i]
				if len(variable) > 3 {
					reset()
				}
			default:
				cur += chars[i]
				if isMultiply {
					reset()
				}
			}
		}
	}

	return slices.Sum(results)
}

func Two(input string) int {
	input = strings.Trim(input, "\n")

	results := []int{}

	isMultiply, enabled := false, true
	var cur, function, variable string
	var x, y int
	for _, line := range strings.Split(input, "\n") {
		chars := strings.Split(line, "")

		reset := func() {
			function, variable = "", ""
			x, y = 0, 0
			isMultiply = false
		}

		for i := 0; i < len(chars); i++ {
			switch chars[i] {
			case "(":
				switch {
				case len(cur) >= 3 && cur[len(cur)-3:] == "mul":
					isMultiply = true
					function = "mul"
				case len(cur) >= 2 && cur[len(cur)-2:] == "do":
					enabled = true
					function = "do"
				case len(cur) >= 5 && cur[len(cur)-5:] == "don't":
					enabled = false
					function = "don't"
				default:
					function = ""
				}

				cur = ""
			case ")":
				if function == "" {
					continue
				}

				if function != "mul" {
					reset()
					continue
				}

				if len(variable) != 0 && x != 0 {
					y, _ = strconv.Atoi(variable)
					if enabled {
						results = append(results, x*y)
					}
				}

				reset()
			case ",":
				if !isMultiply {
					continue
				}
				x, _ = strconv.Atoi(variable)
				variable = ""
			case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
				if !isMultiply {
					continue
				}
				variable += chars[i]
				if len(variable) > 3 {
					reset()
				}
			default:
				cur += chars[i]
				if isMultiply {
					reset()
				}
			}
		}
	}

	return slices.Sum(results)
}
