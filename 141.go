package code

func Execuate_141(head *ListNode) bool {
	return hasCycle(head)
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return false
	}
	//通过快慢指针，如果有环，快指针最终会和慢指针在某一点相遇
	fast, slow := head, head
	for fast != nil && fast.Next != nil && slow != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
