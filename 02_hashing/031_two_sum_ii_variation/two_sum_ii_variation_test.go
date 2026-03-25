package twosumiivariation

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name string
		numbers []int
		target int
		expected []int
	}{
		{name: "Example 1", numbers: []int{2, 7, 11, 15}, target: 9, expected: []int{1, 2}},
		{name: "Example 2", numbers: []int{2, 3, 4}, target: 6, expected: []int{1, 3}},
		{name: "Example 3", numbers: []int{-1, 0}, target: -1, expected: []int{1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := twoSum(tt.numbers, tt.target)
			if len(result) != 2 || result[0] != tt.expected[0] || result[1] != tt.expected[1] {
				t.Errorf("twoSum(%v, %v) = %v, want %v", tt.numbers, tt.target, result, tt.expected)
			}
		})
	}
}
