package problems

import (
	"reflect"
	"testing"
)

func TestSubarrayBitwiseORs(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{
			name: "Example 1",
			arr:  []int{0},
			want: 1,
		},
		{
			name: "Example 2",
			arr:  []int{1, 1, 2},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := subarrayBitwiseORs(tt.arr)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subarrayBitwiseORs() = %v, want %v", got, tt.want)
			}
		})
	}
}
