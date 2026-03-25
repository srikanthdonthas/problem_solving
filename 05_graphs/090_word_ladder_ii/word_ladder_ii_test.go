package wordladderii

import (
	"testing"
)

func TestFindLadders(t *testing.T) {
	tests := []struct {
		name string
		beginWord string
		endWord string
		wordList []string
	}{
		{name: "Example 1", beginWord: "hit", endWord: "cog", wordList: []string{"hot","dot","dog","lot","log","cog"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findLadders(tt.beginWord, tt.endWord, tt.wordList)
			_ = result
		})
	}
}
