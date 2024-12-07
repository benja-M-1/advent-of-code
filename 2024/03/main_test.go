package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPuzzleOne(t *testing.T) {
	var input = `
xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))
`
	assert.Equal(t, 161, One(input))
}

func TestPuzzleTwo(t *testing.T) {
	var input = `
xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))
`
	assert.Equal(t, 48, Two(input))
}
