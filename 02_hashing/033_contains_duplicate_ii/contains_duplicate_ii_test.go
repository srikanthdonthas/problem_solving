package containsduplicateii

import (
	"testing"
)

func TestContainsNearbyDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k int
		expected bool
	}{
		{name: "Example 1", nums: []int{1, 2, 3, 1}, k: 3, expected: true},
		{name: "Example 2", nums: []int{1, 0, 1, 1}, k: 1, expected: true},
		{name: "Example 3", nums: []int{1, 2, 3, 1, 2, 3}, k: 2, expected: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := containsNearbyDuplicate(tt.nums, tt.k)
			if result != tt.expected {
				t.Errorf("containsNearbyDuplicate(%v, %v) = %v, want %v", tt.nums, tt.k, result, tt.expected)
			}
		})
	}
}
