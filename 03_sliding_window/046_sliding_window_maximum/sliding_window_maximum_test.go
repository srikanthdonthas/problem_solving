package slidingwindowmaximum

import (
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k int
		expected []int
	}{
		{name: "Example 1", nums: []int{1, 3, -1, -3, 5, 3, 6, 7}, k: 3, expected: []int{3, 3, 5, 5, 6, 7}},
		{name: "Example 2", nums: []int{1}, k: 1, expected: []int{1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxSlidingWindow(tt.nums, tt.k)
			if len(result) != len(tt.expected) {
				t.Errorf("expected len %v, got %v for result %v", len(tt.expected), len(result), result)
				return
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("maxSlidingWindow() = %v, want %v", result, tt.expected)
					break
				}
			}
		})
	}
}
