package code

import (
	"testing"
)

func TestQueue(t *testing.T) {
	// 定义测试用例和期望的结果
	s := NewQueue()
	//为空
	if !s.IsEmpty() {
		t.Errorf("s should be empty")
	}
	s.Push(1)
	l := s.Len()

	if l != 1 {
		t.Errorf("Queue len should be 1")
	}

	v, ok := s.Pop()

	if !ok {
		t.Errorf("pop should be have data")
	}

	if v != 1 {
		t.Errorf("pop data=%d, but=%d", 1, v)
	}

	v, ok = s.Pop()

	if ok {
		t.Errorf("pop data should be empty, becuase Queue is empty, but=%d", v)
	}

	s.Push(2)

	v, ok = s.Peek()

	if !ok {
		t.Errorf("peek should be have data")
	}

	if v != 2 {
		t.Errorf("peek data=%d, but=%d", 1, v)
	}

	l = s.Len()

	if l != 1 {
		t.Errorf("Queue len should be 1, but=%d", l)
	}

	s.Pop()
	data := [3]int{3, 4, 5}
	for _, v := range data {
		s.Push(v)
	}

	for _, v := range data {
		res, _ := s.Pop()
		if res != v {
			t.Errorf("push data=%d, but=%d", v, res)
		}
	}

}
