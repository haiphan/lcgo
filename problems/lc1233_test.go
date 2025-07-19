package problems

import (
	"reflect"
	"testing"
)

func TestRemoveSubfolders(t *testing.T) {
	tests := []struct {
		name   string
		folder []string
		want   []string
	}{
		{
			name:   "Example 1",
			folder: []string{"/a/b/c", "/a/b/ca", "/a/b/d"},
			want:   []string{"/a/b/c", "/a/b/ca", "/a/b/d"},
		},
		{
			name:   "Example 2",
			folder: []string{"/a", "/a/b/c", "/a/b/d"},
			want:   []string{"/a"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeSubfolders(tt.folder)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeSubfolders() = %v, want %v", got, tt.want)
			}
		})
	}
}
