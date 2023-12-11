package main

import (
	"bufio"
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
	var overlaps []string

	s := bufio.NewScanner(strings.NewReader(input))
	for s.Scan() {
		pairs := strings.Split(s.Text(), ",")
		elf1 := stringsToInt(strings.Split(pairs[0], "-"))
		elf2 := stringsToInt(strings.Split(pairs[1], "-"))

		if (isIncluded(elf2[0], elf1[0], elf1[1]) && isIncluded(elf2[1], elf1[0], elf1[1])) ||
			(isIncluded(elf1[0], elf2[0], elf2[1]) && isIncluded(elf1[1], elf2[0], elf2[1])) {
			overlaps = append(overlaps, s.Text())
		}
	}

	return strconv.Itoa(len(overlaps))
}

func stringsToInt(s []string) []int {
	var i []int
	for _, v := range s {
		n, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		i = append(i, n)
	}
	return i
}

func isIncluded(n int, left int, right int) bool {
	return n >= left && n <= right
}

func Two(input string) string {
	input = strings.Trim(input, "\n")
	var overlaps []string

	s := bufio.NewScanner(strings.NewReader(input))
	for s.Scan() {
		pairs := strings.Split(s.Text(), ",")
		elf1 := stringsToInt(strings.Split(pairs[0], "-"))
		elf2 := stringsToInt(strings.Split(pairs[1], "-"))

		if isIncluded(elf2[0], elf1[0], elf1[1]) || isIncluded(elf2[1], elf1[0], elf1[1]) ||
			isIncluded(elf1[0], elf2[0], elf2[1]) || isIncluded(elf1[1], elf2[0], elf2[1]) {
			overlaps = append(overlaps, s.Text())
		}
	}

	fmt.Println(overlaps)

	return strconv.Itoa(len(overlaps))
}
