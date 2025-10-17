package problems

import (
	"reflect"
	"testing"
)

func TestMaxPartitionsAfterOperations(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want int
	}{
		{
			name: "Example 1",
			s:    "accca",
			k:    2,
			want: 3,
		},
		{
			name: "Example 2",
			s:    "aabaab",
			k:    3,
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxPartitionsAfterOperations(tt.s, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxPartitionsAfterOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
