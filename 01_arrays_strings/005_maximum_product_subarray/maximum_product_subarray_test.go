package maximumproductsubarray

import (
	"testing"
)

func TestMaxProduct(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		expected int
	}{
		{name: "Example 1", nums: []int{2, 3, -2, 4}, expected: 6},
		{name: "Example 2", nums: []int{-2, 0, -1}, expected: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxProduct(tt.nums)
			if result != tt.expected {
				t.Errorf("maxProduct(%v) = %v, want %v", tt.nums, result, tt.expected)
			}
		})
	}
}
