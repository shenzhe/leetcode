package code

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	type person struct {
		name string
		age  int
	}
	lessFunc := func(p1, p2 any) bool {
		return p1.(person).age < p2.(person).age
	}
	pq := NewPriorityQueue(lessFunc)
	heap.Init(pq)
	persons := []person{
		{
			"bob1",
			13,
		},
		{
			"bob2",
			10,
		},
		{
			"bob3",
			16,
		},
		{
			"bob4",
			8,
		},
	}
	for _, p := range persons {
		heap.Push(pq, p)
	}
	fmt.Printf("l=%d\n", pq.Len())
	for pq.Len() > 0 {
		fmt.Printf("p=%v\n", heap.Pop(pq).(person))
	}

}
