package flattenbinarytreetolinkedlist

import (
	"testing"
)

func TestFlatten(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
	}{
		{name: "Example 1", root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 5}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flatten(tt.root)
			// verification of the flattened tree goes here
		})
	}
}
