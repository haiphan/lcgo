package problems

import (
	"reflect"
	"testing"
)

func TestNumSquares(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "Example 1",
			n:    4,
			want: 1,
		},
		{
			name: "Example 2",
			n:    12,
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numSquares(tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
