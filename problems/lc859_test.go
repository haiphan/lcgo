package problems

import (
	"testing"
)

func TestFreqStack(t *testing.T) {
	fstack := Constructor()
	t.Run("test 1", func(t *testing.T) {
		fstack.Push(5)
		fstack.Push(7)
		fstack.Push(5)
		fstack.Push(7)
		fstack.Push(4)
		fstack.Push(5)
		if got := fstack.Pop(); got != 5 {
			t.Errorf("Pop() = %v, want %v", got, 5)
		}
		if got := fstack.Pop(); got != 7 {
			t.Errorf("Pop() = %v, want %v", got, 7)
		}
		if got := fstack.Pop(); got != 5 {
			t.Errorf("Pop() = %v, want %v", got, 5)
		}
		if got := fstack.Pop(); got != 4 {
			t.Errorf("Pop() = %v, want %v", got, 4)
		}
	})
}
