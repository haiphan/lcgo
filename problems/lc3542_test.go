package problems

import (
	"reflect"
	"testing"
)

func TestMinOperations3542(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{0, 2},
			want: 1,
		},
		{
			name: "Example 2",
			nums: []int{3, 1, 2, 1},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minOperations3542(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minOperations3542() = %v, want %v", got, tt.want)
			}
		})
	}
}
