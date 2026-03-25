package redundantconnection

import (
	"testing"
)

func TestFindRedundantConnection(t *testing.T) {
	tests := []struct {
		name string
		edges [][]int
		expected []int
	}{
		{name: "Example 1", edges: [][]int{{1,2},{1,3},{2,3}}, expected: []int{2,3}},
		{name: "Example 2", edges: [][]int{{1,2},{2,3},{3,4},{1,4},{1,5}}, expected: []int{1,4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findRedundantConnection(tt.edges)
			if len(result) != 2 || result[0] != tt.expected[0] || result[1] != tt.expected[1] {
				t.Errorf("findRedundantConnection() = %v, want %v", result, tt.expected)
			}
		})
	}
}
