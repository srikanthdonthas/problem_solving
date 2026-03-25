package surroundedregions

import (
	"testing"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		name string
		board [][]byte
		expected [][]byte
	}{
		{name: "Example 1", board: [][]byte{{'X','X','X','X'},{'X','O','O','X'},{'X','X','O','X'},{'X','O','X','X'}}, expected: [][]byte{{'X','X','X','X'},{'X','X','X','X'},{'X','X','X','X'},{'X','O','X','X'}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solve(tt.board)
			// check if board matches expected here
		})
	}
}
