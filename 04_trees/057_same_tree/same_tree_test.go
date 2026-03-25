package sametree

import (
	"testing"
)

func TestIsSameTree(t *testing.T) {
	tests := []struct {
		name string
		p *TreeNode
		q *TreeNode
		expected bool
	}{
		{name: "Example 1", p: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}, q: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}, expected: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isSameTree(tt.p, tt.q)
			if result != tt.expected {
				t.Errorf("isSameTree() = %v, want %v", result, tt.expected)
			}
		})
	}
}
