package problems

import (
	"reflect"
	"testing"
)

func TestNumSubmat(t *testing.T) {
	tests := []struct {
		name string
		mat  [][]int
		want int
	}{
		{
			name: "Example 1",
			mat:  [][]int{{0, 1, 1, 0}, {0, 1, 1, 1}, {1, 1, 1, 0}},
			want: 24,
		},
		{
			name: "Example 2",
			mat:  [][]int{{1, 0, 1}, {1, 1, 0}, {1, 1, 0}},
			want: 13,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numSubmat(tt.mat)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numSubmat() = %v, want %v", got, tt.want)
			}
		})
	}
}
