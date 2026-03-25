package cheapestflightswithinkstops

import (
	"testing"
)

func TestFindCheapestPrice(t *testing.T) {
	tests := []struct {
		name string
		n int
		flights [][]int
		src int
		dst int
		k int
		expected int
	}{
		{name: "Example 1", n: 3, flights: [][]int{{0,1,100},{1,2,100},{0,2,500}}, src: 0, dst: 2, k: 1, expected: 200},
		{name: "Example 2", n: 3, flights: [][]int{{0,1,100},{1,2,100},{0,2,500}}, src: 0, dst: 2, k: 0, expected: 500},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findCheapestPrice(tt.n, tt.flights, tt.src, tt.dst, tt.k)
			if result != tt.expected {
				t.Errorf("findCheapestPrice() = %v, want %v", result, tt.expected)
			}
		})
	}
}
