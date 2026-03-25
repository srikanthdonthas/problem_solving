package subarrayswithkdifferentintegers

import (
	"testing"
)

func TestSubarraysWithKDistinct(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k int
		expected int
	}{
		{name: "Example 1", nums: []int{1, 2, 1, 2, 3}, k: 2, expected: 7},
		{name: "Example 2", nums: []int{1, 2, 1, 3, 4}, k: 3, expected: 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := subarraysWithKDistinct(tt.nums, tt.k)
			if result != tt.expected {
				t.Errorf("subarraysWithKDistinct() = %v, want %v", result, tt.expected)
			}
		})
	}
}
