package code

var sum int

func Execuate_538(root *TreeNode) *TreeNode {
	return convertBST(root)
}

func convertBST(root *TreeNode) *TreeNode {
	sum = 0
	traverse_538(root)
	return root
}

func traverse_538(root *TreeNode) {
	if root == nil {
		return
	}
	traverse_538(root.Right)
	sum += root.Val
	root.Val = sum
	traverse_538(root.Left)
}
