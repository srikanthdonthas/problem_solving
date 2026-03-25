package lowestcommonancestorbst

import (
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		p *TreeNode
		q *TreeNode
		expected int
	}{
		{name: "Example 1", root: &TreeNode{Val: 6, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 8}}, p: &TreeNode{Val: 2}, q: &TreeNode{Val: 8}, expected: 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lowestCommonAncestor(tt.root, tt.p, tt.q)
			if result != nil && result.Val != tt.expected {
				t.Errorf("lowestCommonAncestor() = %v, want %v", result.Val, tt.expected)
			}
		})
	}
}
