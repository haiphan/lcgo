package problems

import (
	"reflect"
	"testing"
)

func TestGetDescentPeriods(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int64
	}{
		{
			name:   "Example 1",
			prices: []int{3, 2, 1, 4},
			want:   7,
		},
		{
			name:   "Example 2",
			prices: []int{8, 6, 7, 7},
			want:   4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getDescentPeriods(tt.prices)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDescentPeriods() = %v, want %v", got, tt.want)
			}
		})
	}
}
