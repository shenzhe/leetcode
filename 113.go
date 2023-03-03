package code

var ans [][]int

func Execuate_113(root *TreeNode, targetSum int) [][]int {
	return pathSum(root, targetSum)
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	ans = make([][]int, 0)
	if root == nil {
		return ans
	}
	path := make([]int, 0)
	traverse_113(root, path, targetSum)
	return ans
}

func traverse_113(root *TreeNode, path []int, targetSum int) []int {
	if root == nil {
		return path
	}
	path = append(path, root.Val)
	if root.Left == nil && root.Right == nil {
		if targetSum == root.Val {
			//加入ans
			newpath := make([]int, len(path))
			copy(newpath, path)
			ans = append(ans, newpath)
			return path[:len(path)-1]
		}
	}
	path = traverse_113(root.Left, path, targetSum-root.Val)
	path = traverse_113(root.Right, path, targetSum-root.Val)
	path = path[:len(path)-1]
	return path
}
