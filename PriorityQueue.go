package code

type PriorityQueue struct {
	queue    []any
	lessFunc func(any, any) bool
}

func NewPriorityQueue(lessFunc func(any, any) bool) *PriorityQueue {
	pq := &PriorityQueue{}
	pq.queue = make([]any, 0)
	pq.lessFunc = lessFunc
	return pq
}

func (pq PriorityQueue) Len() int {
	return len(pq.queue)
}

func (pq *PriorityQueue) Less(i, j int) bool {
	return pq.lessFunc(pq.queue[i], pq.queue[j])
}

func (pq *PriorityQueue) Swap(i, j int) {
	pq.queue[i], pq.queue[j] = pq.queue[j], pq.queue[i]
}

func (pq *PriorityQueue) Push(x any) {
	pq.queue = append(pq.queue, x)
}

func (pq *PriorityQueue) Pop() any {
	l := len(pq.queue)
	if l < 1 {
		return nil
	}
	x := pq.queue[l-1]
	pq.queue = pq.queue[:l-1]
	return x
}

func (pq *PriorityQueue) Peek() any {
	l := len(pq.queue)
	if l < 1 {
		return nil
	}
	return pq.queue[l-1]
}
