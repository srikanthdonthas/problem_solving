package bipartitegraph

import (
	"testing"
)

func TestIsBipartite(t *testing.T) {
	tests := []struct {
		name string
		graph [][]int
		expected bool
	}{
		{name: "Example 1", graph: [][]int{{1,2,3},{0,2},{0,1,3},{0,2}}, expected: false},
		{name: "Example 2", graph: [][]int{{1,3},{0,2},{1,3},{0,2}}, expected: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isBipartite(tt.graph)
			if result != tt.expected {
				t.Errorf("isBipartite() = %v, want %v", result, tt.expected)
			}
		})
	}
}
