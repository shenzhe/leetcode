package code

import "container/heap"

func Execuate_23(lists []*ListNode) *ListNode {
	return mergeKLists(lists)
}

func mergeKLists(lists []*ListNode) *ListNode {
	l := len(lists)
	if l == 0 {
		return nil
	}

	if l == 1 {
		return lists[0]
	}

	ans := &ListNode{}
	minHeap := &MinlHeap{}
	heap.Init(minHeap)
	for _, node := range lists {
		for node != nil {
			heap.Push(minHeap, node.Val)
			node = node.Next
		}
	}
	p := ans
	for minHeap.Len() > 0 {
		p.Next = &ListNode{
			Val: heap.Pop(minHeap).(int),
		}
		p = p.Next
	}

	return ans.Next
}
