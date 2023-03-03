package code

func Execuate_1361(n int, leftChild []int, rightChild []int) bool {
	return validateBinaryTreeNodes(n, leftChild, rightChild)
}

func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	indegree := make([]int, n)
	for i := 0; i < n; i++ {
		if leftChild[i] != -1 {
			indegree[leftChild[i]]++ //入度++
		}
		if rightChild[i] != -1 {
			indegree[rightChild[i]]++ //入度++
		}
	}

	// fmt.Printf("indegree=%v\n", indegree)

	//统计入度为0的节点
	roots := make([]int, 0)
	for i, v := range indegree {
		if v == 0 {
			roots = append(roots, i)
		}
	}
	// fmt.Printf("roots=%v\n", roots)
	//如果root的个数为0 ，说明成环，不是btree
	if len(roots) == 0 {
		return false
	}

	//如果root的个数超过1 ，不只一个btree
	if len(roots) > 1 {
		return false
	}

	//BFS遍历，看是否为完成的btree,条件是节点数=n
	queue := make([]int, 0)
	queue = append(queue, roots[0])

	count := 0
	for len(queue) > 0 {
		// fmt.Printf("q=%v\n", queue)
		root := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		// fmt.Printf("r=%d, q=%v\n", root, queue)
		count++
		if count > n {
			break
		}
		if leftChild[root] != -1 {
			queue = append(queue, leftChild[root])
		}
		if rightChild[root] != -1 {
			queue = append(queue, rightChild[root])
		}
		// fmt.Printf("r2=%d, q=%v\n", root, queue)
	}
	//不成环,不分裂，count==n
	return count == n
}
