package code

import "fmt"

func Execuate_210(numCourses int, prerequisites [][]int) []int {
	return findOrder(numCourses, prerequisites)
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make([][]int, numCourses)

	//出度表
	outdegree := make([]int, numCourses)

	for _, v := range prerequisites {
		graph[v[1]] = append(graph[v[1]], v[0])
		fmt.Printf("out=%d, v=%v\n", v[0], v)
		outdegree[v[0]] = outdegree[v[0]] + 1
	}
	fmt.Printf("outdegree=%v\n", outdegree)

	//找出出度表为0的节点，即为初始修的课程
	queue := make([]int, 0)
	for k, v := range outdegree {
		if v == 0 {
			queue = append(queue, k)
		}
	}

	//遍历q
	order := make([]int, 0)
	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]
		order = append(order, course)
		//从graph找出他的入节点
		for _, v := range graph[course] {
			//出度-1
			outdegree[v]--
			if outdegree[v] == 0 {
				//加入到队列·
				queue = append(queue, v)
			}
		}
	}

	fmt.Printf("order=%v, outdegree=%v, graph=%v\n", order, outdegree, graph)

	if len(order) < numCourses {
		return []int{}
	}
	return order
}
