package problems

import (
	"reflect"
	"testing"
)

func TestGetSneakyNumbers(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "Example 1",
			nums: []int{0, 1, 1, 0},
			want: []int{1, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getSneakyNumbers(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSneakyNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
