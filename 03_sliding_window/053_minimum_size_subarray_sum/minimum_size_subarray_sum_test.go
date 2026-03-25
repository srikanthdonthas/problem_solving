package minimumsizesubarraysum

import (
	"testing"
)

func TestMinSubArrayLen(t *testing.T) {
	tests := []struct {
		name string
		target int
		nums []int
		expected int
	}{
		{name: "Example 1", target: 7, nums: []int{2, 3, 1, 2, 4, 3}, expected: 2},
		{name: "Example 2", target: 4, nums: []int{1, 4, 4}, expected: 1},
		{name: "Example 3", target: 11, nums: []int{1, 1, 1, 1, 1, 1, 1, 1}, expected: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minSubArrayLen(tt.target, tt.nums)
			if result != tt.expected {
				t.Errorf("minSubArrayLen() = %v, want %v", result, tt.expected)
			}
		})
	}
}
