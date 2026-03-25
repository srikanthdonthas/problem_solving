package convertsortedarraytobst

import (
	"testing"
)

func TestSortedArrayToBST(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		expectedRoot int
	}{
		{name: "Example 1", nums: []int{-10, -3, 0, 5, 9}, expectedRoot: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sortedArrayToBST(tt.nums)
			if result != nil && result.Val != tt.expectedRoot {
				t.Errorf("sortedArrayToBST() root = %v, want %v", result.Val, tt.expectedRoot)
			}
		})
	}
}
