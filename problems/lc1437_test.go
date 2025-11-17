package problems

import (
	"reflect"
	"testing"
)

func TestKLengthApart(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want bool
	}{
		{
			name: "Example 1",
			nums: []int{1, 0, 0, 0, 1, 0, 0, 1},
			k:    2,
			want: true,
		},
		{
			name: "Example 2",
			nums: []int{1, 0, 0, 1, 0, 1},
			k:    2,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := kLengthApart(tt.nums, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kLengthApart() = %v, want %v", got, tt.want)
			}
		})
	}
}
