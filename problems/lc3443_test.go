package problems

import (
	"reflect"
	"testing"
)

func TesmaxDistance(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want int
	}{
		{
			name: "Example 1",
			s:    "NSWWEW",
			k:    3,
			want: 6,
		},
		{
			name: "Example 2",
			s:    "NWSE",
			k:    1,
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxDistance(tt.s, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
