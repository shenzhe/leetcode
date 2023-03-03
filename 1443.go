package code

func Execuate_1443(n int, edges [][]int, hasApple []bool) int {
	return minTime(n, edges, hasApple)
}

var graph_1443 [][]int
var visited_1443 map[int]bool
var path_1443 []int

func minTime(n int, edges [][]int, hasApple []bool) int {
	l := len(edges)
	if l == 0 {
		return 0
	}
	appNumber_1443 := 0
	for _, b := range hasApple {
		if b {
			appNumber_1443++
		}
	}

	if appNumber_1443 == 0 {
		return 0
	}
	if appNumber_1443 == 1 && hasApple[0] {
		return 0
	}

	//转化为一个无向图
	graph_1443 = make([][]int, n)
	// fmt.Println("len", l)
	for _, edge := range edges {
		graph_1443[edge[0]] = append(graph_1443[edge[0]], edge[1])
		graph_1443[edge[1]] = append(graph_1443[edge[1]], edge[0])
	}
	visited_1443 = make(map[int]bool)
	// fmt.Printf("graph=%v\n", graph_1443)
	//得到一个包含所有apple的路径
	path_1443 = make([]int, 0)
	traverse_1443(0, hasApple)
	return (len(path_1443) - 1) * 2
}

func traverse_1443(n int, hasApple []bool) {
	// fmt.Printf("n=%d, visited=%v\n", n, visited_1443)
	if _, ok := visited_1443[n]; ok {
		return
	}
	visited_1443[n] = true
	path_1443 = append(path_1443, n)

	for _, v := range graph_1443[n] {
		traverse_1443(v, hasApple)
	}
	// fmt.Println("check", n)
	if !hasApple[n] {
		pl := len(path_1443) - 1
		if path_1443[pl] == n {
			path_1443 = path_1443[:pl]
		}
	}
	// fmt.Printf("n=%d, path=%v\n", n, path_1443)
}
