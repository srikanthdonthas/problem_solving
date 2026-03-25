package binarytreerightsideview

import (
	"testing"
)

func TestRightSideView(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		expected []int
	}{
		{name: "Example 1", root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}, expected: []int{1, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := rightSideView(tt.root)
			_ = result
		})
	}
}
