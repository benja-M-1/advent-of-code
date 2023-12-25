package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = `
1,0,1~1,2,1
0,0,2~2,0,2
0,2,3~2,2,3
0,0,4~0,2,4
2,0,5~2,2,5
0,1,6~2,1,6
1,1,8~1,1,9
`

func TestPuzzleOne(t *testing.T) {
	assert.Equal(t, 5, One(input))
}

func TestPuzzleTwo(t *testing.T) {
	assert.Equal(t, 0, Two(input))
}
