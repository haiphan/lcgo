package problems

import (
	"reflect"
	"testing"
)

func TestCanAliceWin(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "Example 1",
			nums: []int{1, 2, 3, 4, 10},
			want: false,
		},
		{
			name: "Example 2",
			nums: []int{1, 2, 3, 4, 5, 14},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canAliceWin(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("canAliceWin() = %v, want %v", got, tt.want)
			}
		})
	}
}
