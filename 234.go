package code

func Execuate_234(head *ListNode) bool {
	return isPalindrome(head)
}

func isPalindrome(head *ListNode) bool {
	//一个节点，为true
	if head.Next == nil {
		return true
	}
	//两个节点，判断两个数是否相等
	if head.Next.Next == nil {
		return head.Next.Val == head.Val
	}

	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			//说明奇数
			break
		}
		fast = fast.Next.Next
	}

	//反转后面链表
	for slow != nil {
		next := slow.Next
		slow.Next = fast
		fast = slow
		slow = next
	}

	//首尾遍历，比较值是否相关
	for fast != nil {
		if head.Val != fast.Val {
			return false
		}
		head = head.Next
		fast = fast.Next
	}

	return true
}
