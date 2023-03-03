package code

func Execuate_19(head *ListNode, n int) *ListNode {
	return removeNthFromEnd(head, n)
}

func Execuate_19_2(head *ListNode, n int) *ListNode {
	return removeNthFromEnd2(head, n)
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//处理只有一个node
	if head.Next == nil {
		return nil
	}
	//处量只有二个node
	if head.Next.Next == nil {
		if n == 2 {
			return head.Next
		} else {
			head.Next = nil
			return head
		}
	}

	fast := head
	// fmt.Println("start", fast.Val, n)
	for n > 1 {
		// fmt.Println("-", fast.Val, n)
		fast = fast.Next
		n--
	}
	// fmt.Println("fast", fast.Val, n)
	//表示要删除第一个节点,直接返回第二个节点
	if fast.Next == nil {
		return head.Next
	}
	pre := head
	//找到倒数K的前一个节点
	for fast.Next.Next != nil {
		pre = pre.Next
		fast = fast.Next
	}
	// fmt.Println("fast2", pre.Val, fast.Val)
	//倒数K的前一个节点的next指向倒数K的next，完成倒数K的删除
	pre.Next = pre.Next.Next
	return head
}

/**/
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	//处理只有一个node
	if head.Next == nil {
		return nil
	}
	//处量只有二个node
	if head.Next.Next == nil {
		if n == 2 {
			return head.Next
		} else {
			head.Next = nil
			return head
		}
	}
	//用一个虚拟节点，可以节省很多边缘的判断（对比上一个解答）
	dummy := &ListNode{}
	dummy.Next = head
	fast := dummy
	//快指针向前走N格
	for i := 0; i <= n; i++ {
		if fast == nil {
			return dummy.Next
		}
		fast = fast.Next
	}
	//快指针走到尾时，慢指针正好走到倒数第N格的前一格
	slow := dummy
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	//慢指针的next=next.next,完成倒数第N格的删除
	slow.Next = slow.Next.Next

	return dummy.Next
}
