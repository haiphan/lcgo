package problems

import (
	"reflect"
	"testing"
)

func TestMaxValue(t *testing.T) {
	tests := []struct {
		name   string
		events [][]int
		k      int
		want   int
	}{
		{
			name:   "Example 1",
			events: [][]int{{1, 2, 4}, {3, 4, 3}, {2, 3, 1}},
			k:      2,
			want:   7,
		},
		{
			name:   "Example 2",
			events: [][]int{{1, 2, 4}, {3, 4, 3}, {2, 3, 10}},
			k:      2,
			want:   10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxValue(tt.events, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
