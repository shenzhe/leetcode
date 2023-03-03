package code

import (
	"container/heap"
	"fmt"
)

type MedianFinder struct {
	sh *MinlHeap
	bh *MaxHeap
}

func NewFindMedian() MedianFinder {
	sh := &MinlHeap{}
	bh := &MaxHeap{}
	mf := MedianFinder{
		sh: sh,
		bh: bh,
	}
	heap.Init(mf.sh)
	heap.Init(mf.bh)

	return mf

}

func (mf *MedianFinder) AddNum(num int) {
	//如果小堆的数量大于大堆的数量
	if mf.sh.Len() >= mf.bh.Len() {
		//最小堆里插入
		heap.Push(mf.sh, num)
		//把最小堆的最大值插入到大堆
		heap.Push(mf.bh, heap.Pop(mf.sh).(int))
	} else {
		heap.Push(mf.bh, num)
		heap.Push(mf.sh, heap.Pop(mf.bh).(int))
	}
}

func (mf *MedianFinder) PopAll() {
	for mf.sh.Len() > 0 {
		fmt.Println("sh", heap.Pop(mf.sh))
	}

	for mf.bh.Len() > 0 {
		fmt.Println("bh", heap.Pop(mf.bh))
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	shLen := mf.sh.Len()
	bhLen := mf.bh.Len()
	fmt.Println("lens", shLen, bhLen, mf.sh, mf.bh)
	if shLen > bhLen {
		return float64(mf.sh.Peek().(int))
	}

	if shLen < bhLen {
		return float64(mf.bh.Peek().(int))
	}

	return (float64(mf.sh.Peek().(int)) + float64(mf.bh.Peek().(int))) / 2.0
}
