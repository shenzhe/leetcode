package code

import (
	"fmt"
	"testing"
)

func TestUnionFind(t *testing.T) {
	n := 9
	uf := NewUnionFind(n)

	if uf.Count() != n {
		t.Errorf("want=%d, but=%d\n", n, uf.Count())
	}
	uf.Union(1, 2)
	n--
	if uf.Count() != n {
		t.Errorf("want=%d, but=%d\n", n, uf.Count())
	}
	if !uf.Connected(1, 2) {
		t.Errorf("want=%d and %d is connected,but not\n", 1, 2)
	}
	fmt.Printf("parent=%v\n", uf.GetParent())
	uf.Union(2, 3)
	n--
	if uf.Count() != n {
		t.Errorf("want=%d, but=%d\n", n, uf.Count())
	}
	if !uf.Connected(1, 3) {
		t.Errorf("want=%d and %d is connected,but not\n", 1, 3)
	}

	fmt.Printf("parent=%v\n", uf.GetParent())

}
