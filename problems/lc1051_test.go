package problems

import (
	"reflect"
	"testing"
)

func TestHeightChecker(t *testing.T) {
	tests := []struct {
		name    string
		heights []int
		want    int
	}{
		{
			name:    "Example 1",
			heights: []int{1, 1, 4, 2, 1, 3},
			want:    3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := heightChecker(tt.heights)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("heightChecker() = %v, want %v", got, tt.want)
			}
		})
	}
}
