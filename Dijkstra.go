package code

import (
	"container/heap"
	"fmt"
	"math"
)

type Dijkstra struct {
	graph      [][][]int
	len        int
	disTo      []int
	pq         *PriorityQueue
	getWeight  func(any, any) any
	checkFunc  func(any, any) bool
	flatFlag   bool
	disToFloat []float64
	graphFloat [][][]float64
}
type status struct {
	s  int
	w  int
	fw float64
}

func NewDijkstra(graph [][][]int, n int) *Dijkstra {
	dk := &Dijkstra{
		graph:    graph,
		flatFlag: false,
		getWeight: func(i1, i2 any) any {
			return i1.(int) + i2.(int)
		},
		checkFunc: func(a1, a2 any) bool {
			return a1.(int) > a2.(int)
		},
	}
	dk.len = n
	dk.disTo = make([]int, n)
	for i := 0; i < n; i++ {
		dk.disTo[i] = math.MaxInt
	}
	dk.pq = NewPriorityQueue(func(a1, a2 any) bool {
		return a1.(status).w < a2.(status).w
	})
	heap.Init(dk.pq)
	return dk
}

func NewDijkstraFloat(graph [][][]float64, n int) *Dijkstra {
	dk := &Dijkstra{
		graphFloat: graph,
		flatFlag:   true,
		getWeight: func(i1, i2 any) any {
			return i1.(float64) + i2.(float64)
		},
		checkFunc: func(a1, a2 any) bool {
			return a1.(float64) < a2.(float64)
		},
	}
	dk.len = n
	dk.disToFloat = make([]float64, n)
	for i := 0; i < n; i++ {
		dk.disToFloat[i] = -1
	}
	dk.pq = NewPriorityQueue(func(a1, a2 any) bool {
		return a1.(status).fw > a2.(status).fw
	})
	heap.Init(dk.pq)
	return dk
}

func (dk *Dijkstra) SetFloat() {
	dk.flatFlag = true
}

func (dk *Dijkstra) SetWeightFunc(f func(i1, i2 any) any) {
	dk.getWeight = f
}

func (dk *Dijkstra) SetCheckFunc(f func(i1, i2 any) bool) {
	dk.checkFunc = f
}

func (dk *Dijkstra) Run(start, end int) []int {
	//起点到起点，权重为0
	dk.disTo[start] = 0
	heap.Push(dk.pq, status{
		s: start,
		w: 0,
	})
	for dk.pq.Len() > 0 {
		cur := heap.Pop(dk.pq).(status)
		if cur.s == end {
			return []int{cur.w}
		}
		//当前节点的距离大于目前节点，不是最短，跳过
		if dk.checkFunc(cur.w, dk.disTo[cur.s]) {
			continue
		}
		for _, n := range dk.graph[cur.s] {
			//到下一个节点的距离=当前距离+边长
			nw := dk.getWeight(cur.w, n[1]).(int)
			//小于，则更新
			if nw < dk.disTo[n[0]] {
				dk.disTo[n[0]] = nw
				//放入队列
				heap.Push(dk.pq, status{
					s: n[0],
					w: nw,
				})
			}
		}
	}
	return dk.disTo
}

func (dk *Dijkstra) RunFloat(start, end int) []float64 {
	fmt.Printf("g=%v\n", dk.graphFloat)
	//起点到起点，权重为0
	dk.disToFloat[start] = 1
	heap.Push(dk.pq, status{
		s:  start,
		fw: 1,
	})
	for dk.pq.Len() > 0 {
		cur := heap.Pop(dk.pq).(status)
		fmt.Printf("cur=%v, disto=%v \n", cur, dk.disToFloat)
		if cur.s == end {
			return dk.disToFloat
		}
		//当前节点的距离大于目前节点，不是最短，跳过
		if dk.checkFunc(cur.fw, dk.disToFloat[cur.s]) {
			continue
		}
		for _, n := range dk.graphFloat[cur.s] {
			//到下一个节点的距离=当前距离+边长
			nw := dk.getWeight(dk.disToFloat[cur.s], n[1]).(float64)
			fmt.Printf("cur=%v, n=%v, nw=%f, disto=%v\n", cur, n, nw, dk.disToFloat)
			//大于，则更新
			// if nw > dk.disToFloat[int(n[0])] {
			if dk.checkFunc(dk.disToFloat[int(n[0])], nw) {
				dk.disToFloat[int(n[0])] = nw
				//放入队列
				heap.Push(dk.pq, status{
					s:  int(n[0]),
					fw: nw,
				})
			}
		}
	}
	return dk.disToFloat
}
