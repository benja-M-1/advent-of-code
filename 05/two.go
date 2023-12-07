package main

import (
	"sort"
	"strconv"
	"strings"
)

func PuzzleTwo(input string) int {
	input = strings.Trim(input, "\n")

	var (
		location  int
		intervals Intervals
	)
	currentMap := -1
	maps := make([][][]int, 7)
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "seeds: ") {
			parts := strings.Fields(line)
			for i := 1; i < len(parts[1:]); i += 2 {
				start, _ := strconv.Atoi(parts[i])
				length, _ := strconv.Atoi(parts[i+1])
				intervals = append(intervals, Interval{start, start + length - 1})
			}

			continue
		}

		if strings.Contains(line, "map:") {
			currentMap++
			continue
		}

		parts := strings.Fields(line)
		destinationStart, _ := strconv.Atoi(parts[0])
		sourceStart, _ := strconv.Atoi(parts[1])
		rangeLength, _ := strconv.Atoi(parts[2])

		maps[currentMap] = append(maps[currentMap], []int{destinationStart, sourceStart, rangeLength})
	}

	intersections := intervals.Intersections()
	var currentIntervals Intervals
	for _, interval := range intersections {
		currentIntervals = Intervals{interval}
		for _, m := range maps {
			nextIntervals := Intervals{}
			for c := 0; c < len(currentIntervals); c++ {
				currentInterval := currentIntervals[c]
				keep := false
				for _, line := range m {
					dst, src, rg := line[0], line[1], line[2]
					srcInterval := Interval{src, src + rg - 1}

					// srcInterval : 	|--|        or        |--|
					// currentInterval:	     |----|    |----|
					if srcInterval.IsBefore(currentInterval) || srcInterval.IsAfter(currentInterval) {
						keep = true
						continue
					}

					// Contains a subset the interval
					// srcInterval : 	|----------|
					// currentInterval:	   |----|
					// result: 			   |----|
					if srcInterval.Contains(currentInterval) {
						keep = false
						nextIntervals = append(nextIntervals, Interval{currentInterval.Start + dst - src, currentInterval.End + dst - src})
						break
					}

					// Contains a subset of the interval
					// srcInterval : 	|----------|
					// currentInterval:		|----------|
					// nextInterval: 		|------|
					// new currentInterval: 	   |---|
					if srcInterval.Start < currentInterval.Start && srcInterval.End <= currentInterval.End {
						keep = false
						nextIntervals = append(nextIntervals, Interval{currentInterval.Start + dst - src, srcInterval.End + dst - src})
						currentInterval = Interval{srcInterval.End + 1, currentInterval.End}
						continue
					}

					// Contains a subset of the interval
					// srcInterval : 			|----------|
					// currentInterval:		|---------|
					// nextInterval: 			|-----|
					// new currentInterval: |---|
					if srcInterval.Start >= currentInterval.Start && srcInterval.End > currentInterval.End {
						keep = false
						nextIntervals = append(nextIntervals, Interval{srcInterval.Start + dst - src, currentInterval.End + dst - src})
						currentInterval = Interval{currentInterval.Start, srcInterval.Start - 1}
						continue
					}

					// srcInterval: 	|----|
					// currentInterval:	|------|
					if currentInterval.Contains(srcInterval) && currentInterval.Start == srcInterval.Start {
						keep = false
						nextIntervals = append(nextIntervals, Interval{currentInterval.Start + dst - src, srcInterval.End + dst - src})
						currentInterval = Interval{srcInterval.End + 1, currentInterval.End}
						continue
					}

					// srcInterval: 	  |----|
					// currentInterval:	|------|
					if currentInterval.Contains(srcInterval) && currentInterval.End == srcInterval.End {
						keep = false
						nextIntervals = append(nextIntervals, Interval{srcInterval.Start + dst - src, srcInterval.End + dst - src})
						currentInterval = Interval{currentInterval.End, srcInterval.Start - 1}
						continue
					}

					// Contained in the interval
					// srcInterval : 		|---|
					// currentInterval:	|-----------|
					// result: 			|---|---|---|
					if currentInterval.Contains(srcInterval) {
						keep = false
						nextIntervals = append(nextIntervals, Interval{srcInterval.Start + dst - src, srcInterval.End + dst - src})
						currentIntervals = append(currentIntervals, Interval{srcInterval.End + 1, currentInterval.End})
						currentInterval = Interval{currentInterval.Start, srcInterval.Start - 1}

						continue
					}
				}

				if keep {
					nextIntervals = append(nextIntervals, currentInterval)
				}
			}

			if len(nextIntervals) > 0 {
				currentIntervals = nextIntervals
			}
			sort.Sort(currentIntervals)
		}

		if location == 0 || location > currentIntervals[0].Start {
			location = currentIntervals[0].Start
		}
	}

	return location
}
