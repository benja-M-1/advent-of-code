package matrix

import (
	"slices"

	"golang.org/x/exp/maps"
)

type Matrix map[int]Row
type Row map[int]string

func (m Matrix) Has(x, y int) bool {
	_, ok := m[y][x]

	return ok
}

func (m Matrix) Get(x, y int) string {
	return m[y][x]
}

func (m Matrix) Set(x, y int, tile string) {
	if _, ok := m[y]; !ok {
		m[y] = Row{}
	}
	m[y][x] = tile
}

func (m Matrix) String() string {
	var s string

	ykeys := maps.Keys(m)
	slices.Sort(ykeys)

	for _, y := range ykeys {
		xkeys := maps.Keys(m[y])
		slices.Sort(xkeys)

		for _, x := range xkeys {
			s += m[y][x]
		}
		s += "\n"
	}

	return s
}
