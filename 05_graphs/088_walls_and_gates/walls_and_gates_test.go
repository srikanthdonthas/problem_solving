package wallsandgates

import (
	"testing"
)

func TestWallsAndGates(t *testing.T) {
	tests := []struct {
		name string
		rooms [][]int
		expected [][]int
	}{
		{name: "Example 1", rooms: [][]int{{2147483647,-1,0,2147483647},{2147483647,2147483647,2147483647,-1},{2147483647,-1,2147483647,-1},{0,-1,2147483647,2147483647}}, expected: [][]int{{3,-1,0,1},{2,2,1,-1},{1,-1,2,-1},{0,-1,3,4}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wallsAndGates(tt.rooms)
			// verification of the board goes here
		})
	}
}
