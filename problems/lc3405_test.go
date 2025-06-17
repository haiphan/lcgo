package problems

import (
	"reflect"
	"testing"
)

func TestCountGoodArrays(t *testing.T) {
	tests := []struct {
		name string
		n    int
		m    int
		k    int
		want int
	}{
		{
			name: "Example 1",
			n:    4,
			m:    2,
			k:    2,
			want: 6,
		},
		{
			name: "Example 2",
			n:    5,
			m:    2,
			k:    0,
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countGoodArrays(tt.n, tt.m, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countGoodArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
