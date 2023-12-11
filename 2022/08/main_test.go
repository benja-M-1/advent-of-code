package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = `
30373
25512
65332
33549
35390
`

func TestOne(t *testing.T) {
	assert.Equal(t, 21, One(input))
}

func TestTwo(t *testing.T) {
	assert.Equal(t, 8, Two(input))
}
