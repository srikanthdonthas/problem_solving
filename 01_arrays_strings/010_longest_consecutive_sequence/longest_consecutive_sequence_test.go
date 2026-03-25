package longestconsecutivesequence

import (
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		expected int
	}{
		{name: "Example 1", nums: []int{100, 4, 200, 1, 3, 2}, expected: 4},
		{name: "Example 2", nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, expected: 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := longestConsecutive(tt.nums)
			if result != tt.expected {
				t.Errorf("longestConsecutive(%v) = %v, want %v", tt.nums, result, tt.expected)
			}
		})
	}
}
