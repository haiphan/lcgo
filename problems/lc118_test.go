package problems

import (
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want [][]int
	}{
		{
			name: "Example 1",
			n:    5,
			want: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}},
		},
		{
			name: "Example 2",
			n:    1,
			want: [][]int{{1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generate(tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generate() = %v, want %v", got, tt.want)
			}
		})
	}
}
