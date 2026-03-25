package jumpgameii

import (
	"testing"
)

func TestJump(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		expected int
	}{
		{name: "Example 1", nums: []int{2, 3, 1, 1, 4}, expected: 2},
		{name: "Example 2", nums: []int{2, 3, 0, 1, 4}, expected: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := jump(tt.nums)
			if result != tt.expected {
				t.Errorf("jump(%v) = %v, want %v", tt.nums, result, tt.expected)
			}
		})
	}
}
