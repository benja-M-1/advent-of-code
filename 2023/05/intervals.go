package main

import (
	"sort"
)

type Interval struct {
	Start int
	End   int
}

func (i Interval) Contains(j Interval) bool {
	return i.Start <= j.Start && i.End >= j.End
}

func (i Interval) IsAfter(j Interval) bool {
	return i.Start > j.End
}

func (i Interval) IsBefore(j Interval) bool {
	return i.End < j.Start
}

func (i Interval) Overlap(j Interval) Interval {
	overlap := Interval{i.Start, i.End}

	if j.Start < i.Start {
		overlap.Start = j.Start
	}

	if j.End > i.End {
		overlap.End = j.End
	}

	return overlap
}

type Intervals []Interval

func (i Intervals) Len() int {
	return len(i)
}

func (i Intervals) Less(a, b int) bool {
	return i[a].Start < i[b].Start
}

func (i Intervals) Swap(a, b int) {
	i[a], i[b] = i[b], i[a]
}

func (i Intervals) Intersections() Intervals {
	var intersections Intervals

	sorted := append(Intervals{}, i...)
	sort.Sort(sorted)

	for _, itvl := range sorted {
		if len(intersections) == 0 {
			intersections = append(intersections, itvl)
			continue
		}

	loop:
		for i, itsc := range intersections {
			switch {
			case itsc.Contains(itvl), itvl.IsBefore(itsc) || itvl.IsAfter(itsc):
				intersections = append(intersections, itvl)
				break loop

			case itvl.Contains(itsc):
				itvl.Start = itsc.Start
				itvl.End = itsc.End

			default:
				ovlp := itvl.Overlap(itsc)
				itsc.Start = ovlp.Start
				itsc.End = ovlp.End
			}

			intersections[i] = itsc
		}
	}

	return intersections
}
