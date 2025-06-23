package problems

import (
	"reflect"
	"testing"
)

func TestBulbSwitch(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "Example 1",
			n:    3,
			want: 1,
		},
		{
			name: "Example 2",
			n:    1,
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := bulbSwitch(tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bulbSwitch() = %v, want %v", got, tt.want)
			}
		})
	}
}
