package problems

import (
	"lcgo/utils"
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		name string
		l1   *ListNode
		l2   *ListNode
		want *ListNode
	}{
		{
			name: "Example 1",
			l1:   utils.CreateLinkedList([]int{2, 4, 3}),
			l2:   utils.CreateLinkedList([]int{5, 6, 4}),
			want: utils.CreateLinkedList([]int{7, 0, 8}),
		},
		{
			name: "Example 2",
			l1:   utils.CreateLinkedList([]int{0}),
			l2:   utils.CreateLinkedList([]int{0}),
			want: utils.CreateLinkedList([]int{0}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addTwoNumbers(tt.l1, tt.l2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
