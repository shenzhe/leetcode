package code

import (
	"testing"
)

func TestBlacklist(t *testing.T) {
	blacklist := []int{3, 6, 8}
	mapping := make(map[int]int)

	for _, v := range blacklist {
		mapping[v] = -1
	}
	fbl := NewFilterBlackList(10, blacklist)
	n := 10
	for n > 0 {
		v := fbl.Pick()
		if _, ok := mapping[v]; ok {
			t.Errorf("get blacklist value=%d", v)
		}
		n--
	}
}
