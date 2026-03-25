package maximumsumsubarrayofsizek

import (
	"testing"
)

func TestMaximumSumSubarray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k int
		expected int
	}{
		{name: "Example 1", nums: []int{100, 200, 300, 400}, k: 2, expected: 700},
		{name: "Example 2", nums: []int{1, 4, 2, 10, 2, 3, 1, 0, 20}, k: 4, expected: 24},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maximumSumSubarray(tt.nums, tt.k)
			if result != tt.expected {
				t.Errorf("maximumSumSubarray() = %v, want %v", result, tt.expected)
			}
		})
	}
}
