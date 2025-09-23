package problems

import (
	"reflect"
	"testing"
)

func TestCompareVersion(t *testing.T) {
	tests := []struct {
		name     string
		version1 string
		version2 string
		want     int
	}{
		{
			name:     "Example 1",
			version1: "1.01",
			version2: "1.001",
			want:     0,
		},
		{
			name:     "Example 2",
			version1: "1.2",
			version2: "1.010",
			want:     -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := compareVersion(tt.version1, tt.version2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("compareVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
