package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPuzzleOne(t *testing.T) {
	tests := []struct {
		input string
		push  int
		want  int
	}{
		{
			input: `
broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a
`,
			push: 1,
			want: 32,
		},
		{
			input: `
broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a
`,
			push: 1000,
			want: 32000000,
		},
		{
			input: `
broadcaster -> a
%a -> inv, con
&inv -> b
%b -> con
&con -> output
`,
			push: 4,
			want: 17 * 11,
		},
		{
			input: `
broadcaster -> a
%a -> inv, con
&inv -> b
%b -> con
&con -> output
`,
			push: 1000,
			want: 11687500,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.want, One(tt.input, tt.push))
	}
}
