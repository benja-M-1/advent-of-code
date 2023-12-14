package main

import (
	"embed"
	"fmt"
	"strings"

	aocstrconv "adventofcode/pkg/strconv"

	"github.com/kofalt/go-memoize"
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

	sum := 0
	for _, l := range strings.Split(input, "\n") {
		parts := strings.Fields(l)
		count := walk(parts[0], aocstrconv.StringstoI(strings.Split(parts[1], ",")), 0, 0, memoize.NewMemoizer(-1, -1))
		fmt.Println(parts[0], parts[1], "===>", count)

		sum += count
	}

	return sum
}

// Courtesy of Dadid Gageot!
// I couldn't figure out how to build the walk correclty.
// My first attempt was to recursively replace the / by a # and check if the result was valid and do the same with a ".".
// It worked well for Part 1. And then in Part 2, the compute time was too long.
// I tried to mathematically guess how many possibilities existed in the unfolded string:
// - x the number of solutions for the initial string
// - y the number of solutions for the initial string prefixed with a ?
// Then the result would be x * y^4. But it does not work for the example #7.
// I could not figure out a avoid unnecessary calls.
// eg with `???.### 1,1,3` `##?.###` is already invalid and thus there is no need to try that value
func walk(conditions string, criteria []int, ci, cv int, memo *memoize.Memoizer) int {
	if len(conditions) == 0 {
		if ci == len(criteria)-1 && cv == criteria[ci] {
			return 1
		}

		if ci == len(criteria) && cv == 0 {
			return 1
		}

		return 0
	}

	count := 0
	first, rest := string(conditions[0]), conditions[1:]
	if first == "." || first == "?" {
		if cv == 0 {
			k := fmt.Sprintf("%d-%d-%d", len(conditions), ci, 0)
			c, _, _ := memo.Memoize(k, func() (interface{}, error) {
				return walk(rest, criteria, ci, 0, memo), nil
			})
			count += c.(int)
		} else if criteria[ci] == cv {
			k := fmt.Sprintf("%d-%d-%d", len(conditions), ci+1, 0)
			c, _, _ := memo.Memoize(k, func() (interface{}, error) {
				return walk(rest, criteria, ci+1, 0, memo), nil
			})
			count += c.(int)
		}
	}

	if first == "#" || first == "?" {
		if ci < len(criteria) && cv < criteria[ci] {
			k := fmt.Sprintf("%d-%d-%d", len(conditions), ci, cv+1)
			c, _, _ := memo.Memoize(k, func() (interface{}, error) {
				return walk(rest, criteria, ci, cv+1, memo), nil
			})
			count += c.(int)
		}
	}

	return count
}

func unfold(conditions string, criteria []int) (string, []int) {
	conditions = strings.Repeat(conditions+"?", 4) + conditions

	var unfoldedCriteria []int
	for i := 0; i < 5; i++ {
		unfoldedCriteria = append(unfoldedCriteria, criteria...)
	}

	return conditions, unfoldedCriteria
}

func Two(input string) int {
	input = strings.Trim(input, "\n")

	sum := 0
	for _, l := range strings.Split(input, "\n") {
		parts := strings.Fields(l)
		conditions, criteria := parts[0], aocstrconv.StringstoI(strings.Split(parts[1], ","))
		conditions, criteria = unfold(conditions, criteria)

		count := walk(conditions, criteria, 0, 0, memoize.NewMemoizer(-1, -1))

		fmt.Println(parts[0], parts[1], "===>", count)

		sum += count
	}

	return sum
}
