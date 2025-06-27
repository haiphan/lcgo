package problems

import (
	"lcgo/utils"
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		name  string
		list1 *ListNode
		list2 *ListNode
		want  *ListNode
	}{
		{
			name:  "Example 1",
			list1: utils.CreateLinkedList([]int{1, 2, 4}),
			list2: utils.CreateLinkedList([]int{1, 3, 4}),
			want:  utils.CreateLinkedList([]int{1, 1, 2, 3, 4, 4}),
		},
		{
			name:  "Example 2",
			list1: utils.CreateLinkedList([]int{}),
			list2: utils.CreateLinkedList([]int{0}),
			want:  utils.CreateLinkedList([]int{0}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeTwoLists(tt.list1, tt.list2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
