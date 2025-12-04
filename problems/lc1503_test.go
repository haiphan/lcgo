package problems

import (
	"reflect"
	"testing"
)

func TestGetLastMoment(t *testing.T) {
	tests := []struct {
		name        string
		n           int
		left, right []int
		want        int
	}{
		{
			name:  "Example 1",
			n:     4,
			left:  []int{4, 3},
			right: []int{0, 1},
			want:  4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getLastMoment(tt.n, tt.left, tt.right)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLastMoment() = %v, want %v", got, tt.want)
			}
		})
	}
}
