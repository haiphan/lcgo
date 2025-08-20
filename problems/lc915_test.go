package problems

import (
	"reflect"
	"testing"
)

func TestPartitionDisjoint(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{5, 0, 3, 8, 6},
			want: 3,
		},
		{
			name: "Example 2",
			nums: []int{1, 1, 1, 0, 6, 12},
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := partitionDisjoint(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partitionDisjoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
