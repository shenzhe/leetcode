package code

var use_bsf bool

func Execuate_886(n int, dislikes [][]int, is_bsf bool) bool {
	use_bsf = is_bsf
	return possibleBipartition(n, dislikes)
}

func possibleBipartition(n int, dislikes [][]int) bool {
	graph := make([][]int, n)
	color := make([]bool, n)
	visited := make([]bool, n)

	for _, dislike := range dislikes {
		i, j := dislike[0]-1, dislike[1]-1
		graph[i] = append(graph[i], j)
		graph[j] = append(graph[j], i)
	}
	if use_bsf {
		for i := 0; i < n; i++ {
			if visited[i] {
				continue
			}
			if !bfs_886(graph, i, color, visited) {
				return false
			}
		}
	} else {
		for i := 0; i < n; i++ {
			if visited[i] {
				continue
			}
			if !dfs_886(graph, i, color, visited) {
				return false
			}
		}
	}
	return true
}

func dfs_886(graph [][]int, n int, color, visited []bool) bool {
	visited[n] = true
	for _, v := range graph[n] {
		//如果访问过，把
		if visited[v] {
			if color[v] == color[n] {
				return false
			}
		} else {
			color[v] = !color[n]
			if !dfs_886(graph, v, color, visited) {
				return false
			}
		}
	}
	return true
}

func bfs_886(graph [][]int, n int, color, visited []bool) bool {
	queue := make([]int, 0)
	queue = append(queue, n)

	for len(queue) > 0 {
		k := queue[0]
		queue = queue[1:]
		for _, v := range graph[k] {
			//如果访问过，把
			if visited[v] {
				if color[v] == color[k] {
					return false
				}
			} else {
				color[v] = !color[k]
				visited[v] = true
				queue = append(queue, v)
			}
		}
	}
	return true
}
