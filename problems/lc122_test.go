package problems

import (
	"reflect"
	"testing"
)

func TestMaxProfit122(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{
			name:   "Example 1",
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   7,
		},
		{
			name:   "Example 2",
			prices: []int{1, 2, 3, 4, 5},
			want:   4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxProfit122(tt.prices)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxProfit122() = %v, want %v", got, tt.want)
			}
		})
	}
}
