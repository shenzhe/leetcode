package code

type commonStack struct {
	list []interface{}
}

func NewCommonStack() commonStack {
	s := commonStack{
		list: make([]interface{}, 0),
	}
	return s
}

func (s *commonStack) Push(item interface{}) {
	// fmt.Println("push", item)
	s.list = append(s.list, item)
}

func (s *commonStack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	l := len(s.list)
	res := s.list[l-1]
	s.list = s.list[0 : l-1]
	return res
}

func (s *commonStack) Peek() interface{} {
	if s.IsEmpty() {
		return nil
	}
	l := len(s.list)
	res := s.list[l-1]
	return res
}

func (s *commonStack) Len() int {
	return len(s.list)
}

func (s *commonStack) IsEmpty() bool {
	return len(s.list) == 0
}
