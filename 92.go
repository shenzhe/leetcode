package code

import "fmt"

func Execuate_92(head *ListNode, left int, right int) *ListNode {
	return reverseBetween(head, left, right)
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	//只有一个节点，不用翻转
	if head.Next == nil {
		return head
	}

	//left和right指向同一节点，不用翻转
	if left == right {
		return head
	}
	//只有两个节点，且right=2，直接翻转
	if head.Next.Next == nil && right == 2 {
		next := head.Next
		head.Next = nil
		next.Next = head
		return next
	}

	return reverseBetweenByRecursion(head, left, right)
}

func reverseBetweenByRecursion(head *ListNode, left, right int) *ListNode {
	if left == 1 {
		//当递归到left=1，表示接下来的链表要开始反转了
		return reverse_92(head, right)
	}
	//递归head.next，right需要-1，因为要保持right-left值不变
	head.Next = reverseBetweenByRecursion(head.Next, left-1, right-1)
	return head
}

// 记录right之后的节点，right个节点完成之后，把right节点指向next,完成整链的链接
var next *ListNode

func reverse_92(head *ListNode, right int) *ListNode {
	if right == 1 {
		//表示，前right个已反转完成，下一个节点存入next,返回当前节点
		next = head.Next
		return head
	}
	//递归head.Next，一个一个反转
	node := reverse_92(head.Next, right-1)
	head.Next.Next = head
	head.Next = next
	return node
}

func reverseBetweenByIteration(head *ListNode, left, right int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	p := dummy
	//反转前的节点
	preHead := dummy
	//记录反转后的最后节点（即为反转前的第一个节点）
	last := p.Next
	//缓存的上一个节点，当遍历到right时，即为反转后的第一个节点
	var pre *ListNode
	for p != nil {
		//开始反转前，记录preHead的节点
		if left == 1 {
			preHead = p
		}
		if left < 1 {
			//开始反转
			if left == 0 {
				//反转前的第一个节点即为反转链表的最后一个结点
				last = p
			}
			//取出下一个结点
			next := p.Next
			//当前结点指向pre
			p.Next = pre
			//pre保存为当前结点
			pre = p
			//指向下一个结点，继续遍历
			p = next
			left--
			right--
			//停止反转
			if right < 0 {
				break
			}
		} else {
			p = p.Next
			left--
			right--
		}
	}
	fmt.Println("pre=", preHead.Val, "f=", pre.Val, "last=", last.Val)
	preHead.Next = pre
	last.Next = p

	return dummy.Next
}
