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

	return checksum(compressByBlock(layout(input)))
}

func compressByBlock(layout []string) []string {
	for slices.Index(layout, ".") != -1 {
		i := slices.Index(layout, ".")
		layout[i], layout[len(layout)-1] = layout[len(layout)-1], layout[i]
		layout = layout[0 : len(layout)-1]
	}
	return layout
}

func layout(input string) []string {
	l := []string{}
	id := 0
	for i, char := range input {
		v, _ := strconv.Atoi(string(char))
		s := "."
		if i%2 == 0 {
			s = strconv.Itoa(id)
			id++
		}

		for ; v > 0; v-- {
			l = append(l, s)
		}
	}
	return l
}

func checksum(layout []string) int {
	checksum := 0
	for i := 0; i < len(layout); i++ {
		v, _ := strconv.Atoi(layout[i])
		checksum += i * v
	}

	return checksum
}

func Two(input string) int {
	input = strings.Trim(input, "\n")

	addresses := [][]string{}
	id := 0
	for i, char := range input {
		v, _ := strconv.Atoi(string(char))
		s := "."
		if i%2 == 0 {
			s = strconv.Itoa(id)
			id++
		}

		if v == 0 {
			continue
		}

		l := []string{}
		for ; v > 0; v-- {
			l = append(l, s)
		}
		addresses = append(addresses, l)
	}

	addresses = compressByFile(addresses)
	layout := []string{}
	for i := range addresses {
		for j := range addresses[i] {
			layout = append(layout, addresses[i][j])
		}
	}

	return checksum(layout)
}

func compressByFile(addresses [][]string) [][]string {
	for i := 0; i < len(addresses); i++ {
		k := addresses[i]

		if k[0] != "." {
			continue
		}

		rev := slices.Clone(addresses)
		slices.Reverse(rev)

		for j, l := range rev {
			i2 := len(addresses) - 1 - j

			if i2 < i || l[0] == "." || len(l) > len(k) {
				continue
			}

			addresses[i] = l
			addresses[i2] = k[:len(l)]
			if len(k)-len(l) > 0 {
				addresses = append(addresses[:i+2], addresses[i+1:]...)
				addresses[i+1] = strings.Split(strings.Repeat(".", len(k)-len(l)), "")
			}
			break
		}
	}

	return addresses
}
