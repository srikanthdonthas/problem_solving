package constructbinarytreefrompreorderinorder

import (
	"testing"
)

func TestBuildTree(t *testing.T) {
	tests := []struct {
		name string
		preorder []int
		inorder []int
		expected int
	}{
		{name: "Example 1", preorder: []int{3, 9, 20, 15, 7}, inorder: []int{9, 3, 15, 20, 7}, expected: 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := buildTree(tt.preorder, tt.inorder)
			if result != nil && result.Val != tt.expected {
				t.Errorf("buildTree() root = %v, want %v", result.Val, tt.expected)
			}
		})
	}
}
