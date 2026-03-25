package recoverbinarysearchtree

import (
	"testing"
)

func TestRecoverTree(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
	}{
		{name: "Example 1", root: &TreeNode{Val: 1, Left: &TreeNode{Val: 3, Right: &TreeNode{Val: 2}}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recoverTree(tt.root)
			// Verify root here
		})
	}
}
