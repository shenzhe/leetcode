package code

import "fmt"

func Execuate_785(graph [][]int) bool {
	return isBipartite(graph)
}

var visited_785 []bool
var color_785 []bool

func isBipartite(graph [][]int) bool {
	l := len(graph)
	visited_785 = make([]bool, l)
	color_785 = make([]bool, l)
	fmt.Printf("v=%v, c=%v\n", visited_785, color_785)
	for i := 0; i < l; i++ {
		if !visited_785[i] {
			if !traverse_785(graph, i) {
				fmt.Printf("v=%v, c=%v\n", visited_785, color_785)
				return false
			}
		}
	}
	// fmt.Printf("v=%v, c=%v\n", visited_785, color_785)
	return true
}

func traverse_785(graph [][]int, n int) bool {
	// if visited_785[n] {
	// 	return true
	// }
	visited_785[n] = true
	// fmt.Printf("n=%d\n", n)
	for _, v := range graph[n] {
		//如果已存在，则判断是否颜色不同
		if visited_785[v] {
			if color_785[v] == color_785[n] {
				fmt.Printf("false, %d, %d\n", v, n)
				return false
			}
		} else {
			//否则染上不同的颜色
			color_785[v] = !color_785[n]
			if !traverse_785(graph, v) {
				return false
			}
		}
	}

	return true
}
