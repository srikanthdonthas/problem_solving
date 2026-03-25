package minimumwindowsubstring

import (
	"testing"
)

func TestMinWindow(t *testing.T) {
	tests := []struct {
		name string
		s string
		tStr string
		expected string
	}{
		{name: "Example 1", s: "ADOBECODEBANC", tStr: "ABC", expected: "BANC"},
		{name: "Example 2", s: "a", tStr: "a", expected: "a"},
		{name: "Example 3", s: "a", tStr: "aa", expected: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minWindow(tt.s, tt.tStr)
			if result != tt.expected {
				t.Errorf("minWindow(%q, %q) = %v, want %v", tt.s, tt.tStr, result, tt.expected)
			}
		})
	}
}
