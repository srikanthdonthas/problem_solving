package wordpattern

import (
	"testing"
)

func TestWordPattern(t *testing.T) {
	tests := []struct {
		name string
		pattern string
		s string
		expected bool
	}{
		{name: "Example 1", pattern: "abba", s: "dog cat cat dog", expected: true},
		{name: "Example 2", pattern: "abba", s: "dog cat cat fish", expected: false},
		{name: "Example 3", pattern: "aaaa", s: "dog cat cat dog", expected: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := wordPattern(tt.pattern, tt.s)
			if result != tt.expected {
				t.Errorf("wordPattern(%q, %q) = %v, want %v", tt.pattern, tt.s, result, tt.expected)
			}
		})
	}
}
