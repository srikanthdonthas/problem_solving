package networkdelaytime

import (
	"testing"
)

func TestNetworkDelayTime(t *testing.T) {
	tests := []struct {
		name string
		times [][]int
		n int
		k int
		expected int
	}{
		{name: "Example 1", times: [][]int{{2,1,1},{2,3,1},{3,4,1}}, n: 4, k: 2, expected: 2},
		{name: "Example 2", times: [][]int{{1,2,1}}, n: 2, k: 1, expected: 1},
		{name: "Example 3", times: [][]int{{1,2,1}}, n: 2, k: 2, expected: -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := networkDelayTime(tt.times, tt.n, tt.k)
			if result != tt.expected {
				t.Errorf("networkDelayTime() = %v, want %v", result, tt.expected)
			}
		})
	}
}
