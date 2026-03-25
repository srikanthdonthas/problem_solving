package countgoodnodesinbinarytree

import (
	"testing"
)

func TestGoodNodes(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		expected int
	}{
		{name: "Example 1", root: &TreeNode{Val: 3, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 5}}}, expected: 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := goodNodes(tt.root)
			if result != tt.expected {
				t.Errorf("goodNodes() = %v, want %v", result, tt.expected)
			}
		})
	}
}
