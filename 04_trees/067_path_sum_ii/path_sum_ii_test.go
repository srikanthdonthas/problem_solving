package pathsumii

import (
	"testing"
)

func TestPathSum(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		targetSum int
		expected [][]int
	}{
		{name: "Example 1", root: &TreeNode{Val: 5, Left: &TreeNode{Val: 4}}, targetSum: 9, expected: [][]int{{5, 4}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := pathSum(tt.root, tt.targetSum)
			_ = result
		})
	}
}
