package kthsmallestelementinbst

import (
	"testing"
)

func TestKthSmallest(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		k int
		expected int
	}{
		{name: "Example 1", root: &TreeNode{Val: 3, Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 4}}, k: 1, expected: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := kthSmallest(tt.root, tt.k)
			if result != tt.expected {
				t.Errorf("kthSmallest() = %v, want %v", result, tt.expected)
			}
		})
	}
}
