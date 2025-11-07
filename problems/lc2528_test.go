package problems

import (
	"reflect"
	"testing"
)

func TestMaxPower(t *testing.T) {
	tests := []struct {
		name     string
		stations []int
		r, k     int
		want     int64
	}{
		{
			name:     "Example 1",
			stations: []int{1, 2, 4, 5, 0},
			r:        1,
			k:        2,
			want:     5,
		},
		{
			name:     "Example 2",
			stations: []int{4, 4, 4, 4},
			r:        0,
			k:        3,
			want:     4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxPower(tt.stations, tt.r, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxPower() = %v, want %v", got, tt.want)
			}
		})
	}
}
