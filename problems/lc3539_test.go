package problems

import (
	"reflect"
	"testing"
)

func TestMagicalSum(t *testing.T) {
	tests := []struct {
		name string
		m, k int
		nums []int
		want int
	}{
		{
			name: "Example 1",
			m:    5,
			k:    5,
			nums: []int{1, 10, 100, 10000, 1000000},
			want: 991600007,
		},
		{
			name: "Example 2",
			m:    2,
			k:    2,
			nums: []int{5, 4, 3, 2, 1},
			want: 170,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := magicalSum(tt.m, tt.k, tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("magicalSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
