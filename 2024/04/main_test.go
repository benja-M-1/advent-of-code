package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = `
MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX
`

/*
*
....XXMAS.
.SAMXMS...
...S..A...
..A.A.MS.X
XMASAMX.MM
X.....XA.A
S.S.S.S.SS
.A.A.A.A.A
..M.M.M.MM
.X.X.XMASX
*/
func TestPuzzleOne(t *testing.T) {
	assert.Equal(t, 18, One(input))
}

/*
*
.M.S......
..A..MSMS.
.M.S.MAA..
..A.ASMSM.
.M.S.M....
..........
S.S.S.S.S.
.A.A.A.A..
M.M.M.M.M.
..........
*/
func TestPuzzleTwo(t *testing.T) {
	assert.Equal(t, 9, Two(input))
}
