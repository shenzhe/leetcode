package code

type myQueue struct {
	list []int
}

func NewQueue() myQueue {
	s := myQueue{
		list: make([]int, 0),
	}
	return s
}

func (s *myQueue) Push(item int) {
	s.list = append(s.list, item)
}

func (s *myQueue) Pop() (int, bool) {
	if s.IsEmpty() {
		return -1, false
	}
	res := s.list[0]
	s.list = s.list[1:]
	return res, true
}

func (s *myQueue) Peek() (int, bool) {
	if s.IsEmpty() {
		return -1, false
	}
	return s.list[0], true
}

func (s *myQueue) Len() int {
	return len(s.list)
}

func (s *myQueue) IsEmpty() bool {
	return len(s.list) == 0
}
