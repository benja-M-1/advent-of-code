package main

import (
	"slices"
	"sort"
	"strings"

	"golang.org/x/exp/maps"
)

const (
	HighCard = iota
	OnePair
	TwoPair
	ThreeOneKind
	FullHouse
	FourOneKind
	FiveOneKind
)

var regularCards = []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
var weakJCards = []string{"J", "2", "3", "4", "5", "6", "7", "8", "9", "T", "Q", "K", "A"}

type Play struct {
	Hand string
	Bid  int
	Kind int
}

func NewPlay(hand string, bid int) Play {
	play := Play{
		Hand: hand,
		Bid:  bid,
	}
	play.Kind = play.kind()

	return play
}

type WildcardPlay struct {
	Play
	Alternative *Play
}

func NewWildcardPlay(hand string, bid int) WildcardPlay {
	play := WildcardPlay{
		Play: NewPlay(hand, bid),
	}

	if play.HasWildcard() {
		alt := NewPlay(play.buildAlternative(), bid)
		play.Alternative = &alt
	}

	return play
}

func (p *WildcardPlay) buildAlternative() string {
	// If JJJJJ then KKKKK
	if p.Kind == FiveOneKind && p.HasWildcard() {
		return "KKKKK"
	}

	g := p.Group()
	keys := maps.Keys(g)

	// Sort the group by the number of cards and strength
	sort.Slice(keys, func(i, j int) bool {
		if g[keys[i]] == g[keys[j]] {
			return slices.Index(weakJCards, keys[i]) > slices.Index(weakJCards, keys[j])
		}

		return g[keys[i]] > g[keys[j]]
	})

	// Do not use J if it is the most frequent card
	card := keys[0]
	if card == "J" {
		card = keys[1]
	}

	// Replace J by the most frequent card
	return strings.ReplaceAll(p.Hand, "J", card)
}

func (p *WildcardPlay) HasWildcard() bool {
	return strings.Contains(p.Hand, "J")
}

func (p *Play) Group() map[string]int {
	groups := map[string]int{}
	for _, r := range p.Hand {
		card := string(r)
		if _, ok := groups[card]; !ok {
			groups[card] = 1
		} else {
			groups[card] += 1
		}
	}

	return groups
}

func (p *Play) kind() int {
	g := p.Group()
	keys := maps.Keys(g)

	sort.Slice(keys, func(i, j int) bool {
		return g[keys[i]] > g[keys[j]]
	})

	switch {
	case len(g) == 1:
		return FiveOneKind
	case len(g) == 2 && g[keys[0]] == 4:
		return FourOneKind
	case len(g) == 2 && g[keys[0]] == 3:
		return FullHouse
	case len(g) == 3 && g[keys[0]] == 3:
		return ThreeOneKind
	case len(g) == 3 && g[keys[0]] == 2:
		return TwoPair
	case len(g) == 4:
		return OnePair
	default:
		return HighCard
	}
}

type ByWildcard []WildcardPlay

func (p ByWildcard) Len() int      { return len(p) }
func (p ByWildcard) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p ByWildcard) Less(i, j int) bool {
	playi := p[i]
	playj := p[j]

	iKind := playi.Kind
	jKind := playj.Kind

	if playi.Alternative != nil {
		iKind = playi.Alternative.Kind
	}

	if playj.Alternative != nil {
		jKind = playj.Alternative.Kind
	}

	if iKind != jKind {
		return iKind < jKind
	}

	for k := 0; k < 5; k++ {
		ics := slices.Index(weakJCards, string(playi.Hand[k]))
		jcs := slices.Index(weakJCards, string(playj.Hand[k]))

		if ics != jcs {
			return ics < jcs
		}
	}

	return false
}

type ByRegular []Play

func (p ByRegular) Len() int      { return len(p) }
func (p ByRegular) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p ByRegular) Less(i, j int) bool {
	playi := p[i]
	playj := p[j]

	if playi.Kind == playj.Kind {
		for k := 0; k < 5; k++ {
			ics := slices.Index(regularCards, string(playi.Hand[k]))
			jcs := slices.Index(regularCards, string(playj.Hand[k]))

			if ics != jcs {
				return ics < jcs
			}

		}
	}

	return playi.Kind < playj.Kind
}
