package problems

import (
	"reflect"
	"testing"
)

func TestCountPartitions3432(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{10, 10, 3, 7, 6},
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countPartitions3432(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countPartitions3432() = %v, want %v", got, tt.want)
			}
		})
	}
}
