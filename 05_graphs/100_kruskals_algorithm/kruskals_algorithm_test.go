package kruskalsalgorithm

import (
	"testing"
)

func TestKruskal(t *testing.T) {
	tests := []struct {
		name string
		n int
		edges [][]int
		expected int
	}{
		{name: "Example 1", n: 4, edges: [][]int{{0,1,1},{0,2,2},{1,2,2},{1,3,3},{2,3,4}}, expected: 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := kruskal(tt.n, tt.edges)
			if result != tt.expected {
				t.Errorf("kruskal() = %v, want %v", result, tt.expected)
			}
		})
	}
}
