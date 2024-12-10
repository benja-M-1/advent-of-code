package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPuzzleOne(t *testing.T) {
	for _, test := range []struct {
		name   string
		input  string
		output int
	}{
		{
			"example 1",
			`
...0...
...1...
...2...
6543456
7.....7
8.....8
9.....9
`,
			2,
		},
		{
			"example 2",
			`
..90..9
...1.98
...2..7
6543456
765.987
876....
987....
`,
			4,
		},
		{
			"example 3",
			`
10..9..
2...8..
3...7..
4567654
...8..3
...9..2
.....01
`,
			3,
		},
		{
			"example 4",
			`
89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732
`,
			36,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.output, One(test.input))
		})
	}
}

func TestPuzzleTwo(t *testing.T) {
	for _, test := range []struct {
		name   string
		input  string
		output int
	}{
		{
			"example 1",
			`
.....0.
..4321.
..5..2.
..6543.
..7..4.
..8765.
..9....
`,
			3,
		},
		{
			"example 2",
			`
..90..9
...1.98
...2..7
6543456
765.987
876....
987....
`,
			13,
		},
		{
			"example 3",
			`
012345
123456
234567
345678
4.6789
56789.
`,
			227,
		},
		{
			"example 4",
			`
89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`,
			81,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.output, Two(test.input))
		})
	}
}
