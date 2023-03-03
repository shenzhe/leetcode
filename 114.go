package code

func Execuate_114(root *TreeNode) {
	flatten(root)
}
func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)
	flatten(root.Right)

	//左边拉平，右边接到左边后面
	left := root.Left
	right := root.Right

	//左边为nil
	root.Left = nil
	//右边接左拉平树
	root.Right = left

	//遍历到新右树的最后一个节点
	p := root
	for p.Right != nil {
		p = p.Right
	}
	//原右边接到后面
	p.Right = right
}
