package problems

import (
	"reflect"
	"testing"
)

func TestNew21Game(t *testing.T) {
	tests := []struct {
		name   string
		n      int
		k      int
		maxPts int
		want   float64
	}{
		{
			name:   "Example 1",
			n:      10,
			k:      1,
			maxPts: 10,
			want:   1.0,
		},
		{
			name:   "Example 2",
			n:      6,
			k:      1,
			maxPts: 10,
			want:   0.6,
		},
	}
	epsilon := 1e-6
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := new21Game(tt.n, tt.k, tt.maxPts)
			diff := got - tt.want
			if diff < 0 {
				diff = -diff
			}
			if !reflect.DeepEqual(true, diff < epsilon) {
				t.Errorf("new21Game() = %v, want %v", got, tt.want)
			}
		})
	}
}
