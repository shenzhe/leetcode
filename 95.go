package code

import "fmt"

func Execuate_95(n int) []*TreeNode {
	return generateTrees(n)
}

var cache_95 map[string][]*TreeNode

func generateTrees(n int) []*TreeNode {

	if n == 0 {
		return []*TreeNode{}
	}

	if n == 1 {
		return []*TreeNode{
			{
				Val: 1,
			},
		}
	}
	cache_95 = make(map[string][]*TreeNode)
	return buildTree_95(1, n)
}

func buildTree_95(lo, hi int) []*TreeNode {
	res := make([]*TreeNode, 0)
	if lo > hi {
		res = append(res, nil)
		return res
	}
	key := fmt.Sprintf("%d_%d", lo, hi)
	if v, ok := cache_95[key]; ok {
		return v
	}
	for i := lo; i <= hi; i++ {
		leftTrees := buildTree_95(lo, i-1)
		rightTrees := buildTree_95(i+1, hi)

		for _, leftTree := range leftTrees {
			for _, rightTree := range rightTrees {
				root := &TreeNode{
					Val: i,
				}
				root.Left = leftTree
				root.Right = rightTree
				res = append(res, root)
			}
		}
	}

	cache_95[key] = res
	return res
}

/**
loop 1 3
loop 1 0
loop 2 3
loop 2 1
loop 3 3
loop 3 2
loop 4 3
loop 2 2
loop 2 1
loop 3 2
loop 4 3
loop 1 1
loop 1 0
loop 2 1
loop 3 3
loop 3 2
loop 4 3
loop 1 2
loop 1 0
loop 2 2
loop 2 1
loop 3 2
loop 1 1
loop 1 0
loop 2 1
loop 3 2
loop 4 3

**/
