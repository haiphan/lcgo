package problems

import (
	"lcgo/utils"
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		want *ListNode
	}{
		{
			name: "Example 1",
			head: utils.CreateLinkedList([]int{1, 2, 3}),
			want: utils.CreateLinkedList([]int{3, 2, 1}),
		},
		{
			name: "Example 2",
			head: utils.CreateLinkedList([]int{1, 2, 3, 4}),
			want: utils.CreateLinkedList([]int{4, 3, 2, 1}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseList(tt.head)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
