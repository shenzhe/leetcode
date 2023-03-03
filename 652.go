package code

import "strconv"

var treeMap map[string]int
var res []*TreeNode

func Execuate_652(root *TreeNode) []*TreeNode {
	return findDuplicateSubtrees(root)
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	treeMap = make(map[string]int)
	res = make([]*TreeNode, 0)
	traverse652(root)
	return res
}

func traverse652(root *TreeNode) string {
	if root == nil {
		return "#"
	}

	left := traverse652(root.Left)
	right := traverse652(root.Right)

	key := left + "," + right + "," + strconv.Itoa(root.Val)

	//看是否在treemap里
	if val, ok := treeMap[key]; ok {
		//仅第二次，表示第一次重复，加入返回数组中
		if val == 1 {
			res = append(res, root)
		}
		treeMap[key] = val + 1
	} else {
		treeMap[key] = 1
	}

	return key
}
