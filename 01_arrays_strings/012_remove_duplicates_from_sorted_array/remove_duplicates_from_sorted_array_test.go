package removeduplicatesfromsortedarray

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		expectedK int
		expectedNums []int
	}{
		{name: "Example 1", nums: []int{1, 1, 2}, expectedK: 2, expectedNums: []int{1, 2}},
		{name: "Example 2", nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, expectedK: 5, expectedNums: []int{0, 1, 2, 3, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numsCopy := make([]int, len(tt.nums))
			copy(numsCopy, tt.nums)
			k := removeDuplicates(numsCopy)
			if k != tt.expectedK {
				t.Errorf("removeDuplicates() returned k=%d, want %d", k, tt.expectedK)
			}
			for i := 0; i < k; i++ {
				if numsCopy[i] != tt.expectedNums[i] {
					t.Errorf("removeDuplicates() array got %v, want prefix %v", numsCopy[:k], tt.expectedNums)
					break
				}
			}
		})
	}
}
