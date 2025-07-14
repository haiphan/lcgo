package problems

import (
	"lcgo/utils"
	"reflect"
	"testing"
)

func TestGetDecimalValue(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		want int
	}{
		{
			name: "Example 1",
			head: utils.CreateLinkedList([]int{1, 0, 1}),
			want: 5,
		},
		{
			name: "Example 2",
			head: utils.CreateLinkedList([]int{0}),
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getDecimalValue(tt.head)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDecimalValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
