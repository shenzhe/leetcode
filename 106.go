package code

var inmap106 = make(map[int]int)

func Execuate106(inorder []int, postorder []int) *TreeNode {
	return buildTree(inorder, postorder)
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	l := len(inorder)
	//构建map，以便快速查询max所对应的位置
	for i := 0; i < l; i++ {
		inmap106[inorder[i]] = i
	}

	return traverse106(inorder, postorder, 0, l-1, 0, l-1)
}

func traverse106(inorder, postorder []int, ilo, ihi, plo, phi int) *TreeNode {
	if plo > phi {
		return nil
	}
	//根节点为postorder的最后一个值
	rootVal := postorder[phi]
	//得到inorder的根节点的位置
	idx := inmap106[rootVal]
	root := &TreeNode{rootVal, nil, nil}
	/*
		postorder的
			左子树的启始位置=原启始位置
			左子树的结束位置为 = 原启始位置+左子树的长度-1
			右子树的启始位置=左子树的结束位置+1
			右子树的结束位置=数组的最后位置-1
		inorder的
			左子树的启始位置=原启始位置
			左子树的结束位置为 = 原启始位置+左子树的长度
			右子树的启始位置=根节点+1
			右子树的结束位置=数组的最后位置
	*/
	leftLen := idx - ilo
	root.Left = traverse106(inorder, postorder, ilo, idx-1, plo, plo+leftLen-1)
	root.Right = traverse106(inorder, postorder, idx+1, ihi, plo+leftLen, phi-1)
	return root
}
