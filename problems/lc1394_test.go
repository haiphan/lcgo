package problems

import (
	"reflect"
	"testing"
)

func TestFindLucky(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{
			name: "Example 1",
			arr:  []int{2, 2, 3, 4},
			want: 2,
		},
		{
			name: "Example 2",
			arr:  []int{1, 2, 2, 3, 3, 3},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findLucky(tt.arr)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findLucky() = %v, want %v", got, tt.want)
			}
		})
	}
}
