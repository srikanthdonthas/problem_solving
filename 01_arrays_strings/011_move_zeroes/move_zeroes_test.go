package movezeroes

import (
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		expected []int
	}{
		{name: "Example 1", nums: []int{0, 1, 0, 3, 12}, expected: []int{1, 3, 12, 0, 0}},
		{name: "Example 2", nums: []int{0}, expected: []int{0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numsCopy := make([]int, len(tt.nums))
			copy(numsCopy, tt.nums)
			moveZeroes(numsCopy)
			for i := range numsCopy {
				if numsCopy[i] != tt.expected[i] {
					t.Errorf("moveZeroes() got %v, want %v", numsCopy, tt.expected)
					break
				}
			}
		})
	}
}
