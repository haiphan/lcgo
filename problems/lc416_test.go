package problems

import (
	"reflect"
	"testing"
)

func TestCanPartition(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "Example 1",
			nums: []int{1, 5, 11, 5},
			want: true,
		},
		{
			name: "Example 2",
			nums: []int{1, 2, 3, 5},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canPartition(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("canPartition() = %v, want %v", got, tt.want)
			}
		})
	}
}
