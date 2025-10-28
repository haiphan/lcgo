package problems

import (
	"reflect"
	"testing"
)

func TestCountValidSelections(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{1, 0, 2, 0, 3},
			want: 2,
		},
		{
			name: "Example 2",
			nums: []int{2, 3, 4, 0, 4, 1, 0},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countValidSelections(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countValidSelections() = %v, want %v", got, tt.want)
			}
		})
	}
}
