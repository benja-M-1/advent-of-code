package main

import (
	_ "embed"
	"strconv"
	"strings"
	"unicode"
)

func PuzzleTwo(input string) int {
	input = strings.Trim(input, "\n")
	matrix := map[int]map[int]string{}
	for r, line := range strings.Split(input, "\n") {
		matrix[r] = map[int]string{}
		n := ""
		for i := 0; i < len(line); i++ {
			c := line[i]

			if unicode.IsNumber(rune(c)) {
				n += string(c)

				if i == len(line)-1 && n != "" {
					matrix[r][i+1-len(n)] = n
				}

				continue
			}

			if n != "" {
				matrix[r][i-len(n)] = n
				n = ""
			}

			if string(c) != "*" {
				continue
			}

			matrix[r][i] = string(c)
		}
	}

	sum := 0
	for ri, row := range matrix {
		for ci, cell := range row {
			if cell != "*" {
				continue
			}

			rows := []map[int]string{
				row,
			}
			if ri > 0 {
				rows = append(rows, matrix[ri-1])
			}
			if ri < len(matrix)-1 {
				rows = append(rows, matrix[ri+1])
			}

			adj := []string{}
			for _, r := range rows {
				for ni, nv := range r {
					if ni > ci+1 || ni+len(nv) < ci || nv == "*" {
						continue
					}

					adj = append(adj, nv)
				}
			}

			if len(adj) == 2 {
				ratio := 1
				for _, a := range adj {
					v, _ := strconv.Atoi(a)
					ratio = ratio * v
				}
				sum += ratio
			}
		}
	}

	return sum
}
