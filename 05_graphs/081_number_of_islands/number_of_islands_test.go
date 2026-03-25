package numberofislands

import (
	"testing"
)

func TestNumIslands(t *testing.T) {
	tests := []struct {
		name string
		grid [][]byte
		expected int
	}{
		{name: "Example 1", grid: [][]byte{{'1','1','1','1','0'},{'1','1','0','1','0'},{'1','1','0','0','0'},{'0','0','0','0','0'}}, expected: 1},
		{name: "Example 2", grid: [][]byte{{'1','1','0','0','0'},{'1','1','0','0','0'},{'0','0','1','0','0'},{'0','0','0','1','1'}}, expected: 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := numIslands(tt.grid)
			if result != tt.expected {
				t.Errorf("numIslands() = %v, want %v", result, tt.expected)
			}
		})
	}
}
