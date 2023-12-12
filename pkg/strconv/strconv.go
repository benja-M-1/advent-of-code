package strconv

import (
	"strconv"
)

func StringstoI(input []string) []int {
	var is []int
	for _, s := range input {
		i, _ := strconv.Atoi(s)
		is = append(is, i)
	}
	return is
}
