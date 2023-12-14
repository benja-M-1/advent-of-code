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

func Sum[S ~[]E, E ~int](x S) E {
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
