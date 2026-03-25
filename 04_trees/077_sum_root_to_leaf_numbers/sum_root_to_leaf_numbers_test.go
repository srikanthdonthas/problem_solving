package sumroottoleafnumbers

import (
	"testing"
)

func TestSumNumbers(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		expected int
	}{
		{name: "Example 1", root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}, expected: 25},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sumNumbers(tt.root)
			if result != tt.expected {
				t.Errorf("sumNumbers() = %v, want %v", result, tt.expected)
			}
		})
	}
}
