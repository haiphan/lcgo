package problems

import (
	"reflect"
	"testing"
)

func TestMaxOperations(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "Example 1",
			s:    "1001101",
			want: 4,
		},
		{
			name: "Example 2",
			s:    "00111",
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxOperations(tt.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
