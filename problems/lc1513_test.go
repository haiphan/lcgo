package problems

import (
	"reflect"
	"testing"
)

func TestNumSub(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "Example 1",
			s:    "0110111",
			want: 9,
		},
		{
			name: "Example 2",
			s:    "101",
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numSub(tt.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numSub() = %v, want %v", got, tt.want)
			}
		})
	}
}
