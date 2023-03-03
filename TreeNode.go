package code

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TreeNode2 struct {
	Val   int
	Left  *TreeNode2
	Right *TreeNode2
	Next  *TreeNode2
}

func PreTraverseTreeNode(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Println(node.Val)
	PreTraverseTreeNode(node.Left)
	PreTraverseTreeNode(node.Right)
}

func InTraverseTreeNode(node *TreeNode) {
	if node == nil {
		return
	}
	InTraverseTreeNode(node.Left)
	fmt.Println(node.Val)
	InTraverseTreeNode(node.Right)
}

func PostTraverseTreeNode(node *TreeNode) {
	if node == nil {
		return
	}
	PostTraverseTreeNode(node.Left)
	PostTraverseTreeNode(node.Right)
	fmt.Println(node.Val)
}

func BuildBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	node := &TreeNode{Val: nums[mid]}
	node.Left = BuildBST(nums[:mid])
	node.Right = BuildBST(nums[mid+1:])
	return node
}
