package besttimetobuyandsellstock

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name string
		prices []int
		expected int
	}{
		{name: "Example 1", prices: []int{7, 1, 5, 3, 6, 4}, expected: 5},
		{name: "Example 2", prices: []int{7, 6, 4, 3, 1}, expected: 0},
		{name: "Single day", prices: []int{5}, expected: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxProfit(tt.prices)
			if result != tt.expected {
				t.Errorf("maxProfit(%v) = %v, want %v", tt.prices, result, tt.expected)
			}
		})
	}
}
