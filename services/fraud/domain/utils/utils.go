package utils

import "math"

func IsFibonacci(n int) bool {
	if n < 0 {
		return false
	}

	check1 := 5*n*n + 4
	sqrt1 := int(math.Sqrt(float64(check1)))
	isPerfectSquare1 := sqrt1*sqrt1 == check1

	check2 := 5*n*n - 4
	sqrt2 := int(math.Sqrt(float64(check2)))
	isPerfectSquare2 := sqrt2*sqrt2 == check2

	return isPerfectSquare1 || isPerfectSquare2
}
