package problems

import (
	"reflect"
	"testing"
)

func TesKnightProbability(t *testing.T) {
	tests := []struct {
		name string
		n    int
		k    int
		r    int
		c    int
		want float64
	}{
		{
			name: "Example 1",
			n:    3,
			k:    2,
			r:    0,
			c:    0,
			want: 0.0625,
		},
		{
			name: "Example 2",
			n:    1,
			k:    0,
			r:    0,
			c:    0,
			want: 1.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := knightProbability(tt.n, tt.k, tt.r, tt.c)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("knightProbability() = %v, want %v", got, tt.want)
			}
		})
	}
}
