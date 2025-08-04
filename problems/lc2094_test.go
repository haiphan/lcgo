package problems

import (
	"reflect"
	"testing"
)

func TestFindEvenNumbers(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "Example 1",
			nums: []int{2, 1, 3, 0},
			want: []int{102, 120, 130, 132, 210, 230, 302, 310, 312, 320},
		},
		{
			name: "Example 2",
			nums: []int{2, 2, 8, 8, 2},
			want: []int{222, 228, 282, 288, 822, 828, 882},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findEvenNumbers(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findEvenNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
