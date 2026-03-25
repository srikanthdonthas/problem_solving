package validanagram

import (
	"testing"
)

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name string
		s string
		tStr string
		expected bool
	}{
		{name: "Example 1", s: "anagram", tStr: "nagaram", expected: true},
		{name: "Example 2", s: "rat", tStr: "car", expected: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isAnagram(tt.s, tt.tStr)
			if result != tt.expected {
				t.Errorf("isAnagram(%v, %v) = %v, want %v", tt.s, tt.tStr, result, tt.expected)
			}
		})
	}
}
