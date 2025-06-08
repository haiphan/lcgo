package problems

import (
	"reflect"
	"testing"
)

func TestLexicalOrder(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		{
			name: "Example 1",
			n:    13,
			want: []int{1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "Example 2",
			n:    20,
			want: []int{1, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 2, 20, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lexicalOrder(tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lexicalOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
