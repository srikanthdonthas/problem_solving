package candy

import (
	"testing"
)

func TestCandy(t *testing.T) {
	tests := []struct {
		name string
		ratings []int
		expected int
	}{
		{name: "Example 1", ratings: []int{1, 0, 2}, expected: 5},
		{name: "Example 2", ratings: []int{1, 2, 2}, expected: 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := candy(tt.ratings)
			if result != tt.expected {
				t.Errorf("candy(%v) = %v, want %v", tt.ratings, result, tt.expected)
			}
		})
	}
}
