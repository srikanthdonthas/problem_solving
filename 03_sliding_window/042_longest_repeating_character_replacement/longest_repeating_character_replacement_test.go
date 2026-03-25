package longestrepeatingcharacterreplacement

import (
	"testing"
)

func TestCharacterReplacement(t *testing.T) {
	tests := []struct {
		name string
		s string
		k int
		expected int
	}{
		{name: "Example 1", s: "ABAB", k: 2, expected: 4},
		{name: "Example 2", s: "AABABBA", k: 1, expected: 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := characterReplacement(tt.s, tt.k)
			if result != tt.expected {
				t.Errorf("characterReplacement(%q, %v) = %v, want %v", tt.s, tt.k, result, tt.expected)
			}
		})
	}
}
