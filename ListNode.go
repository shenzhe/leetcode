package code

type ListNode struct {
	Val  int
	Next *ListNode
}

func BuildListNode(nums []int) *ListNode {
	res := &ListNode{}
	p := res
	for _, v := range nums {
		p.Next = &ListNode{
			Val: v,
		}
		p = p.Next
	}
	return res.Next
}

func BuildCycleListNode(nums []int, k int) *ListNode {
	res := &ListNode{}
	p := res
	cycle := &ListNode{}
	l := len(nums)
	for i := 0; i < l; i++ {
		p.Next = &ListNode{
			Val: nums[i],
		}
		if k == i {
			cycle.Next = p
		}
		p = p.Next
	}
	if cycle.Next != nil {
		p.Next = cycle.Next
	}
	return res.Next
}

func ListNodeToArr(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	res := make([]int, 0)
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}
