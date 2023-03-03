package code

import "fmt"

func Execuate_2049(parents []int) int {
	return countHighestScoreNodes(parents)
}

var graph_2049 [][]int
var visited_2049 map[int]int

func countHighestScoreNodes(parents []int) int {
	fmt.Printf("parents=%v\n", parents)
	//build graph
	l := len(parents)
	if l == 2 {
		return 2
	}
	graph_2049 = make([][]int, l)
	visited_2049 = make(map[int]int)

	for k, v := range parents {
		if v >= 0 {
			graph_2049[v] = append(graph_2049[v], k)
		}
	}
	traverse_2049(0)
	fmt.Printf("graph=%v\n", graph_2049)
	fmt.Printf("visited=%v\n", visited_2049)
	//计算分数
	max := 0
	res := 0
	for k, v := range parents {
		//求分数
		sum := 1
		if v == -1 {
			//顶点，=左右两边
			for _, sv := range graph_2049[k] {
				sum *= visited_2049[sv]
			}
		} else {
			//最未结点，直接根结点-1
			sum = visited_2049[0] - visited_2049[k]
			for _, sv := range graph_2049[k] {
				sum *= visited_2049[sv]
			}
		}
		fmt.Printf("n=%d, sum=%d, max=%d\n", k, sum, max)
		if sum > max {
			res = 1
			max = sum
		} else if sum == max {
			res++
		}
	}
	fmt.Println("res:", res)
	return res
}

func traverse_2049(n int) int {
	// fmt.Println("ns=", n)
	// if _, ok := visited_2049[n]; ok {
	// 	return visited_2049[n]
	// }
	visited_2049[n] = 1
	for _, v := range graph_2049[n] {
		visited_2049[n] += traverse_2049(v)
	}
	fmt.Println("n=", n, visited_2049[n])
	return visited_2049[n]

}
