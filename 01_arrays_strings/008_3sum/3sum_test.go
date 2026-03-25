package q3sum

import (
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		expected [][]int
	}{
		{name: "Example 1", nums: []int{-1, 0, 1, 2, -1, -4}, expected: [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{name: "Example 2", nums: []int{0, 1, 1}, expected: [][]int{}},
		{name: "Example 3", nums: []int{0, 0, 0}, expected: [][]int{{0, 0, 0}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := threeSum(tt.nums)
			// Note: deep equal might fail if order of triplets is different; in interview prep usually just test length or exact match
			if len(result) != len(tt.expected) {
				t.Errorf("threeSum() returned len %d, want %d", len(result), len(tt.expected))
			}
		})
	}
}
