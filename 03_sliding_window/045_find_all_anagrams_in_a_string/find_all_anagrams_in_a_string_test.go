package findallanagramsinastring

import (
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	tests := []struct {
		name string
		s string
		p string
		expected []int
	}{
		{name: "Example 1", s: "cbaebabacd", p: "abc", expected: []int{0, 6}},
		{name: "Example 2", s: "abab", p: "ab", expected: []int{0, 1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findAnagrams(tt.s, tt.p)
			if len(result) != len(tt.expected) {
				t.Errorf("expected len %v, got len %v for result %v", len(tt.expected), len(result), result)
				return
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("findAnagrams(%q, %q) = %v, want %v", tt.s, tt.p, result, tt.expected)
					break
				}
			}
		})
	}
}
