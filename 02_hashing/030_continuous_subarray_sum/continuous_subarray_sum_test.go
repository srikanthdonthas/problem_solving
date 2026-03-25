package continuoussubarraysum

import (
	"testing"
)

func TestCheckSubarraySum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k int
		expected bool
	}{
		{name: "Example 1", nums: []int{23, 2, 4, 6, 7}, k: 6, expected: true},
		{name: "Example 2", nums: []int{23, 2, 6, 4, 7}, k: 6, expected: true},
		{name: "Example 3", nums: []int{23, 2, 6, 4, 7}, k: 13, expected: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := checkSubarraySum(tt.nums, tt.k)
			if result != tt.expected {
				t.Errorf("checkSubarraySum(%v, %v) = %v, want %v", tt.nums, tt.k, result, tt.expected)
			}
		})
	}
}
