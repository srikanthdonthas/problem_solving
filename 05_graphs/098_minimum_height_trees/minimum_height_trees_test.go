package minimumheighttrees

import (
	"testing"
)

func TestFindMinHeightTrees(t *testing.T) {
	tests := []struct {
		name string
		n int
		edges [][]int
		expected []int
	}{
		{name: "Example 1", n: 4, edges: [][]int{{1,0},{1,2},{1,3}}, expected: []int{1}},
		{name: "Example 2", n: 6, edges: [][]int{{3,0},{3,1},{3,2},{3,4},{5,4}}, expected: []int{3,4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findMinHeightTrees(tt.n, tt.edges)
			if len(result) != len(tt.expected) {
				t.Errorf("findMinHeightTrees() len = %v, want %v", len(result), len(tt.expected))
			}
		})
	}
}
