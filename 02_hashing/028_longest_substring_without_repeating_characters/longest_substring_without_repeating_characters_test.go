package longestsubstringwithoutrepeatingcharacters

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s string
		expected int
	}{
		{name: "Example 1", s: "abcabcbb", expected: 3},
		{name: "Example 2", s: "bbbbb", expected: 1},
		{name: "Example 3", s: "pwwkew", expected: 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lengthOfLongestSubstring(tt.s)
			if result != tt.expected {
				t.Errorf("lengthOfLongestSubstring(%q) = %v, want %v", tt.s, result, tt.expected)
			}
		})
	}
}
