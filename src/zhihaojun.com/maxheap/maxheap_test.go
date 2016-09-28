package maxheap

import (
	"testing"
)

func TestBuild(t *testing.T) {
	n := []int{4, 5, 1, 8, 10, 2, 0, 3}
	mh := NewMaxHeap(n)
	if mh.Extract() != 10 {
		t.Error("max heap error")
	}
	if mh.Extract() != 8 {
		t.Error("max heap error")
	}
}
