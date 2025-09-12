package problems

import (
	"reflect"
	"testing"
)

func TestDoesAliceWin(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "Example 1",
			s:    "bob",
			want: true,
		},
		{
			name: "Example 2",
			s:    "khgf",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := doesAliceWin(tt.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("doesAliceWin() = %v, want %v", got, tt.want)
			}
		})
	}
}
