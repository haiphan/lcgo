package problems

import (
	"reflect"
	"testing"
)

func TestMaxBottlesDrunk(t *testing.T) {
	tests := []struct {
		name        string
		numBottles  int
		numExchange int
		want        int
	}{
		{
			name:        "Example 1",
			numBottles:  13,
			numExchange: 6,
			want:        15,
		},
		{
			name:        "Example 2",
			numBottles:  10,
			numExchange: 3,
			want:        13,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxBottlesDrunk(tt.numBottles, tt.numExchange)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxBottlesDrunk() = %v, want %v", got, tt.want)
			}
		})
	}
}
