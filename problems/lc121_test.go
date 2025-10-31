package problems

import (
	"reflect"
	"testing"
)

func TestMaxProfit121(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{
			name:   "Example 1",
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   5,
		},
		{
			name:   "Example 2",
			prices: []int{7, 6, 4, 3, 1},
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxProfit121(tt.prices)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxProfit121() = %v, want %v", got, tt.want)
			}
		})
	}
}
