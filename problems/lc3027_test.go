package problems

import (
	"reflect"
	"testing"
)

func TestNumberOfPairs2(t *testing.T) {
	tests := []struct {
		name   string
		points [][]int
		want   int
	}{
		{
			name:   "Example 1",
			points: [][]int{{1, 1}, {2, 2}, {3, 3}},
			want:   0,
		},
		{
			name:   "Example 2",
			points: [][]int{{6, 2}, {4, 4}, {2, 6}},
			want:   2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numberOfPairs2(tt.points)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numberOfPairs2() = %v, want %v", got, tt.want)
			}
		})
	}
}
