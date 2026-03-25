package isomorphicstrings

import (
	"testing"
)

func TestIsIsomorphic(t *testing.T) {
	tests := []struct {
		name string
		s string
		tStr string
		expected bool
	}{
		{name: "Example 1", s: "egg", tStr: "add", expected: true},
		{name: "Example 2", s: "foo", tStr: "bar", expected: false},
		{name: "Example 3", s: "paper", tStr: "title", expected: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isIsomorphic(tt.s, tt.tStr)
			if result != tt.expected {
				t.Errorf("isIsomorphic(%q, %q) = %v, want %v", tt.s, tt.tStr, result, tt.expected)
			}
		})
	}
}
