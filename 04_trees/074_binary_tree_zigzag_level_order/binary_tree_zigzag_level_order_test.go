package binarytreezigzaglevelorder

import (
	"testing"
)

func TestZigzagLevelOrder(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		expected [][]int
	}{
		{name: "Example 1", root: &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20}}, expected: [][]int{{3}, {20, 9}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := zigzagLevelOrder(tt.root)
			_ = result
		})
	}
}
