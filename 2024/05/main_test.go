package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = `
`

func TestPuzzleOne(t *testing.T) {
	assert.Equal(t, 0, One(input))
}

func TestPuzzleTwo(t *testing.T) {
	assert.Equal(t, 0, Two(input))
}
