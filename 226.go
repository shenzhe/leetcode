package code

func InvertTree(root *TreeNode) *TreeNode {

	if root == nil {
		return root
	}

	InvertTree(root.Left)
	InvertTree(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}
