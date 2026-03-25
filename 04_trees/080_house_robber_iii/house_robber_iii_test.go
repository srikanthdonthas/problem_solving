package houserobberiii

import (
	"testing"
)

func TestRob(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		expected int
	}{
		{name: "Example 1", root: &TreeNode{Val: 3, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 1}}}, expected: 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := rob(tt.root)
			if result != tt.expected {
				t.Errorf("rob() = %v, want %v", result, tt.expected)
			}
		})
	}
}
