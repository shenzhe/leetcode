package code

func Connect(root *TreeNode2) *TreeNode2 {
	if root == nil {
		return nil
	}
	queue := NewCommonQueue()
	queue.Push(root)
	for !queue.IsEmpty() {
		//弹出根节点
		l := queue.Len()
		preNode := &TreeNode2{}
		for i := 0; i < l; i++ {
			node := queue.Pop().(*TreeNode2)
			if preNode.Next == nil {
				preNode.Next = node
			} else {
				preNode.Next.Next = node
			}
			if node.Left != nil {
				queue.Push(node.Left)
			}
			if node.Right != nil {
				queue.Push(node.Right)
			}
		}
		preNode.Next = nil

	}
	return root
}
