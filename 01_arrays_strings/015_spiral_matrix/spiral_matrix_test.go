package spiralmatrix

import (
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	tests := []struct {
		name string
		matrix [][]int
		expected []int
	}{
		{name: "Example 1", matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, expected: []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		{name: "Example 2", matrix: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}, expected: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := spiralOrder(tt.matrix)
			if len(result) != len(tt.expected) {
				t.Errorf("expected len %d, got %d", len(tt.expected), len(result))
				return
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("spiralOrder() = %v, want %v", result, tt.expected)
					break
				}
			}
		})
	}
}
