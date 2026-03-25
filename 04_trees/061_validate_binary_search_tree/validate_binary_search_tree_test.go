package validatebinarysearchtree

import (
	"testing"
)

func TestIsValidBST(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		expected bool
	}{
		{name: "Example 1", root: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, expected: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValidBST(tt.root)
			if result != tt.expected {
				t.Errorf("isValidBST() = %v, want %v", result, tt.expected)
			}
		})
	}
}
