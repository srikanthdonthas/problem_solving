package trappingrainwater

import (
	"testing"
)

func TestTrap(t *testing.T) {
	tests := []struct {
		name string
		height []int
		expected int
	}{
		{name: "Example 1", height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, expected: 6},
		{name: "Example 2", height: []int{4, 2, 0, 3, 2, 5}, expected: 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := trap(tt.height)
			if result != tt.expected {
				t.Errorf("trap(%v) = %v, want %v", tt.height, result, tt.expected)
			}
		})
	}
}
