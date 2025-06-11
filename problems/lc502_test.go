package problems

import (
	"reflect"
	"testing"
)

func TestFindMaximizedCapital(t *testing.T) {
	tests := []struct {
		name    string
		k       int
		w       int
		profits []int
		capital []int
		want    int
	}{
		{
			name:    "Example 1",
			k:       2,
			w:       0,
			profits: []int{1, 2, 3},
			capital: []int{0, 1, 1},
			want:    4,
		},
		{
			name:    "Example 2",
			k:       3,
			w:       0,
			profits: []int{1, 2, 3},
			capital: []int{0, 1, 2},
			want:    6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMaximizedCapital(tt.k, tt.w, tt.profits, tt.capital)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMaximizedCapital() = %v, want %v", got, tt.want)
			}
		})
	}
}
