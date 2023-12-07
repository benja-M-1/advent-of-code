package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntervals_Sort(t *testing.T) {
	intervals := Intervals{
		Interval{3, 4},
		Interval{2, 3},
		Interval{1, 2},
	}

	sort.Sort(intervals)

	assert.Equal(t, 1, intervals[0].Start)
	assert.Equal(t, 2, intervals[1].Start)
	assert.Equal(t, 3, intervals[2].Start)
}

func TestIntervals_Intersections(t *testing.T) {
	intervals := Intervals{
		Interval{3, 7},
		Interval{4, 6},
		Interval{1, 5},
		Interval{9, 10},
	}

	intersections := intervals.Intersections()

	assert.Equal(t, 2, len(intersections))
	assert.Equal(t, 1, intersections[0].Start)
	assert.Equal(t, 7, intersections[0].End)
	assert.Equal(t, 9, intersections[1].Start)
	assert.Equal(t, 10, intersections[1].End)
}

func TestIntervals_Intersections2(t *testing.T) {
	intervals := Intervals{
		Interval{1310704671, 1310704671 + 312415190 - 1},  // 1623119860
		Interval{1034820096, 1034820096 + 106131293 - 1},  // 1140951388
		Interval{682397438, 682397438 + 30365957 - 1},     // 712763394
		Interval{2858337556, 2858337556 + 1183890307 - 1}, // 4042227862
		Interval{665754577, 665754577 + 13162298 - 1},     // 678916874
		Interval{2687187253, 2687187253 + 74991378 - 1},   // 2762178630
		Interval{1782124901, 1782124901 + 3190497 - 1},    // 1785315397
		Interval{208902075, 208902075 + 226221606 - 1},    // 435123681
		Interval{4116455504, 4116455504 + 87808390 - 1},   // 4204263893
		Interval{2403629707, 2403629707 + 66592398 - 1},   // 2470222104
	}

	//   208 902 075 -   435 123 681
	//   665 754 577 -   678 916 874
	//   682 397 438 -   712 763 394
	// 1 034 820 096 - 1 140 951 388
	// 1 310 704 671 - 1 623 119 860
	// 1 782 124 901 - 1 785 315 397
	// 2 403 629 707 - 2 470 222 104
	// 2 687 187 253 - 2 762 178 630
	// 2 858 337 556 - 4 042 227 862
	// 4 116 455 504 - 4 204 263 893

	intersections := intervals.Intersections()

	assert.Equal(t, 10, len(intersections))
}
