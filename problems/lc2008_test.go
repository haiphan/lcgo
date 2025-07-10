package problems

import (
	"reflect"
	"testing"
)

func TestMaxTaxiEarnings(t *testing.T) {
	tests := []struct {
		name  string
		n     int
		rides [][]int
		want  int64
	}{
		{
			name:  "Example 1",
			n:     20,
			rides: [][]int{{1, 6, 1}, {3, 10, 2}, {10, 12, 3}, {11, 12, 2}, {12, 15, 2}, {13, 18, 1}},
			want:  20,
		},
		{
			name:  "Example 2",
			n:     5,
			rides: [][]int{{2, 5, 4}, {1, 5, 1}},
			want:  7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxTaxiEarnings(tt.n, tt.rides)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxTaxiEarnings() = %v, want %v", got, tt.want)
			}
		})
	}
}
