package firstmissingpositive

import (
	"testing"
)

func TestFirstMissingPositive(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		expected int
	}{
		{name: "Example 1", nums: []int{1, 2, 0}, expected: 3},
		{name: "Example 2", nums: []int{3, 4, -1, 1}, expected: 2},
		{name: "Example 3", nums: []int{7, 8, 9, 11, 12}, expected: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := firstMissingPositive(tt.nums)
			if result != tt.expected {
				t.Errorf("firstMissingPositive(%v) = %v, want %v", tt.nums, result, tt.expected)
			}
		})
	}
}
