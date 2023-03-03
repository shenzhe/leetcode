package code

func Execuate_142(head *ListNode) *ListNode {
	return detectCycle(head)
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}
	fast, slow := head, head
	for fast != nil && fast.Next != nil && slow != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}

	for head != nil {
		if head == slow {
			return head
		}
		head = head.Next
		slow = slow.Next
	}
	return nil
}
