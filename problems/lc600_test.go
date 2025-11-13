package problems

import (
	"reflect"
	"testing"
)

func TestFindIntegers(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "Example 1",
			n:    5,
			want: 5,
		},
		{
			name: "Example 2",
			n:    1,
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findIntegers(tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findIntegers() = %v, want %v", got, tt.want)
			}
		})
	}
}
