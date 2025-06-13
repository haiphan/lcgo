package problems

import (
	"reflect"
	"testing"
)

func TestMinimizeMax(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		p    int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{10, 1, 2, 7, 1, 3},
			p:    2,
			want: 1,
		},
		{
			name: "Example 2",
			nums: []int{4, 2, 1, 2},
			p:    1,
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minimizeMax(tt.nums, tt.p)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minimizeMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
