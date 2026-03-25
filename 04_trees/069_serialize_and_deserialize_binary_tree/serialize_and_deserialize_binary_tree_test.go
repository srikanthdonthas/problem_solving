package serializeanddeserializebinarytree

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
	}{
		{name: "Example 1", root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ser := Constructor()
			deser := Constructor()
			data := ser.serialize(tt.root)
			ans := deser.deserialize(data)
			_ = ans
		})
	}
}
