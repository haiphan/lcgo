package problems

import (
	"reflect"
	"testing"
)

func TestMaximumProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		k      int
		want   int64
	}{
		{
			name:   "Example 1",
			prices: []int{1, 7, 9, 8, 2},
			k:      2,
			want:   14,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximumProfit(tt.prices, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maximumProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
