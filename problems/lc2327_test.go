package problems

import (
	"reflect"
	"testing"
)

func TestPeopleAwareOfSecret(t *testing.T) {
	tests := []struct {
		name   string
		n      int
		delay  int
		forget int
		want   int
	}{
		{
			name:   "Example 1",
			n:      6,
			delay:  2,
			forget: 4,
			want:   5,
		},
		{
			name:   "Example 2",
			n:      4,
			delay:  1,
			forget: 3,
			want:   6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := peopleAwareOfSecret(tt.n, tt.delay, tt.forget)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("peopleAwareOfSecret() = %v, want %v", got, tt.want)
			}
		})
	}
}
