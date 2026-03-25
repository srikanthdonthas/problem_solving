package nextpermutation

import (
	"testing"
)

func TestNextPermutation(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		expected []int
	}{
		{name: "Example 1", nums: []int{1, 2, 3}, expected: []int{1, 3, 2}},
		{name: "Example 2", nums: []int{3, 2, 1}, expected: []int{1, 2, 3}},
		{name: "Example 3", nums: []int{1, 1, 5}, expected: []int{1, 5, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numsCopy := make([]int, len(tt.nums))
			copy(numsCopy, tt.nums)
			nextPermutation(numsCopy)
			for i := range numsCopy {
				if numsCopy[i] != tt.expected[i] {
					t.Errorf("nextPermutation() got %v, want %v", numsCopy, tt.expected)
					break
				}
			}
		})
	}
}
