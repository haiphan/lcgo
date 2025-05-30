package problems

import (
	"reflect"
	"testing"
)

func TestClosestMeetingNode(t *testing.T) {
	tests := []struct {
		name  string
		edges []int
		node1 int
		node2 int
		want  int
	}{
		{
			name:  "Example 1",
			edges: []int{2, 2, 3, -1},
			node1: 0,
			node2: 1,
			want:  2,
		},
		{
			name:  "Example 2",
			edges: []int{1, 2, -1},
			node1: 0,
			node2: 2,
			want:  2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := closestMeetingNode(tt.edges, tt.node1, tt.node2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("closestMeetingNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
