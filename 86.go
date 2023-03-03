package code

func Execuate_86(head *ListNode, x int) *ListNode {
	return partition(head, x)
}

func partition(head *ListNode, x int) *ListNode {
	small := &ListNode{}
	big := &ListNode{}
	p := head
	s, b := small, big
	for p != nil {
		if p.Val < x {
			s.Next = p
			s = s.Next
			// fmt.Println("small", p.Val)
		} else {
			b.Next = p
			b = b.Next
			// fmt.Println("big", p.Val)
		}
		p = p.Next
	}
	// fmt.Println(s.Val, big.Next.Val)
	if big.Next != nil {
		b.Next = nil
		s.Next = big.Next
	}
	return small.Next
}
