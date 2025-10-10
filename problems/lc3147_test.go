package problems

import (
	"reflect"
	"testing"
)

func TestMaximumEnergy(t *testing.T) {
	tests := []struct {
		name   string
		energy []int
		k      int
		want   int
	}{
		{
			name:   "Example 1",
			energy: []int{5, 2, -10, -5, 1},
			k:      3,
			want:   3,
		},
		{
			name:   "Example 2",
			energy: []int{-2, -3, -1},
			k:      2,
			want:   -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximumEnergy(tt.energy, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maximumEnergy() = %v, want %v", got, tt.want)
			}
		})
	}
}
