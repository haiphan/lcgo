package problems

import (
	"reflect"
	"testing"
)

func TestMinBitwiseArray3315(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "Example 1",
			nums: []int{2, 3, 5, 7},
			want: []int{-1, 1, 4, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minBitwiseArray3315(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minBitwiseArray3315() = %v, want %v", got, tt.want)
			}
		})
	}
}
