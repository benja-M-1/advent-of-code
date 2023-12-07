package main

import (
	"bufio"
	"embed"
	_ "embed"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func DayOne(input string) int {
	var sum int

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		c, err := calibration(scanner.Text())
		if err == nil {
			sum += c
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error occurred: %v\n", err)
	}

	return sum
}

func calibration(input string) (int, error) {
	var (
		digits []int
		s      string
	)
	for _, i := range strings.Split(input, "") {
		v, err := strconv.Atoi(i)
		if err == nil {
			s = ""
			digits = append(digits, v)
			continue
		}

		s += i
		d, err := digit(s)
		if err != nil {
			continue
		}

		digits = append(digits, d)
		s = i
	}

	if len(digits) == 0 {
		return 0, errors.New("no digits found")
	}

	a, b := 0, 0
	if len(digits) > 1 {
		b = len(digits) - 1
	}

	return digits[a]*10 + digits[b], nil
}

func digit(input string) (int, error) {
	numbers := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	for n, v := range numbers {
		if strings.Contains(input, n) {
			return v, nil
		}
	}

	return 0, errors.New("invalid digit")
}

//go:embed *.txt
var f embed.FS

func main() {
	inputPuzzleOne, _ := f.ReadFile("1.txt")
	r := DayOne(string(inputPuzzleOne))

	fmt.Println(r)
	// answer is 53340
}
