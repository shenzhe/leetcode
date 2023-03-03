package code

type MinlHeap []int

func (sh MinlHeap) Len() int {
	return len(sh)
}

func (sh MinlHeap) Less(i, j int) bool {
	return sh[i] < sh[j]
}

func (sh MinlHeap) Swap(i, j int) {
	sh[i], sh[j] = sh[j], sh[i]
}

func (sh *MinlHeap) Push(i interface{}) {
	*sh = append(*sh, i.(int))
}

func (sh *MinlHeap) Pop() interface{} {
	old := *sh
	n := len(old)
	num := old[n-1]
	*sh = old[0 : n-1]
	return num
}

func (sh MinlHeap) Peek() interface{} {
	if len(sh) > 0 {
		return sh[0]
	}
	return 0
}

type MaxHeap []int

func (bh MaxHeap) Len() int {
	return len(bh)
}

func (bh MaxHeap) Less(i, j int) bool {
	return bh[i] > bh[j]
}

func (bh MaxHeap) Swap(i, j int) {
	bh[i], bh[j] = bh[j], bh[i]
}

func (bh *MaxHeap) Push(i interface{}) {
	*bh = append(*bh, i.(int))
}

func (bh *MaxHeap) Pop() interface{} {
	old := *bh
	n := len(old)
	num := old[n-1]
	*bh = old[0 : n-1]
	return num
}

func (bh MaxHeap) Peek() interface{} {
	if len(bh) > 0 {
		return bh[0]
	}
	return 0
}
