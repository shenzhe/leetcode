package code

func ConstructMaximumBinaryTree(nums []int) *TreeNode {
	return traverse(nums, 0, len(nums)-1)
}

func traverse(nums []int, lo, hi int) *TreeNode {
	if lo > hi {
		return nil
	}

	//找到最大值以及所在的位置
	idx, max := -1, -1

	for i := lo; i <= hi; i++ {
		//找到最大值并更新max和idx
		if nums[i] > max {
			max = nums[i]
			idx = i
		}
	}

	//根据最大值构建根节点
	root := &TreeNode{max, nil, nil}
	//左树=最大值左边
	root.Left = traverse(nums, 0, idx-1)
	//右树=最大值的右边
	root.Right = traverse(nums, idx+1, hi)
	return root
}
