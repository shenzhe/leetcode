package code

func Execuate_1373(root *TreeNode) int {
	return maxSumBST(root)
}

var maxNum_1373 int
var def_1373, nilDef_1373, res_1373 [4]int

func maxSumBST(root *TreeNode) int {
	maxNum_1373 = 0
	nilDef_1373 = [4]int{1, 100000, -100000, 0}
	traverse_1373(root)
	return maxNum_1373
}

func traverse_1373(root *TreeNode) [4]int {
	if root == nil {
		/*
			int[0] 记录以 root 为根的⼆叉树是否是 BST，若为 1 则说明是 BST，若为 0 则说明不是 BST；
			int[1] 记录以 root 为根的⼆叉树所有节点中的最⼩值；
			int[2] 记录以 root 为根的⼆叉树所有节点中的最⼤值；
			int[3] 记录以 root 为根的⼆叉树所有节点值之和。
		*/
		return nilDef_1373
	}

	left_1373 := traverse_1373(root.Left)
	right_1373 := traverse_1373(root.Right)

	/*
		形成新的BST,需要满足的条件是：
		1、左右子树都是BST
		2、root.Val比左树的最大值要大，比右树的最小值要小
	*/
	if left_1373[0] == 1 && right_1373[0] == 1 && root.Val > left_1373[2] && root.Val < right_1373[1] {
		min, max := root.Val, root.Val
		//计算以此root为根的BST的最小值
		if min > left_1373[1] {
			min = left_1373[1]
		}
		//计算以此root为根的BST的最大值
		if max < right_1373[2] {
			max = right_1373[2]
		}
		//计算以此root为根的BST的累积和
		sum := root.Val + left_1373[3] + right_1373[3]
		if sum > maxNum_1373 {
			maxNum_1373 = sum
		}
		res_1373[0] = 1
		res_1373[1] = min
		res_1373[2] = max
		res_1373[3] = sum
		return res_1373
	}

	return def_1373

}
