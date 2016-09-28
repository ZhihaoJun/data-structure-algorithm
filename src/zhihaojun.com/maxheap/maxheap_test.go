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

func intsEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestSort(t *testing.T) {
	n := []int{4, 5, 1, 8, 10, 2, 0, 3}
	s := []int{10, 8, 5, 4, 3, 2, 1, 0}
	mh := NewMaxHeap(n)
	sorted := mh.Sorted()
	if intsEqual(sorted, s) == false {
		t.Error("max heap sort error")
	}
}
