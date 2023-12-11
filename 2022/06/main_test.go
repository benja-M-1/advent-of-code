package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDaySixPuzzleOne(t *testing.T) {
	tests := []struct {
		args string
		want int
	}{
		{
			args: "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			want: 7,
		},
		{
			args: "bvwbjplbgvbhsrlpgdmjqwftvncz",
			want: 5,
		},
		{
			args: "nppdvjthqldpwncqszvftbrmjlhg",
			want: 6,
		},
		{
			args: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			want: 10,
		},
		{
			args: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			want: 11,
		},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("%s: first marker after character %d", tt.args, tt.want)
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, One(tt.args))
		})
	}
}

func TestDaySixPuzzleTwo(t *testing.T) {
	tests := []struct {
		args string
		want int
	}{
		{
			args: "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			want: 19,
		},
		{
			args: "bvwbjplbgvbhsrlpgdmjqwftvncz",
			want: 23,
		},
		{
			args: "nppdvjthqldpwncqszvftbrmjlhg",
			want: 23,
		},
		{
			args: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			want: 29,
		},
		{
			args: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			want: 26,
		},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("%s: first marker after character %d", tt.args, tt.want)
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, Two(tt.args))
		})
	}
}
