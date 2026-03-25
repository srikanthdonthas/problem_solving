package subtreeofanothertree

import (
	"testing"
)

func TestIsSubtree(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		subRoot *TreeNode
		expected bool
	}{
		{name: "Example 1", root: &TreeNode{Val: 3, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 5}}, subRoot: &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}}, expected: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isSubtree(tt.root, tt.subRoot)
			if result != tt.expected {
				t.Errorf("isSubtree() = %v, want %v", result, tt.expected)
			}
		})
	}
}
