package balancedbinarytree

import (
	"testing"
)

func TestIsBalanced(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		expected bool
	}{
		{name: "Example 1", root: &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20}}, expected: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isBalanced(tt.root)
			if result != tt.expected {
				t.Errorf("isBalanced() = %v, want %v", result, tt.expected)
			}
		})
	}
}
