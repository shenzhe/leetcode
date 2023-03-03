package code

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var pre *ListNode
	pre = nil

	for head != nil {
		//找到下个节点
		next := head.Next
		//当前节点指向pre
		head.Next = pre
		//把当前节点缓存到pre
		pre = head
		head = next
	}

	return pre
}
