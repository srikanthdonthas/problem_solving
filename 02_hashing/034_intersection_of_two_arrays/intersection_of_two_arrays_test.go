package intersectionoftwoarrays

import (
	"testing"
)

func TestIntersection(t *testing.T) {
	tests := []struct {
		name string
		nums1 []int
		nums2 []int
		expected []int
	}{
		{name: "Example 1", nums1: []int{1, 2, 2, 1}, nums2: []int{2, 2}, expected: []int{2}},
		{name: "Example 2", nums1: []int{4, 9, 5}, nums2: []int{9, 4, 9, 8, 4}, expected: []int{9, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := intersection(tt.nums1, tt.nums2)
			if len(result) != len(tt.expected) {
				t.Errorf("expected len %d, got %d", len(tt.expected), len(result))
			}
		})
	}
}
