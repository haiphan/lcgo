package problems

import (
	"reflect"
	"testing"
)

func TestMinimumCost2976(t *testing.T) {
	tests := []struct {
		name              string
		source, target    string
		original, changed []byte
		cost              []int
		want              int64
	}{
		{
			name:     "Example 1",
			source:   "abc",
			target:   "bcd",
			original: []byte{'a', 'b', 'c'},
			changed:  []byte{'b', 'c', 'd'},
			cost:     []int{1, 1, 1},
			want:     3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minimumCost2976(tt.source, tt.target, tt.original, tt.changed, tt.cost)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minimumCost2976() = %v, want %v", got, tt.want)
			}
		})
	}
}
