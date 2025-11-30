package problems

import (
	"reflect"
	"testing"
)

func TestMinSubarray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		p    int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{3, 1, 4, 2},
			p:    6,
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minSubarray(tt.nums, tt.p)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
