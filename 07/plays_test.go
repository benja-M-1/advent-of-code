package main

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func PlayPtr(p Play) *Play {
	return &p
}

func TestNewWildcardPlay_BuildAlternative(t *testing.T) {
	type args struct {
		hand string
		bid  int
	}

	bid := 1

	tests := []struct {
		name string
		args args
		want WildcardPlay
	}{
		{
			name: "High card: 32T3K -> 32T3K",
			args: args{hand: "32T3K", bid: bid},
			want: WildcardPlay{
				Play:        NewPlay("32T3K", bid),
				Alternative: nil,
			},
		},
		{
			name: "High card: 6983J -> 69839",
			args: args{hand: "6983J", bid: bid},
			want: WildcardPlay{
				Play:        NewPlay("6983J", bid),
				Alternative: PlayPtr(NewPlay("69839", bid)),
			},
		},
		{
			name: "One pair: J6396 -> 66396",
			args: args{hand: "J6396", bid: bid},
			want: WildcardPlay{
				Play:        NewPlay("J6396", bid),
				Alternative: PlayPtr(NewPlay("66396", bid)),
			},
		},
		{
			name: "One pair: T55J2 -> T5552",
			args: args{hand: "T55J2", bid: bid},
			want: WildcardPlay{
				Play:        NewPlay("T55J2", bid),
				Alternative: PlayPtr(NewPlay("T5552", bid)),
			},
		},
		{
			name: "Two pair: KTJJT -> KTTTT",
			args: args{hand: "KTJJT", bid: bid},
			want: WildcardPlay{
				Play:        NewPlay("KTJJT", bid),
				Alternative: PlayPtr(NewPlay("KTTTT", bid)),
			},
		},
		{
			name: "Three one kind: KKKTJ -> KKKTK",
			args: args{hand: "KKKTJ", bid: bid},
			want: WildcardPlay{
				Play:        NewPlay("KKKTJ", bid),
				Alternative: PlayPtr(NewPlay("KKKTK", bid)),
			},
		},
		{
			name: "Three one kind: KKKJJ -> KKKKK",
			args: args{hand: "KKKJJ", bid: bid},
			want: WildcardPlay{
				Play:        NewPlay("KKKJJ", bid),
				Alternative: PlayPtr(NewPlay("KKKKK", bid)),
			},
		},
		{
			name: "Four one kind: KKKKJ -> KKKKK",
			args: args{hand: "KKKKJ", bid: bid},
			want: WildcardPlay{
				Play:        NewPlay("KKKKJ", bid),
				Alternative: PlayPtr(NewPlay("KKKKK", bid)),
			},
		},
		{
			name: "Five one kind: JJJJJJ -> KKKKK",
			args: args{hand: "JJJJJJ", bid: bid},
			want: WildcardPlay{
				Play:        NewPlay("JJJJJJ", bid),
				Alternative: PlayPtr(NewPlay("KKKKK", bid)),
			},
		},
		{
			name: "High card with more J: 2QJJ4 -> 2QQQ4",
			args: args{hand: "2QJJ4", bid: bid},
			want: WildcardPlay{
				Play:        NewPlay("2QJJ4", bid),
				Alternative: PlayPtr(NewPlay("2QQQ4", bid)),
			},
		},
		{
			name: "Four J: J8JJJ -> 88888",
			args: args{hand: "J8JJJ", bid: bid},
			want: WildcardPlay{
				Play:        NewPlay("J8JJJ", bid),
				Alternative: PlayPtr(NewPlay("88888", bid)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			play := NewWildcardPlay(tt.args.hand, tt.args.bid)
			assert.Equalf(t, tt.want.Alternative, play.Alternative, "NewWildcardPlay(%v, %v)", tt.args.hand, tt.args.bid)
		})
	}
}

func TestByWildcard_Less(t *testing.T) {
	type args struct {
		hand1 string
		hand2 string
	}

	tests := []args{
		{hand1: "JKKK2", hand2: "QQQQ2"},
		{hand1: "282A2", hand2: "2TT5J"},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("%s < %s", tt.hand1, tt.hand2)
		t.Run(name, func(t *testing.T) {
			p1 := NewWildcardPlay(tt.hand1, 1)
			p2 := NewWildcardPlay(tt.hand2, 1)

			p := ByWildcard{p1, p2}
			sort.Sort(p)

			assert.Equal(t, p1, p[0])
		})
	}
}
