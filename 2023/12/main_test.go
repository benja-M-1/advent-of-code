package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = `
???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1
`

func TestPuzzleOne(t *testing.T) {
	assert.Equal(t, 21, One(input))
}

func TestPuzzleTwo(t *testing.T) {
	assert.Equal(t, 525152, Two(input))
}
