package problems

import (
	"reflect"
	"testing"
)

func TestDistanceTraveled(t *testing.T) {
	tests := []struct {
		name           string
		mainTank       int
		additionalTank int
		want           int
	}{
		{
			name:           "Example 1",
			mainTank:       5,
			additionalTank: 1,
			want:           60,
		},
		{
			name:           "Example 2",
			mainTank:       1,
			additionalTank: 2,
			want:           10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := distanceTraveled(tt.mainTank, tt.additionalTank)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("distanceTraveled() = %v, want %v", got, tt.want)
			}
		})
	}
}
