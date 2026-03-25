package topkfrequentelements

import (
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k int
		expected []int
	}{
		{name: "Example 1", nums: []int{1, 1, 1, 2, 2, 3}, k: 2, expected: []int{1, 2}},
		{name: "Example 2", nums: []int{1}, k: 1, expected: []int{1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := topKFrequent(tt.nums, tt.k)
			if len(result) != len(tt.expected) {
				t.Errorf("expected len %d, got %d", len(tt.expected), len(result))
			}
		})
	}
}
