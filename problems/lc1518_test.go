package problems

import (
	"reflect"
	"testing"
)

func TestNumWaterBottles(t *testing.T) {
	tests := []struct {
		name            string
		numWaterBottles int
		numExchange     int
		want            int
	}{
		{
			name:            "Example 1",
			numWaterBottles: 9,
			numExchange:     3,
			want:            13,
		},
		{
			name:            "Example 2",
			numWaterBottles: 15,
			numExchange:     4,
			want:            19,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numWaterBottles(tt.numWaterBottles, tt.numExchange)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numWaterBottles() = %v, want %v", got, tt.want)
			}
		})
	}
}
