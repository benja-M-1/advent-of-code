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
	fmt.Printf("puzzle 1: %v\n", r1) // -> 432`
	r2 := Two(string(input))
	fmt.Printf("puzzle 2: %v\n", r2) // -> 488
}

func One(input string) int {
	input = strings.Trim(input, "\n")

	safe := 0
	for _, line := range strings.Split(input, "\n") {
		isSafe := true
		increasing := true
		row := strings.Split(line, " ")
	line:
		for k := 1; k < len(row); k++ {
			level, _ := strconv.Atoi(row[k])
			prev, _ := strconv.Atoi(row[k-1])

			if k == 1 {
				if level < prev {
					increasing = false
				}
			} else {
				if increasing && level < prev {
					isSafe = false
					break line
				}
			}

			diff := level - prev
			if !increasing {
				diff = prev - level
			}

			if diff < 1 || diff > 3 {
				isSafe = false
				break line
			}
		}

		if isSafe {
			safe++
		}
	}

	return safe
}

func Two(input string) int {
	input = strings.Trim(input, "\n")

	safe := 0
	for _, line := range strings.Split(input, "\n") {
		row := strings.Split(line, " ")
		report := make([]int, len(row))
		for k, r := range row {
			level, _ := strconv.Atoi(r)
			report[k] = level
		}

		ok, failed := isSafe(report)
		if ok {
			safe++
			continue
		}

		a := append([]int{}, report[:failed]...)
		a = append(a, report[failed+1:]...)
		if ok, _ := isSafe(a); ok {
			safe++
			continue
		}

		b := append([]int{}, report[:failed+1]...)
		b = append(b, report[failed+2:]...)
		if ok, _ := isSafe(b); ok {
			safe++
			continue
		}

		if failed > 0 {
			c := append([]int{}, report[:failed-1]...)
			c = append(c, report[failed:]...)
			if ok, _ := isSafe(c); ok {
				fmt.Println(report, "safe")
				safe++
				continue
			}
		}
	}

	return safe
}

func isSafe(report []int) (bool, int) {
	increasing := true

	for k := 1; k < len(report); k++ {
		level := report[k]
		prev := report[k-1]

		if k == 1 {
			if level < prev {
				increasing = false
			}
		} else {
			if increasing && level < prev {
				return false, k - 1
			}
		}

		diff := level - prev
		if !increasing {
			diff = prev - level
		}

		if diff < 1 || diff > 3 {
			return false, k - 1
		}
	}

	return true, 0
}
