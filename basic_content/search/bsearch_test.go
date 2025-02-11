package search

import (
	"math"
	"testing"
)

// TestSquareRoot 单元测试
func TestSquareRoot(t *testing.T) {
	tests := []struct {
		input    int
		expected float64
	}{
		{input: 0, expected: 0},
		{input: 1, expected: 1},
		{input: 2, expected: 1.414214}, // Expected value for sqrt(2)
		{input: 4, expected: 2},
		{input: 9, expected: 3},
		{input: 16, expected: 4},
		{input: 25, expected: 5},
		{input: 100, expected: 10},
		{input: 50, expected: 7.071068}, // Expected value for sqrt(50)
	}

	for _, test := range tests {
		t.Run("Testing square root", func(t *testing.T) {
			result := squareRoot(test.input)
			// 比较保留到六位小数的结果
			if math.Abs(result-test.expected) > 1e-6 {
				t.Errorf("For input %d, expected %.6f but got %.6f", test.input, test.expected, result)
			} else {
				t.Logf("For input %d, expected %.6f", test.input, test.expected)
			}
		})
	}
}
