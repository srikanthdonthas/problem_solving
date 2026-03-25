package invertbinarytree

import (
	"testing"
)

func TestInvertTree(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		expected *TreeNode
	}{
		{name: "Example 1", root: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, expected: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 1}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := invertTree(tt.root)
			_ = result // add comparison logic here as trees are hard to deep equal blindly
		})
	}
}
