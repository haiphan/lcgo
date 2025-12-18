package problems

import (
	"reflect"
	"testing"
)

func TestMaxProfit3652(t *testing.T) {
	tests := []struct {
		name             string
		prices, strategy []int
		k                int
		want             int64
	}{
		{
			name:     "Example 1",
			prices:   []int{4, 2, 8},
			strategy: []int{-1, 0, 1},
			k:        2,
			want:     10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxProfit3652(tt.prices, tt.strategy, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxProfit3652() = %v, want %v", got, tt.want)
			}
		})
	}
}
