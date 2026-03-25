package clonegraph

import (
	"testing"
)

func TestCloneGraph(t *testing.T) {
	tests := []struct {
		name string
		node *Node
	}{
		{name: "Example 1", node: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := cloneGraph(tt.node)
			if result != nil {
				t.Errorf("cloneGraph() expected nil for nil input, got %v", result)
			}
		})
	}
}
