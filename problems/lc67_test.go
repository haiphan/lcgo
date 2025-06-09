package problems

import (
	"reflect"
	"testing"
)

func TestAddBinary(t *testing.T) {
	tests := []struct {
		name string
		a    string
		b    string
		want string
	}{
		{
			name: "Example 1",
			a:    "11",
			b:    "1",
			want: "100",
		},
		{
			name: "Example 2",
			a:    "1010",
			b:    "1011",
			want: "10101",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addBinary(tt.a, tt.b)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
