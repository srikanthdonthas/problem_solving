package countnumberofnicesubarrays

import (
	"testing"
)

func TestNumberOfSubarrays(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k int
		expected int
	}{
		{name: "Example 1", nums: []int{1, 1, 2, 1, 1}, k: 3, expected: 2},
		{name: "Example 2", nums: []int{2, 4, 6}, k: 1, expected: 0},
		{name: "Example 3", nums: []int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}, k: 2, expected: 16},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := numberOfSubarrays(tt.nums, tt.k)
			if result != tt.expected {
				t.Errorf("numberOfSubarrays() = %v, want %v", result, tt.expected)
			}
		})
	}
}
