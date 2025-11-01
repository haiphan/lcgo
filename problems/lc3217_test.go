package problems

import (
	"lcgo/utils"
	"testing"
)

func TestModifiedList(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		head *ListNode
		want []int
	}{
		{
			name: "Example 1",
			nums: []int{1, 2, 3},
			head: utils.CreateLinkedList([]int{1, 2, 3, 4, 5}),
			want: []int{4, 5},
		},
		{
			name: "Example 2",
			nums: []int{1},
			head: utils.CreateLinkedList([]int{1, 2, 1, 2, 1, 2}),
			want: []int{2, 2, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := modifiedList(tt.nums, tt.head)
			i := 0
			for got != nil {
				if i > len(tt.want) || got.Val != tt.want[i] {
					t.Errorf("modifiedList() = %v, want %v", got.Val, tt.want[i])
				}
				got = got.Next
				i++
			}
		})
	}
}
