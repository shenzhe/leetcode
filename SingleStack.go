package code

type mySingleStack struct {
	list []int
}

func NewSingleStack() myStack {
	s := myStack{
		list: make([]int, 0),
	}
	return s
}

func (s *mySingleStack) Push(item int) {
	for !s.IsEmpty() {
		v, ok := s.Peek()
		if ok && v > item {
			break
		}
		s.Pop()
	}
	s.list = append(s.list, item)
}

func (s *mySingleStack) Pop() (int, bool) {
	if s.IsEmpty() {
		return -1, false
	}
	l := len(s.list)
	res := s.list[l-1]
	s.list = s.list[0 : l-1]
	return res, true
}

func (s *mySingleStack) Peek() (int, bool) {
	if s.IsEmpty() {
		return -1, false
	}
	l := len(s.list)
	res := s.list[l-1]
	return res, true
}

func (s *mySingleStack) Len() int {
	return len(s.list)
}

func (s *mySingleStack) IsEmpty() bool {
	return len(s.list) == 0
}
