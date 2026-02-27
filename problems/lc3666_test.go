package problems

import (
	"reflect"
	"testing"
)

func TestMinOperations3666(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want int
	}{
		{
			name: "Example 1",
			s:    "110",
			k:    1,
			want: 1,
		},
		{
			name: "Example 2",
			s:    "0101",
			k:    3,
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minOperations3666(tt.s, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minOperations3666() = %v, want %v", got, tt.want)
			}
		})
	}
}
