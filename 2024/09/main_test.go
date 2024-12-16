package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = "2333133121414131402"

func TestPuzzleOne(t *testing.T) {
	assert.Equal(t, 1928, One(input))
}

func TestPuzzleTwo(t *testing.T) {
	assert.Equal(t, 2858, Two(input))
}
