package problems

import (
	"reflect"
	"testing"
)

func TestCombinationSum4(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "Example 1",
			nums:   []int{1, 2, 3},
			target: 4,
			want:   7,
		},
		{
			name:   "Example 2",
			nums:   []int{9},
			target: 3,
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combinationSum4(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum4() = %v, want %v", got, tt.want)
			}
		})
	}
}
