package code

import (
	"container/heap"
	"sort"
)

func Execuate_1584(points [][]int) int {
	return minCostConnectPoints(points)
}

func minCostConnectPoints(points [][]int) int {
	//获得所有边
	sides := make([]int, 0)
	sideMap := make(map[int][][2]int)
	l := len(points)
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			d := getDistance(points[i], points[j])
			sides = append(sides, d)
			sideMap[d] = append(sideMap[d], [2]int{i, j})
		}
	}
	// fmt.Printf("sides=%v, sideMap=%v\n", sides, sideMap)

	// sides = MergeSort(sides)
	sort.Ints(sides)
	// fmt.Printf("sides=%v, sideMap=%v\n", sides, sideMap)
	minCost := 0
	uf := NewUnionFind(l)
	for _, cost := range sides {
		p1, p2 := sideMap[cost][0][0], sideMap[cost][0][1]
		// fmt.Printf("p1=%d, p2=%d, cost=%d\n", p1, p2, cost)
		sideMap[cost] = sideMap[cost][1:]
		if uf.Connected(p1, p2) {
			continue
		}
		minCost += cost
		// fmt.Printf("c:p1=%d, p2=%d, cost=%d, minCost=%d\n", p1, p2, cost, minCost)
		uf.Union(p1, p2)

	}
	return minCost
}

func getDistance(p1, p2 []int) int {
	//|xi - xj| + |yi - yj|
	xd := p1[0] - p2[0]
	if xd < 0 {
		xd *= -1
	}
	yd := p1[1] - p2[1]
	if yd < 0 {
		yd *= -1
	}
	return xd + yd
}

func Execuate_1584_1(points [][]int) int {
	return prim(points)
}

func prim(points [][]int) int {
	l := len(points)
	graph := make([][][2]int, l)
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			d := getDistance(points[i], points[j])
			graph[i] = append(graph[i], [2]int{j, d})
			graph[j] = append(graph[j], [2]int{i, d})
		}
	}
	// fmt.Printf("g=%v\n", graph)
	// queue := make([][2]int, 0)
	inMst := make([]bool, l)
	// queue = append(queue, graph[0]...)
	// sort.Slice(queue, func(i, j int) bool {
	// 	return queue[i][1] < queue[j][1]
	// })
	queue := NewPriorityQueue(func(a1, a2 any) bool {
		return a1.([2]int)[1] < a2.([2]int)[1]
	})
	for _, v := range graph[0] {
		heap.Push(queue, v)
	}
	// fmt.Printf("q=%v\n", queue)
	inMst[0] = true
	minCost := 0
	for queue.Len() > 0 {
		v := heap.Pop(queue).([2]int)
		if inMst[v[0]] {
			continue
		}
		inMst[v[0]] = true
		minCost += v[1]
		// fmt.Printf("v=%v, minCost=%d\n", v, minCost)
		for _, v2 := range graph[v[0]] {
			if inMst[v2[0]] {
				continue
			}
			heap.Push(queue, v2)
			// queue = append(queue, v2)

		}
		// sort.Slice(queue, func(i, j int) bool {
		// 	return queue[i][1] < queue[j][1]
		// })
	}

	//确保包括所有的结点
	for _, v := range inMst {
		if !v {
			return -1
		}
	}

	return minCost
}
