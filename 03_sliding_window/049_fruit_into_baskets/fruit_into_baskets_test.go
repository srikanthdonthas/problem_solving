package fruitintobaskets

import (
	"testing"
)

func TestTotalFruit(t *testing.T) {
	tests := []struct {
		name string
		fruits []int
		expected int
	}{
		{name: "Example 1", fruits: []int{1, 2, 1}, expected: 3},
		{name: "Example 2", fruits: []int{0, 1, 2, 2}, expected: 3},
		{name: "Example 3", fruits: []int{1, 2, 3, 2, 2}, expected: 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := totalFruit(tt.fruits)
			if result != tt.expected {
				t.Errorf("totalFruit() = %v, want %v", result, tt.expected)
			}
		})
	}
}
