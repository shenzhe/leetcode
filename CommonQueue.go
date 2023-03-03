package code

type commonQueue struct {
	list []interface{}
}

func NewCommonQueue() commonQueue {
	s := commonQueue{
		list: make([]interface{}, 0),
	}
	return s
}

func (s *commonQueue) Push(item interface{}) {
	s.list = append(s.list, item)
}

func (s *commonQueue) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	res := s.list[0]
	s.list = s.list[1:]
	return res
}

func (s *commonQueue) Peek() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.list[0]
}

func (s *commonQueue) Len() int {
	return len(s.list)
}

func (s *commonQueue) IsEmpty() bool {
	return len(s.list) == 0
}
