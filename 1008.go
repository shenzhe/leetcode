package code

func Execuate1008(preorder []int) *TreeNode {
	return bstFromPreorder(preorder)
}

func bstFromPreorder(preorder []int) *TreeNode {
	l := len(preorder)

	if l < 1 {
		return nil
	}

	if l == 1 {
		return &TreeNode{preorder[0], nil, nil}
	}

	return traverse1008(preorder, 0, l-1)
}

func traverse1008(preorder []int, lo, hi int) *TreeNode {

	if lo > hi {
		return nil
	}

	rootVal := preorder[lo]

	root := &TreeNode{rootVal, nil, nil}

	if preorder[lo+1] > rootVal {
		root.Left = traverse1008(preorder, lo+1, hi)
	} else {
		root.Right = traverse1008(preorder, lo+1, hi)
	}
	return root
}
