package longestcontinuoussubarraywithabsolutedifflimit

import (
	"testing"
)

func TestLongestSubarray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		limit int
		expected int
	}{
		{name: "Example 1", nums: []int{8, 2, 4, 7}, limit: 4, expected: 2},
		{name: "Example 2", nums: []int{10, 1, 2, 4, 7, 2}, limit: 5, expected: 4},
		{name: "Example 3", nums: []int{4, 2, 2, 2, 4, 4, 2, 2}, limit: 0, expected: 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := longestSubarray(tt.nums, tt.limit)
			if result != tt.expected {
				t.Errorf("longestSubarray() = %v, want %v", result, tt.expected)
			}
		})
	}
}
