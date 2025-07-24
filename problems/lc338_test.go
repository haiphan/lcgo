package problems

import (
	"reflect"
	"testing"
)

func TestCountBits(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		{
			name: "Example 1",
			n:    2,
			want: []int{0, 1, 1},
		},
		{
			name: "Example 2",
			n:    5,
			want: []int{0, 1, 1, 2, 1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countBits(tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
