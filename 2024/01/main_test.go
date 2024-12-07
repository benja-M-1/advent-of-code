package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = `
3   4
4   3
2   5
1   3
3   9
3   3
`

func TestPuzzleOne(t *testing.T) {
	assert.Equal(t, 11, One(input))
}

func TestPuzzleTwo(t *testing.T) {
	assert.Equal(t, 31, Two(input))
}
