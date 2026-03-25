package groupanagrams

import (
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		expected [][]string
	}{
		{name: "Example 1", strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"}, expected: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}},
		{name: "Example 2", strs: []string{""}, expected: [][]string{{""}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := groupAnagrams(tt.strs)
			if len(result) != len(tt.expected) {
				t.Errorf("expected len %d, got %d", len(tt.expected), len(result))
			}
		})
	}
}
