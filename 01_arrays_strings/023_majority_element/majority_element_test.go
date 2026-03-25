package majorityelement

import (
	"testing"
)

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		expected int
	}{
		{name: "Example 1", nums: []int{3, 2, 3}, expected: 3},
		{name: "Example 2", nums: []int{2, 2, 1, 1, 1, 2, 2}, expected: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := majorityElement(tt.nums)
			if result != tt.expected {
				t.Errorf("majorityElement(%v) = %v, want %v", tt.nums, result, tt.expected)
			}
		})
	}
}
