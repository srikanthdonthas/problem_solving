package gasstation

import (
	"testing"
)

func TestCanCompleteCircuit(t *testing.T) {
	tests := []struct {
		name string
		gas []int
		cost []int
		expected int
	}{
		{name: "Example 1", gas: []int{1, 2, 3, 4, 5}, cost: []int{3, 4, 5, 1, 2}, expected: 3},
		{name: "Example 2", gas: []int{2, 3, 4}, cost: []int{3, 4, 3}, expected: -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := canCompleteCircuit(tt.gas, tt.cost)
			if result != tt.expected {
				t.Errorf("canCompleteCircuit(%v, %v) = %v, want %v", tt.gas, tt.cost, result, tt.expected)
			}
		})
	}
}
