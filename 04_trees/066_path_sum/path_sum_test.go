package pathsum

import (
	"testing"
)

func TestHasPathSum(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		targetSum int
		expected bool
	}{
		{name: "Example 1", root: &TreeNode{Val: 5, Left: &TreeNode{Val: 4}}, targetSum: 9, expected: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := hasPathSum(tt.root, tt.targetSum)
			if result != tt.expected {
				t.Errorf("hasPathSum() = %v, want %v", result, tt.expected)
			}
		})
	}
}
