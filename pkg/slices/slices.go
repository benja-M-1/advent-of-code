package slices

import (
	"slices"
)

// CompareFunc returns true if any element in s satisfies cmp function.
func CompareFunc[S ~[]E, E ~int](s S, v E, cmp func(a, b E) bool) bool {
	for _, x := range s {
		if cmp(x, v) {
			return true
		}
	}
	return false
}

// ContainsGreater returns true if any element in s is greater than v.
func ContainsGreater[S ~[]E, E ~int](s S, v E) bool {
	return CompareFunc(s, v, func(a, b E) bool { return a > b })
}

// ContainsGreaterOrEqual returns true if any element in s is greater or equal than v.
func ContainsGreaterOrEqual[S ~[]E, E ~int](s S, v E) bool {
	return CompareFunc(s, v, func(a, b E) bool { return a >= b })
}

func Product[S ~[]E, E ~int](x S) E {
	var p E = 1
	for _, v := range x {
		p *= v
	}
	return p
}

func Sum[S ~[]E, E ~int | ~float64](x S) E {
	var p E = 0
	for _, v := range x {
		p += v
	}
	return p
}

func Find[S ~[]E, E any](s S, f func(E) bool) S {
	var r S
	for _, x := range s {
		if f(x) {
			r = append(r, x)
		}
	}
	return r
}

func Diff[S ~[]E, E comparable](s1 S, s2 S) S {
	var d S
	for i := range s1 {
		if s1[i] == s2[i] {
			continue
		}
		d = append(d, s1[i])
	}

	return d
}

func CountIf[S ~[]E, E comparable](s S, cmp func(E) bool) int {
	return len(Find(s, cmp))
}

func Includes[S ~[]E, E comparable](s1, s2 S) bool {
	if len(s1) <= len(s2) {
		return false
	}

	i := slices.Index(s1, s2[0])
	if i == -1 || len(s2) > len(s1[i:]) {
		return false
	}

	for j := 0; j < len(s2); j++ {
		if s2[j] != s1[i+j] {
			return false
		}
	}

	return true
}
