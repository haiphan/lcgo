package problems

import (
	"reflect"
	"testing"
)

func TestMinPairSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{3, 5, 2, 3},
			want: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minPairSum(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minPairSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
