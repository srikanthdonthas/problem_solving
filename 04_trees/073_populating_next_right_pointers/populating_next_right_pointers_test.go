package populatingnextrightpointers

import (
	"testing"
)

func TestConnect(t *testing.T) {
	tests := []struct {
		name string
		root *Node
	}{
		{name: "Example 1", root: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := connect(tt.root)
			_ = result
		})
	}
}
