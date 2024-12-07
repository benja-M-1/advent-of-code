package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = `
7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
`

func TestPuzzleOne(t *testing.T) {
	assert.Equal(t, 2, One(input))
}

func TestPuzzleTwo(t *testing.T) {
	assert.Equal(t, 4, Two(input))
}
