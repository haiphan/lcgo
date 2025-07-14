package problems

import (
	"lcgo/utils"
	"reflect"
	"testing"
)

func TestInsertGreatestCommonDivisors(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		want *ListNode
	}{
		{
			name: "Example 1",
			head: utils.CreateLinkedList([]int{18, 6, 10, 3}),
			want: utils.CreateLinkedList([]int{18, 6, 6, 2, 10, 1, 3}),
		},
		{
			name: "Example 2",
			head: utils.CreateLinkedList([]int{7}),
			want: utils.CreateLinkedList([]int{7}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := insertGreatestCommonDivisors(tt.head)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertGreatestCommonDivisors() = %v, want %v", got, tt.want)
			}
		})
	}
}
