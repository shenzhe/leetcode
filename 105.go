package code

// 构建inorder的map,方便通过根节点的索引位置
var inmap = make(map[int]int)

func BuildTree(preorder []int, inorder []int) *TreeNode {
	l := len(inorder)

	for i := 0; i < l; i++ {
		inmap[inorder[i]] = i
	}

	return traverse105(preorder, inorder, 0, l-1, 0, l-1)
}

func traverse105(preorder, inorder []int, plo, phi, ilo, ihi int) *TreeNode {
	if plo > phi {
		return nil
	}

	//根节点在preorder的第一个位置
	rootVal := preorder[plo]
	//找出根节点在inorder的位置
	iIdx := inmap[rootVal]

	//通过inorder算出左子树的长度 = 根节点位置-起始值
	leftLen := iIdx - ilo
	// rightLen := len(preorder) - iIdx - 1
	root := &TreeNode{rootVal, nil, nil}
	/*
		对于preorder来说：
			左子树的起始位置=第一个位置+1(去掉根节点之后)， 结束位置=启始位置+左子树的长度
			右子树的超始位置=左子树的结束位置+1，结束位置为最后一个位置
		对于inorder来说 （根节点的左边为左子树，右边为右子树）：
			左子树的超始位置=第一个位置，结束位置=根节点所在的位置-1
			右子树的超始位置=根节点所在的位置+1， 结束位置=最后一个位置
	*/
	leftTree := traverse105(preorder, inorder, plo+1, plo+leftLen, ilo, iIdx-1)
	rightTree := traverse105(preorder, inorder, plo+leftLen+1, phi, iIdx+1, ihi)
	root.Left = leftTree
	root.Right = rightTree
	return root
}
