package code

import "math"

func Execuate_743(times [][]int, n int, k int) int {
	return networkDelayTime(times, n, k)
}

func networkDelayTime(times [][]int, n int, k int) int {
	graph := make([][][]int, n)
	for _, t := range times {
		graph[t[0]-1] = append(graph[t[0]-1], []int{t[1] - 1, t[2]})
	}

	dk := NewDijkstra(graph, n)
	ans := dk.Run(k-1, -1)

	max := 0
	for _, v := range ans {
		if v == math.MaxInt {
			return -1
		}
		if v > max {
			max = v
		}
	}
	return max
}
