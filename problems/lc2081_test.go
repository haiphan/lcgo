package problems

import (
	"reflect"
	"testing"
)

func TestKMirror(t *testing.T) {
	tests := []struct {
		name string
		k    int
		n    int
		want int64
	}{
		{
			name: "Example 1",
			k:    2,
			n:    5,
			want: 25,
		},
		{
			name: "Example 2",
			k:    3,
			n:    7,
			want: 499,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := kMirror(tt.k, tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kMirror() = %v, want %v", got, tt.want)
			}
		})
	}
}
