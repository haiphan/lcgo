package problems

import (
	"reflect"
	"testing"
)

func TestFindKthNumber(t *testing.T) {
	tests := []struct {
		name string
		n    int
		k    int
		want int
	}{
		{
			name: "Example 1",
			n:    13,
			k:    2,
			want: 10,
		},
		{
			name: "Example 2",
			n:    1,
			k:    1,
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findKthNumber(tt.n, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findKthNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
