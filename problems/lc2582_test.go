package problems

import (
	"reflect"
	"testing"
)

func TestPassThePillow(t *testing.T) {
	tests := []struct {
		name string
		n    int
		time int
		want int
	}{
		{
			name: "Example 1",
			n:    4,
			time: 5,
			want: 2,
		},
		{
			name: "Example 2",
			n:    3,
			time: 2,
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := passThePillow(tt.n, tt.time)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("passThePillow() = %v, want %v", got, tt.want)
			}
		})
	}
}
