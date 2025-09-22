package problems

import (
	"reflect"
	"testing"
)

func TestMaxFrequencyElements(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{1, 2, 3},
			want: 3,
		},
		{
			name: "Example 2",
			nums: []int{1, 2, 1, 2, 3},
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxFrequencyElements(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxFrequencyElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
