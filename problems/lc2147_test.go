package problems

import (
	"reflect"
	"testing"
)

func TestNumberOfWays2147(t *testing.T) {
	tests := []struct {
		name     string
		corridor string
		want     int
	}{
		{
			name:     "Example 1",
			corridor: "SSPPSPS",
			want:     3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numberOfWays2147(tt.corridor)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numberOfWays2147() = %v, want %v", got, tt.want)
			}
		})
	}
}
