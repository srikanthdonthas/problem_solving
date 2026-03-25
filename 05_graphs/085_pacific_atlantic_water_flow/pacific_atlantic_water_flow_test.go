package pacificatlanticwaterflow

import (
	"testing"
)

func TestPacificAtlantic(t *testing.T) {
	tests := []struct {
		name string
		heights [][]int
		expected [][]int
	}{
		{name: "Example 1", heights: [][]int{{1,2,2,3,5},{3,2,3,4,4},{2,4,5,3,1},{6,7,1,4,5},{5,1,1,2,4}}, expected: [][]int{{0,4},{1,3},{1,4},{2,2},{3,0},{3,1},{4,0}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := pacificAtlantic(tt.heights)
			if len(result) != len(tt.expected) {
				t.Errorf("pacificAtlantic() len = %v, want %v", len(result), len(tt.expected))
			}
		})
	}
}
