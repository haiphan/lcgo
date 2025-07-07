package problems

import (
	"reflect"
	"testing"
)

func TestMaxEvents(t *testing.T) {
	tests := []struct {
		name   string
		events [][]int
		want   int
	}{
		{
			name:   "Example 1",
			events: [][]int{{1, 2}, {2, 3}, {3, 4}},
			want:   3,
		},
		{
			name:   "Example 2",
			events: [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 2}},
			want:   4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxEvents(tt.events)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxEvents() = %v, want %v", got, tt.want)
			}
		})
	}
}
