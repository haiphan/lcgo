package problems

import (
	"reflect"
	"testing"
)

func TestPrefixesDivBy5(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []bool
	}{
		{
			name: "Example 1",
			nums: []int{0, 1, 1},
			want: []bool{true, false, false},
		},
		{
			name: "Example 2",
			nums: []int{1, 1, 1},
			want: []bool{false, false, false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := prefixesDivBy5(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prefixesDivBy5() = %v, want %v", got, tt.want)
			}
		})
	}
}
