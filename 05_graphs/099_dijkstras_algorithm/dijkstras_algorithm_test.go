package dijkstrasalgorithm

import (
	"testing"
)

func TestDijkstra(t *testing.T) {
	tests := []struct {
		name string
		n int
		edges [][]int
		start int
		expected []int
	}{
		{name: "Example 1", n: 3, edges: [][]int{{0,1,4},{0,2,1},{1,2,2}}, start: 0, expected: []int{0, 3, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := dijkstra(tt.n, tt.edges, tt.start)
			if len(result) != len(tt.expected) {
				t.Errorf("dijkstra() len = %v, want %v", len(result), len(tt.expected))
			}
		})
	}
}
