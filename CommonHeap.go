package code

import "container/heap"

type CommonMinHeap []int
type CommonMaxHeap []int
type CommonHeap struct {
	minHeap *CommonMinHeap
	maxHeap *CommonMaxHeap
}

func NewCommonHeap(minHeap *CommonMinHeap, maxHeap *CommonMaxHeap) *CommonHeap {
	commonHeap := &CommonHeap{
		minHeap: minHeap,
		maxHeap: maxHeap,
	}
	if commonHeap.minHeap != nil {
		heap.Init(commonHeap.minHeap)
	}
	if commonHeap.maxHeap != nil {
		heap.Init(commonHeap.maxHeap)
	}
	return commonHeap
}

func (ch CommonHeap) GetMinHeap() *CommonMinHeap {
	return ch.minHeap
}

func (ch CommonHeap) GetMinLen() int {
	return ch.minHeap.Len()
}

func (ch *CommonHeap) MinPush(i interface{}) {
	heap.Push(ch.minHeap, i)
}

func (ch *CommonHeap) MinPop() interface{} {
	return heap.Pop(ch.minHeap)
}

func (ch *CommonHeap) MinPeek() interface{} {
	return ch.minHeap.Peek()
}

func (ch CommonHeap) GetMaxHeap() *CommonMaxHeap {
	return ch.maxHeap
}

func (ch CommonHeap) GetMaxLen() int {
	return ch.maxHeap.Len()
}

func (ch *CommonHeap) MaxPush(i interface{}) {
	heap.Push(ch.maxHeap, i)
}

func (ch *CommonHeap) MaxPop() interface{} {
	return heap.Pop(ch.maxHeap)
}

func (ch *CommonHeap) MaxPeek() interface{} {
	return ch.maxHeap.Peek()
}

func (minh CommonMinHeap) Len() int {
	return len(minh)
}

func (minh CommonMinHeap) Less(i, j int) bool {
	return minh[i] < minh[j]
}

func (minh CommonMinHeap) Swap(i, j int) {
	minh[i], minh[j] = minh[j], minh[i]
}

func (minh *CommonMinHeap) Push(i interface{}) {
	*minh = append(*minh, i.(int))
}

func (minh *CommonMinHeap) Pop() interface{} {
	old := *minh
	n := len(old)
	num := old[n-1]
	*minh = old[0 : n-1]
	return num
}

func (minh CommonMinHeap) Peek() interface{} {
	if len(minh) > 0 {
		return minh[0]
	}
	return 0
}

func (maxh CommonMaxHeap) Len() int {
	return len(maxh)
}

func (maxh CommonMaxHeap) Less(i, j int) bool {
	return maxh[i] > maxh[j]
}

func (maxh CommonMaxHeap) Swap(i, j int) {
	maxh[i], maxh[j] = maxh[j], maxh[i]
}

func (maxh *CommonMaxHeap) Push(i interface{}) {
	*maxh = append(*maxh, i.(int))
}

func (maxh *CommonMaxHeap) Pop() interface{} {
	old := *maxh
	n := len(old)
	num := old[n-1]
	*maxh = old[0 : n-1]
	return num
}

func (maxh CommonMaxHeap) Peek() interface{} {
	if len(maxh) > 0 {
		return maxh[0]
	}
	return 0
}
