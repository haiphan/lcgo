package problems

import (
	"reflect"
	"testing"
)

func TestFindFinalValue(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		original int
		want     int
	}{
		{
			name:     "Example 1",
			nums:     []int{5, 3, 6, 1, 12},
			original: 3,
			want:     24,
		},
		{
			name:     "Example 2",
			nums:     []int{2, 7, 9},
			original: 4,
			want:     4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findFinalValue(tt.nums, tt.original)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findFinalValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
