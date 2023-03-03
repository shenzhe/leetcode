package code

func Execuate_230(root *TreeNode, k int) int {
	return kthSmallest(root, k)
}

var rank230 int
var res230 int

func kthSmallest(root *TreeNode, k int) int {
	rank230 = 0
	res230 = -1
	traverse230(root, k)
	return res230
}

func traverse230(root *TreeNode, k int) {
	if root == nil {
		return
	}
	traverse230(root.Left, k)
	rank230++
	if rank230 == k {
		res230 = root.Val
		return
	}
	// fmt.Println(root.Val)
	traverse230(root.Right, k)
}
