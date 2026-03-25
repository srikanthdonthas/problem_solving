package maxconsecutiveonesiii

import (
	"testing"
)

func TestLongestOnes(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k int
		expected int
	}{
		{name: "Example 1", nums: []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, k: 2, expected: 6},
		{name: "Example 2", nums: []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, k: 3, expected: 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := longestOnes(tt.nums, tt.k)
			if result != tt.expected {
				t.Errorf("longestOnes() = %v, want %v", result, tt.expected)
			}
		})
	}
}
