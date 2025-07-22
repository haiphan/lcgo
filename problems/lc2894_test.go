package problems

import (
	"reflect"
	"testing"
)

func TestDifferenceOfSums(t *testing.T) {
	tests := []struct {
		name string
		n    int
		m    int
		want int
	}{
		{
			name: "Example 1",
			n:    10,
			m:    3,
			want: 19,
		},
		{
			name: "Example 2",
			n:    5,
			m:    6,
			want: 15,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := differenceOfSums(tt.n, tt.m)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("differenceOfSums() = %v, want %v", got, tt.want)
			}
		})
	}
}
