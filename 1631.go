package code

func Execuate_1631(heights [][]int) int {
	return minimumEffortPath(heights)
}

func minimumEffortPath(heights [][]int) int {
	m, n := len(heights), len(heights[0])
	graph := make([][][]int, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			//上
			v := TwodToOned(i, j, n)
			if j > 0 {
				graph[v] = append(graph[v], []int{TwodToOned(i, j-1, n), getEffortNum(heights, i, j, i, j-1)})
			}
			//下
			if j < n-1 {
				graph[v] = append(graph[v], []int{TwodToOned(i, j+1, n), getEffortNum(heights, i, j, i, j+1)})
			}
			//左
			if i > 0 {
				graph[v] = append(graph[v], []int{TwodToOned(i-1, j, n), getEffortNum(heights, i, j, i-1, j)})
			}
			//右
			if i < m-1 {
				graph[v] = append(graph[v], []int{TwodToOned(i+1, j, n), getEffortNum(heights, i, j, i+1, j)})
			}
		}
	}
	dk := NewDijkstra(graph, m*n)
	dk.getWeight = func(a1, a2 any) any {
		w1, w2 := a1.(int), a2.(int)
		if w1 > w2 {
			return w1
		}
		return w2
	}
	ans := dk.Run(0, TwodToOned(m-1, n-1, n))
	return ans[0]

}

func TwodToOned(x, y, n int) int {
	return x*n + y
}

func getEffortNum(heights [][]int, x1, y1, x2, y2 int) int {
	return abs(heights[x1][y1] - heights[x2][y2])
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
