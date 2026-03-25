package diameterofbinarytree

import (
	"testing"
)

func TestDiameterOfBinaryTree(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		expected int
	}{
		{name: "Example 1", root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}}, expected: 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := diameterOfBinaryTree(tt.root)
			if result != tt.expected {
				t.Errorf("diameterOfBinaryTree() = %v, want %v", result, tt.expected)
			}
		})
	}
}
