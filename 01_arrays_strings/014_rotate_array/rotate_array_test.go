package rotatearray

import (
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k int
		expected []int
	}{
		{name: "Example 1", nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 3, expected: []int{5, 6, 7, 1, 2, 3, 4}},
		{name: "Example 2", nums: []int{-1, -100, 3, 99}, k: 2, expected: []int{3, 99, -1, -100}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numsCopy := make([]int, len(tt.nums))
			copy(numsCopy, tt.nums)
			rotate(numsCopy, tt.k)
			for i := range numsCopy {
				if numsCopy[i] != tt.expected[i] {
					t.Errorf("rotate() got %v, want %v", numsCopy, tt.expected)
					break
				}
			}
		})
	}
}
