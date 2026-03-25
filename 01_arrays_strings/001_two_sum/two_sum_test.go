package twosum

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{name: "Example 1", nums: []int{2, 7, 11, 15}, target: 9, expected: []int{0, 1}},
		{name: "Example 2", nums: []int{3, 2, 4}, target: 6, expected: []int{1, 2}},
		{name: "Example 3", nums: []int{3, 3}, target: 6, expected: []int{0, 1}},
		{name: "Example 4", nums: []int{2, 7, 11, 15}, target: 22, expected: []int{1, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := twoSum(tt.nums, tt.target)
			if len(result) == 2 && len(tt.expected) == 2 {
				if (result[0] != tt.expected[0] || result[1] != tt.expected[1]) && (result[0] != tt.expected[1] || result[1] != tt.expected[0]) {
					t.Errorf("twoSum(%v, %v) = %v, want %v", tt.nums, tt.target, result, tt.expected)
				}
			} else if len(result) != len(tt.expected) {
				t.Errorf("twoSum(%v, %v) = %v, want %v", tt.nums, tt.target, result, tt.expected)
			}
		})
	}
}
