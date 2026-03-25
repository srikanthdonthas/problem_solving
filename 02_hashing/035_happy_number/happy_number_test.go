package happynumber

import (
	"testing"
)

func TestIsHappy(t *testing.T) {
	tests := []struct {
		name string
		n int
		expected bool
	}{
		{name: "Example 1", n: 19, expected: true},
		{name: "Example 2", n: 2, expected: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isHappy(tt.n)
			if result != tt.expected {
				t.Errorf("isHappy(%v) = %v, want %v", tt.n, result, tt.expected)
			}
		})
	}
}
