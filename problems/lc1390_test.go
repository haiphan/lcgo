package problems

import (
	"reflect"
	"testing"
)

func TestSumFourDivisors(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{21, 4, 7},
			want: 32,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sumFourDivisors(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sumFourDivisors() = %v, want %v", got, tt.want)
			}
		})
	}
}
