package problems

import (
	"testing"
)

func TestStockSpanner(t *testing.T) {
	tests := []struct {
		name  string
		price []int
		want  []int
	}{
		{
			name:  "Example 1",
			price: []int{100, 80, 60, 70, 60, 75, 85},
			want:  []int{1, 1, 1, 2, 1, 4, 6},
		},
		{
			name:  "Example 2",
			price: []int{30, 40, 50, 60, 70, 80, 90},
			want:  []int{1, 2, 3, 4, 5, 6, 7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			obj := StockSpannerConstructor()
			for i, p := range tt.price {
				v := obj.Next(p)
				if v != tt.want[i] {
					t.Errorf("StockSpanner.Next(%d) = %d, want %d", p, v, tt.want[i])
				}
			}
		})
	}
}
