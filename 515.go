package code

import "math"

func LargestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := make([]int, 0)
	queue := NewCommonQueue()
	queue.Push(root)
	// res = append(res, root.Val)

	for !queue.IsEmpty() {
		//弹出根节点
		max := math.MinInt32
		l := queue.Len()
		for i := 0; i < l; i++ {
			node := queue.Pop().(*TreeNode)
			if node.Val > max {
				max = node.Val
			}
			if node.Left != nil {
				queue.Push(node.Left)
			}
			if node.Right != nil {
				queue.Push(node.Right)
			}
		}
		res = append(res, max)
	}
	return res
}
