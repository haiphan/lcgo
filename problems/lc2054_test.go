package problems

import (
	"reflect"
	"testing"
)

func TestMaxTwoEvents(t *testing.T) {
	tests := []struct {
		name   string
		events [][]int
		want   int
	}{
		{
			name:   "Example 1",
			events: [][]int{{1, 3, 2}, {4, 5, 2}, {2, 4, 3}},
			want:   4,
		},
		{
			name:   "Example 2",
			events: [][]int{{1, 3, 2}, {4, 5, 2}, {1, 5, 5}},
			want:   5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxTwoEvents(tt.events)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxTwoEvents() = %v, want %v", got, tt.want)
			}
		})
	}
}
