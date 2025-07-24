package utils_test

import (
	"hexabank/services/fraud/domain/utils"
	"testing"
)

func TestIsFibonacci(t *testing.T) {
	testCases := []struct {
		input    int
		expected bool
		desc     string
	}{
		{-1, false, "negative number"},
		{-5, false, "negative number"},
		{0, true, "first Fibonacci number"},
		{1, true, "second and third Fibonacci number"},
		{2, true, "fourth Fibonacci number"},
		{3, true, "fifth Fibonacci number"},
		{1000000, false, "large non-Fibonacci number"},
		{832040, true, "large Fibonacci number (30th)"},
	}

	for _, tc := range testCases {
		if result := utils.IsFibonacci(tc.input); result != tc.expected {
			t.Errorf("IsFibonacci(%d) = %v; want %v (%s)", tc.input, result, tc.expected, tc.desc)
		}
	}
}
