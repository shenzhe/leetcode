package code

func Execuate_450(root *TreeNode, key int) *TreeNode {
	return deleteNode(root, key)
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key == root.Val {
		//找到，删除
		if root.Left == nil {
			//左子树为空，直接返回右子树替换root
			return root.Right
		}
		if root.Right == nil {
			//右子树为空，直接返回左子树替换root
			return root.Left
		}
		//还有多个子节点，找到右子树的最小节点，以便替换当前的root
		minNode := getMinNode(root.Right)
		//删除右子数的minNode
		root.Right = deleteNode(root.Right, minNode.Val)
		//把MinNode交换到当前点，完成真正的删除
		minNode.Left = root.Left
		minNode.Right = root.Right
		return minNode
	} else if key > root.Val {
		//往右子数找
		root.Right = deleteNode(root.Right, key)
	} else {
		//往左子树找
		root.Left = deleteNode(root.Left, key)
	}

	return root
}

func getMinNode(root *TreeNode) *TreeNode {
	for root.Left != nil {
		root = root.Left
	}
	return root
}
