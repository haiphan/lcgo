package problems

import (
	"reflect"
	"testing"
)

func TestSimplifyPath(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
	}{
		{
			name: "Example 1",
			path: "/home/",
			want: "/home",
		},
		{
			name: "Example 2",
			path: "/home//foo/",
			want: "/home/foo",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := simplifyPath(tt.path)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("simplifyPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
