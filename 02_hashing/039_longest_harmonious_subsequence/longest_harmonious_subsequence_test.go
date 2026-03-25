package longestharmonioussubsequence

import (
	"testing"
)

func TestFindLHS(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		expected int
	}{
		{name: "Example 1", nums: []int{1, 3, 2, 2, 5, 2, 3, 7}, expected: 5},
		{name: "Example 2", nums: []int{1, 2, 3, 4}, expected: 2},
		{name: "Example 3", nums: []int{1, 1, 1, 1}, expected: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findLHS(tt.nums)
			if result != tt.expected {
				t.Errorf("findLHS(%v) = %v, want %v", tt.nums, result, tt.expected)
			}
		})
	}
}
