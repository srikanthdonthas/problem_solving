package binarysubarrayswithsum

import (
	"testing"
)

func TestNumSubarraysWithSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		goal int
		expected int
	}{
		{name: "Example 1", nums: []int{1, 0, 1, 0, 1}, goal: 2, expected: 4},
		{name: "Example 2", nums: []int{0, 0, 0, 0, 0}, goal: 0, expected: 15},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := numSubarraysWithSum(tt.nums, tt.goal)
			if result != tt.expected {
				t.Errorf("numSubarraysWithSum() = %v, want %v", result, tt.expected)
			}
		})
	}
}
