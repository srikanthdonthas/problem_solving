package subarraysumequalsk

import (
	"testing"
)

func TestSubarraySum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k int
		expected int
	}{
		{name: "Example 1", nums: []int{1, 1, 1}, k: 2, expected: 2},
		{name: "Example 2", nums: []int{1, 2, 3}, k: 3, expected: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := subarraySum(tt.nums, tt.k)
			if result != tt.expected {
				t.Errorf("subarraySum(%v, %v) = %v, want %v", tt.nums, tt.k, result, tt.expected)
			}
		})
	}
}
