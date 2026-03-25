package sortcolorsdutchflag

import (
	"testing"
)

func TestSortColors(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		expected []int
	}{
		{name: "Example 1", nums: []int{2, 0, 2, 1, 1, 0}, expected: []int{0, 0, 1, 1, 2, 2}},
		{name: "Example 2", nums: []int{2, 0, 1}, expected: []int{0, 1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numsCopy := make([]int, len(tt.nums))
			copy(numsCopy, tt.nums)
			sortColors(numsCopy)
			for i := range numsCopy {
				if numsCopy[i] != tt.expected[i] {
					t.Errorf("sortColors() got %v, want %v", numsCopy, tt.expected)
					break
				}
			}
		})
	}
}
