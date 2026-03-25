package numberofconnectedcomponents

import (
	"testing"
)

func TestCountComponents(t *testing.T) {
	tests := []struct {
		name string
		n int
		edges [][]int
		expected int
	}{
		{name: "Example 1", n: 5, edges: [][]int{{0,1},{1,2},{3,4}}, expected: 2},
		{name: "Example 2", n: 5, edges: [][]int{{0,1},{1,2},{2,3},{3,4}}, expected: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countComponents(tt.n, tt.edges)
			if result != tt.expected {
				t.Errorf("countComponents() = %v, want %v", result, tt.expected)
			}
		})
	}
}
