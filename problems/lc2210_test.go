package problems

import (
	"reflect"
	"testing"
)

func TestCountHillValley(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{2, 4, 1, 1, 6, 5},
			want: 3,
		},
		{
			name: "Example 2",
			nums: []int{6, 6, 5, 5, 4, 1},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countHillValley(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countHillValley() = %v, want %v", got, tt.want)
			}
		})
	}
}
