package problems

import (
	"reflect"
	"testing"
)

func TestFindSmallestInteger(t *testing.T) {
	tests := []struct {
		name  string
		nums  []int
		value int
		want  int
	}{
		{
			name:  "Example 1",
			nums:  []int{1, -10, 7, 13, 6, 8},
			value: 5,
			want:  4,
		},
		{
			name:  "Example 2",
			nums:  []int{1, -10, 7, 13, 6, 8},
			value: 7,
			want:  2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findSmallestInteger(tt.nums, tt.value)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSmallestInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}
