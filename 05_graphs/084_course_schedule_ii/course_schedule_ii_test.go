package coursescheduleii

import (
	"testing"
)

func TestFindOrder(t *testing.T) {
	tests := []struct {
		name string
		numCourses int
		prerequisites [][]int
		expected []int
	}{
		{name: "Example 1", numCourses: 2, prerequisites: [][]int{{1,0}}, expected: []int{0,1}},
		{name: "Example 2", numCourses: 4, prerequisites: [][]int{{1,0},{2,0},{3,1},{3,2}}, expected: []int{0,1,2,3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findOrder(tt.numCourses, tt.prerequisites)
			if len(result) != len(tt.expected) {
				t.Errorf("findOrder() len = %v, want %v", len(result), len(tt.expected))
			}
		})
	}
}
