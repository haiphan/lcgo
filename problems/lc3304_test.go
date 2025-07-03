package problems

import (
	"reflect"
	"testing"
)

func TestKthCharacter(t *testing.T) {
	tests := []struct {
		name string
		k    int
		want byte
	}{
		{
			name: "Example 1",
			k:    5,
			want: 'b',
		},
		{
			name: "Example 2",
			k:    10,
			want: 'c',
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := kthCharacter(tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kthCharacter() = %v, want %v", got, tt.want)
			}
		})
	}
}
