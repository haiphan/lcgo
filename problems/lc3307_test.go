package problems

import (
	"reflect"
	"testing"
)

func TestKthCharacter2(t *testing.T) {
	tests := []struct {
		name       string
		k          int64
		operations []int
		want       byte
	}{
		{
			name:       "Example 1",
			k:          12145134613,
			operations: []int{0, 0, 0, 0, 1, 0, 0, 0, 1, 1, 1, 1, 1, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 1, 0, 1, 0, 0, 1, 1, 1, 1, 1},
			want:       'i',
		},
		{
			name:       "Example 2",
			k:          5,
			operations: []int{0, 0, 0},
			want:       'a',
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := kthCharacter2(tt.k, tt.operations)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kthCharacter2() = %v, want %v", got, tt.want)
			}
		})
	}
}
