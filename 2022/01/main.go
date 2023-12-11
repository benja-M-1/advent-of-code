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
	calories := strings.Split(input, "\n")
	max := 0
	count := 0
	for _, calorie := range calories {
		calorie := strings.TrimSpace(calorie)
		if calorie == "" {
			if count > max {
				max = count
			}
			count = 0
			continue
		}

		c, err := strconv.Atoi(calorie)
		if err != nil {
			panic(err)
		}

		count += c
	}

	return strconv.Itoa(max)
}

func Two(input string) string {
	calories := strings.Split(input, "\n")
	max := 0
	max2 := 0
	max3 := 0
	count := 0
	for _, calorie := range calories {
		calorie := strings.TrimSpace(calorie)
		if calorie != "" {
			c, err := strconv.Atoi(calorie)
			if err != nil {
				panic(err)
			}
			count += c

			continue
		}

		if count > max3 {
			if count > max2 {
				max3 = max2

				if count > max {
					max2 = max
					max = count
				} else {
					max2 = count
				}
			} else {
				max3 = count
			}
		}

		count = 0
	}

	return strconv.Itoa(max + max2 + max3)
}
