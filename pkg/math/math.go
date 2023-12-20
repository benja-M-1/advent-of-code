package math

// GreatestCommonDivisor returns the greatest common divisor of two integers.
func GreatestCommonDivisor(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

// LeastCommonMultiplied returns the least common multiple of two integers.
func LeastCommonMultiplied(a, b int, n ...int) int {
	result := a * b / GreatestCommonDivisor(a, b)
	for i := 0; i < len(n); i++ {
		result = LeastCommonMultiplied(result, n[i])
	}

	return result
}
