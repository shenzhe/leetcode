package code

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	// 定义测试用例和期望的结果
	lru := NewLRUCache(3)
	cases := []struct {
		a, b, want int
	}{
		{1, 1, 1},
		{2, 1, 1},
		{2, 2, -1},
		{1, 2, 2},
		{1, 3, 3},
		{1, 4, 4},
		{2, 3, 3},
		{1, 3, 4},
		{2, 3, 4},
	}

	// 遍历测试用例，逐一进行测试
	for _, c := range cases {
		if c.a == 1 {
			lru.Put(c.b, c.want)
		} else {
			got := lru.Get(c.b)
			if got != c.want {
				t.Errorf("Get(%d) == %d, want %d", c.b, got, c.want)
			}
		}
	}
}
