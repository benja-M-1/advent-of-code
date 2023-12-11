package slices

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
