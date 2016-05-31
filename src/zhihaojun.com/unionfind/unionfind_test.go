package unionfind

import (
	"testing"
)

func TestNew(t *testing.T) {
	n := 10
	uf := New(n)
	for i := 0; i < n; i++ {
		if i != uf.Find(i) {
			t.Error("error in new sets")
		}
	}
}

func Test() {

}
