package permutationinstring

import (
	"testing"
)

func TestCheckInclusion(t *testing.T) {
	tests := []struct {
		name string
		s1 string
		s2 string
		expected bool
	}{
		{name: "Example 1", s1: "ab", s2: "eidbaooo", expected: true},
		{name: "Example 2", s1: "ab", s2: "eidboaoo", expected: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := checkInclusion(tt.s1, tt.s2)
			if result != tt.expected {
				t.Errorf("checkInclusion(%q, %q) = %v, want %v", tt.s1, tt.s2, result, tt.expected)
			}
		})
	}
}
