package findminimuminrotatedsortedarray

import (
	"testing"
)

func TestFindMin(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		expected int
	}{
		{name: "Example 1", nums: []int{3, 4, 5, 1, 2}, expected: 1},
		{name: "Example 2", nums: []int{4, 5, 6, 7, 0, 1, 2}, expected: 0},
		{name: "Example 3", nums: []int{11, 13, 15, 17}, expected: 11},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findMin(tt.nums)
			if result != tt.expected {
				t.Errorf("findMin(%v) = %v, want %v", tt.nums, result, tt.expected)
			}
		})
	}
}
