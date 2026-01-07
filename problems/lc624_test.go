package problems

import (
	"reflect"
	"testing"
)

func TestMaxDistance624(t *testing.T) {
	tests := []struct {
		name   string
		arrays [][]int
		want   int
	}{
		{
			name:   "Example 1",
			arrays: [][]int{{1, 2, 3}, {4, 5}, {1, 2, 3}},
			want:   4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxDistance624(tt.arrays)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxDistance624() = %v, want %v", got, tt.want)
			}
		})
	}
}
