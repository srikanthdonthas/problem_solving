package setmatrixzeroes

import (
	"testing"
)

func TestSetZeroes(t *testing.T) {
	tests := []struct {
		name string
		matrix [][]int
		expected [][]int
	}{
		{name: "Example 1", matrix: [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}, expected: [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}},
		{name: "Example 2", matrix: [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}, expected: [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setZeroes(tt.matrix)
			for i := range tt.matrix {
				for j := range tt.matrix[i] {
					if tt.matrix[i][j] != tt.expected[i][j] {
						t.Errorf("setZeroes() got %v, want %v", tt.matrix, tt.expected)
						goto done
					}
				}
			}
			done:
		})
	}
}
