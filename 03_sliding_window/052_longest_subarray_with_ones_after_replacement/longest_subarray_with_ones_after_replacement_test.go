package longestsubarraywithonesafterreplacement

import (
	"testing"
)

func TestLongestSubarray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		expected int
	}{
		{name: "Example 1", nums: []int{1, 1, 0, 1}, expected: 3},
		{name: "Example 2", nums: []int{0, 1, 1, 1, 0, 1, 1, 0, 1}, expected: 5},
		{name: "Example 3", nums: []int{1, 1, 1}, expected: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := longestSubarray(tt.nums)
			if result != tt.expected {
				t.Errorf("longestSubarray() = %v, want %v", result, tt.expected)
			}
		})
	}
}
