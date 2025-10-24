package problems

import (
	"reflect"
	"testing"
)

func TestNextBeautifulNumber(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "Example 1",
			n:    1,
			want: 22,
		},
		{
			name: "Example 2",
			n:    1000,
			want: 1333,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := nextBeautifulNumber(tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextBeautifulNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
