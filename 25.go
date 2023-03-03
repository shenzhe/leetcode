package code

func Execuate_25(head *ListNode, k int) *ListNode {
	return reverseKGroup(head, k)
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	start, end := head, head

	for i := 0; i < k; i++ {
		if end == nil {
			return head
		}
		end = end.Next
	}
	newHead := reverseK(start, end)
	start.Next = reverseKGroup(end, k)

	return newHead

}

func reverseK(head *ListNode, end *ListNode) *ListNode {
	var pre *ListNode
	pre = nil
	for head != end {
		//取出next
		next := head.Next
		//当前节点指向pre
		head.Next = pre
		//把当前节点缓存到pre
		pre = head
		head = next
	}
	return pre
}

func reverseKGroup2(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	//计算count
	count := 0
	p := head
	for p != nil {
		count++
		p = p.Next
	}
	//反转次数为0
	loopNum := count / k
	if loopNum == 0 {
		return head
	}
	start, newHead, last := head, head, head
	var preLast *ListNode
	preLast = nil
	for i := 0; i < loopNum; i++ {
		newHead, start, last = reverseK2(start, k)
		if p == nil {
			//第一个newHead就是反转后的head
			p = newHead
		}
		if preLast != nil {
			//把上一个反转后的链表最后一个node链到下一个反转链表的头node，完成两个链表的连接
			preLast.Next = newHead
		}
		//把当前反转后的最后一个node缓存，以备链接到下一个链表
		preLast = last
		// fmt.Println("reverse", head.Val, newHead.Val, start.Val, last.Val)
	}

	//还有剩余不足K个的链表，直接链接到最后一个反转链表的尾node
	if start != nil {
		preLast.Next = start
	}

	return p

}

func reverseK2(head *ListNode, k int) (*ListNode, *ListNode, *ListNode) {
	var pre, last *ListNode
	pre = nil
	last = nil
	for k > 0 {
		if last == nil {
			//第一个head就是反转后的尾node
			last = head
		}
		//取出next
		next := head.Next
		//当前节点指向pre
		head.Next = pre
		//把当前节点缓存到pre,也就是反转后的head
		pre = head
		//需要反转下一个链表的head
		head = next
		k--
	}
	return pre, head, last
}
