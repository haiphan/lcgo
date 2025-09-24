package problems

import (
	"reflect"
	"testing"
)

func TestMaxDifferenceII(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want int
	}{
		{
			name: "Example 1",
			s:    "12233",
			k:    4,
			want: -1,
		},
		{
			name: "Example 2",
			s:    "1122211",
			k:    3,
			want: 1,
		},
		{
			name: "Example 3",
			s:    "110",
			k:    3,
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxDifferenceII(tt.s, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxDifferenceII() = %v, want %v", got, tt.want)
			}
		})
	}
}
