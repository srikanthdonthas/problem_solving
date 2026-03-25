package searchinrotatedsortedarray

import (
	"testing"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		target int
		expected int
	}{
		{name: "Example 1", nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 0, expected: 4},
		{name: "Example 2", nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 3, expected: -1},
		{name: "Example 3", nums: []int{1}, target: 0, expected: -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := search(tt.nums, tt.target)
			if result != tt.expected {
				t.Errorf("search(%v, %d) = %v, want %v", tt.nums, tt.target, result, tt.expected)
			}
		})
	}
}
