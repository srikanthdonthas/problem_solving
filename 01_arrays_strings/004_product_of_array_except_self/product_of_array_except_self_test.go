package productofarrayexceptself

import (
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		expected []int
	}{
		{name: "Example 1", nums: []int{1, 2, 3, 4}, expected: []int{24, 12, 8, 6}},
		{name: "Example 2", nums: []int{-1, 1, 0, -3, 3}, expected: []int{0, 0, 9, 0, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := productExceptSelf(tt.nums)
			if len(result) != len(tt.expected) {
				t.Errorf("expected len %d, got %d", len(tt.expected), len(result))
				return
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("productExceptSelf() = %v, want %v", result, tt.expected)
					break
				}
			}
		})
	}
}
