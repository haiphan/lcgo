package problems

import (
	"reflect"
	"testing"
)

func TestMinimumCost3013(t *testing.T) {
	tests := []struct {
		name    string
		nums    []int
		k, dist int
		want    int64
	}{
		{
			name: "Example 1",
			nums: []int{1, 3, 2, 6, 4, 2},
			k:    3,
			dist: 3,
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minimumCost3013(tt.nums, tt.k, tt.dist)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minimumCost3013() = %v, want %v", got, tt.want)
			}
		})
	}
}
