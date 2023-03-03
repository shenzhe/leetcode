package code

type myStack struct {
	list []int
}

func NewStack() myStack {
	s := myStack{
		list: make([]int, 0),
	}
	return s
}

func (s *myStack) Push(item int) {
	// fmt.Println("push", item)
	s.list = append(s.list, item)
}

func (s *myStack) Pop() (int, bool) {
	if s.IsEmpty() {
		return -1, false
	}
	l := len(s.list)
	res := s.list[l-1]
	s.list = s.list[0 : l-1]
	return res, true
}

func (s *myStack) Peek() (int, bool) {
	if s.IsEmpty() {
		return -1, false
	}
	l := len(s.list)
	res := s.list[l-1]
	return res, true
}

func (s *myStack) Len() int {
	return len(s.list)
}

func (s *myStack) IsEmpty() bool {
	return len(s.list) == 0
}

func (s *myStack) Sum() int {
	num := 0
	for _, v := range s.list {
		num += v
	}
	return num
}
