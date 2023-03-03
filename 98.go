package code

import "math"

var num_98 int
var res_98 bool

func Execuate_98(root *TreeNode) bool {
	return isValidBST(root)
}

func isValidBST(root *TreeNode) bool {
	num_98 = math.MinInt
	res_98 = true
	traverse_98(root)
	return res_98
}

func traverse_98(root *TreeNode) {
	if root == nil {
		return
	}

	traverse_98(root.Left)
	if root.Val < num_98 {
		res_98 = false
		return
	}
	num_98 = root.Val
	traverse_98(root.Right)
}
